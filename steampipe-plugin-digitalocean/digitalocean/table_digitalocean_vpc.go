package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanVPC(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_vpc",
		Description: "VPCs (virtual private clouds) are virtual networks containing resources that can communicate with each other in full isolation using private IP addresses.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanVPC,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanVPC,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VPC.ID"),
				Description: "A unique ID that can be used to identify and reference the VPC."},
			{Name: "name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VPC.Name"),
				Description: "The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VPC.CreatedAt"),
				Description: "A time value given in ISO8601 combined date and time format."},
			{Name: "description", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VPC.Description"),
				Description: "A free-form text field for describing the VPC's purpose. It may be a maximum of 255 characters."},
			{Name: "ip_range", Type: proto.ColumnType_CIDR,
				Transform:   transform.FromField("Description.VPC.IPRange"),
				Description: "The range of IP addresses in the VPC in CIDR notation."},
			// Rename to avoid conflict with default keyword in postgres
			{Name: "is_default", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.VPC.Default"), Description: "A boolean value indicating whether or not the VPC is the default network for the region. All applicable resources are placed into the default VPC network unless otherwise specified during their creation. The `default` field cannot be unset from `true`. If you want to set a new default VPC network, update the `default` field of another VPC network in the same region. The previous network's `default` field will be set to `false` when a new default VPC has been defined."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.VPC.RegionSlug"), Description: "The slug identifier for the region where the VPC will be created."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.VPC.URN"), Description: "The uniform resource name (URN) for the VPC."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.VPC.URN").Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.VPC.Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}
