package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	// plug in plugins here, for example:
	// _ "import/path/here"
	_ "github.com/aablinov/caddy-geoip"     // http.genip
	_ "github.com/epicagency/caddy-expires" // http.expire
	_ "github.com/miekg/caddy-prometheus"   // http.prometheus
	_ "github.com/nicolasazrak/caddy-cache" // http.cache
)

func main() {
	// optional: disable telemetry
	//caddymain.EnableTelemetry = false
	caddymain.EnableTelemetry = true
	caddymain.Run()
}
