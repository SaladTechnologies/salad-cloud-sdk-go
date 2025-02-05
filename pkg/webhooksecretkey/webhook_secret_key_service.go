package webhooksecretkey

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"time"
)

type WebhookSecretKeyService struct {
	manager *configmanager.ConfigManager
}

func NewWebhookSecretKeyService() *WebhookSecretKeyService {
	return &WebhookSecretKeyService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

func (api *WebhookSecretKeyService) WithConfigManager(manager *configmanager.ConfigManager) *WebhookSecretKeyService {
	api.manager = manager
	return api
}

func (api *WebhookSecretKeyService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetWebhookSecretKey()
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
func (api *WebhookSecretKeyService) GetWebhookSecretKey(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[WebhookSecretKey], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/webhook-secret-key").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebhookSecretKey](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[WebhookSecretKey](err)
	}

	return shared.NewSaladCloudSdkResponse[WebhookSecretKey](resp), nil
}

// Updates the webhook secret key
func (api *WebhookSecretKeyService) UpdateWebhookSecretKey(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[WebhookSecretKey], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/webhook-secret-key").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebhookSecretKey](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[WebhookSecretKey](err)
	}

	return shared.NewSaladCloudSdkResponse[WebhookSecretKey](resp), nil
}
