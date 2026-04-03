package configmanager

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"time"
)

// ConfigManager manages configuration across all services with synchronized updates.
// Provides centralized configuration management and OAuth token handling for multiple services.
type ConfigManager struct {
	ContainerGroups    saladcloudsdkconfig.Config
	SystemLogs         saladcloudsdkconfig.Config
	Queues             saladcloudsdkconfig.Config
	Quotas             saladcloudsdkconfig.Config
	InferenceEndpoints saladcloudsdkconfig.Config
	OrganizationData   saladcloudsdkconfig.Config
	WebhookSecretKey   saladcloudsdkconfig.Config
	Logs               saladcloudsdkconfig.Config
	Organizations      saladcloudsdkconfig.Config
}

// NewConfigManager creates a new configuration manager with the provided config and optional OAuth token service.
// Initializes service-specific configs and sets up OAuth token management if enabled.
func NewConfigManager(config saladcloudsdkconfig.Config) *ConfigManager {
	return &ConfigManager{
		ContainerGroups:    config,
		SystemLogs:         config,
		Queues:             config,
		Quotas:             config,
		InferenceEndpoints: config,
		OrganizationData:   config,
		WebhookSecretKey:   config,
		Logs:               config,
		Organizations:      config,
	}
}

// SetBaseUrl updates the BaseUrl configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.ContainerGroups.SetBaseUrl(baseUrl)
	c.SystemLogs.SetBaseUrl(baseUrl)
	c.Queues.SetBaseUrl(baseUrl)
	c.Quotas.SetBaseUrl(baseUrl)
	c.InferenceEndpoints.SetBaseUrl(baseUrl)
	c.OrganizationData.SetBaseUrl(baseUrl)
	c.WebhookSecretKey.SetBaseUrl(baseUrl)
	c.Logs.SetBaseUrl(baseUrl)
	c.Organizations.SetBaseUrl(baseUrl)
}

// SetTimeout updates the Timeout configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.ContainerGroups.SetTimeout(timeout)
	c.SystemLogs.SetTimeout(timeout)
	c.Queues.SetTimeout(timeout)
	c.Quotas.SetTimeout(timeout)
	c.InferenceEndpoints.SetTimeout(timeout)
	c.OrganizationData.SetTimeout(timeout)
	c.WebhookSecretKey.SetTimeout(timeout)
	c.Logs.SetTimeout(timeout)
	c.Organizations.SetTimeout(timeout)
}

// SetApiKey updates the ApiKey configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetApiKey(apiKey string) {
	c.ContainerGroups.SetApiKey(apiKey)
	c.SystemLogs.SetApiKey(apiKey)
	c.Queues.SetApiKey(apiKey)
	c.Quotas.SetApiKey(apiKey)
	c.InferenceEndpoints.SetApiKey(apiKey)
	c.OrganizationData.SetApiKey(apiKey)
	c.WebhookSecretKey.SetApiKey(apiKey)
	c.Logs.SetApiKey(apiKey)
	c.Organizations.SetApiKey(apiKey)
}

// GetContainerGroups returns the configuration for the ContainerGroups service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetContainerGroups() *saladcloudsdkconfig.Config {
	return &c.ContainerGroups
}

// GetSystemLogs returns the configuration for the SystemLogs service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSystemLogs() *saladcloudsdkconfig.Config {
	return &c.SystemLogs
}

// GetQueues returns the configuration for the Queues service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetQueues() *saladcloudsdkconfig.Config {
	return &c.Queues
}

// GetQuotas returns the configuration for the Quotas service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetQuotas() *saladcloudsdkconfig.Config {
	return &c.Quotas
}

// GetInferenceEndpoints returns the configuration for the InferenceEndpoints service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetInferenceEndpoints() *saladcloudsdkconfig.Config {
	return &c.InferenceEndpoints
}

// GetOrganizationData returns the configuration for the OrganizationData service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetOrganizationData() *saladcloudsdkconfig.Config {
	return &c.OrganizationData
}

// GetWebhookSecretKey returns the configuration for the WebhookSecretKey service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWebhookSecretKey() *saladcloudsdkconfig.Config {
	return &c.WebhookSecretKey
}

// GetLogs returns the configuration for the Logs service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetLogs() *saladcloudsdkconfig.Config {
	return &c.Logs
}

// GetOrganizations returns the configuration for the Organizations service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetOrganizations() *saladcloudsdkconfig.Config {
	return &c.Organizations
}
