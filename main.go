package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	
	// Plugins
  _ "github.com/caddyserver/dnsproviders/cloudflare"
	_ "github.com/epicagency/caddy-expires"
  _ "github.com/captncraig/caddy-realip"
)

func main() {
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
