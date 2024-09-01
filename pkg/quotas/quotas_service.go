package quotas

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

type QuotasService struct {
	manager *configmanager.ConfigManager
}

func NewQuotasService(manager *configmanager.ConfigManager) *QuotasService {
	return &QuotasService{
		manager: manager,
	}
}

func (api *QuotasService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetQuotas()
}

func (api *QuotasService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *QuotasService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the organization quotas
func (api *QuotasService) GetQuotas(ctx context.Context, organizationName string) (*shared.SaladCloudSdkResponse[Quotas], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Quotas](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/quotas", config)

	request.SetPathParam("organization_name", organizationName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[Quotas](err)
	}

	return shared.NewSaladCloudSdkResponse[Quotas](resp), nil
}
