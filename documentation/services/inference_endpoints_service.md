# InferenceEndpointsService

A list of all methods in the `InferenceEndpointsService` service. Click on the method name to view detailed information about that method.

| Methods                                                   | Description                                    |
| :-------------------------------------------------------- | :--------------------------------------------- |
| [ListInferenceEndpoints](#listinferenceendpoints)         | Gets the list of inference endpoints           |
| [GetInferenceEndpoint](#getinferenceendpoint)             | Gets an inference endpoint                     |
| [GetInferenceEndpointJobs](#getinferenceendpointjobs)     | Retrieves a list of an inference endpoint jobs |
| [CreateInferenceEndpointJob](#createinferenceendpointjob) | Creates a new job                              |
| [GetInferenceEndpointJob](#getinferenceendpointjob)       | Retrieves a job in an inference endpoint       |
| [DeleteInferenceEndpointJob](#deleteinferenceendpointjob) | Deletes an inference endpoint job              |

## ListInferenceEndpoints

Gets the list of inference endpoints

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/inference-endpoints`

**Parameters**

| Name             | Type                                  | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------------------------------------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | `Context`                             | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | `string`                              | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| params           | `ListInferenceEndpointsRequestParams` | ✅       | Additional request parameters                                                                                                                                                                                                                       |

**Return Type**

`InferenceEndpointsList`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


params := inferenceendpoints.ListInferenceEndpointsRequestParams{}


response, err := client.InferenceEndpoints.ListInferenceEndpoints(context.Background(), "organizationName", params)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## GetInferenceEndpoint

Gets an inference endpoint

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}`

**Parameters**

| Name                  | Type      | Required | Description                                                                                                                                                                                                                                         |
| :-------------------- | :-------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                   | `Context` | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName      | `string`  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| inferenceEndpointName | `string`  | ✅       | The unique inference endpoint name                                                                                                                                                                                                                  |

**Return Type**

`InferenceEndpoint`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.InferenceEndpoints.GetInferenceEndpoint(context.Background(), "organizationName", "inferenceEndpointName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## GetInferenceEndpointJobs

Retrieves a list of an inference endpoint jobs

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs`

**Parameters**

| Name                  | Type                                    | Required | Description                                                                                                                                                                                                                                         |
| :-------------------- | :-------------------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                   | `Context`                               | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName      | `string`                                | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| inferenceEndpointName | `string`                                | ✅       | The unique inference endpoint name                                                                                                                                                                                                                  |
| params                | `GetInferenceEndpointJobsRequestParams` | ✅       | Additional request parameters                                                                                                                                                                                                                       |

**Return Type**

`InferenceEndpointJobList`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


params := inferenceendpoints.GetInferenceEndpointJobsRequestParams{}


response, err := client.InferenceEndpoints.GetInferenceEndpointJobs(context.Background(), "organizationName", "inferenceEndpointName", params)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateInferenceEndpointJob

Creates a new job

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs`

**Parameters**

| Name                       | Type                         | Required | Description                                                                                                                                                                                                                                         |
| :------------------------- | :--------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                        | `Context`                    | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName           | `string`                     | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| inferenceEndpointName      | `string`                     | ✅       | The unique inference endpoint name                                                                                                                                                                                                                  |
| createInferenceEndpointJob | `CreateInferenceEndpointJob` | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`InferenceEndpointJob`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


request := inferenceendpoints.CreateInferenceEndpointJob{}
request.SetInput(any)

response, err := client.InferenceEndpoints.CreateInferenceEndpointJob(context.Background(), "organizationName", "inferenceEndpointName", request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## GetInferenceEndpointJob

Retrieves a job in an inference endpoint

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs/{inference_endpoint_job_id}`

**Parameters**

| Name                   | Type      | Required | Description                                                                                                                                                                                                                                         |
| :--------------------- | :-------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                    | `Context` | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName       | `string`  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| inferenceEndpointName  | `string`  | ✅       | The unique inference endpoint name                                                                                                                                                                                                                  |
| inferenceEndpointJobId | `string`  | ✅       | The unique job id                                                                                                                                                                                                                                   |

**Return Type**

`InferenceEndpointJob`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.InferenceEndpoints.GetInferenceEndpointJob(context.Background(), "organizationName", "inferenceEndpointName", "inferenceEndpointJobId")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## DeleteInferenceEndpointJob

Deletes an inference endpoint job

- HTTP Method: `DELETE`
- Endpoint: `/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs/{inference_endpoint_job_id}`

**Parameters**

| Name                   | Type      | Required | Description                                                                                                                                                                                                                                         |
| :--------------------- | :-------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                    | `Context` | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName       | `string`  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| inferenceEndpointName  | `string`  | ✅       | The unique inference endpoint name                                                                                                                                                                                                                  |
| inferenceEndpointJobId | `string`  | ✅       | The unique job id                                                                                                                                                                                                                                   |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.InferenceEndpoints.DeleteInferenceEndpointJob(context.Background(), "organizationName", "inferenceEndpointName", "inferenceEndpointJobId")
if err != nil {
  panic(err)
}

fmt.Print(response)
```
