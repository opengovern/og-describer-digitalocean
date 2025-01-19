package maps
import (
	"github.com/opengovern/og-describer-digitalocean/discovery/describers"
	"github.com/opengovern/og-describer-digitalocean/discovery/provider"
	"github.com/opengovern/og-describer-digitalocean/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
	model "github.com/opengovern/og-describer-digitalocean/discovery/pkg/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"DigitalOcean::Account": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Account",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanAccount),
		GetDescriber:         nil,
	},

	"DigitalOcean::Action": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Action",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanAction),
		GetDescriber:         nil,
	},

	"DigitalOcean::AlertPolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::AlertPolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanAlertPolicy),
		GetDescriber:         nil,
	},

	"DigitalOcean::App": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::App",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanApp),
		GetDescriber:         nil,
	},

	"DigitalOcean::Balance": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Balance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanBalance),
		GetDescriber:         nil,
	},

	"DigitalOcean::Bill": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Bill",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanBill),
		GetDescriber:         nil,
	},

	"DigitalOcean::ContainerRegistry": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::ContainerRegistry",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanContainerRegistry),
		GetDescriber:         nil,
	},

	"DigitalOcean::Database": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Database",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanDatabase),
		GetDescriber:         nil,
	},

	"DigitalOcean::Domain": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Domain",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanDomain),
		GetDescriber:         nil,
	},

	"DigitalOcean::Droplet": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Droplet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanDroplet),
		GetDescriber:         nil,
	},

	"DigitalOcean::Firewall": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Firewall",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanFirewall),
		GetDescriber:         nil,
	},

	"DigitalOcean::FloatingIP": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::FloatingIP",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanFloatingIP),
		GetDescriber:         nil,
	},

	"DigitalOcean::Image": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Image",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanImage),
		GetDescriber:         nil,
	},

	"DigitalOcean::Key": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Key",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanKey),
		GetDescriber:         nil,
	},

	"DigitalOcean::KubernetesCluster": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::KubernetesCluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanKubernetesCluster),
		GetDescriber:         nil,
	},

	"DigitalOcean::KubernetesNodePool": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::KubernetesNodePool",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanKubernetesNodePool),
		GetDescriber:         nil,
	},

	"DigitalOcean::LoadBalancer": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::LoadBalancer",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanLoadBalancer),
		GetDescriber:         nil,
	},

	"DigitalOcean::Project": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Project",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanProject),
		GetDescriber:         nil,
	},

	"DigitalOcean::Region": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Region",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanRegion),
		GetDescriber:         nil,
	},

	"DigitalOcean::Size": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Size",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanSize),
		GetDescriber:         nil,
	},

	"DigitalOcean::Snapshot": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Snapshot",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanSnapshot),
		GetDescriber:         nil,
	},

	"DigitalOcean::Tag": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Tag",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanTag),
		GetDescriber:         nil,
	},

	"DigitalOcean::Volume": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::Volume",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanVolume),
		GetDescriber:         nil,
	},

	"DigitalOcean::VPC": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "DigitalOcean::VPC",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeByIntegration(describers.DigitalOceanVPC),
		GetDescriber:         nil,
	},
}


var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"DigitalOcean::Account": {
		Name:         "DigitalOcean::Account",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Action": {
		Name:         "DigitalOcean::Action",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::AlertPolicy": {
		Name:         "DigitalOcean::AlertPolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::App": {
		Name:         "DigitalOcean::App",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Balance": {
		Name:         "DigitalOcean::Balance",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Bill": {
		Name:         "DigitalOcean::Bill",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::ContainerRegistry": {
		Name:         "DigitalOcean::ContainerRegistry",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Database": {
		Name:         "DigitalOcean::Database",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Domain": {
		Name:         "DigitalOcean::Domain",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Droplet": {
		Name:         "DigitalOcean::Droplet",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Firewall": {
		Name:         "DigitalOcean::Firewall",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::FloatingIP": {
		Name:         "DigitalOcean::FloatingIP",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Image": {
		Name:         "DigitalOcean::Image",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Key": {
		Name:         "DigitalOcean::Key",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::KubernetesCluster": {
		Name:         "DigitalOcean::KubernetesCluster",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::KubernetesNodePool": {
		Name:         "DigitalOcean::KubernetesNodePool",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::LoadBalancer": {
		Name:         "DigitalOcean::LoadBalancer",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Project": {
		Name:         "DigitalOcean::Project",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Region": {
		Name:         "DigitalOcean::Region",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Size": {
		Name:         "DigitalOcean::Size",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Snapshot": {
		Name:         "DigitalOcean::Snapshot",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Tag": {
		Name:         "DigitalOcean::Tag",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::Volume": {
		Name:         "DigitalOcean::Volume",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"DigitalOcean::VPC": {
		Name:         "DigitalOcean::VPC",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},
}


var ResourceTypesList = []string{
  "DigitalOcean::Account",
  "DigitalOcean::Action",
  "DigitalOcean::AlertPolicy",
  "DigitalOcean::App",
  "DigitalOcean::Balance",
  "DigitalOcean::Bill",
  "DigitalOcean::ContainerRegistry",
  "DigitalOcean::Database",
  "DigitalOcean::Domain",
  "DigitalOcean::Droplet",
  "DigitalOcean::Firewall",
  "DigitalOcean::FloatingIP",
  "DigitalOcean::Image",
  "DigitalOcean::Key",
  "DigitalOcean::KubernetesCluster",
  "DigitalOcean::KubernetesNodePool",
  "DigitalOcean::LoadBalancer",
  "DigitalOcean::Project",
  "DigitalOcean::Region",
  "DigitalOcean::Size",
  "DigitalOcean::Snapshot",
  "DigitalOcean::Tag",
  "DigitalOcean::Volume",
  "DigitalOcean::VPC",
}