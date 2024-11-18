package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_account",
		Description: "Retrieve information about your current account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanAccount,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Account.Email"), Description: "The email address used by the current user to register for DigitalOcean."},
			// Other columns
			{Name: "droplet_limit", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Account.DropletLimit"), Description: "The total number of Droplets the current user or team may have at one time."},
			{Name: "email_verified", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Account.EmailVerified"), Description: "If true, the user has verified their account via email. False otherwise."},
			{Name: "floating_ip_limit", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Account.FloatingIPLimit"), Description: "The total number of floating IPs the current user or team may have."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Account.Status"), Description: "This value is one of \"active\", \"warning\" or \"locked\"."},
			{Name: "status_message", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Account.StatusMessage"), Description: "A human-readable message giving more details about the status of the account."},
			{Name: "uuid", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Account.UUID"), Description: "The unique universal identifier for the current user."},
			{Name: "volume_limit", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Account.VolumeLimit"), Description: "The total number of volumes the current user or team may have."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(accountAkas), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Account.Email"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func accountAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	a := d.Value.(opengovernance.DigitalOceanAccount).Description.Account
	return []string{"do:account:" + a.UUID}, nil
}
