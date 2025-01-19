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

func tableDigitalOceanKubernetesNodePool(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_kubernetes_node_pool",
		Description: "DigitalOcean Kubernetes Node Pool",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanKubernetesNodePool,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "cluster_id"}),
			Hydrate:    opengovernance.GetDigitalOceanKubernetesNodePool,
		},
		Columns: integrationColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique universal identifier of this node pool.",
				Transform:   transform.FromField("Description.NodePool.ID"),
			},
			{
				Name:        "cluster_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique universal identifier of the associated cluster.",
				Transform:   transform.FromField("Description.ClusterID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The globally unique human-readable name for the node pool.",
				Transform:   transform.FromField("Description.NodePool.Name"),
			},
			{
				Name:        "auto_scale",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean value indicating whether the node pool has autoscaling enabled.",
				Transform:   transform.FromField("Description.NodePool.AutoScale"),
			},
			{
				Name:        "count",
				Type:        proto.ColumnType_INT,
				Description: "The number of nodes in the node pool.",
				Transform:   transform.FromField("Description.NodePool.Count"),
			},
			{
				Name:        "max_nodes",
				Type:        proto.ColumnType_INT,
				Description: "The maximum number of nodes allowed in the node pool.",
				Transform:   transform.FromField("Description.NodePool.MaxNodes"),
			},
			{
				Name:        "min_nodes",
				Type:        proto.ColumnType_INT,
				Description: "The minimum number of nodes allowed in the node pool.",
				Transform:   transform.FromField("Description.NodePool.MinNodes"),
			},
			{
				Name:        "size",
				Type:        proto.ColumnType_STRING,
				Description: "The size of the node pool.",
				Transform:   transform.FromField("Description.NodePool.Size"),
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the node pool.",
				Transform:   transform.FromValue().Transform(nodePoolToURN),
			},

			{
				Name:        "labels",
				Type:        proto.ColumnType_JSON,
				Description: "The labels for the node pool.",
				Transform:   transform.FromField("Description.NodePool.Labels"),
			},
			{
				Name:        "nodes",
				Type:        proto.ColumnType_JSON,
				Description: "The nodes available in the node pool.",
				Transform:   transform.FromField("Description.NodePool.Nodes"),
			},
			{
				Name:        "taints",
				Type:        proto.ColumnType_JSON,
				Description: "The taints of the node pool.",
				Transform:   transform.FromField("Description.NodePool.Taints"),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("tags"),
				Transform:   transform.FromField("Description.NodePool.Tags").Transform(labelsToTagsMap),
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: resourceInterfaceDescription("title"),
				Transform:   transform.FromField("Description.NodePool.Name"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(nodePoolToURN).Transform(ensureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func nodePoolToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(opengovernance.DigitalOceanKubernetesNodePool).Description.NodePool
	return fmt.Sprintf("do:kubernetesNodePool:%s", i.ID), nil
}
