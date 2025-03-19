package queues

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"time"
)

type QueuesService struct {
	manager *configmanager.ConfigManager
}

func NewQueuesService() *QueuesService {
	return &QueuesService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

func (api *QueuesService) WithConfigManager(manager *configmanager.ConfigManager) *QueuesService {
	api.manager = manager
	return api
}

func (api *QueuesService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetQueues()
}

func (api *QueuesService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *QueuesService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *QueuesService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the list of queues in the given project.
func (api *QueuesService) ListQueues(ctx context.Context, organizationName string, projectName string) (*shared.SaladCloudSdkResponse[QueueCollection], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/projects/{project_name}/queues").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[QueueCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[QueueCollection](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueCollection](resp), nil
}

// Creates a new queue in the given project.
func (api *QueuesService) CreateQueue(ctx context.Context, organizationName string, projectName string, queuePrototype QueuePrototype) (*shared.SaladCloudSdkResponse[Queue], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/projects/{project_name}/queues").
		WithConfig(config).
		WithBody(queuePrototype).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Queue](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[Queue](err)
	}

	return shared.NewSaladCloudSdkResponse[Queue](resp), nil
}

// Gets an existing queue in the given project.
func (api *QueuesService) GetQueue(ctx context.Context, organizationName string, projectName string, queueName string) (*shared.SaladCloudSdkResponse[Queue], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("queue_name", queueName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Queue](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[Queue](err)
	}

	return shared.NewSaladCloudSdkResponse[Queue](resp), nil
}

// Updates an existing queue in the given project.
func (api *QueuesService) UpdateQueue(ctx context.Context, organizationName string, projectName string, queueName string, queuePatch QueuePatch) (*shared.SaladCloudSdkResponse[Queue], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PATCH").
		WithPath("/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}").
		WithConfig(config).
		WithBody(queuePatch).
		AddHeader("CONTENT-TYPE", "application/merge-patch+json").
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("queue_name", queueName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Queue](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[Queue](err)
	}

	return shared.NewSaladCloudSdkResponse[Queue](resp), nil
}

// Deletes an existing queue in the given project.
func (api *QueuesService) DeleteQueue(ctx context.Context, organizationName string, projectName string, queueName string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("queue_name", queueName).
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

// Gets the list of jobs in a queue
func (api *QueuesService) ListQueueJobs(ctx context.Context, organizationName string, projectName string, queueName string, params ListQueueJobsRequestParams) (*shared.SaladCloudSdkResponse[QueueJobCollection], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("queue_name", queueName).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[QueueJobCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[QueueJobCollection](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueJobCollection](resp), nil
}

// Creates a new job
func (api *QueuesService) CreateQueueJob(ctx context.Context, organizationName string, projectName string, queueName string, queueJobPrototype QueueJobPrototype) (*shared.SaladCloudSdkResponse[QueueJob], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs").
		WithConfig(config).
		WithBody(queueJobPrototype).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("queue_name", queueName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[QueueJob](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[QueueJob](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueJob](resp), nil
}

// Gets a job in a queue
func (api *QueuesService) GetQueueJob(ctx context.Context, organizationName string, projectName string, queueName string, queueJobId string) (*shared.SaladCloudSdkResponse[QueueJob], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id}").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("queue_name", queueName).
		AddPathParam("queue_job_id", queueJobId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[QueueJob](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[QueueJob](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueJob](resp), nil
}

// Cancels a job in a queue
func (api *QueuesService) DeleteQueueJob(ctx context.Context, organizationName string, projectName string, queueName string, queueJobId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id}").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("project_name", projectName).
		AddPathParam("queue_name", queueName).
		AddPathParam("queue_job_id", queueJobId).
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
