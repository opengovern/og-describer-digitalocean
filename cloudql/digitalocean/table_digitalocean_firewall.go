package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
	opengovernance "github.com/opengovern/og-describer-digitalocean/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanFirewall(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_firewall",
		Description: "DigitalOcean Cloud Firewalls are a network-based, stateful firewall service for Droplets provided at no additional cost. Cloud firewalls block all traffic that isn’t expressly permitted by a rule.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanFirewall,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanFirewall,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Firewall.ID"),
				Description: "The unique universal identifier of this firewall.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Firewall.Name"),
				Description: "The name of the Firewall.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Firewall.Created"),
				Description: "A time value given in ISO8601 combined date and time format that represents when the Firewall was created.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Firewall.Status"),
				Description: "A status string indicating the current state of the Firewall.",
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the Firewall.",
				Transform:   transform.FromValue().Transform(firewallToURN),
			},
			{
				Name:        "droplet_ids",
				Type:        proto.ColumnType_JSON,
				Description: "The list of the IDs of the Droplets assigned to the Firewall.",
				Transform:   transform.FromField("Description.Firewall.DropletIDs"),
			},
			{
				Name:        "inbound_rules",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Firewall.InboundRules"),
				Description: "The inbound access rule block for the Firewall.",
			},
			{
				Name:        "outbound_rules",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Firewall.OutboundRules"),
				Description: "The outbound access rule block for the Firewall.",
			},
			{
				Name:        "pending_changes",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Firewall.PendingChanges"),
				Description: "An list of object containing the fields, `droplet_id`, `removing`, and `status`. It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.",
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("tags"),
				Transform:   transform.FromField("Description.Firewall.Tags").Transform(labelsToTagsMap),
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: resourceInterfaceDescription("title"),
				Transform:   transform.FromField("Description.Firewall.Name"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(firewallToURN).Transform(ensureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func firewallToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var firewall godo.Firewall
	switch d.Value.(type) {
	case *opengovernance.DigitalOceanFirewall:
		firewall = d.Value.(*opengovernance.DigitalOceanFirewall).Description.Firewall
	case opengovernance.DigitalOceanFirewall:
		firewall = d.Value.(opengovernance.DigitalOceanFirewall).Description.Firewall
	}
	return "do:firewall:" + firewall.ID, nil
}
