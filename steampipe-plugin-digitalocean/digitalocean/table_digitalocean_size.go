package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanSize(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_size",
		Description: "The sizes objects represent different packages of hardware resources that can be used for Droplets. This includes the amount of RAM, the number of virtual CPUs, disk space, and transfer. The size object also includes the pricing details and the regions that the size is available in.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanSize,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "slug", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Size.Slug"),
				Description: "A unique slug identifier for the size."},
			// Other columns
			{Name: "available", Type: proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Size.Available"),
				Description: "This is a boolean value that represents whether new Droplets can be created with this size."},
			{Name: "disk", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Size.Disk"),
				Description: "The amount of disk space set aside for Droplets of this size. The value is represented in gigabytes."},
			{Name: "memory", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Size.Memory"),
				Description: "The amount of RAM allocated to Droplets created of this size. The value is represented in megabytes."},
			{Name: "price_hourly", Type: proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Size.PriceHourly"),
				Description: "This describes the price of the Droplet size as measured hourly. The value is measured in US dollars."},
			{Name: "price_monthly", Type: proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Size.PriceMonthly"),
				Description: "This attribute describes the monthly cost of this Droplet size if the Droplet is kept for an entire month. The value is measured in US dollars."},
			{Name: "regions", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Size.Regions"),
				Description: "An array containing the region slugs where this size is available for Droplet creates."},
			{Name: "transfer", Type: proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Size.Transfer"),
				Description: "The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes."},
			{Name: "vcpus", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Size.Vcpus"),
				Description: "The integer of number CPUs allocated to Droplets of this size."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(sizeToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Size.Slug"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func sizeToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(opengovernance.DigitalOceanSize).Description.Size
	return "do:size:" + i.Slug, nil
}
