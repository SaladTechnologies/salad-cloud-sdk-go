package queues

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

// QueuesService provides methods to interact with QueuesService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type QueuesService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewQueuesService() *QueuesService {
	return &QueuesService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *QueuesService) WithConfigManager(manager *configmanager.ConfigManager) *QueuesService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *QueuesService) WithHook(hook hooks.Hook) *QueuesService {
	api.hook = hook
	return api
}

func (api *QueuesService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetQueues()
}

func (api *QueuesService) getHook() hooks.Hook {
	return api.hook
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
func (api *QueuesService) ListQueues(ctx context.Context, organizationName string, projectName string) (*shared.SaladCloudSdkResponse[QueueCollection], *shared.SaladCloudSdkError[[]byte]) {
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

	client := restClient.NewRestClient[QueueCollection, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueCollection](resp), nil
}

// Creates a new queue in the given project.
func (api *QueuesService) CreateQueue(ctx context.Context, organizationName string, projectName string, queuePrototype QueuePrototype) (*shared.SaladCloudSdkResponse[Queue], *shared.SaladCloudSdkError[[]byte]) {
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

	client := restClient.NewRestClient[Queue, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[Queue](resp), nil
}

// Gets an existing queue in the given project.
func (api *QueuesService) GetQueue(ctx context.Context, organizationName string, projectName string, queueName string) (*shared.SaladCloudSdkResponse[Queue], *shared.SaladCloudSdkError[[]byte]) {
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

	client := restClient.NewRestClient[Queue, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[Queue](resp), nil
}

// Updates an existing queue in the given project.
func (api *QueuesService) UpdateQueue(ctx context.Context, organizationName string, projectName string, queueName string, queuePatch QueuePatch) (*shared.SaladCloudSdkResponse[Queue], *shared.SaladCloudSdkError[[]byte]) {
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

	client := restClient.NewRestClient[Queue, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[Queue](resp), nil
}

// Deletes an existing queue in the given project.
func (api *QueuesService) DeleteQueue(ctx context.Context, organizationName string, projectName string, queueName string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError[[]byte]) {
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

	client := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}

// Gets the list of jobs in a queue
func (api *QueuesService) ListQueueJobs(ctx context.Context, organizationName string, projectName string, queueName string, params ListQueueJobsRequestParams) (*shared.SaladCloudSdkResponse[QueueJobCollection], *shared.SaladCloudSdkError[[]byte]) {
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

	client := restClient.NewRestClient[QueueJobCollection, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueJobCollection](resp), nil
}

// Creates a new job
func (api *QueuesService) CreateQueueJob(ctx context.Context, organizationName string, projectName string, queueName string, queueJobPrototype QueueJobPrototype) (*shared.SaladCloudSdkResponse[QueueJob], *shared.SaladCloudSdkError[[]byte]) {
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

	client := restClient.NewRestClient[QueueJob, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueJob](resp), nil
}

// Gets a job in a queue
func (api *QueuesService) GetQueueJob(ctx context.Context, organizationName string, projectName string, queueName string, queueJobId string) (*shared.SaladCloudSdkResponse[QueueJob], *shared.SaladCloudSdkError[[]byte]) {
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

	client := restClient.NewRestClient[QueueJob, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[QueueJob](resp), nil
}

// Cancels a job in a queue
func (api *QueuesService) DeleteQueueJob(ctx context.Context, organizationName string, projectName string, queueName string, queueJobId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError[[]byte]) {
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

	client := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[[]byte](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}
