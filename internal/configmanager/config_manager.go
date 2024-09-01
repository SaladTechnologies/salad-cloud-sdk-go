package configmanager

import "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"

type ConfigManager struct {
	ContainerGroups    saladcloudsdkconfig.Config
	WorkloadErrors     saladcloudsdkconfig.Config
	Queues             saladcloudsdkconfig.Config
	Quotas             saladcloudsdkconfig.Config
	InferenceEndpoints saladcloudsdkconfig.Config
	OrganizationData   saladcloudsdkconfig.Config
	WebhookSecretKey   saladcloudsdkconfig.Config
}

func NewConfigManager(config saladcloudsdkconfig.Config) *ConfigManager {
	return &ConfigManager{
		ContainerGroups:    config,
		WorkloadErrors:     config,
		Queues:             config,
		Quotas:             config,
		InferenceEndpoints: config,
		OrganizationData:   config,
		WebhookSecretKey:   config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.ContainerGroups.SetBaseUrl(baseUrl)
	c.WorkloadErrors.SetBaseUrl(baseUrl)
	c.Queues.SetBaseUrl(baseUrl)
	c.Quotas.SetBaseUrl(baseUrl)
	c.InferenceEndpoints.SetBaseUrl(baseUrl)
	c.OrganizationData.SetBaseUrl(baseUrl)
	c.WebhookSecretKey.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) SetApiKey(apiKey string) {
	c.ContainerGroups.SetApiKey(apiKey)
	c.WorkloadErrors.SetApiKey(apiKey)
	c.Queues.SetApiKey(apiKey)
	c.Quotas.SetApiKey(apiKey)
	c.InferenceEndpoints.SetApiKey(apiKey)
	c.OrganizationData.SetApiKey(apiKey)
	c.WebhookSecretKey.SetApiKey(apiKey)
}

func (c *ConfigManager) GetContainerGroups() *saladcloudsdkconfig.Config {
	return &c.ContainerGroups
}
func (c *ConfigManager) GetWorkloadErrors() *saladcloudsdkconfig.Config {
	return &c.WorkloadErrors
}
func (c *ConfigManager) GetQueues() *saladcloudsdkconfig.Config {
	return &c.Queues
}
func (c *ConfigManager) GetQuotas() *saladcloudsdkconfig.Config {
	return &c.Quotas
}
func (c *ConfigManager) GetInferenceEndpoints() *saladcloudsdkconfig.Config {
	return &c.InferenceEndpoints
}
func (c *ConfigManager) GetOrganizationData() *saladcloudsdkconfig.Config {
	return &c.OrganizationData
}
func (c *ConfigManager) GetWebhookSecretKey() *saladcloudsdkconfig.Config {
	return &c.WebhookSecretKey
}
