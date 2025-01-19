package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
	opengovernance "github.com/opengovern/og-describer-digitalocean/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDigitalOceanAlertPolicy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_alert_policy",
		Description: "DigitalOcean Alert Policy",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanAlertPolicy,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("uuid"),
			Hydrate:    opengovernance.GetDigitalOceanAlertPolicy,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:        "uuid",
				Type:        proto.ColumnType_STRING,
				Description: "UUID of the alert policy.",
				Transform:   transform.FromField("Description.AlertPolicy.UUID"),
			},
			{
				Name:        "compare",
				Type:        proto.ColumnType_STRING,
				Description: "The compare parameter for the metric in alert policy.",
				Transform:   transform.FromField("Description.AlertPolicy.Compare"),
			},
			{
				Name:        "enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Alert Policy enabled or not.",
				Transform:   transform.FromField("Description.AlertPolicy.Enabled"),
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Alert Policy type.",
				Transform:   transform.FromField("Description.AlertPolicy.Type"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the alert policy.",
				Transform:   transform.FromField("Description.AlertPolicy.Description"),
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_INT,
				Description: "The value of the metric threshold in alert policy.",
				Transform:   transform.FromField("Description.AlertPolicy.Value"),
			},
			{
				Name:        "interval",
				Type:        proto.ColumnType_STRING,
				Description: "The interval time of the metric in alert policy.",
				Transform:   transform.FromField("Description.AlertPolicy.Window"),
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the alert policy.",
				Transform:   transform.FromValue().Transform(alertPolicyToURN),
			},
			{
				Name:        "alerts",
				Type:        proto.ColumnType_JSON,
				Description: "The notification details where alert details will be send.",
				Transform:   transform.FromField("Description.AlertPolicy.Alerts"),
			},
			{
				Name:        "entities",
				Type:        proto.ColumnType_JSON,
				Description: "An array of entities of the alert policy.",
				Transform:   transform.FromField("Description.AlertPolicy.Entities"),
			},
			{
				Name:        "tags_src",
				Type:        proto.ColumnType_JSON,
				Description: "An array of tags for the resource.",
				Transform:   transform.FromField("Description.AlertPolicy.Tags"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: resourceInterfaceDescription("title"),
				Transform:   transform.FromField("Description.AlertPolicy.UUID"),
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("tags"),
				Transform:   transform.FromField("Description.AlertPolicy.Tags").Transform(labelsToTagsMap),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(alertPolicyToURN).Transform(ensureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func alertPolicyToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var policy godo.AlertPolicy
	switch d.Value.(type) {
	case *opengovernance.DigitalOceanAlertPolicy:
		policy = d.Value.(*opengovernance.DigitalOceanAlertPolicy).Description.AlertPolicy
	case opengovernance.DigitalOceanAlertPolicy:
		policy = d.Value.(opengovernance.DigitalOceanAlertPolicy).Description.AlertPolicy
	}
	return "do:alertPolicy:" + policy.UUID, nil
}
