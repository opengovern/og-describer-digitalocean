package provider
import (
	"github.com/opengovern/og-describer-digitalocean/provider/describer"
	"github.com/opengovern/og-describer-digitalocean/provider/configs"
	model "github.com/opengovern/og-describer-digitalocean/pkg/sdk/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"DigitalOcean::Account": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Account",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanAccount),
		GetDescriber:         nil,
	},

	"DigitalOcean::Action": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Action",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanAction),
		GetDescriber:         nil,
	},

	"DigitalOcean::AlertPolicy": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::AlertPolicy",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanAlertPolicy),
		GetDescriber:         nil,
	},

	"DigitalOcean::App": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::App",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanApp),
		GetDescriber:         nil,
	},

	"DigitalOcean::Balance": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Balance",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanBalance),
		GetDescriber:         nil,
	},

	"DigitalOcean::Bill": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Bill",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanBill),
		GetDescriber:         nil,
	},

	"DigitalOcean::ContainerRegistry": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::ContainerRegistry",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanContainerRegistry),
		GetDescriber:         nil,
	},

	"DigitalOcean::Database": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Database",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanDatabase),
		GetDescriber:         nil,
	},

	"DigitalOcean::Domain": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Domain",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanDomain),
		GetDescriber:         nil,
	},

	"DigitalOcean::Droplet": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Droplet",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanDroplet),
		GetDescriber:         nil,
	},

	"DigitalOcean::Firewall": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Firewall",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanFirewall),
		GetDescriber:         nil,
	},

	"DigitalOcean::FloatingIP": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::FloatingIP",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanFloatingIP),
		GetDescriber:         nil,
	},

	"DigitalOcean::Image": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Image",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanImage),
		GetDescriber:         nil,
	},

	"DigitalOcean::Key": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Key",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanKey),
		GetDescriber:         nil,
	},

	"DigitalOcean::KubernetesCluster": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::KubernetesCluster",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanKubernetesCluster),
		GetDescriber:         nil,
	},

	"DigitalOcean::KubernetesNodePool": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::KubernetesNodePool",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanKubernetesNodePool),
		GetDescriber:         nil,
	},

	"DigitalOcean::LoadBalancer": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::LoadBalancer",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanLoadBalancer),
		GetDescriber:         nil,
	},

	"DigitalOcean::Project": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Project",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanProject),
		GetDescriber:         nil,
	},

	"DigitalOcean::Region": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Region",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanRegion),
		GetDescriber:         nil,
	},

	"DigitalOcean::Size": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Size",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanSize),
		GetDescriber:         nil,
	},

	"DigitalOcean::Snapshot": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Snapshot",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanSnapshot),
		GetDescriber:         nil,
	},

	"DigitalOcean::Tag": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Tag",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanTag),
		GetDescriber:         nil,
	},

	"DigitalOcean::Volume": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::Volume",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanVolume),
		GetDescriber:         nil,
	},

	"DigitalOcean::VPC": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "DigitalOcean::VPC",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeByIntegration(describer.DigitalOceanVPC),
		GetDescriber:         nil,
	},
}
