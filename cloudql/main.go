package main

import (
	"github.com/opengovern/og-describer-digitalocean/cloudql/openai"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: openai.Plugin})
}
