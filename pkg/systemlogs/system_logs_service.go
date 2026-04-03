package systemlogs

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

// SystemLogsService provides methods to interact with SystemLogsService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type SystemLogsService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewSystemLogsService() *SystemLogsService {
	return &SystemLogsService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *SystemLogsService) WithConfigManager(manager *configmanager.ConfigManager) *SystemLogsService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *SystemLogsService) WithHook(hook hooks.Hook) *SystemLogsService {
	api.hook = hook
	return api
}

func (api *SystemLogsService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetSystemLogs()
}

func (api *SystemLogsService) getHook() hooks.Hook {
	return api.hook
}

func (api *SystemLogsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *SystemLogsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *SystemLogsService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the System Logs
func (api *SystemLogsService) GetSystemLogs(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[SystemLogList], *shared.SaladCloudSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/system-logs").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SystemLogList, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[SystemLogList](resp), nil
}
