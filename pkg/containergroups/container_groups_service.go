package containergroups

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"time"
)

type ContainerGroupsService struct {
	manager *configmanager.ConfigManager
}

func NewContainerGroupsService() *ContainerGroupsService {
	return &ContainerGroupsService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

func (api *ContainerGroupsService) WithConfigManager(manager *configmanager.ConfigManager) *ContainerGroupsService {
	api.manager = manager
	return api
}

func (api *ContainerGroupsService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetContainerGroups()
}

func (api *ContainerGroupsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *ContainerGroupsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *ContainerGroupsService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the list of container groups
func (api *ContainerGroupsService) ListContainerGroups(ctx context.Context, organizationName string, projectName string) (*shared.SaladCloudSdkResponse[ContainerGroupCollection], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ContainerGroupCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[ContainerGroupCollection](err)
	}

	return shared.NewSaladCloudSdkResponse[ContainerGroupCollection](resp), nil
}

// Creates a new container group
func (api *ContainerGroupsService) CreateContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupCreationRequest ContainerGroupCreationRequest) (*shared.SaladCloudSdkResponse[shared.ContainerGroup], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers").
		WithConfig(config).
		WithBody(containerGroupCreationRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[shared.ContainerGroup](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[shared.ContainerGroup](err)
	}

	return shared.NewSaladCloudSdkResponse[shared.ContainerGroup](resp), nil
}

// Gets a container group
func (api *ContainerGroupsService) GetContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[shared.ContainerGroup], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[shared.ContainerGroup](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[shared.ContainerGroup](err)
	}

	return shared.NewSaladCloudSdkResponse[shared.ContainerGroup](resp), nil
}

// Updates a container group
func (api *ContainerGroupsService) UpdateContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupName string, containerGroupPatch ContainerGroupPatch) (*shared.SaladCloudSdkResponse[shared.ContainerGroup], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PATCH").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}").
		WithConfig(config).
		WithBody(containerGroupPatch).
		AddHeader("CONTENT-TYPE", "application/merge-patch+json").
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[shared.ContainerGroup](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[shared.ContainerGroup](err)
	}

	return shared.NewSaladCloudSdkResponse[shared.ContainerGroup](resp), nil
}

// Deletes a container group
func (api *ContainerGroupsService) DeleteContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Starts a container group
func (api *ContainerGroupsService) StartContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/start").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Stops a container group
func (api *ContainerGroupsService) StopContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/stop").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Gets the list of container group instances
func (api *ContainerGroupsService) ListContainerGroupInstances(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[ContainerGroupInstanceCollection], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ContainerGroupInstanceCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[ContainerGroupInstanceCollection](err)
	}

	return shared.NewSaladCloudSdkResponse[ContainerGroupInstanceCollection](resp), nil
}

// Gets a container group instance
func (api *ContainerGroupsService) GetContainerGroupInstance(ctx context.Context, organizationName string, projectName string, containerGroupName string, containerGroupInstanceId string) (*shared.SaladCloudSdkResponse[ContainerGroupInstance], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		AddPathParam("container_group_instance_id", containerGroupInstanceId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ContainerGroupInstance](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[ContainerGroupInstance](err)
	}

	return shared.NewSaladCloudSdkResponse[ContainerGroupInstance](resp), nil
}

// Updates a container group instance
func (api *ContainerGroupsService) UpdateContainerGroupInstance(ctx context.Context, organizationName string, projectName string, containerGroupName string, containerGroupInstanceId string, containerGroupInstancePatch ContainerGroupInstancePatch) (*shared.SaladCloudSdkResponse[ContainerGroupInstance], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PATCH").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}").
		WithConfig(config).
		WithBody(containerGroupInstancePatch).
		AddHeader("CONTENT-TYPE", "application/merge-patch+json").
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		AddPathParam("container_group_instance_id", containerGroupInstanceId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ContainerGroupInstance](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[ContainerGroupInstance](err)
	}

	return shared.NewSaladCloudSdkResponse[ContainerGroupInstance](resp), nil
}

// Reallocates a container group instance to run on a different Salad Node
func (api *ContainerGroupsService) ReallocateContainerGroupInstance(ctx context.Context, organizationName string, projectName string, containerGroupName string, containerGroupInstanceId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/reallocate").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		AddPathParam("container_group_instance_id", containerGroupInstanceId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Stops a container, destroys it, and starts a new one without requiring the image to be downloaded again on a new Salad Node
func (api *ContainerGroupsService) RecreateContainerGroupInstance(ctx context.Context, organizationName string, projectName string, containerGroupName string, containerGroupInstanceId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/recreate").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		AddPathParam("container_group_instance_id", containerGroupInstanceId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Stops a container and restarts it on the same Salad Node
func (api *ContainerGroupsService) RestartContainerGroupInstance(ctx context.Context, organizationName string, projectName string, containerGroupName string, containerGroupInstanceId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/restart").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("container_group_name", containerGroupName).
		AddPathParam("container_group_instance_id", containerGroupInstanceId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}
