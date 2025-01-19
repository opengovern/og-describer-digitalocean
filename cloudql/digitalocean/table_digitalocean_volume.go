package digitalocean

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-digitalocean/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_volume",
		Description: "DigitalOcean Block Storage Volumes provide expanded storage capacity for your Droplets and can be moved between Droplets within a specific region. Volumes function as raw block devices, meaning they appear to the operating system as locally attached storage which can be formatted using any file system supported by the OS.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanVolume,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanVolume,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.ID"),
				Description: "The unique identifier for the block storage volume."},
			{Name: "name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.Name"),
				Description: "A human-readable name for the block storage volume. Must be lowercase and be composed only of numbers, letters and \"-\", up to a limit of 64 characters. The name must begin with a letter."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Volume.CreatedAt"),
				Description: "Time when the block storage volume was created."},
			{Name: "description", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.Description"),
				Description: "An optional free-form text field to describe a block storage volume."},
			{Name: "droplet_ids", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Volume.DropletIDs"),
				Description: "An array containing the IDs of the Droplets the volume is attached to. Note that at this time, a volume can only be attached to a single Droplet."},
			{Name: "filesystem_label", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.FilesystemLabel"),
				Description: "The label currently applied to the filesystem."},
			{Name: "filesystem_type", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Volume.FilesystemType"),
				Description: "The type of filesystem currently in-use on the volume."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Volume.Region.Slug"), Description: "The unique slug identifier for the region the volume is deployed in."},
			{Name: "region_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Volume.Region.Name"), Description: "The name of the region the volume is deployed in."},
			{Name: "size_gigabytes", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Volume.SizeGigaBytes"),
				Description: "The size of the block storage volume in GiB (1024^3)."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Volume.Tags"), Description: "An array of tags the volume has been tagged with"},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.URN"), Description: "The uniform resource name (URN) for the volume."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.URN").Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Volume.Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Volume.Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}
