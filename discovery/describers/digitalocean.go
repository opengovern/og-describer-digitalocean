package describers

import (
	"context"
	"errors"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/opengovern/og-describer-digitalocean/discovery/pkg/models"
	model "github.com/opengovern/og-describer-digitalocean/discovery/provider"
	"strings"
)

func DigitalOceanAccount(ctx context.Context, creds *models.IntegrationCredentials, triggerType string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	account, _, err := client.Account.Get(ctx)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, nil
	}
	resource := models.Resource{
		ID:   account.UUID,
		Name: account.Name,
		Description: model.DigitalOceanAccountDescription{
			Account: *account,
		},
	}

	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	}

	return []models.Resource{resource}, nil
}

func getAccountUUID(ctx context.Context, creds *models.IntegrationCredentials) (string, error) {
	client := godo.NewFromToken(creds.AuthToken)
	account, _, err := client.Account.Get(ctx)
	if err != nil {
		return "", err
	}
	if account == nil {
		return "", errors.New("account not found")
	}
	return account.UUID, nil
}

func DigitalOceanAction(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}
	resources := make([]models.Resource, 0)

	opt := &godo.ListOptions{}
	for {
		actions, resp, err := client.Actions.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if actions == nil {
			return nil, nil
		}

		for _, action := range actions {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:action:%d", accountUUID, action.ID),
				Name: fmt.Sprintf("%d", action.ID),
				Description: model.DigitalOceanActionDescription{
					Action: action,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanAlertPolicy(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}
	resources := make([]models.Resource, 0)

	opt := &godo.ListOptions{}
	for {
		policies, resp, err := client.Monitoring.ListAlertPolicies(ctx, opt)
		if err != nil {
			return nil, err
		}
		if policies == nil {
			return nil, nil
		}

		for _, policy := range policies {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:alert_policy:%s", accountUUID, policy.UUID),
				Name: policy.UUID,
				Description: model.DigitalOceanAlertPolicyDescription{
					AlertPolicy: policy,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanApp(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		apps, resp, err := client.Apps.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if apps == nil {
			return nil, nil
		}

		for _, app := range apps {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:app:%s", accountUUID, app.ID),
				Name: app.ID,
				Description: model.DigitalOceanAppDescription{
					App: *app,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanBalance(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	balance, _, err := client.Balance.Get(ctx)
	if err != nil {
		return nil, err
	}
	if balance == nil {
		return nil, nil
	}

	resource := models.Resource{
		ID:   fmt.Sprintf("do:%s:balance", accountUUID),
		Name: "Balance",
		Description: model.DigitalOceanBalanceDescription{
			Balance: *balance,
		},
	}

	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	}

	return []models.Resource{resource}, nil
}

func DigitalOceanBill(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		bills, resp, err := client.BillingHistory.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if bills == nil {
			return nil, nil
		}

		for _, bill := range bills.BillingHistory {
			id := bill.Date.Format("2006-01-02")
			if bill.InvoiceID != nil {
				id = *bill.InvoiceID
			} else if bill.InvoiceUUID != nil {
				id = *bill.InvoiceUUID
			}
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:bill:%s", accountUUID, id),
				Name: id,
				Description: model.DigitalOceanBillDescription{
					Bill: bill,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanContainerRegistry(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	registry, _, err := client.Registry.Get(ctx)
	if err != nil {
		return nil, err
	}
	if registry == nil {
		return nil, nil
	}

	resource := models.Resource{
		ID:   fmt.Sprintf("do:%s:container_registry", accountUUID),
		Name: registry.Name,
		Description: model.DigitalOceanContainerRegistryDescription{
			ContainerRegistry: *registry,
		},
	}

	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	}

	return []models.Resource{resource}, nil
}

func DigitalOceanDatabase(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		databases, resp, err := client.Databases.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if databases == nil {
			return nil, nil
		}

		for _, database := range databases {
			description := model.DigitalOceanDatabaseDescription{
				Database: database,
			}

			userOpts := &godo.ListOptions{}
			for {
				users, resp, err := client.Databases.ListUsers(ctx, database.ID, userOpts)
				if err != nil {
					if strings.Contains(err.Error(), ": 404") {
						break
					}
					return nil, err
				}
				if users != nil {
					description.DatabaseUsers = append(description.DatabaseUsers, users...)
				}

				if resp.Links == nil || resp.Links.IsLastPage() {
					break
				}
				page, err := resp.Links.CurrentPage()
				if err != nil {
					return nil, err
				}
				userOpts.Page = page + 1
			}

			dbOpts := &godo.ListOptions{}
			for {
				names, resp, err := client.Databases.ListDBs(ctx, database.ID, dbOpts)
				if err != nil {
					if strings.Contains(err.Error(), ": 404") {
						break
					}
					return nil, err
				}
				if names != nil {
					description.DatabaseNames = append(description.DatabaseNames, names...)
				}

				if resp.Links == nil || resp.Links.IsLastPage() {
					break
				}
				page, err := resp.Links.CurrentPage()
				if err != nil {
					return nil, err
				}
				dbOpts.Page = page + 1
			}

			firewallRules, _, err := client.Databases.GetFirewallRules(ctx, database.ID)
			if err != nil {
				if strings.Contains(err.Error(), ": 404") {
					break
				}
				return nil, err
			}
			if firewallRules != nil {
				description.DatabaseFirewallRules = firewallRules
			}

			resource := models.Resource{
				ID:          fmt.Sprintf("do:%s:database:%s", accountUUID, database.ID),
				Name:        database.Name,
				Description: description,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanDomain(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		domains, resp, err := client.Domains.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if domains == nil {
			return nil, nil
		}

		for _, domain := range domains {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:domain:%s", accountUUID, domain.Name),
				Name: domain.Name,
				Description: model.DigitalOceanDomainDescription{
					Domain: domain,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanDroplet(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := client.Droplets.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if droplets == nil {
			return nil, nil
		}

		for _, droplet := range droplets {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:droplet:%d", accountUUID, droplet.ID),
				Name: fmt.Sprintf("%d", droplet.ID),
				Description: model.DigitalOceanDropletDescription{
					URN:     droplet.URN(),
					Droplet: droplet,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanFirewall(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		firewalls, resp, err := client.Firewalls.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if firewalls == nil {
			return nil, nil
		}

		for _, firewall := range firewalls {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:firewall:%s", accountUUID, firewall.ID),
				Name: firewall.Name,
				Description: model.DigitalOceanFirewallDescription{
					Firewall: firewall,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanFloatingIP(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		floatingIPs, resp, err := client.FloatingIPs.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if floatingIPs == nil {
			return nil, nil
		}

		for _, floatingIP := range floatingIPs {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:floating_ip:%s", accountUUID, strings.ReplaceAll(floatingIP.IP, ".", "_")),
				Name: floatingIP.IP,
				Description: model.DigitalOceanFloatingIPDescription{
					URN:        floatingIP.URN(),
					FloatingIP: floatingIP,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanImage(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		images, resp, err := client.Images.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if images == nil {
			return nil, nil
		}

		for _, image := range images {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:image:%d", accountUUID, image.ID),
				Name: image.Name,
				Description: model.DigitalOceanImageDescription{
					Image: image,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanKey(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		keys, resp, err := client.Keys.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if keys == nil {
			return nil, nil
		}

		for _, key := range keys {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:key:%d", accountUUID, key.ID),
				Name: key.Name,
				Description: model.DigitalOceanKeyDescription{
					Key: key,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanKubernetesCluster(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		clusters, resp, err := client.Kubernetes.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if clusters == nil {
			return nil, nil
		}

		for _, cluster := range clusters {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:kubernetes_cluster:%s", accountUUID, cluster.ID),
				Name: cluster.Name,
				Description: model.DigitalOceanKubernetesClusterDescription{
					Cluster: *cluster,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanKubernetesNodePool(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		clusters, resp, err := client.Kubernetes.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if clusters == nil {
			return nil, nil
		}

		for _, cluster := range clusters {
			nodePoolOpts := &godo.ListOptions{}
			for {
				nodePools, npResp, err := client.Kubernetes.ListNodePools(ctx, cluster.ID, nodePoolOpts)
				if err != nil {
					return nil, err
				}
				if nodePools == nil {
					return nil, nil
				}

				for _, nodePool := range nodePools {
					resource := models.Resource{
						ID:   fmt.Sprintf("do:%s:kubernetes_cluster:%s:node_pool:%s", accountUUID, cluster.ID, nodePool.Name),
						Name: nodePool.Name,
						Description: model.DigitalOceanKubernetesNodePoolDescription{
							ClusterID: cluster.ID,
							NodePool:  *nodePool,
						},
					}
					if stream != nil {
						if err := (*stream)(resource); err != nil {
							return nil, err
						}
					} else {
						resources = append(resources, resource)
					}
				}

				if npResp.Links == nil || npResp.Links.IsLastPage() {
					break
				}
				page, err := npResp.Links.CurrentPage()
				if err != nil {
					return nil, err
				}
				nodePoolOpts.Page = page + 1
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanLoadBalancer(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		loadBalancers, resp, err := client.LoadBalancers.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if loadBalancers == nil {
			return nil, nil
		}

		for _, lb := range loadBalancers {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:load_balancer:%s", accountUUID, lb.ID),
				Name: lb.Name,
				Description: model.DigitalOceanLoadBalancerDescription{
					LoadBalancer: lb,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanProject(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}
	resources := make([]models.Resource, 0)

	opt := &godo.ListOptions{}
	for {
		projects, resp, err := client.Projects.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if projects == nil {
			return nil, nil
		}

		for _, project := range projects {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:project:%s", accountUUID, project.ID),
				Name: project.Name,
				Description: model.DigitalOceanProjectDescription{
					Project: project,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanRegion(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}
	resources := make([]models.Resource, 0)

	opt := &godo.ListOptions{}
	for {
		regions, resp, err := client.Regions.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if regions == nil {
			return nil, nil
		}

		for _, region := range regions {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:region:%s", accountUUID, region.Slug),
				Name: region.Name,
				Description: model.DigitalOceanRegionDescription{
					Region: region,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanSize(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}
	resources := make([]models.Resource, 0)

	opt := &godo.ListOptions{}
	for {
		sizes, resp, err := client.Sizes.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if sizes == nil {
			return nil, nil
		}

		for _, size := range sizes {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:size:%s", accountUUID, size.Slug),
				Name: size.Slug,
				Description: model.DigitalOceanSizeDescription{
					Size: size,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanSnapshot(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}
	resources := make([]models.Resource, 0)

	opt := &godo.ListOptions{}
	for {
		snapshots, resp, err := client.Snapshots.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if snapshots == nil {
			return nil, nil
		}

		for _, snapshot := range snapshots {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:snapshot:%s", accountUUID, snapshot.ID),
				Name: snapshot.Name,
				Description: model.DigitalOceanSnapshotDescription{
					Snapshot: snapshot,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanTag(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)
	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}
	resources := make([]models.Resource, 0)

	opt := &godo.ListOptions{}
	for {
		tags, resp, err := client.Tags.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if tags == nil {
			return nil, nil
		}

		for _, tag := range tags {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:tag:%s", accountUUID, tag.Name),
				Name: tag.Name,
				Description: model.DigitalOceanTagDescription{
					Tag: tag,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanVolume(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)

	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListVolumeParams{
		ListOptions: &godo.ListOptions{},
	}
	for {
		volumes, resp, err := client.Storage.ListVolumes(ctx, opt)
		if err != nil {
			return nil, err
		}
		if volumes == nil {
			return nil, nil
		}

		for _, volume := range volumes {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:volume:%s", accountUUID, volume.ID),
				Name: volume.Name,
				Description: model.DigitalOceanVolumeDescription{
					URN:    volume.URN(),
					Volume: volume,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.ListOptions.Page = page + 1
	}
	return resources, nil
}

func DigitalOceanVPC(ctx context.Context, creds *models.IntegrationCredentials, _ string, stream *models.StreamSender) ([]models.Resource, error) {
	client := godo.NewFromToken(creds.AuthToken)

	accountUUID, err := getAccountUUID(ctx, creds)
	if err != nil {
		return nil, err
	}

	resources := make([]models.Resource, 0)
	opt := &godo.ListOptions{}
	for {
		vpcs, resp, err := client.VPCs.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		if vpcs == nil {
			return nil, nil
		}

		for _, vpc := range vpcs {
			resource := models.Resource{
				ID:   fmt.Sprintf("do:%s:vpc:%s", accountUUID, vpc.ID),
				Name: vpc.Name,
				Description: model.DigitalOceanVPCDescription{
					VPC: *vpc,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, resource)
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.Page = page + 1
	}
	return resources, nil
}
