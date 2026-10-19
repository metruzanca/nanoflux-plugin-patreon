// Command nanoflux-plugin-patreon serves the Patreon plugin over gRPC so
// nanoflux can load it as an external plugin.
//
// Build it and drop the binary into the host's plugins directory
// (NF_PLUGINS_DIR, ./plugins by default):
//
//	GOOS=linux CGO_ENABLED=0 go build -o ../nanoflux-plugin-patreon .
package main

import (
	goplugin "github.com/hashicorp/go-plugin"

	patreon "nanoflux-plugin-patreon/patreon"

	"github.com/metruzanca/nanoflux/pluginapi"
)

func main() {
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: pluginapi.Handshake,
		Plugins:         pluginapi.PluginSet(patreon.Plugin{}),
		GRPCServer:      goplugin.DefaultGRPCServer,
	})
}
