package digitalocean

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	opengovernance "github.com/opengovern/og-describer-digitalocean/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_key",
		Description: "DigitalOcean allows you to add SSH public keys to the interface so that you can embed your public key into a Droplet at the time of creation. Only the public key is required to take advantage of this functionality.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "fingerprint"}),
			Hydrate:    opengovernance.GetDigitalOceanKey,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Key.ID"),
				Description: "This is a unique identification number for the key."},
			{Name: "name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.Name"),
				Description: "The human-readable display name for the given SSH key."},
			// Other columns
			{Name: "fingerprint", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.Fingerprint"),
				Description: "The fingerprint value that is generated from the public key."},
			{Name: "public_key", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Key.PublicKey"),
				Description: "The entire public key string that was uploaded.  This is what is embedded into the root user's authorized_keys file if you choose to include this SSH key during Droplet creation."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(keyToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Key.Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func keyToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var i godo.Key
	switch d.Value.(type) {
	case *opengovernance.DigitalOceanKey:
		i = d.Value.(*opengovernance.DigitalOceanKey).Description.Key
	case opengovernance.DigitalOceanKey:
		i = d.Value.(opengovernance.DigitalOceanKey).Description.Key
	}
	return fmt.Sprintf("do:key:%d", i.ID), nil
}
