package saladcloudsdk

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/containergroups"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/organizationdata"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/quotas"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/webhooksecretkey"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/workloaderrors"
)

type SaladCloudSdk struct {
	ContainerGroups    *containergroups.ContainerGroupsService
	WorkloadErrors     *workloaderrors.WorkloadErrorsService
	Queues             *queues.QueuesService
	Quotas             *quotas.QuotasService
	InferenceEndpoints *inferenceendpoints.InferenceEndpointsService
	OrganizationData   *organizationdata.OrganizationDataService
	WebhookSecretKey   *webhooksecretkey.WebhookSecretKeyService
	manager            *configmanager.ConfigManager
}

func NewSaladCloudSdk(config saladcloudsdkconfig.Config) *SaladCloudSdk {
	manager := configmanager.NewConfigManager(config)
	return &SaladCloudSdk{
		ContainerGroups:    containergroups.NewContainerGroupsService(manager),
		WorkloadErrors:     workloaderrors.NewWorkloadErrorsService(manager),
		Queues:             queues.NewQueuesService(manager),
		Quotas:             quotas.NewQuotasService(manager),
		InferenceEndpoints: inferenceendpoints.NewInferenceEndpointsService(manager),
		OrganizationData:   organizationdata.NewOrganizationDataService(manager),
		WebhookSecretKey:   webhooksecretkey.NewWebhookSecretKeyService(manager),
		manager:            manager,
	}
}

func (s *SaladCloudSdk) SetBaseUrl(baseUrl string) {
	s.manager.SetBaseUrl(baseUrl)
}

func (s *SaladCloudSdk) SetApiKey(apiKey string) {
	s.manager.SetApiKey(apiKey)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
