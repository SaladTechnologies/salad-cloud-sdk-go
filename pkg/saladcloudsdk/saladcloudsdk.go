package saladcloudsdk

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/containergroups"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/organizationdata"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/quotas"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/systemlogs"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/webhooksecretkey"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/workloaderrors"
	"time"
)

type SaladCloudSdk struct {
	ContainerGroups    *containergroups.ContainerGroupsService
	WorkloadErrors     *workloaderrors.WorkloadErrorsService
	SystemLogs         *systemlogs.SystemLogsService
	Queues             *queues.QueuesService
	Quotas             *quotas.QuotasService
	InferenceEndpoints *inferenceendpoints.InferenceEndpointsService
	OrganizationData   *organizationdata.OrganizationDataService
	WebhookSecretKey   *webhooksecretkey.WebhookSecretKeyService
	manager            *configmanager.ConfigManager
}

func NewSaladCloudSdk(config saladcloudsdkconfig.Config) *SaladCloudSdk {
	containerGroups := containergroups.NewContainerGroupsService()
	workloadErrors := workloaderrors.NewWorkloadErrorsService()
	systemLogs := systemlogs.NewSystemLogsService()
	queues := queues.NewQueuesService()
	quotas := quotas.NewQuotasService()
	inferenceEndpoints := inferenceendpoints.NewInferenceEndpointsService()
	organizationData := organizationdata.NewOrganizationDataService()
	webhookSecretKey := webhooksecretkey.NewWebhookSecretKeyService()

	manager := configmanager.NewConfigManager(config)
	containerGroups.WithConfigManager(manager)
	workloadErrors.WithConfigManager(manager)
	systemLogs.WithConfigManager(manager)
	queues.WithConfigManager(manager)
	quotas.WithConfigManager(manager)
	inferenceEndpoints.WithConfigManager(manager)
	organizationData.WithConfigManager(manager)
	webhookSecretKey.WithConfigManager(manager)

	return &SaladCloudSdk{
		ContainerGroups:    containerGroups,
		WorkloadErrors:     workloadErrors,
		SystemLogs:         systemLogs,
		Queues:             queues,
		Quotas:             quotas,
		InferenceEndpoints: inferenceEndpoints,
		OrganizationData:   organizationData,
		WebhookSecretKey:   webhookSecretKey,
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
