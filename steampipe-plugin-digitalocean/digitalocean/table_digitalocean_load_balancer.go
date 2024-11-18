package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanLoadBalancer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_load_balancer",
		Description: "DigitalOcean Load Balancers provide a way to distribute traffic across multiple Droplets.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanLoadBalancer,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanLoadBalancer,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.ID"),
				Description: "A unique ID that can be used to identify and reference a load balancer."},
			{Name: "name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Name"),
				Description: "A human-readable name for a load balancer instance."},
			// Other columns
			{Name: "algorithm", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Algorithm"),
				Description: "The load balancing algorithm used to determine which backend Droplet will be selected by a client. It must be either \"round_robin\" or \"least_connections\"."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LoadBalancer.Created"),
				Description: "Time when the load balancer was created."},
			{Name: "droplet_ids", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.DropletIDs"),
				Description: "An array containing the IDs of the Droplets assigned to the load balancer."},
			{Name: "enable_backend_keepalive", Type: proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.LoadBalancer.EnableBackendKeepalive"),
				Description: "A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets."},
			{Name: "enable_proxy_protocol", Type: proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.LoadBalancer.EnableProxyProtocol"),
				Description: "A boolean value indicating whether PROXY Protocol is in use."},
			{Name: "forwarding_rules", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.ForwardingRules"),
				Description: "An object specifying the forwarding rules for a load balancer."},
			{Name: "health_check_healthy_threshold", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.LoadBalancer.HealthCheck.HealthyThreshold"), Description: "The number of times a health check must pass for a backend Droplet to be marked \"healthy\" and be re-added to the pool."},
			{Name: "health_check_interval_seconds", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.LoadBalancer.HealthCheck.CheckIntervalSeconds"), Description: "The number of seconds between between two consecutive health checks."},
			{Name: "health_check_path", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LoadBalancer.HealthCheck.Path"), Description: "The path on the backend Droplets to which the load balancer instance will send a request."},
			{Name: "health_check_port", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.LoadBalancer.HealthCheck.Port"), Description: "An integer representing the port on the backend Droplets on which the health check will attempt a connection."},
			{Name: "health_check_protocol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LoadBalancer.HealthCheck.Protocol"), Description: "The protocol used for health checks sent to the backend Droplets. The possible values are \"http\", \"https\", or \"tcp\"."},
			{Name: "health_check_response_timeout_seconds", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.LoadBalancer.HealthCheck.ResponseTimeoutSeconds"), Description: "The number of seconds the load balancer instance will wait for a response until marking a health check as failed."},
			{Name: "health_check_unhealthy_threshold", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.LoadBalancer.HealthCheck.UnhealthyThreshold"), Description: "The number of times a health check must fail for a backend Droplet to be marked \"unhealthy\" and be removed from the pool."},
			{Name: "ip", Type: proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.LoadBalancer.IP"),
				Description: "An attribute containing the public-facing IP address of the load balancer."},
			{Name: "redirect_http_to_https", Type: proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.LoadBalancer.RedirectHttpToHttps"),
				Description: "A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LoadBalancer.Region.Slug"), Description: "The unique slug identifier for the region the load balancer is deployed in."},
			{Name: "region_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LoadBalancer.Region.Name"), Description: "The name of the region the load balancer is deployed in."},
			{Name: "size_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LoadBalancer.SizeSlug"), Description: "The size of the load balancer. The available sizes are \"lb-small\", \"lb-medium\", or \"lb-large\"."},
			{Name: "status", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Status"),
				Description: "A status string indicating the current state of the load balancer. This can be \"new\", \"active\", or \"errored\"."},
			{Name: "sticky_sessions_cookie_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LoadBalancer.StickySessions.CookieName").NullIfZero(), Description: "The name of the cookie sent to the client. This attribute is only returned when using \"cookies\" for the sticky sessions type."},
			{Name: "sticky_sessions_cookie_ttl_seconds", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.LoadBalancer.StickySessions.CookieTtlSeconds").NullIfZero(), Description: "The number of seconds until the cookie set by the load balancer expires. This attribute is only returned when using \"cookies\" for the sticky sessions type."},
			{Name: "sticky_sessions_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LoadBalancer.StickySessions.Type"), Description: "An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are \"cookies\" or \"none\"."},
			{Name: "tag", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Tag"),
				Description: "The name of a Droplet tag corresponding to Droplets assigned to the load balancer."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(loadBalancerToURN), Description: "The uniform resource name (URN) for the load balancer."},
			{Name: "vpc_uuid", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.VPCUUID"),
				Description: "A string specifying the UUID of the VPC to which the load balancer is assigned."},

			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(loadBalancerToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.LoadBalancer.Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LoadBalancer.Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func loadBalancerToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(opengovernance.DigitalOceanLoadBalancer).Description.LoadBalancer
	return "do:loadBalancer:" + i.ID, nil
}
