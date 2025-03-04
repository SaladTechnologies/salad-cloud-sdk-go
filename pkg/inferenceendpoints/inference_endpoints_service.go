package inferenceendpoints

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"time"
)

type InferenceEndpointsService struct {
	manager *configmanager.ConfigManager
}

func NewInferenceEndpointsService() *InferenceEndpointsService {
	return &InferenceEndpointsService{
		manager: configmanager.NewConfigManager(saladcloudsdkconfig.Config{}),
	}
}

func (api *InferenceEndpointsService) WithConfigManager(manager *configmanager.ConfigManager) *InferenceEndpointsService {
	api.manager = manager
	return api
}

func (api *InferenceEndpointsService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetInferenceEndpoints()
}

func (api *InferenceEndpointsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *InferenceEndpointsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *InferenceEndpointsService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Lists inference endpoints.
func (api *InferenceEndpointsService) ListInferenceEndpoints(ctx context.Context, organizationName string, params ListInferenceEndpointsRequestParams) (*shared.SaladCloudSdkResponse[InferenceEndpointList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/inference-endpoints").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InferenceEndpointList](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[InferenceEndpointList](err)
	}

	return shared.NewSaladCloudSdkResponse[InferenceEndpointList](resp), nil
}

// Gets an inference endpoint.
func (api *InferenceEndpointsService) GetInferenceEndpoint(ctx context.Context, organizationName string, inferenceEndpointName string) (*shared.SaladCloudSdkResponse[InferenceEndpoint], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("inference_endpoint_name", inferenceEndpointName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InferenceEndpoint](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[InferenceEndpoint](err)
	}

	return shared.NewSaladCloudSdkResponse[InferenceEndpoint](resp), nil
}

// Lists inference endpoint jobs.
func (api *InferenceEndpointsService) ListInferenceEndpointJobs(ctx context.Context, organizationName string, inferenceEndpointName string, params ListInferenceEndpointJobsRequestParams) (*shared.SaladCloudSdkResponse[InferenceEndpointJobList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("inference_endpoint_name", inferenceEndpointName).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InferenceEndpointJobList](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[InferenceEndpointJobList](err)
	}

	return shared.NewSaladCloudSdkResponse[InferenceEndpointJobList](resp), nil
}

// Creates a new inference endpoint job.
func (api *InferenceEndpointsService) CreateInferenceEndpointJob(ctx context.Context, organizationName string, inferenceEndpointName string, createInferenceEndpointJob CreateInferenceEndpointJob) (*shared.SaladCloudSdkResponse[InferenceEndpointJob], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs").
		WithConfig(config).
		WithBody(createInferenceEndpointJob).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("organization_name", organizationName).
		AddPathParam("inference_endpoint_name", inferenceEndpointName).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InferenceEndpointJob](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[InferenceEndpointJob](err)
	}

	return shared.NewSaladCloudSdkResponse[InferenceEndpointJob](resp), nil
}

// Gets an inference endpoint job.
func (api *InferenceEndpointsService) GetInferenceEndpointJob(ctx context.Context, organizationName string, inferenceEndpointName string, inferenceEndpointJobId string) (*shared.SaladCloudSdkResponse[InferenceEndpointJob], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs/{inference_endpoint_job_id}").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("inference_endpoint_name", inferenceEndpointName).
		AddPathParam("inference_endpoint_job_id", inferenceEndpointJobId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InferenceEndpointJob](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[InferenceEndpointJob](err)
	}

	return shared.NewSaladCloudSdkResponse[InferenceEndpointJob](resp), nil
}

// Cancels an inference endpoint job.
func (api *InferenceEndpointsService) CancelInferenceEndpointJob(ctx context.Context, organizationName string, inferenceEndpointName string, inferenceEndpointJobId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs/{inference_endpoint_job_id}").
		WithConfig(config).
		AddPathParam("organization_name", organizationName).
		AddPathParam("inference_endpoint_name", inferenceEndpointName).
		AddPathParam("inference_endpoint_job_id", inferenceEndpointJobId).
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
