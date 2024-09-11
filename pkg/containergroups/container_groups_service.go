package containergroups

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

type ContainerGroupsService struct {
	manager *configmanager.ConfigManager
}

func NewContainerGroupsService(manager *configmanager.ConfigManager) *ContainerGroupsService {
	return &ContainerGroupsService{
		manager: manager,
	}
}

func (api *ContainerGroupsService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetContainerGroups()
}

func (api *ContainerGroupsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *ContainerGroupsService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the list of container groups
func (api *ContainerGroupsService) ListContainerGroups(ctx context.Context, organizationName string, projectName string) (*shared.SaladCloudSdkResponse[ContainerGroupList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ContainerGroupList](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/projects/{project_name}/containers", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[ContainerGroupList](err)
	}

	return shared.NewSaladCloudSdkResponse[ContainerGroupList](resp), nil
}

// Creates a new container group
func (api *ContainerGroupsService) CreateContainerGroup(ctx context.Context, organizationName string, projectName string, createContainerGroup CreateContainerGroup) (*shared.SaladCloudSdkResponse[shared.ContainerGroup], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[shared.ContainerGroup](config)

	request := httptransport.NewRequest(ctx, "POST", "/organizations/{organization_name}/projects/{project_name}/containers", config)
	request.Headers["Content-Type"] = "application/json"

	request.Body = createContainerGroup

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[shared.ContainerGroup](err)
	}

	return shared.NewSaladCloudSdkResponse[shared.ContainerGroup](resp), nil
}

// Gets a container group
func (api *ContainerGroupsService) GetContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[shared.ContainerGroup], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[shared.ContainerGroup](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[shared.ContainerGroup](err)
	}

	return shared.NewSaladCloudSdkResponse[shared.ContainerGroup](resp), nil
}

// Updates a container group
func (api *ContainerGroupsService) UpdateContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupName string, updateContainerGroup UpdateContainerGroup) (*shared.SaladCloudSdkResponse[shared.ContainerGroup], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[shared.ContainerGroup](config)

	request := httptransport.NewRequest(ctx, "PATCH", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}", config)
	request.Headers["Content-Type"] = "application/merge-patch+json"

	request.Body = updateContainerGroup

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[shared.ContainerGroup](err)
	}

	return shared.NewSaladCloudSdkResponse[shared.ContainerGroup](resp), nil
}

// Deletes a container group
func (api *ContainerGroupsService) DeleteContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Starts a container group
func (api *ContainerGroupsService) StartContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "POST", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/start", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Stops a container group
func (api *ContainerGroupsService) StopContainerGroup(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "POST", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/stop", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Retrieves a list of container group instances
func (api *ContainerGroupsService) ListContainerGroupInstances(ctx context.Context, organizationName string, projectName string, containerGroupName string) (*shared.SaladCloudSdkResponse[ContainerGroupInstances], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ContainerGroupInstances](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[ContainerGroupInstances](err)
	}

	return shared.NewSaladCloudSdkResponse[ContainerGroupInstances](resp), nil
}

// Retrieves the details of a single instance within a container group by instance ID
func (api *ContainerGroupsService) GetContainerGroupInstance(ctx context.Context, organizationName string, projectName string, containerGroupName string, containerGroupInstanceId string) (*shared.SaladCloudSdkResponse[ContainerGroupInstance], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ContainerGroupInstance](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)
	request.SetPathParam("container_group_instance_id", containerGroupInstanceId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[ContainerGroupInstance](err)
	}

	return shared.NewSaladCloudSdkResponse[ContainerGroupInstance](resp), nil
}

// Remove a node from a workload and reallocate the workload to a different node
func (api *ContainerGroupsService) ReallocateContainerGroupInstance(ctx context.Context, organizationName string, projectName string, containerGroupName string, containerGroupInstanceId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "POST", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/reallocate", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)
	request.SetPathParam("container_group_instance_id", containerGroupInstanceId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Stops a container, destroys it, creates a new one without requiring the image to be downloaded again on a different node
func (api *ContainerGroupsService) RecreateContainerGroupInstance(ctx context.Context, organizationName string, projectName string, containerGroupName string, containerGroupInstanceId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "POST", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/recreate", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)
	request.SetPathParam("container_group_instance_id", containerGroupInstanceId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Restarts a workload on a node without reallocating it
func (api *ContainerGroupsService) RestartContainerGroupInstance(ctx context.Context, organizationName string, projectName string, containerGroupName string, containerGroupInstanceId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "POST", "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/restart", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("container_group_name", containerGroupName)
	request.SetPathParam("container_group_instance_id", containerGroupInstanceId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}
