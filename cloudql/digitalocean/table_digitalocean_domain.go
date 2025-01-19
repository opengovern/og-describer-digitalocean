package digitalocean

import (
	"context"
	"github.com/digitalocean/godo"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDigitalOceanDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_domain",
		Description: "DigitalOcean Domain",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanDomain,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetDigitalOceanDomain,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Name"),
				Description: "The globally unique human-readable name for the domain.",
			},
			{
				Name:        "ttl",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Domain.TTL"),
				Description: "TTL value of domain.",
			},
			{
				Name:        "zone_file",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.ZoneFile"),
				Description: "It contains the DNS record details.",
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the domain.",
				Transform:   transform.FromValue().Transform(domainToURN),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: resourceInterfaceDescription("title"),
				Transform:   transform.FromField("Description.Domain.Name"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(domainToURN).Transform(ensureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func domainToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var i godo.Domain
	switch d.Value.(type) {
	case *opengovernance.DigitalOceanDomain:
		i = d.Value.(*opengovernance.DigitalOceanDomain).Description.Domain
	case opengovernance.DigitalOceanDomain:
		i = d.Value.(opengovernance.DigitalOceanDomain).Description.Domain
	}
	return "do:domain:" + i.Name, nil
}
