package maps
import (
	"github.com/opengovern/og-describer-openai/discovery/describers"
	"github.com/opengovern/og-describer-openai/discovery/provider"
	"github.com/opengovern/og-describer-openai/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
	model "github.com/opengovern/og-describer-openai/discovery/pkg/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"OpenAI/Project": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "OpenAI/Project",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByOpenAI(describers.ListProjects),
		GetDescriber:         provider.DescribeSingleByOpenAI(describers.GetProject),
	},

	"OpenAI/Project/ApiKey": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "OpenAI/Project/ApiKey",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByOpenAI(describers.ListProjectAPIKeys),
		GetDescriber:         nil,
	},

	"OpenAI/Project/RateLimit": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "OpenAI/Project/RateLimit",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByOpenAI(describers.ListProjectRateLimits),
		GetDescriber:         nil,
	},

	"OpenAI/Project/ServiceAccount": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "OpenAI/Project/ServiceAccount",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByOpenAI(describers.ListProjectServiceAccounts),
		GetDescriber:         nil,
	},

	"OpenAI/Project/User": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "OpenAI/Project/User",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByOpenAI(describers.ListProjectUsers),
		GetDescriber:         nil,
	},

	"OpenAI/Model": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "OpenAI/Model",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByOpenAI(describers.ListModels),
		GetDescriber:         provider.DescribeSingleByOpenAI(describers.GetModel),
	},

	"OpenAI/File": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "OpenAI/File",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByOpenAI(describers.ListFiles),
		GetDescriber:         nil,
	},

	"OpenAI/VectorStore": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "OpenAI/VectorStore",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByOpenAI(describers.ListVectorStores),
		GetDescriber:         provider.DescribeSingleByOpenAI(describers.GetVectorStore),
	},

	"OpenAI/Assistant": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "OpenAI/Assistant",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByOpenAI(describers.ListAssistants),
		GetDescriber:         provider.DescribeSingleByOpenAI(describers.GetAssistant),
	},
}


var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"OpenAI/Project": {
		Name:         "OpenAI/Project",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"OpenAI/Project/ApiKey": {
		Name:         "OpenAI/Project/ApiKey",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"OpenAI/Project/RateLimit": {
		Name:         "OpenAI/Project/RateLimit",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"OpenAI/Project/ServiceAccount": {
		Name:         "OpenAI/Project/ServiceAccount",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"OpenAI/Project/User": {
		Name:         "OpenAI/Project/User",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"OpenAI/Model": {
		Name:         "OpenAI/Model",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"OpenAI/File": {
		Name:         "OpenAI/File",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"OpenAI/VectorStore": {
		Name:         "OpenAI/VectorStore",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"OpenAI/Assistant": {
		Name:         "OpenAI/Assistant",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},
}


var ResourceTypesList = []string{
  "OpenAI/Project",
  "OpenAI/Project/ApiKey",
  "OpenAI/Project/RateLimit",
  "OpenAI/Project/ServiceAccount",
  "OpenAI/Project/User",
  "OpenAI/Model",
  "OpenAI/File",
  "OpenAI/VectorStore",
  "OpenAI/Assistant",
}