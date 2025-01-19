package digitalocean

import (
	"context"
	"fmt"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDigitalOceanApp(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_app",
		Description: "DigitalOcean App",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanApp,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanApp,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The id of the app.",
				Transform:   transform.FromField("Description.App.ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the app",
				Transform:   transform.FromField("Description.App.Spec.Name"),
			},
			{
				Name:        "owner_uuid",
				Type:        proto.ColumnType_STRING,
				Description: "OwnerUUID of the app.",
				Transform:   transform.FromField("Description.App.OwnerUUID"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the app was created.",
				Transform:   transform.FromField("Description.App.CreatedAt"),
			},
			{
				Name:        "last_deployment_created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the app last deployed.",
				Transform:   transform.FromField("Description.App.LastDeploymentCreatedAt"),
			},
			{
				Name:        "live_url",
				Type:        proto.ColumnType_STRING,
				Description: "The live URL of the app.",
				Transform:   transform.FromField("Description.App.LiveURL"),
			},
			{
				Name:        "live_url_base",
				Type:        proto.ColumnType_STRING,
				Description: "The live URL base of the app.",
				Transform:   transform.FromField("Description.App.LiveURLBase"),
			},
			{
				Name:        "live_domain",
				Type:        proto.ColumnType_STRING,
				Description: "The live domain of the app.",
				Transform:   transform.FromField("Description.App.LiveDomain"),
			},
			{
				Name:        "tier_slug",
				Type:        proto.ColumnType_STRING,
				Description: "Tier slug of the app",
				Transform:   transform.FromField("Description.App.TierSlug"),
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the app was updated.",
				Transform:   transform.FromField("Description.App.UpdatedAt"),
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the app.",
				Transform:   transform.FromValue().Transform(appToURN),
			},
			{
				Name:        "active_deployment",
				Type:        proto.ColumnType_JSON,
				Description: "The app's currently active deployment.",
				Transform:   transform.FromField("Description.App.ActiveDeployment"),
			},
			{
				Name:        "in_progress_deployment",
				Type:        proto.ColumnType_JSON,
				Description: "In progress deployment of the app.",
				Transform:   transform.FromField("Description.App.InProgressDeployment"),
			},
			{
				Name:        "region",
				Type:        proto.ColumnType_JSON,
				Description: "The DigitalOcean data center region hosting the app.",
				Transform:   transform.FromField("Description.App.Region"),
			},
			{
				Name:        "spec",
				Type:        proto.ColumnType_JSON,
				Description: "A DigitalOcean App spec describing the app.",
				Transform:   transform.FromField("Description.App.Spec"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: resourceInterfaceDescription("title"),
				Transform:   transform.FromField("Description.App.Spec.Name"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(appToURN).Transform(ensureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func appToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(opengovernance.DigitalOceanApp).Description.App
	return fmt.Sprintf("do:app:%s", i.ID), nil
}
