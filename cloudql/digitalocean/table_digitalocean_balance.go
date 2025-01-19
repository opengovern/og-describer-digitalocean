package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDigitalOceanBalance(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_balance",
		Description: "Balance information for the current account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanBalance,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:      "account_balance",
				Transform: transform.FromField("Description.Balance.AccountBalance"),
				Type:      proto.ColumnType_DOUBLE, Description: "Current balance of the customer's most recent billing activity. Does not reflect month_to_date_usage."},
			{
				Name:      "generated_at",
				Transform: transform.FromField("Description.Balance.GeneratedAt"),
				Type:      proto.ColumnType_TIMESTAMP, Description: "The time at which balances were most recently generated."},
			{
				Name:      "month_to_date_balance",
				Transform: transform.FromField("Description.Balance.MonthToDateBalance"),
				Type:      proto.ColumnType_DOUBLE, Description: "Balance as of the generated_at time. This value includes the account_balance and month_to_date_usage."},
			{
				Name:      "month_to_date_usage",
				Transform: transform.FromField("Description.Balance.MonthToDateUsage"),
				Type:      proto.ColumnType_DOUBLE, Description: "Amount used in the current billing period as of the generated_at time."},
		}),
	}
}
