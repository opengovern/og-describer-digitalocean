//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json -pluginPath ../../steampipe-plugin-digitalocean/digitalocean

// Implement types for each resource

package provider

import "github.com/digitalocean/godo"

type Metadata struct {
}

type DigitalOceanAccountDescription struct {
	Account godo.Account `json:"account"`
}

type DigitalOceanActionDescription struct {
	Action godo.Action `json:"action"`
}

type DigitalOceanAlertPolicyDescription struct {
	AlertPolicy godo.AlertPolicy `json:"alert_policy"`
}

type DigitalOceanAppDescription struct {
	App godo.App `json:"app"`
}

type DigitalOceanBalanceDescription struct {
	Balance godo.Balance `json:"balance"`
}

type DigitalOceanBillDescription struct {
	Bill godo.BillingHistoryEntry `json:"bill"`
}

type DigitalOceanContainerRegistryDescription struct {
	ContainerRegistry godo.Registry `json:"container_registry"`
}

type DigitalOceanDatabaseDescription struct {
	Database              godo.Database               `json:"database"`
	DatabaseUsers         []godo.DatabaseUser         `json:"database_users"`
	DatabaseNames         []godo.DatabaseDB           `json:"database_names"`
	DatabaseFirewallRules []godo.DatabaseFirewallRule `json:"database_firewall_rules"`
}

type DigitalOceanDomainDescription struct {
	Domain godo.Domain `json:"domain"`
}

type DigitalOceanDropletDescription struct {
	URN     string       `json:"urn"`
	Droplet godo.Droplet `json:"droplet"`
}

type DigitalOceanFirewallDescription struct {
	Firewall godo.Firewall `json:"firewall"`
}

type DigitalOceanFloatingIPDescription struct {
	URN        string          `json:"urn"`
	FloatingIP godo.FloatingIP `json:"floating_ip"`
}

type DigitalOceanImageDescription struct {
	Image godo.Image `json:"image"`
}

type DigitalOceanKeyDescription struct {
	Key godo.Key `json:"key"`
}

type DigitalOceanKubernetesClusterDescription struct {
	Cluster godo.KubernetesCluster `json:"cluster"`
}

type DigitalOceanKubernetesNodePoolDescription struct {
	ClusterID string                  `json:"cluster_id"`
	NodePool  godo.KubernetesNodePool `json:"node_pool"`
}

type DigitalOceanLoadBalancerDescription struct {
	LoadBalancer godo.LoadBalancer `json:"load_balancer"`
}

type DigitalOceanProjectDescription struct {
	Project godo.Project `json:"project"`
}

type DigitalOceanRegionDescription struct {
	Region godo.Region `json:"region"`
}

type DigitalOceanSizeDescription struct {
	Size godo.Size `json:"size"`
}

type DigitalOceanSnapshotDescription struct {
	Snapshot godo.Snapshot `json:"snapshot"`
}

type DigitalOceanTagDescription struct {
	Tag godo.Tag `json:"tag"`
}

type DigitalOceanVolumeDescription struct {
	URN    string      `json:"urn"`
	Volume godo.Volume `json:"volume"`
}

type DigitalOceanVPCDescription struct {
	VPC godo.VPC `json:"vpc"`
}
