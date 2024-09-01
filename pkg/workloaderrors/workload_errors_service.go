package workloaderrors

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

type WorkloadErrorsService struct {
	manager *configmanager.ConfigManager
}

func NewWorkloadErrorsService(manager *configmanager.ConfigManager) *WorkloadErrorsService {
	return &WorkloadErrorsService{
		manager: manager,
	}
}

func (api *WorkloadErrorsService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetWorkloadErrors()
}

func (api *WorkloadErrorsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *WorkloadErrorsService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the workload errors
func (api *WorkloadErrorsService) GetWorkloadErrors(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[WorkloadErrorList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[WorkloadErrorList](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/errors", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[WorkloadErrorList](err)
	}

	return shared.NewSaladCloudSdkResponse[WorkloadErrorList](resp), nil
}
