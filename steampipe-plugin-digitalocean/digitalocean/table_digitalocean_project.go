package digitalocean

import (
	"context"
	"github.com/digitalocean/godo"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_project",
		Description: "Projects allow you to organize your resources into groups that fit the way you work. You can group resources (like Droplets, Spaces, load balancers, domains, and floating IPs) in ways that align with the applications you host on DigitalOcean.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanProject,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanProject,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.ID"),
				Description: "The unique universal identifier of this project."},
			{Name: "name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Name"),
				Description: "The globally unique human-readable name for the project."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Project.CreatedAt"),
				Description: "Time when the project was created."},
			{Name: "description", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Description"),
				Description: "The description of the project."},
			{Name: "environment", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Environment"),
				Description: "The environment of the project's resources."},
			{Name: "is_default", Type: proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Project.IsDefault"),
				Description: "If true, all resources will be added to this project if no project is specified."},
			{Name: "owner_id", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Project.OwnerID"),
				Description: "The integer id of the project owner."},
			{Name: "owner_uuid", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.OwnerUUID"),
				Description: "The unique universal identifier of the project owner."},
			{Name: "purpose", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Project.Purpose"),
				Description: "The purpose of the project."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Project.UpdatedAt"),
				Description: "Time when the project was updated."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(projectToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Project.Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func projectToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var i godo.Project
	switch d.Value.(type) {
	case *opengovernance.DigitalOceanProject:
		i = d.Value.(*opengovernance.DigitalOceanProject).Description.Project
	case opengovernance.DigitalOceanProject:
		i = d.Value.(opengovernance.DigitalOceanProject).Description.Project
	}
	return "do:project:" + i.ID, nil
}
