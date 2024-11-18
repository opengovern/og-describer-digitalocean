package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDigitalOceanBill(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_bill",
		Description: "Billing history is a record of billing events for your account. For example, entries may include events like payments made, invoices issued, or credits granted.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanBill,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{
				Name:      "date",
				Transform: transform.FromField("Description.Bill.Date"),
				Type:      proto.ColumnType_TIMESTAMP, Description: "Time the billing history entry occured."},
			// Other columns
			{
				Name:      "amount",
				Transform: transform.FromField("Description.Bill.Amount"),
				Type:      proto.ColumnType_STRING, Description: "Amount of the billing history entry."},
			{
				Name:      "description",
				Transform: transform.FromField("Description.Bill.Description"),
				Type:      proto.ColumnType_STRING, Description: "Description of the billing history entry."},
			{
				Name:      "invoice_id",
				Transform: transform.FromField("Description.Bill.InvoiceID"),
				Type:      proto.ColumnType_STRING, Description: "ID of the invoice associated with the billing history entry, if applicable."},
			{
				Name:      "invoice_uuid",
				Transform: transform.FromField("Description.Bill.InvoiceUUID"),
				Type:      proto.ColumnType_STRING, Description: "UUID of the invoice associated with the billing history entry, if applicable."},
			{
				Name:      "type",
				Transform: transform.FromField("Description.Bill.Type"),
				Type:      proto.ColumnType_STRING, Description: "Type of billing history entry."},
		}),
	}
}
