package quotas

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

// QuotasService provides methods to interact with QuotasService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type QuotasService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewQuotasService() *QuotasService {
	return &QuotasService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *QuotasService) WithConfigManager(manager *configmanager.ConfigManager) *QuotasService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *QuotasService) WithHook(hook hooks.Hook) *QuotasService {
	api.hook = hook
	return api
}

func (api *QuotasService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetQuotas()
}

func (api *QuotasService) getHook() hooks.Hook {
	return api.hook
}

func (api *QuotasService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *QuotasService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *QuotasService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the organization quotas
func (api *QuotasService) GetQuotas(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[Quotas], *shared.SaladCloudSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/quotas").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Quotas, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[Quotas](resp), nil
}
