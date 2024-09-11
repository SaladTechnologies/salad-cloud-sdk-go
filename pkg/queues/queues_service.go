package queues

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

type QueuesService struct {
	manager *configmanager.ConfigManager
}

func NewQueuesService(manager *configmanager.ConfigManager) *QueuesService {
	return &QueuesService{
		manager: manager,
	}
}

func (api *QueuesService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetQueues()
}

func (api *QueuesService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *QueuesService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the list of queues in the given project.
func (api *QueuesService) ListQueues(ctx context.Context, organizationName string, projectName string) (*shared.SaladCloudSdkResponse[QueueList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[QueueList](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/projects/{project_name}/queues", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[QueueList](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueList](resp), nil
}

// Creates a new queue in the given project.
func (api *QueuesService) CreateQueue(ctx context.Context, organizationName string, projectName string, createQueue CreateQueue) (*shared.SaladCloudSdkResponse[Queue], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Queue](config)

	request := httptransport.NewRequest(ctx, "POST", "/organizations/{organization_name}/projects/{project_name}/queues", config)
	request.Headers["Content-Type"] = "application/json"

	request.Body = createQueue

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[Queue](err)
	}

	return shared.NewSaladCloudSdkResponse[Queue](resp), nil
}

// Gets an existing queue in the given project.
func (api *QueuesService) GetQueue(ctx context.Context, organizationName string, projectName string, queueName string) (*shared.SaladCloudSdkResponse[Queue], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Queue](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("queue_name", queueName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[Queue](err)
	}

	return shared.NewSaladCloudSdkResponse[Queue](resp), nil
}

// Updates an existing queue in the given project.
func (api *QueuesService) UpdateQueue(ctx context.Context, organizationName string, projectName string, queueName string, updateQueue UpdateQueue) (*shared.SaladCloudSdkResponse[Queue], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Queue](config)

	request := httptransport.NewRequest(ctx, "PATCH", "/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}", config)
	request.Headers["Content-Type"] = "application/merge-patch+json"

	request.Body = updateQueue

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("queue_name", queueName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[Queue](err)
	}

	return shared.NewSaladCloudSdkResponse[Queue](resp), nil
}

// Deletes an existing queue in the given project.
func (api *QueuesService) DeleteQueue(ctx context.Context, organizationName string, projectName string, queueName string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("queue_name", queueName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Gets the list of jobs in a queue
func (api *QueuesService) ListQueueJobs(ctx context.Context, organizationName string, projectName string, queueName string, params ListQueueJobsRequestParams) (*shared.SaladCloudSdkResponse[QueueJobList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[QueueJobList](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs", config)

	request.Options = params

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("queue_name", queueName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[QueueJobList](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueJobList](resp), nil
}

// Creates a new job
func (api *QueuesService) CreateQueueJob(ctx context.Context, organizationName string, projectName string, queueName string, createQueueJob CreateQueueJob) (*shared.SaladCloudSdkResponse[QueueJob], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[QueueJob](config)

	request := httptransport.NewRequest(ctx, "POST", "/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs", config)
	request.Headers["Content-Type"] = "application/json"

	request.Body = createQueueJob

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("queue_name", queueName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[QueueJob](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueJob](resp), nil
}

// Gets a job in a queue
func (api *QueuesService) GetQueueJob(ctx context.Context, organizationName string, projectName string, queueName string, queueJobId string) (*shared.SaladCloudSdkResponse[QueueJob], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[QueueJob](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id}", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("queue_name", queueName)
	request.SetPathParam("queue_job_id", queueJobId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[QueueJob](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueJob](resp), nil
}

// Cancels a job in a queue
func (api *QueuesService) DeleteQueueJob(ctx context.Context, organizationName string, projectName string, queueName string, queueJobId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id}", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("project_name", projectName)
	request.SetPathParam("queue_name", queueName)
	request.SetPathParam("queue_job_id", queueJobId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}
