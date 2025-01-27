package digitalocean

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-digitalocean/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanAction(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_action",
		Description: "Actions are records of events that have occurred on the resources in your account. These can be things like rebooting a Action, or transferring an image to a new region.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanAction,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanAction,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Account.ID"), Description: "A unique numeric ID that can be used to identify and reference an action."},
			{Name: "resource_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Account.ResourceID"), Description: "A unique identifier for the resource that the action is associated with."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Account.Type"), Description: "This is the type of action that the object represents.  For example, this could be \"transfer\" to represent the state of an image transfer action."},
			{Name: "started_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.Account.StartedAt").Transform(timestampToIsoTimestamp), Description: "Time when when the action was initiated."},
			// Other columns
			{Name: "completed_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.Account.CompletedAt").Transform(timestampToIsoTimestamp), Description: "Time when the action was completed."},
			// Skip this object dump for now, they can join on the slug if really necessary
			//{Name: "region", Type: proto.ColumnType_JSON, Transform: transform.FromField("Region"), Description: "A full region object containing information about the region where the action occurred."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Account.RegionSlug"), Description: "The region where the action occurred."},
			{Name: "resource_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Account.ResourceType"), Description: "The type of resource that the action is associated with."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Account.Status"), Description: "The current status of the action.  This can be \"in-progress\", \"completed\", or \"errored\"."},
		}),
	}
}
