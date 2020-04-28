package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/blang/semver"
	"github.com/caddyserver/xcaddy"
)

var (
	CaddyVersion = "v2.0.0-rc.3"

	Minor = "0"
	Patch = "1"
)

var (
	version = flag.Bool("version", false, "print version")
)

func main() {
	flag.Parse()

	ver, err := semver.Make(strings.TrimPrefix(CaddyVersion, "v"))
	if err != nil {
		panic(err)
	}

	Major := fmt.Sprintf("%d%02d%02d", ver.Major, ver.Minor, ver.Patch)

	myOwnCaddyVersion := fmt.Sprintf("%s.%s.%s", Major, Minor, Patch)

	if *version {
		fmt.Print(myOwnCaddyVersion)
		return
	}

	platforms := []xcaddy.Platform{
		{
			OS:   "darwin",
			Arch: "amd64",
		},
		{
			OS:   "linux",
			Arch: "amd64",
		},
		{
			OS:   "windows",
			Arch: "amd64",
		},
	}

	for _, platform := range platforms {
		builder := xcaddy.Builder{
			Compile: xcaddy.Compile{
				Platform: platform,
			},
			CaddyVersion: CaddyVersion,
			Plugins:      []xcaddy.Dependency{},
		}
		ext := ""
		if platform.OS == "windows" {
			ext = ".exe"
		}
		output := fmt.Sprintf("dist/caddy_%v_%v/caddy%v", platform.OS, platform.Arch, ext)
		err := builder.Build(context.Background(), output)
		if err != nil {
			panic(err)
		}
	}
}
