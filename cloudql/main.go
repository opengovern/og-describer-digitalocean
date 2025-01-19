package main

import (
	"github.com/opengovern/og-describer-digitalocean/cloudql/digitalocean"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: digitalocean.Plugin})
}
