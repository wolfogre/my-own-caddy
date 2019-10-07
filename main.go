package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	// plug in plugins here, for example:
	// _ "import/path/here"
	_ "github.com/aablinov/caddy-geoip"      // http.genip
	_ "github.com/captncraig/cors"           // http.cors
	_ "github.com/epicagency/caddy-expires"  // http.expire
	_ "github.com/nicolasazrak/caddy-cache"  // http.cache
	_ "github.com/wolfogre/caddy-prometheus" // http.prometheus
)

func main() {
	// optional: disable telemetry
	//caddymain.EnableTelemetry = false
	caddymain.EnableTelemetry = true
	caddymain.Run()
}
