package maps

import (
	"github.com/opengovern/og-describer-digitalocean/discovery/pkg/es"
)

var ResourceTypesToTables = map[string]string{
  "DigitalOcean::Account": "digitalocean_account",
  "DigitalOcean::Action": "digitalocean_action",
  "DigitalOcean::AlertPolicy": "digitalocean_alert_policy",
  "DigitalOcean::App": "digitalocean_app",
  "DigitalOcean::Balance": "digitalocean_balance",
  "DigitalOcean::Bill": "digitalocean_bill",
  "DigitalOcean::ContainerRegistry": "digitalocean_container_registry",
  "DigitalOcean::Database": "digitalocean_database",
  "DigitalOcean::Domain": "digitalocean_domain",
  "DigitalOcean::Droplet": "digitalocean_droplet",
  "DigitalOcean::Firewall": "digitalocean_firewall",
  "DigitalOcean::FloatingIP": "digitalocean_floating_ip",
  "DigitalOcean::Image": "digitalocean_image",
  "DigitalOcean::Key": "digitalocean_key",
  "DigitalOcean::KubernetesCluster": "digitalocean_kubernetes_cluster",
  "DigitalOcean::KubernetesNodePool": "digitalocean_kubernetes_node_pool",
  "DigitalOcean::LoadBalancer": "digitalocean_load_balancer",
  "DigitalOcean::Project": "digitalocean_project",
  "DigitalOcean::Region": "digitalocean_region",
  "DigitalOcean::Size": "digitalocean_size",
  "DigitalOcean::Snapshot": "digitalocean_snapshot",
  "DigitalOcean::Tag": "digitalocean_tag",
  "DigitalOcean::Volume": "digitalocean_volume",
  "DigitalOcean::VPC": "digitalocean_vpc",
}

var ResourceTypeToDescription = map[string]interface{}{
  "DigitalOcean::Account": opengovernance.DigitalOceanAccount{},
  "DigitalOcean::Action": opengovernance.DigitalOceanAction{},
  "DigitalOcean::AlertPolicy": opengovernance.DigitalOceanAlertPolicy{},
  "DigitalOcean::App": opengovernance.DigitalOceanApp{},
  "DigitalOcean::Balance": opengovernance.DigitalOceanBalance{},
  "DigitalOcean::Bill": opengovernance.DigitalOceanBill{},
  "DigitalOcean::ContainerRegistry": opengovernance.DigitalOceanContainerRegistry{},
  "DigitalOcean::Database": opengovernance.DigitalOceanDatabase{},
  "DigitalOcean::Domain": opengovernance.DigitalOceanDomain{},
  "DigitalOcean::Droplet": opengovernance.DigitalOceanDroplet{},
  "DigitalOcean::Firewall": opengovernance.DigitalOceanFirewall{},
  "DigitalOcean::FloatingIP": opengovernance.DigitalOceanFloatingIP{},
  "DigitalOcean::Image": opengovernance.DigitalOceanImage{},
  "DigitalOcean::Key": opengovernance.DigitalOceanKey{},
  "DigitalOcean::KubernetesCluster": opengovernance.DigitalOceanKubernetesCluster{},
  "DigitalOcean::KubernetesNodePool": opengovernance.DigitalOceanKubernetesNodePool{},
  "DigitalOcean::LoadBalancer": opengovernance.DigitalOceanLoadBalancer{},
  "DigitalOcean::Project": opengovernance.DigitalOceanProject{},
  "DigitalOcean::Region": opengovernance.DigitalOceanRegion{},
  "DigitalOcean::Size": opengovernance.DigitalOceanSize{},
  "DigitalOcean::Snapshot": opengovernance.DigitalOceanSnapshot{},
  "DigitalOcean::Tag": opengovernance.DigitalOceanTag{},
  "DigitalOcean::Volume": opengovernance.DigitalOceanVolume{},
  "DigitalOcean::VPC": opengovernance.DigitalOceanVPC{},
}

var TablesToResourceTypes = map[string]string{
  "digitalocean_account": "DigitalOcean::Account",
  "digitalocean_action": "DigitalOcean::Action",
  "digitalocean_alert_policy": "DigitalOcean::AlertPolicy",
  "digitalocean_app": "DigitalOcean::App",
  "digitalocean_balance": "DigitalOcean::Balance",
  "digitalocean_bill": "DigitalOcean::Bill",
  "digitalocean_container_registry": "DigitalOcean::ContainerRegistry",
  "digitalocean_database": "DigitalOcean::Database",
  "digitalocean_domain": "DigitalOcean::Domain",
  "digitalocean_droplet": "DigitalOcean::Droplet",
  "digitalocean_firewall": "DigitalOcean::Firewall",
  "digitalocean_floating_ip": "DigitalOcean::FloatingIP",
  "digitalocean_image": "DigitalOcean::Image",
  "digitalocean_key": "DigitalOcean::Key",
  "digitalocean_kubernetes_cluster": "DigitalOcean::KubernetesCluster",
  "digitalocean_kubernetes_node_pool": "DigitalOcean::KubernetesNodePool",
  "digitalocean_load_balancer": "DigitalOcean::LoadBalancer",
  "digitalocean_project": "DigitalOcean::Project",
  "digitalocean_region": "DigitalOcean::Region",
  "digitalocean_size": "DigitalOcean::Size",
  "digitalocean_snapshot": "DigitalOcean::Snapshot",
  "digitalocean_tag": "DigitalOcean::Tag",
  "digitalocean_volume": "DigitalOcean::Volume",
  "digitalocean_vpc": "DigitalOcean::VPC",
}
