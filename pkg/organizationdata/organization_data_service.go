package organizationdata

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/hooks"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"time"
)

// OrganizationDataService provides methods to interact with OrganizationDataService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type OrganizationDataService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewOrganizationDataService() *OrganizationDataService {
	return &OrganizationDataService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *OrganizationDataService) WithConfigManager(manager *configmanager.ConfigManager) *OrganizationDataService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *OrganizationDataService) WithHook(hook hooks.Hook) *OrganizationDataService {
	api.hook = hook
	return api
}

func (api *OrganizationDataService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetOrganizationData()
}

func (api *OrganizationDataService) getHook() hooks.Hook {
	return api.hook
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
func (api *OrganizationDataService) ListGpuClasses(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[GpuClassesList], *shared.SaladCloudSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/gpu-classes").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GpuClassesList, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[GpuClassesList](resp), nil
}
