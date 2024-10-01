# SaladCloudSdk Go SDK 0.9.0-alpha.4

Welcome to the SaladCloudSdk SDK documentation. This guide will help you get started with integrating and using the SaladCloudSdk SDK in your project.

## Versions

- API version: `0.9.0-alpha.4`
- SDK version: `0.9.0-alpha.4`

## About the API

The SaladCloud REST API. Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

## Table of Contents

- [Setup & Configuration](#setup--configuration)
  - [Supported Language Versions](#supported-language-versions)
  - [Installation](#installation)
- [Authentication](#authentication)
  - [API Key Authentication](#api-key-authentication)
- [Services](#services)
  - [Response Wrappers](#response-wrappers)
- [Models](#models)
- [License](#license)

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Authentication

### API Key Authentication

The salad-cloud-sdk API uses API keys as a form of authentication. An API key is a unique identifier used to authenticate a user, developer, or a program that is calling the API.

#### Setting the API key

When you initialize the SDK, you can set the API key as follows:

```go
import (
    "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
    "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  )

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("YOUR-TOKEN")

sdk := saladcloudsdk.NewSaladCloudSdk(config)
```

If you need to set or update the API key after initializing the SDK, you can use:

```go
import (
    "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
    "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  )

config := saladcloudsdkconfig.NewConfig()

sdk := saladcloudsdk.NewSaladCloudSdk(config)
sdk.SetApiKey("YOUR-TOKEN")
```

## Services

The SDK provides various services to interact with the API.

<details> 
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                                               |
| :--------------------------------------------------------------------------------- |
| [ContainerGroupsService](documentation/services/container_groups_service.md)       |
| [WorkloadErrorsService](documentation/services/workload_errors_service.md)         |
| [QueuesService](documentation/services/queues_service.md)                          |
| [QuotasService](documentation/services/quotas_service.md)                          |
| [InferenceEndpointsService](documentation/services/inference_endpoints_service.md) |
| [OrganizationDataService](documentation/services/organization_data_service.md)     |
| [WebhookSecretKeyService](documentation/services/webhook_secret_key_service.md)    |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `SaladCloudSdkResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                            | Description                                 |
| :------- | :------------------------------ | :------------------------------------------ |
| Data     | `T`                             | The body of the API response                |
| Metadata | `SaladCloudSdkResponseMetadata` | Status code and headers returned by the API |

#### `SaladCloudSdkError`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                            | Description                                 |
| :------- | :------------------------------ | :------------------------------------------ |
| Err      | `error`                         | The error that occurred                     |
| Body     | `T`                             | The body of the API response                |
| Metadata | `SaladCloudSdkResponseMetadata` | Status code and headers returned by the API |

#### `SaladCloudSdkResponseMetadata`

This struct is shared by both response wrappers and contains the following fields:

| Name       | Type                | Description                                      |
| :--------- | :------------------ | :----------------------------------------------- |
| Headers    | `map[string]string` | A map containing the headers returned by the API |
| StatusCode | `int`               | The status code returned by the API              |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details> 
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                                               | Description                                                              |
| :------------------------------------------------------------------------------------------------- | :----------------------------------------------------------------------- |
| [ContainerGroupList](documentation/models/container_group_list.md)                                 | Represents a list of container groups                                    |
| [CreateContainerGroup](documentation/models/create_container_group.md)                             | Represents a request to create a container group                         |
| [ContainerGroup](documentation/models/container_group.md)                                          | Represents a container group                                             |
| [UpdateContainerGroup](documentation/models/update_container_group.md)                             | Represents a request to update a container group                         |
| [ContainerGroupInstances](documentation/models/container_group_instances.md)                       | Represents a list of container group instances                           |
| [ContainerGroupInstance](documentation/models/container_group_instance.md)                         | Represents the details of a single container group instance              |
| [WorkloadErrorList](documentation/models/workload_error_list.md)                                   | Represents a list of workload errors                                     |
| [QueueList](documentation/models/queue_list.md)                                                    | Represents a list of queues                                              |
| [CreateQueue](documentation/models/create_queue.md)                                                | Represents a request to create a new queue.                              |
| [Queue](documentation/models/queue.md)                                                             | Represents a queue.                                                      |
| [UpdateQueue](documentation/models/update_queue.md)                                                | Represents a request to update an existing queue.                        |
| [QueueJobList](documentation/models/queue_job_list.md)                                             | Represents a list of queue jobs                                          |
| [CreateQueueJob](documentation/models/create_queue_job.md)                                         | Represents a request to create a queue job                               |
| [QueueJob](documentation/models/queue_job.md)                                                      | Represents a queue job                                                   |
| [Quotas](documentation/models/quotas.md)                                                           | Represents the organization quotas                                       |
| [InferenceEndpointsList](documentation/models/inference_endpoints_list.md)                         | Represents a list of inference endpoints                                 |
| [InferenceEndpoint](documentation/models/inference_endpoint.md)                                    | Represents an inference endpoint                                         |
| [InferenceEndpointJobList](documentation/models/inference_endpoint_job_list.md)                    | Represents a list of inference endpoint jobs                             |
| [CreateInferenceEndpointJob](documentation/models/create_inference_endpoint_job.md)                | Represents a request to create a inference endpoint job                  |
| [InferenceEndpointJob](documentation/models/inference_endpoint_job.md)                             | Represents a inference endpoint job                                      |
| [GpuClassesList](documentation/models/gpu_classes_list.md)                                         | Represents a list of GPU classes                                         |
| [WebhookSecretKey](documentation/models/webhook_secret_key.md)                                     | Represents a webhook secret key                                          |
| [Container](documentation/models/container.md)                                                     | Represents a container                                                   |
| [ContainerRestartPolicy](documentation/models/container_restart_policy.md)                         |                                                                          |
| [ContainerGroupState](documentation/models/container_group_state.md)                               | Represents a container group state                                       |
| [CountryCode](documentation/models/country_code.md)                                                |                                                                          |
| [ContainerGroupNetworking](documentation/models/container_group_networking.md)                     | Represents container group networking parameters                         |
| [ContainerGroupLivenessProbe](documentation/models/container_group_liveness_probe.md)              | Represents the container group liveness probe                            |
| [ContainerGroupReadinessProbe](documentation/models/container_group_readiness_probe.md)            | Represents the container group readiness probe                           |
| [ContainerGroupStartupProbe](documentation/models/container_group_startup_probe.md)                | Represents the container group startup probe                             |
| [ContainerGroupQueueConnection](documentation/models/container_group_queue_connection.md)          | Represents container group queue connection                              |
| [QueueAutoscaler](documentation/models/queue_autoscaler.md)                                        | Represents the autoscaling rules for a queue                             |
| [ContainerResourceRequirements](documentation/models/container_resource_requirements.md)           | Represents a container resource requirements                             |
| [ContainerGroupPriority](documentation/models/container_group_priority.md)                         |                                                                          |
| [ContainerGroupStatus](documentation/models/container_group_status.md)                             |                                                                          |
| [ContainerGroupInstanceStatusCount](documentation/models/container_group_instance_status_count.md) | Represents a container group instance status count                       |
| [ContainerNetworkingProtocol](documentation/models/container_networking_protocol.md)               |                                                                          |
| [ContainerGroupProbeTcp](documentation/models/container_group_probe_tcp.md)                        |                                                                          |
| [ContainerGroupProbeHttp](documentation/models/container_group_probe_http.md)                      |                                                                          |
| [ContainerGroupProbeGrpc](documentation/models/container_group_probe_grpc.md)                      |                                                                          |
| [ContainerGroupProbeExec](documentation/models/container_group_probe_exec.md)                      |                                                                          |
| [ContainerProbeHttpScheme](documentation/models/container_probe_http_scheme.md)                    |                                                                          |
| [ContainerGroupProbeHttpHeaders2](documentation/models/container_group_probe_http_headers_2.md)    |                                                                          |
| [CreateContainer](documentation/models/create_container.md)                                        | Represents a container                                                   |
| [CreateContainerGroupNetworking](documentation/models/create_container_group_networking.md)        | Represents container group networking parameters                         |
| [UpdateContainer](documentation/models/update_container.md)                                        | Represents an update container object                                    |
| [UpdateContainerGroupNetworking](documentation/models/update_container_group_networking.md)        | Represents update container group networking parameters                  |
| [WorkloadError](documentation/models/workload_error.md)                                            | Represents a workload error                                              |
| [QueueJobEvent](documentation/models/queue_job_event.md)                                           | Represents an event for queue job                                        |
| [ContainerGroupsQuotas](documentation/models/container_groups_quotas.md)                           |                                                                          |
| [InferenceEndpointJobEvent](documentation/models/inference_endpoint_job_event.md)                  | Represents an event for inference endpoint job                           |
| [GpuClass](documentation/models/gpu_class.md)                                                      | Represents a GPU Class                                                   |
| [GpuClassPrice](documentation/models/gpu_class_price.md)                                           | Represents the price of a GPU class for a given container group priority |

</details>

## License

This SDK is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more details.
