package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanTag(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_tag",
		Description: "A tag is a label that can be applied to a resource (currently Droplets, Images, Volumes, Volume Snapshots, and Database clusters) in order to better organize or facilitate the lookups and actions on it.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanTag,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetDigitalOceanTag,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Tag.Name"),
				Description: "The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per tag."},
			{Name: "resource_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Tag.Resources.Count"), Description: "The number of resources with this tag."},
			{Name: "resources", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Tag.Resources"), Description: "An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a last_tagged_uri attribute set to the last resource tagged with the current tag."},
		}),
	}
}
