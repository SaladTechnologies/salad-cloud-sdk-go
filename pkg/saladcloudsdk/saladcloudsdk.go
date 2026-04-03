package saladcloudsdk

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/hooks"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/containergroups"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/logs"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/organizationdata"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/organizations"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/quotas"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/systemlogs"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/webhooksecretkey"
	"time"
)

// SaladCloudSdk is the main SDK client that provides access to all service endpoints.
// It manages configuration, authentication, and service instances with centralized settings.
type SaladCloudSdk struct {
	ContainerGroups    *containergroups.ContainerGroupsService
	SystemLogs         *systemlogs.SystemLogsService
	Queues             *queues.QueuesService
	Quotas             *quotas.QuotasService
	InferenceEndpoints *inferenceendpoints.InferenceEndpointsService
	OrganizationData   *organizationdata.OrganizationDataService
	WebhookSecretKey   *webhooksecretkey.WebhookSecretKeyService
	Logs               *logs.LogsService
	Organizations      *organizations.OrganizationsService
	manager            *configmanager.ConfigManager
}

func NewSaladCloudSdk(config saladcloudsdkconfig.Config) *SaladCloudSdk {
	containerGroups := containergroups.NewContainerGroupsService()
	systemLogs := systemlogs.NewSystemLogsService()
	queues := queues.NewQueuesService()
	quotas := quotas.NewQuotasService()
	inferenceEndpoints := inferenceendpoints.NewInferenceEndpointsService()
	organizationData := organizationdata.NewOrganizationDataService()
	webhookSecretKey := webhooksecretkey.NewWebhookSecretKeyService()
	logs := logs.NewLogsService()
	organizations := organizations.NewOrganizationsService()

	manager := configmanager.NewConfigManager(config)
	hook := hooks.NewDefaultHook()
	containerGroups.WithConfigManager(manager)
	systemLogs.WithConfigManager(manager)
	queues.WithConfigManager(manager)
	quotas.WithConfigManager(manager)
	inferenceEndpoints.WithConfigManager(manager)
	organizationData.WithConfigManager(manager)
	webhookSecretKey.WithConfigManager(manager)
	logs.WithConfigManager(manager)
	organizations.WithConfigManager(manager)
	containerGroups.WithHook(hook)
	systemLogs.WithHook(hook)
	queues.WithHook(hook)
	quotas.WithHook(hook)
	inferenceEndpoints.WithHook(hook)
	organizationData.WithHook(hook)
	webhookSecretKey.WithHook(hook)
	logs.WithHook(hook)
	organizations.WithHook(hook)

	return &SaladCloudSdk{
		ContainerGroups:    containerGroups,
		SystemLogs:         systemLogs,
		Queues:             queues,
		Quotas:             quotas,
		InferenceEndpoints: inferenceEndpoints,
		OrganizationData:   organizationData,
		WebhookSecretKey:   webhookSecretKey,
		Logs:               logs,
		Organizations:      organizations,
		manager:            manager,
	}
}

func (s *SaladCloudSdk) SetBaseUrl(baseUrl string) {
	s.manager.SetBaseUrl(baseUrl)
}

func (s *SaladCloudSdk) SetTimeout(timeout time.Duration) {
	s.manager.SetTimeout(timeout)
}

func (s *SaladCloudSdk) SetApiKey(apiKey string) {
	s.manager.SetApiKey(apiKey)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
