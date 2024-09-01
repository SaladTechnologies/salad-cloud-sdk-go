package organizationdata

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

type OrganizationDataService struct {
	manager *configmanager.ConfigManager
}

func NewOrganizationDataService(manager *configmanager.ConfigManager) *OrganizationDataService {
	return &OrganizationDataService{
		manager: manager,
	}
}

func (api *OrganizationDataService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetOrganizationData()
}

func (api *OrganizationDataService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *OrganizationDataService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// List the GPU Classes
func (api *OrganizationDataService) ListGpuClasses(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[GpuClassesList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[GpuClassesList](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/gpu-classes", config)

	request.SetPathParam("organization_name", organizationName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[GpuClassesList](err)
	}

	return shared.NewSaladCloudSdkResponse[GpuClassesList](resp), nil
}
