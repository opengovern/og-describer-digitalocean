package digitalocean

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_image",
		Description: "A DigitalOcean image can be used to create a Droplet and may come in a number of flavors. Currently, there are five types of images: snapshots, backups, applications, distributions, and custom images.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanImage,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "slug"}),
			Hydrate:    opengovernance.GetDigitalOceanImage,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Image.ID"),
				Description: "A unique number that can be used to identify and reference a specific image."},
			{Name: "name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Name"),
				Description: "The display name that has been given to an image. This is what is shown in the control panel and is generally a descriptive title for the image in question."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Image.Created"),
				Description: "Time when the image was created."},
			{Name: "description", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Description"),
				Description: "An optional free-form text field to describe an image."},
			{Name: "distribution", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Distribution"),
				Description: "This attribute describes the base distribution used for this image. For custom images, this is user defined."},
			{Name: "error_message", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.ErrorMessage"),
				Description: "A string containing information about errors that may occur when importing a custom image."},
			{Name: "min_disk_size", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Image.MinDiskSize"),
				Description: "The minimum disk size in GB required for a Droplet to use this image."},
			{Name: "public", Type: proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Image.Public"),
				Description: "This is a boolean value that indicates whether the image in question is public or not. An image that is public is available to all accounts. A non-public image is only accessible from your account."},
			{Name: "regions", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.Regions"),
				Description: "Array of region slugs where the image is available."},
			{Name: "size_gigabytes", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Image.SizeGigaBytes"),
				Description: "The size of the image in gigabytes."},
			{Name: "slug", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Slug"),
				Description: "A uniquely identifying string that is associated with each of the DigitalOcean-provided public images. These can be used to reference a public image as an alternative to the numeric id."},
			{Name: "status", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Status"),
				Description: "A status string indicating the state of a custom image. This may be \"NEW\", \"available\", \"pending\", or \"deleted\"."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Image.Tags"), Description: "An array containing the names of the tags the image has been tagged with."},
			{Name: "type", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Type"),
				Description: "Describes the kind of image. It may be one of \"snapshot\", \"backup\", or \"custom\"."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(imageToURN), Description: "The uniform resource name (URN) for the volume."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(imageToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Image.Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Image.Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func imageToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(opengovernance.DigitalOceanImage).Description.Image
	return fmt.Sprintf("do:image:%d", i.ID), nil
}
