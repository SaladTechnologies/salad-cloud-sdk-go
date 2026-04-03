package webhooksecretkey

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

// WebhookSecretKeyService provides methods to interact with WebhookSecretKeyService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type WebhookSecretKeyService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewWebhookSecretKeyService() *WebhookSecretKeyService {
	return &WebhookSecretKeyService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *WebhookSecretKeyService) WithConfigManager(manager *configmanager.ConfigManager) *WebhookSecretKeyService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *WebhookSecretKeyService) WithHook(hook hooks.Hook) *WebhookSecretKeyService {
	api.hook = hook
	return api
}

func (api *WebhookSecretKeyService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetWebhookSecretKey()
}

func (api *WebhookSecretKeyService) getHook() hooks.Hook {
	return api.hook
}

func (api *WebhookSecretKeyService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *WebhookSecretKeyService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *WebhookSecretKeyService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the webhook secret key
func (api *WebhookSecretKeyService) GetWebhookSecretKey(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[WebhookSecretKey], *shared.SaladCloudSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/webhook-secret-key").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebhookSecretKey, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[WebhookSecretKey](resp), nil
}

// Updates the webhook secret key
func (api *WebhookSecretKeyService) UpdateWebhookSecretKey(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[WebhookSecretKey], *shared.SaladCloudSdkError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/webhook-secret-key").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebhookSecretKey, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[WebhookSecretKey](resp), nil
}
