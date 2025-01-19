package digitalocean

import (
	"context"
	"fmt"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDigitalOceanKubernetesCluster(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_kubernetes_cluster",
		Description: "DigitalOcean Kubernetes (DOKS) is a managed Kubernetes service that lets you deploy Kubernetes clusters without the complexities of handling the control plane and containerized infrastructure. Clusters are compatible with standard Kubernetes toolchains and integrate natively with DigitalOcean Load Balancers and block storage volumes.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanKubernetesCluster,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanKubernetesCluster,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ID"),
				Description: "The unique universal identifier of this cluster.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.Name"),
				Description: "The globally unique human-readable name for the cluster.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.",
				Transform:   transform.FromField("Description.Cluster.Status.State"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time when the Kubernetes cluster was created.",
				Transform:   transform.FromField("Description.Cluster.CreatedAt"),
			},
			{
				Name:        "auto_upgrade",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.",
				Transform:   transform.FromField("Description.Cluster.AutoUpgrade"),
			},
			{
				Name:        "cluster_subnet",
				Type:        proto.ColumnType_STRING,
				Description: "The range of IP addresses in the overlay network of the Kubernetes cluster.",
				Transform:   transform.FromField("Description.Cluster.ClusterSubnet"),
			},
			{
				Name:        "endpoint",
				Type:        proto.ColumnType_STRING,
				Description: "The base URL of the API server on the Kubernetes master node.",
				Transform:   transform.FromField("Description.Cluster.Endpoint"),
			},
			{
				Name:        "ipv4",
				Type:        proto.ColumnType_STRING,
				Description: "The public IPv4 address of the Kubernetes master node.",
				Transform:   transform.FromField("Description.Cluster.IPv4"),
			},
			{
				Name:        "region_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug identifier for the region where the Kubernetes cluster will be created.",
				Transform:   transform.FromField("Description.Cluster.RegionSlug"),
			},
			{
				Name:        "registry_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean value indicating whether cluster integrated with container registry.",
				Transform:   transform.FromField("Description.Cluster.RegistryEnabled"),
			},
			{
				Name:        "service_subnet",
				Type:        proto.ColumnType_STRING,
				Description: "The range of assignable IP addresses for services running in the Kubernetes cluster.",
				Transform:   transform.FromField("Description.Cluster.ServiceSubnet"),
			},
			{
				Name:        "surge_upgrade",
				Type:        proto.ColumnType_BOOL,
				Description: "Enable/disable surge upgrades for a cluster.",
				Transform:   transform.FromField("Description.Cluster.SurgeUpgrade"),
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time when the Kubernetes cluster was last updated.",
				Transform:   transform.FromField("Description.Cluster.UpdatedAt"),
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the cluster.",
				Transform:   transform.FromValue().Transform(clusterToURN),
			},
			{
				Name:        "version_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug identifier for the version of Kubernetes used for the cluster.",
				Transform:   transform.FromField("Description.Cluster.VersionSlug"),
			},
			{
				Name:        "vpc_uuid",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the VPC where the Kubernetes cluster will be located.",
				Transform:   transform.FromField("Description.Cluster.VPCUUID"),
			},
			{
				Name:        "maintenance_policy",
				Type:        proto.ColumnType_JSON,
				Description: "A block representing the cluster's maintenance window.",
				Transform:   transform.FromField("Description.Cluster.MaintenancePolicy"),
			},
			{
				Name:        "node_pools",
				Type:        proto.ColumnType_JSON,
				Description: "The cluster's default node pool.",
				Transform:   transform.FromField("Description.Cluster.NodePools"),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("tags"),
				Transform:   transform.FromField("Description.Cluster.Tags").Transform(labelsToTagsMap),
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: resourceInterfaceDescription("title"),
				Transform:   transform.FromField("Description.Cluster.Name"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(clusterToURN).Transform(ensureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func clusterToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(opengovernance.DigitalOceanKubernetesCluster).Description.Cluster
	return fmt.Sprintf("do:kubernetesCluster:%s", i.ID), nil
}
