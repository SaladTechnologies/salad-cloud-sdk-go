package logs

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

// LogsService provides methods to interact with LogsService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type LogsService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewLogsService() *LogsService {
	return &LogsService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *LogsService) WithConfigManager(manager *configmanager.ConfigManager) *LogsService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *LogsService) WithHook(hook hooks.Hook) *LogsService {
	api.hook = hook
	return api
}

func (api *LogsService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetLogs()
}

func (api *LogsService) getHook() hooks.Hook {
	return api.hook
}

func (api *LogsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *LogsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *LogsService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Retrieve a collection of _log entries_ for the _organization_ identified by `{organization_name}` matching the log query.
func (api *LogsService) QueryLogEntries(ctx context.Context, organizationName string, logEntryQuery LogEntryQuery) (*shared.SaladCloudSdkResponse[LogEntryCollection], *shared.SaladCloudSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/log-entries").
		WithConfig(config).
		WithBody(logEntryQuery).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("organization_name", organizationName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[LogEntryCollection, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[LogEntryCollection](resp), nil
}
