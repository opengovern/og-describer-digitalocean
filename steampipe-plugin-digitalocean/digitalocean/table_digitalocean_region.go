package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanRegion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_region",
		Description: "A region in DigitalOcean represents a datacenter where Droplets can be deployed and images can be transferred. Each region represents a specific datacenter in a geographic location. Some geographical locations may have multiple \"regions\" available. This means that there are multiple datacenters available within that area.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanRegion,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "slug", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Region.Slug"),
				Description: "A human-readable string that is used as a unique identifier for each region."},
			{Name: "name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Region.Name"),
				Description: "The display name of the region. This will be a full name that is used in the control panel and other interfaces."},
			// Other columns
			{Name: "available", Type: proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Region.Available"),
				Description: "This is a boolean value that represents whether new Droplets can be created in this region."},
			{Name: "features", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Region.Features"),
				Description: "This attribute is set to an array which contains features available in this region."},
			{Name: "sizes", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Region.Sizes"),
				Description: "This attribute is set to an array which contains the identifying slugs for the sizes available in this region."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(regionToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Region.Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func regionToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(opengovernance.DigitalOceanRegion).Description.Region
	return "do:region:" + i.Slug, nil
}
