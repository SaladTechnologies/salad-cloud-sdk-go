package webhooksecretkey

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

type WebhookSecretKeyService struct {
	manager *configmanager.ConfigManager
}

func NewWebhookSecretKeyService(manager *configmanager.ConfigManager) *WebhookSecretKeyService {
	return &WebhookSecretKeyService{
		manager: manager,
	}
}

func (api *WebhookSecretKeyService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetWebhookSecretKey()
}

func (api *WebhookSecretKeyService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *WebhookSecretKeyService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the webhook secret key
func (api *WebhookSecretKeyService) GetWebhookSecretKey(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[WebhookSecretKey], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[WebhookSecretKey](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/webhook-secret-key", config)

	request.SetPathParam("organization_name", organizationName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[WebhookSecretKey](err)
	}

	return shared.NewSaladCloudSdkResponse[WebhookSecretKey](resp), nil
}

// Updates the webhook secret key
func (api *WebhookSecretKeyService) UpdateWebhookSecretKey(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[WebhookSecretKey], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[WebhookSecretKey](config)

	request := httptransport.NewRequest(ctx, "POST", "/organizations/{organization_name}/webhook-secret-key", config)

	request.SetPathParam("organization_name", organizationName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[WebhookSecretKey](err)
	}

	return shared.NewSaladCloudSdkResponse[WebhookSecretKey](resp), nil
}
