package organizations

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

// OrganizationsService provides methods to interact with OrganizationsService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type OrganizationsService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewOrganizationsService() *OrganizationsService {
	return &OrganizationsService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *OrganizationsService) WithConfigManager(manager *configmanager.ConfigManager) *OrganizationsService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *OrganizationsService) WithHook(hook hooks.Hook) *OrganizationsService {
	api.hook = hook
	return api
}

func (api *OrganizationsService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetOrganizations()
}

func (api *OrganizationsService) getHook() hooks.Hook {
	return api.hook
}

func (api *OrganizationsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *OrganizationsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *OrganizationsService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the CPU availability for the given organization
func (api *OrganizationsService) GetCpuAvailability(ctx context.Context, organizationName string, cpuAvailabilityPrototype CpuAvailabilityPrototype) (*shared.SaladCloudSdkResponse[CpuAvailability], *shared.SaladCloudSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/availability/sce-cpu-availability").
		WithConfig(config).
		WithBody(cpuAvailabilityPrototype).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("organization_name", organizationName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[CpuAvailability, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[CpuAvailability](resp), nil
}

// Gets the GPU availability for the given organization
func (api *OrganizationsService) GetGpuAvailability(ctx context.Context, organizationName string, gpuAvailabilityPrototype GpuAvailabilityPrototype) (*shared.SaladCloudSdkResponse[GpuAvailability], *shared.SaladCloudSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/availability/sce-gpu-availability").
		WithConfig(config).
		WithBody(gpuAvailabilityPrototype).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("organization_name", organizationName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GpuAvailability, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[GpuAvailability](resp), nil
}
