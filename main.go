package main

import (
	"context"
	"fmt"

	"github.com/caddyserver/xcaddy"
)

func main() {
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
			CaddyVersion: "v2.0.0-rc.3",
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
