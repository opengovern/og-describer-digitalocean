package maps
import (

	"github.com/opengovern/og-util/pkg/integration/interfaces"
	model "github.com/opengovern/og-describer-digitalocean/discovery/pkg/models"
)
var ResourceTypes = map[string]model.ResourceType{}

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{
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