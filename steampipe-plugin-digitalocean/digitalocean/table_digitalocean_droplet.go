package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanDroplet(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_droplet",
		Description: "A Droplet is a DigitalOcean virtual machine.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanDroplet,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanDroplet,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Droplet.ID"),
				Description: "A unique identifier for each Droplet instance."},
			{Name: "name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Droplet.Name"),
				Description: "The human-readable name set for the Droplet instance."},
			// Other columns
			{Name: "backup_ids", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Droplet.BackupIDs"),
				Description: "An array of backup IDs of any backups that have been taken of the Droplet instance."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Droplet.Created"),
				Description: "Time when the Droplet was created."},
			{Name: "disk", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Droplet.Disk"),
				Description: "The size of the Droplet's disk in gigabytes."},
			{Name: "features", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Droplet.Features"),
				Description: "An array of features enabled on this Droplet."},
			{Name: "image", Type: proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Droplet.Image"), Description: "Information about the base image used to create the Droplet instance."},
			{Name: "kernel", Type: proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Droplet.Kernel"), Description: "The current kernel. This will initially be set to the kernel of the base image when the Droplet is created."},
			{Name: "locked", Type: proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Droplet.Locked"),
				Description: "A boolean value indicating whether the Droplet has been locked, preventing actions by users."},
			{Name: "memory", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Droplet.Memory"),
				Description: "Memory of the Droplet in megabytes."},
			{Name: "networks", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Droplet.Networks"),
				Description: "The details of the network that are configured for the Droplet instance. This is an object that contains keys for IPv4 and IPv6. The value of each of these is an array that contains objects describing an individual IP resource allocated to the Droplet. These will define attributes like the IP address, netmask, and gateway of the specific network depending on the type of network it is."},
			{Name: "next_backup_window_start", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Droplet.NextBackupWindow.Start").Transform(timestampToIsoTimestamp), Description: "Start time of the window during which the backup will start."},
			{Name: "next_backup_window_end", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Droplet.NextBackupWindow.End").Transform(timestampToIsoTimestamp), Description: "End time of the window during which the backup will start."},
			{Name: "private_ipv4", Type: proto.ColumnType_IPADDR, Transform: transform.FromMethod("Description.Droplet.PrivateIPv4"), Description: "Private IPv4 address of the Droplet."},
			{Name: "public_ipv4", Type: proto.ColumnType_IPADDR, Transform: transform.FromMethod("Description.Droplet.PublicIPv4"), Description: "Public IPv4 address of the Droplet."},
			{Name: "public_ipv6", Type: proto.ColumnType_IPADDR, Transform: transform.FromMethod("Description.Droplet.PublicIPv6"), Description: "Public IPv6 address of the Droplet."},
			{Name: "region", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Droplet.Region"), Description: "Information about region that the Droplet instance is deployed in."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Droplet.Region.Slug"), Description: "The unique slug identifier for the region the Droplet is deployed in."},
			{Name: "size", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Droplet.Size"), Description: "Information about the size of the Droplet. Note: Due to resize operations, the disk column is more accurate than the disk field in this size data."},
			{Name: "size_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Droplet.Size.Slug"), Description: "The unique slug identifier for the size of this Droplet."},
			{Name: "snapshot_ids", Type: proto.ColumnType_JSON, Description: "An array of snapshot IDs of any snapshots created from the Droplet instance."},
			{Name: "status", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Droplet.Status"),
				Description: "A status string indicating the state of the Droplet instance.  This may be \"new\", \"active\", \"off\", or \"archive\"."},
			{Name: "tags_src", Type: proto.ColumnType_JSON,
				Transform: transform.FromField("Description.Droplet.Tags"), Description: "An array of tags the Droplet has been tagged with."},
			{Name: "urn", Type: proto.ColumnType_STRING,
				Transform: transform.FromField("Description.URN"), Description: "The uniform resource name (URN) for the Droplet."},
			{Name: "vcpus",
				Transform: transform.FromField("Description.Droplet.Vcpus"),
				Type:      proto.ColumnType_INT, Description: "The number of virtual CPUs."},
			{Name: "volume_ids", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Droplet.VolumeIDs"),
				Description: "A flat array including the unique identifier for each Block Storage volume attached to the Droplet."},
			{Name: "vpc_uuid", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Droplet.VPCUUID"),
				Description: "A string specifying the UUID of the VPC to which the Droplet is assigned."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.URN").Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Droplet.Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Droplet.Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}
