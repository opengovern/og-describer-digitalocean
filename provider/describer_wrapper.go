package provider

import (
	model "github.com/opengovern/og-describer-digitalocean/pkg/sdk/models"
	"github.com/opengovern/og-describer-digitalocean/provider/configs"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
)

// DescribeByIntegration TODO: implement a wrapper to pass integration authorization to describer functions
func DescribeByIntegration(describe func(context.Context, *configs.IntegrationCredentials, string, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		return describe(ctx, &cfg, string(triggerType), stream)
	}
}
