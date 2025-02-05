package workloaderrors

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"time"
)

type WorkloadErrorsService struct {
	manager *configmanager.ConfigManager
}

func NewWorkloadErrorsService() *WorkloadErrorsService {
	return &WorkloadErrorsService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

func (api *WorkloadErrorsService) WithConfigManager(manager *configmanager.ConfigManager) *WorkloadErrorsService {
	api.manager = manager
	return api
}

func (api *WorkloadErrorsService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetWorkloadErrors()
}

func (api *WorkloadErrorsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *WorkloadErrorsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *WorkloadErrorsService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the workload errors. This has been deprecated and will be replaced by the new System Logs endpoint. See `/system-logs`.
func (api *WorkloadErrorsService) GetWorkloadErrors(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[WorkloadErrorList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/errors").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WorkloadErrorList](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[WorkloadErrorList](err)
	}

	return shared.NewSaladCloudSdkResponse[WorkloadErrorList](resp), nil
}
