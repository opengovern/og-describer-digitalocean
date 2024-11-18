package digitalocean

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-digitalocean/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanDatabase(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_database",
		Description: "DigitalOcean's managed database service simplifies the creation and management of highly available database clusters. Currently, it offers support for PostgreSQL, Redis, and MySQL.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDigitalOceanDatabase,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDigitalOceanDatabase,
		},
		Columns: integrationColumns([]*plugin.Column{
			// Top columns
			{
				Name:      "id",
				Transform: transform.FromField("Description.Database.ID"),
				Type:      proto.ColumnType_STRING, Description: "A unique ID that can be used to identify and reference a database cluster."},
			{
				Name:      "name",
				Transform: transform.FromField("Description.Database.Name"),
				Type:      proto.ColumnType_STRING, Description: "A unique, human-readable name referring to a database cluster."},
			{
				Name:      "engine",
				Transform: transform.FromField("Description.Database.Engine"),
				Type:      proto.ColumnType_STRING, Description: "A slug representing the database engine used for the cluster. The possible values are: \"pg\" for PostgreSQL, \"mysql\" for MySQL, and \"redis\" for Redis."},
			{
				Name:      "version",
				Transform: transform.FromField("Description.Database.Version"),
				Type:      proto.ColumnType_STRING, Description: "A string representing the version of the database engine in use for the cluster."},
			// Other columns
			{Name: "connection_database", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.Connection.Database"), Description: "The name of the default database."},
			{Name: "connection_host", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.Connection.Host"), Description: "A public FQDN pointing to the database cluster's current primary node."},
			{Name: "connection_password", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.Connection.Password"), Description: "The randomly generated password for the default user."},
			{Name: "connection_port", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Database.Connection.Port"), Description: "The port on which the database cluster is listening."},
			{Name: "connection_ssl", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Database.Connection.SSL"), Description: "A boolean value indicating if the connection should be made over SSL."},
			{Name: "connection_uri", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.Connection.URI"), Description: "A connection string in the format accepted by the psql command. This is provided as a convenience and should be able to be constructed by the other attributes."},
			{Name: "connection_user", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.Connection.User"), Description: "The default user for the database."},
			{
				Name:      "created_at",
				Transform: transform.FromField("Description.Database.CreatedAt"),
				Type:      proto.ColumnType_TIMESTAMP, Description: "Time when the database was created."},

			{Name: "db_names", Type: proto.ColumnType_JSON, Description: "An array of strings containing the names of databases created in the database cluster.", Transform: transform.FromField("Description.DatabaseNames")},
			{Name: "firewall_rules", Type: proto.ColumnType_JSON, Description: "A list of rules describing the inbound source to a database.", Transform: transform.FromField("Description.DatabaseFirewallRules")},

			{Name: "maintenance_window_day", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.MaintenanceWindow.Day"), Description: "The day of the week on which to apply maintenance updates (e.g. \"tuesday\")."},
			{Name: "maintenance_window_description", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Database.MaintenanceWindow.Description"), Description: "A list of strings, each containing information about a pending maintenance update."},
			{Name: "maintenance_window_hour", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.MaintenanceWindow.Hour"), Description: "The hour in UTC at which maintenance updates will be applied in 24 hour format (e.g. \"16:00:00\")."},
			{Name: "maintenance_window_pending", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Database.MaintenanceWindow.Pending"), Description: "A boolean value indicating whether any maintenance is scheduled to be performed in the next window."},
			{Name: "num_nodes",
				Transform: transform.FromField("Description.Database.NumNodes"),
				Type:      proto.ColumnType_INT, Description: "The number of nodes in the database cluster."},
			{Name: "private_connection_database", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.PrivateConnection.Database"), Description: "The name of the default database."},
			{Name: "private_connection_host", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.PrivateConnection.Host"), Description: "The private FQDN pointing to the database cluster's current primary node."},
			{Name: "private_connection_password", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.PrivateConnection.Password"), Description: "The randomly generated password for the default user."},
			{Name: "private_connection_port", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Database.PrivateConnection.Port"), Description: "The port on which the database cluster is listening."},
			{Name: "private_connection_ssl", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Database.PrivateConnection.SSL"), Description: "A boolean value indicating if the connection should be made over SSL."},
			{Name: "private_connection_uri", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.PrivateConnection.URI"), Description: "A connection string in the format accepted by the psql command. This is provided as a convenience and should be able to be constructed by the other attributes."},
			{Name: "private_connection_user", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.PrivateConnection.User"), Description: "The default user for the database."},
			{Name: "private_network_uuid",
				Transform: transform.FromField("Description.Database.PrivateNetworkUUID"),
				Type:      proto.ColumnType_STRING, Description: "A string specifying the UUID of the VPC to which the database cluster is assigned."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.RegionSlug"), Description: "The unique slug identifier for the region the database is deployed in."},
			{Name: "size_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.SizeSlug"), Description: "The slug identifier representing the size of the nodes in the database cluster."},
			{Name: "status", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Status"),
				Description: "A string representing the current status of the database cluster. Possible values include creating, online, resizing, and migrating."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Database.Tags"), Description: "An array of tags that have been applied to the database cluster."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(databaseToURN), Description: "The uniform resource name (URN) for the database."},

			{Name: "users", Type: proto.ColumnType_JSON, Description: "An array containing objects describing the database's users.", Transform: transform.FromField("Description.DatabaseUsers")},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(databaseToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Database.Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Database.Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func databaseToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(opengovernance.DigitalOceanDatabase)
	return "do:database:" + i.Description.Database.ID, nil
}
