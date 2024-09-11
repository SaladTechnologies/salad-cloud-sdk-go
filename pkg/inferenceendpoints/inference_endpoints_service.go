package inferenceendpoints

import (
	"context"
	restClient "github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/configmanager"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

type InferenceEndpointsService struct {
	manager *configmanager.ConfigManager
}

func NewInferenceEndpointsService(manager *configmanager.ConfigManager) *InferenceEndpointsService {
	return &InferenceEndpointsService{
		manager: manager,
	}
}

func (api *InferenceEndpointsService) getConfig() *saladcloudsdkconfig.Config {
	return api.manager.GetInferenceEndpoints()
}

func (api *InferenceEndpointsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *InferenceEndpointsService) SetApiKey(apiKey string) {
	config := api.getConfig()
	config.SetApiKey(apiKey)
}

// Gets the list of all inference endpoints
func (api *InferenceEndpointsService) ListInferenceEndpoints(ctx context.Context, organizationName string, params ListInferenceEndpointsRequestParams) (*shared.SaladCloudSdkResponse[InferenceEndpointsList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[InferenceEndpointsList](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/inference-endpoints", config)

	request.Options = params

	request.SetPathParam("organization_name", organizationName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[InferenceEndpointsList](err)
	}

	return shared.NewSaladCloudSdkResponse[InferenceEndpointsList](resp), nil
}

// Gets an inference endpoint
func (api *InferenceEndpointsService) GetInferenceEndpoint(ctx context.Context, organizationName string, inferenceEndpointName string) (*shared.SaladCloudSdkResponse[InferenceEndpoint], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[InferenceEndpoint](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("inference_endpoint_name", inferenceEndpointName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[InferenceEndpoint](err)
	}

	return shared.NewSaladCloudSdkResponse[InferenceEndpoint](resp), nil
}

// Retrieves a list of an inference endpoint jobs
func (api *InferenceEndpointsService) GetInferenceEndpointJobs(ctx context.Context, organizationName string, inferenceEndpointName string, params GetInferenceEndpointJobsRequestParams) (*shared.SaladCloudSdkResponse[InferenceEndpointJobList], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[InferenceEndpointJobList](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs", config)

	request.Options = params

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("inference_endpoint_name", inferenceEndpointName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[InferenceEndpointJobList](err)
	}

	return shared.NewSaladCloudSdkResponse[InferenceEndpointJobList](resp), nil
}

// Creates a new job
func (api *InferenceEndpointsService) CreateInferenceEndpointJob(ctx context.Context, organizationName string, inferenceEndpointName string, createInferenceEndpointJob CreateInferenceEndpointJob) (*shared.SaladCloudSdkResponse[InferenceEndpointJob], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[InferenceEndpointJob](config)

	request := httptransport.NewRequest(ctx, "POST", "/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs", config)
	request.Headers["Content-Type"] = "application/json"

	request.Body = createInferenceEndpointJob

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("inference_endpoint_name", inferenceEndpointName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[InferenceEndpointJob](err)
	}

	return shared.NewSaladCloudSdkResponse[InferenceEndpointJob](resp), nil
}

// Retrieves a job in an inference endpoint
func (api *InferenceEndpointsService) GetInferenceEndpointJob(ctx context.Context, organizationName string, inferenceEndpointName string, inferenceEndpointJobId string) (*shared.SaladCloudSdkResponse[InferenceEndpointJob], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[InferenceEndpointJob](config)

	request := httptransport.NewRequest(ctx, "GET", "/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs/{inference_endpoint_job_id}", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("inference_endpoint_name", inferenceEndpointName)
	request.SetPathParam("inference_endpoint_job_id", inferenceEndpointJobId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[InferenceEndpointJob](err)
	}

	return shared.NewSaladCloudSdkResponse[InferenceEndpointJob](resp), nil
}

// Deletes an inference endpoint job
func (api *InferenceEndpointsService) DeleteInferenceEndpointJob(ctx context.Context, organizationName string, inferenceEndpointName string, inferenceEndpointJobId string) (*shared.SaladCloudSdkResponse[any], *shared.SaladCloudSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs/{inference_endpoint_job_id}", config)

	request.SetPathParam("organization_name", organizationName)
	request.SetPathParam("inference_endpoint_name", inferenceEndpointName)
	request.SetPathParam("inference_endpoint_job_id", inferenceEndpointJobId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSaladCloudSdkError[any](err)
	}

	return shared.NewSaladCloudSdkResponse[any](resp), nil
}
