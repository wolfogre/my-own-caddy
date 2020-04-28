// Copyright 2020 Matthew Holt
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xcaddy

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func (b Builder) newEnvironment(ctx context.Context) (*environment, error) {
	// assume Caddy v2 if no semantic version is provided
	caddyModulePath := defaultCaddyModulePath
	if !strings.HasPrefix(b.CaddyVersion, "v") || !strings.Contains(b.CaddyVersion, ".") {
		caddyModulePath += "/v2"
	}
	caddyModulePath, err := versionedModulePath(caddyModulePath, b.CaddyVersion)
	if err != nil {
		return nil, err
	}

	// clean up any SIV-incompatible module paths real quick
	for i, p := range b.Plugins {
		b.Plugins[i].ModulePath, err = versionedModulePath(p.ModulePath, p.Version)
		if err != nil {
			return nil, err
		}
	}

	// create the context for the main module template
	tplCtx := goModTemplateContext{
		CaddyModule: caddyModulePath,
	}
	for _, p := range b.Plugins {
		tplCtx.Plugins = append(tplCtx.Plugins, p.ModulePath)
	}

	// evaluate the template for the main module
	var buf bytes.Buffer
	tpl, err := template.New("main").Parse(mainModuleTemplate)
	if err != nil {
		return nil, err
	}
	err = tpl.Execute(&buf, tplCtx)
	if err != nil {
		return nil, err
	}

	// create the folder in which the build environment will operate
	tempFolder, err := newTempFolder()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err2 := os.RemoveAll(tempFolder)
			if err2 != nil {
				err = fmt.Errorf("%w; additionally, cleaning up folder: %v", err, err2)
			}
		}
	}()
	log.Printf("[INFO] Temporary folder: %s", tempFolder)

	// write the main module file to temporary folder
	mainPath := filepath.Join(tempFolder, "main.go")
	log.Printf("[INFO] Writing main module: %s", mainPath)
	err = ioutil.WriteFile(mainPath, buf.Bytes(), 0644)
	if err != nil {
		return nil, err
	}

	env := &environment{
		caddyVersion:    b.CaddyVersion,
		plugins:         b.Plugins,
		caddyModulePath: caddyModulePath,
		tempFolder:      tempFolder,
	}

	// initialize the go module
	log.Println("[INFO] Initializing Go module")
	cmd := env.newCommand("go", "mod", "init", "caddy")
	err = env.runCommand(ctx, cmd, 10*time.Second)
	if err != nil {
		return nil, err
	}

	// specify module replacements before pinning versions
	replaced := make(map[string]string)
	for _, r := range b.Replacements {
		log.Printf("[INFO] Replace %s => %s", r.Old, r.New)
		cmd := env.newCommand("go", "mod", "edit",
			"-replace", fmt.Sprintf("%s=%s", r.Old, r.New))
		err := env.runCommand(ctx, cmd, 10*time.Second)
		if err != nil {
			return nil, err
		}
		replaced[r.Old] = r.New
	}

	// check for early abort
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// pin versions by populating go.mod, first for Caddy itself and then plugins
	log.Println("[INFO] Pinning versions")
	err = env.execGoGet(ctx, caddyModulePath, env.caddyVersion)
	if err != nil {
		return nil, err
	}
	for _, p := range b.Plugins {
		// if module is locally available; do not "go get" it
		if replaced[p.ModulePath] != "" {
			continue
		}
		err = env.execGoGet(ctx, p.ModulePath, p.Version)
		if err != nil {
			return nil, err
		}
		// check for early abort
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}

	log.Println("[INFO] Build environment ready")

	return env, nil
}

type environment struct {
	caddyVersion    string
	plugins         []Dependency
	caddyModulePath string
	tempFolder      string
}

// Close cleans up the build environment, including deleting
// the temporary folder from the disk.
func (env environment) Close() error {
	log.Printf("[INFO] Cleaning up temporary folder: %s", env.tempFolder)
	return os.RemoveAll(env.tempFolder)
}

func (env environment) newCommand(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Dir = env.tempFolder
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func (env environment) runCommand(ctx context.Context, cmd *exec.Cmd, timeout time.Duration) error {
	log.Printf("[INFO] exec (timeout=%s): %+v ", timeout, cmd)

	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	// start the command; if it fails to start, report error immediately
	err := cmd.Start()
	if err != nil {
		return err
	}

	// wait for the command in a goroutine; the reason for this is
	// very subtle: if, in our select, we do `case cmdErr := <-cmd.Wait()`,
	// then that case would be chosen immediately, because cmd.Wait() is
	// immediately available (even though it blocks for potentially a long
	// time, it can be evaluated immediately). So we have to remove that
	// evaluation from the `case` statement.
	cmdErrChan := make(chan error)
	go func() {
		cmdErrChan <- cmd.Wait()
	}()

	// unblock either when the command finishes, or when the done
	// channel is closed -- whichever comes first
	select {
	case cmdErr := <-cmdErrChan:
		// process ended; report any error immediately
		return cmdErr
	case <-ctx.Done():
		// context was canceled, either due to timeout or
		// maybe a signal from higher up canceled the parent
		// context; presumably, the OS also sent the signal
		// to the child process, so wait for it to die
		select {
		case <-time.After(15 * time.Second):
			cmd.Process.Kill()
		case <-cmdErrChan:
		}
		return ctx.Err()
	}
}

func (env environment) execGoGet(ctx context.Context, modulePath, moduleVersion string) error {
	mod := modulePath
	if moduleVersion != "" {
		mod += "@" + moduleVersion
	}
	cmd := env.newCommand("go", "get", "-d", "-v", mod)
	return env.runCommand(ctx, cmd, 5*time.Minute)
}

type goModTemplateContext struct {
	CaddyModule string
	Plugins     []string
}

const mainModuleTemplate = `package main

import (
	caddycmd "{{.CaddyModule}}/cmd"

	// plug in Caddy modules here
	_ "{{.CaddyModule}}/modules/standard"
	{{- range .Plugins}}
	_ "{{.}}"
	{{- end}}
)

func main() {
	caddycmd.Main()
}
`
