package digitalocean

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-digitalocean/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanSnapshot(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_snapshot",
		Description: "Snapshots are saved instances of a Droplet or a block storage volume.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanSnapshot,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanSnapshot,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.ID"),
				Description: "The unique identifier for the snapshot."},
			{Name: "name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.Name"),
				Description: "A human-readable name for the snapshot."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Snapshot.Created"),
				Description: "Time when the block storage volume was created."},
			{Name: "min_disk_size", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Snapshot.MinDiskSize"),
				Description: "The minimum size in GB required for a volume or Droplet to use this snapshot."},
			{Name: "regions", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.Regions"),
				Description: "An array of regions the snapshot is available in. The region slug is used."},
			{Name: "resource_id", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.ResourceID"),
				Description: "A unique identifier for the resource that the action is associated with."},
			{Name: "resource_type", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.ResourceType"),
				Description: "The type of resource that the action is associated with."},
			{Name: "size_gigabytes", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Snapshot.SizeGigaBytes"),
				Description: "The billable size of the snapshot in gigabytes."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Snapshot.Tags"), Description: "An array of Tags the snapshot has been tagged with."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(snapshotToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Snapshot.Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Snapshot.Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func snapshotToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(opengovernance.DigitalOceanSnapshot).Description.Snapshot
	return "do:snapshot:" + i.ID, nil
}
