package main

import (
	"github.com/mholt/caddy/caddy/caddymain"
	// plug in plugins here, for example:
	// _ "import/path/here"
	_ "github.com/aablinov/caddy-geoip"
	_ "github.com/nicolasazrak/caddy-cache"
)

func main() {
	// optional: disable telemetry
	//caddymain.EnableTelemetry = false
	caddymain.EnableTelemetry = true
	caddymain.Run()
}
