package digitalocean

import (
	"context"

	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-digitalocean",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: essdk.ConfigInstance,
			Schema:      essdk.ConfigSchema(),
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"digitalocean_account":              tableDigitalOceanAccount(ctx),
			"digitalocean_action":               tableDigitalOceanAction(ctx),
			"digitalocean_alert_policy":         tableDigitalOceanAlertPolicy(ctx),
			"digitalocean_app":                  tableDigitalOceanApp(ctx),
			"digitalocean_balance":              tableDigitalOceanBalance(ctx),
			"digitalocean_bill":                 tableDigitalOceanBill(ctx),
			"digitalocean_container_registry":   tableDigitalOceanContainerRegistry(ctx),
			"digitalocean_database":             tableDigitalOceanDatabase(ctx),
			"digitalocean_domain":               tableDigitalOceanDomain(ctx),
			"digitalocean_droplet":              tableDigitalOceanDroplet(ctx),
			"digitalocean_firewall":             tableDigitalOceanFirewall(ctx),
			"digitalocean_floating_ip":          tableDigitalOceanFloatingIP(ctx),
			"digitalocean_image":                tableDigitalOceanImage(ctx),
			"digitalocean_key":                  tableDigitalOceanKey(ctx),
			"digitalocean_kubernetes_cluster":   tableDigitalOceanKubernetesCluster(ctx),
			"digitalocean_kubernetes_node_pool": tableDigitalOceanKubernetesNodePool(ctx),
			"digitalocean_load_balancer":        tableDigitalOceanLoadBalancer(ctx),
			"digitalocean_project":              tableDigitalOceanProject(ctx),
			"digitalocean_region":               tableDigitalOceanRegion(ctx),
			"digitalocean_size":                 tableDigitalOceanSize(ctx),
			"digitalocean_snapshot":             tableDigitalOceanSnapshot(ctx),
			"digitalocean_tag":                  tableDigitalOceanTag(ctx),
			"digitalocean_volume":               tableDigitalOceanVolume(ctx),
			"digitalocean_vpc":                  tableDigitalOceanVPC(ctx),
		},
	}

	for key, table := range p.TableMap {
		if table == nil {
			continue
		}
		if table.Get != nil && table.Get.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}
		if table.List != nil && table.List.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}

		opengovernanceTable := false
		for _, col := range table.Columns {
			if col != nil && col.Name == "og_account_id" {
				opengovernanceTable = true
			}
		}

		if opengovernanceTable {
			if table.Get != nil {
				table.Get.KeyColumns = append(table.Get.KeyColumns, plugin.OptionalColumns([]string{"og_account_id", "og_resource_id"})...)
			}

			if table.List != nil {
				table.List.KeyColumns = append(table.List.KeyColumns, plugin.OptionalColumns([]string{"og_account_id", "og_resource_id"})...)
			}
		}
	}

	return p
}
