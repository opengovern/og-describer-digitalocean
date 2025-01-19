package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanFloatingIP(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_floating_ip",
		Description: "DigitalOcean Floating IPs are publicly-accessible static IP addresses that can be mapped to one of your Droplets. They can be used to create highly available setups or other configurations requiring movable addresses.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanFloatingIP,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("ip"),
			Hydrate:    opengovernance.GetDigitalOceanFloatingIP,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "ip", Type: proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.FloatingIP.IP"),
				Description: "The public IP address of the floating IP. It also serves as its identifier."},
			// Other columns
			{Name: "droplet", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.FloatingIP.Droplet"), Description: "The Droplet that the floating IP has been assigned to."},
			{Name: "droplet_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.FloatingIP.Droplet.ID"), Description: "ID of the Droplet that the floating IP has been assigned to."},
			{Name: "region", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.FloatingIP.Region"), Description: "The region that the floating IP is reserved to."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.FloatingIP.Region.Slug"), Description: "The slug of the region that the floating IP is reserved to."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.URN"), Description: "The uniform resource name (URN) for the database."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromMethod("Description.URN").Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.FloatingIP.IP"), Description: resourceInterfaceDescription("title")},
		}),
	}
}
