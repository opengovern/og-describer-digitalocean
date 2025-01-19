package global

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "openai"                                    // example: aws, azure
	IntegrationName      = integration.Type("openai_integration")                  // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-openai" // example: github.com/opengovern/og-describer-aws
)


type IntegrationCredentials struct {
	APIKey         string `json:"api_key"`
	ProjectID      string `json:"project_id"`
	ProjectName    string `json:"project_name"`
	OrganizationID string `json:"organization_id"`
}

