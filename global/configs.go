package global

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "digitalocean"                                    // example: aws, azure
	IntegrationName      = integration.Type("digitalocean_team")                  // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-digitalocean" // example: github.com/opengovern/og-describer-aws
)


type IntegrationCredentials struct {
	AuthToken string `json:"auth_token"`
}

