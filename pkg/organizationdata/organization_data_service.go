package organizationdata

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"time"
)

type OrganizationDataService struct {
	manager *configmanager.ConfigManager
}

func NewOrganizationDataService() *OrganizationDataService {
	return &OrganizationDataService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

func (api *OrganizationDataService) WithConfigManager(manager *configmanager.ConfigManager) *OrganizationDataService {
	api.manager = manager
	return api
}

func (api *OrganizationDataService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetOrganizationData()
}

func (api *OrganizationDataService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *OrganizationDataService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *OrganizationDataService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// List the GPU Classes
func (api *OrganizationDataService) ListGpuClasses(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[GpuClassesList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/gpu-classes").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GpuClassesList](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[GpuClassesList](err)
	}

	return shared.NewSaladCloudSdkResponse[GpuClassesList](resp), nil
}
