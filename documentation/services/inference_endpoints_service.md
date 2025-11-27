# InferenceEndpointsService

A list of all methods in the `InferenceEndpointsService` service. Click on the method name to view detailed information about that method.

| Methods                                                   | Description                           |
| :-------------------------------------------------------- | :------------------------------------ |
| [ListInferenceEndpoints](#listinferenceendpoints)         | Lists inference endpoints.            |
| [GetInferenceEndpoint](#getinferenceendpoint)             | Gets an inference endpoint.           |
| [ListInferenceEndpointJobs](#listinferenceendpointjobs)   | Lists inference endpoint jobs.        |
| [CreateInferenceEndpointJob](#createinferenceendpointjob) | Creates a new inference endpoint job. |
| [GetInferenceEndpointJob](#getinferenceendpointjob)       | Gets an inference endpoint job.       |
| [DeleteInferenceEndpointJob](#deleteinferenceendpointjob) | Cancels an inference endpoint job.    |

## ListInferenceEndpoints

Lists inference endpoints.

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/inference-endpoints`

**Parameters**

| Name             | Type                                | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :---------------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context                             | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string                              | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| params           | ListInferenceEndpointsRequestParams | ✅       | Additional request parameters                                                                                                                                                                                                                       |

**Return Type**

`InferenceEndpointCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


params := inferenceendpoints.ListInferenceEndpointsRequestParams{
  Page: util.ToPointer(int64(1)),
  PageSize: util.ToPointer(int64(1)),
}

response, err := client.InferenceEndpoints.ListInferenceEndpoints(context.Background(), "acme-corp", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetInferenceEndpoint

Gets an inference endpoint.

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}`

**Parameters**

| Name                  | Type    | Required | Description                                                                                                                                                                                                                                         |
| :-------------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                   | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName      | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| inferenceEndpointName | string  | ✅       | The inference endpoint name.                                                                                                                                                                                                                        |

**Return Type**

`InferenceEndpoint`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.InferenceEndpoints.GetInferenceEndpoint(context.Background(), "acme-corp", "transcribe")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListInferenceEndpointJobs

Lists inference endpoint jobs.

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs`

**Parameters**

| Name                  | Type                                   | Required | Description                                                                                                                                                                                                                                         |
| :-------------------- | :------------------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                   | Context                                | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName      | string                                 | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| inferenceEndpointName | string                                 | ✅       | The inference endpoint name.                                                                                                                                                                                                                        |
| params                | ListInferenceEndpointJobsRequestParams | ✅       | Additional request parameters                                                                                                                                                                                                                       |

**Return Type**

`InferenceEndpointJobCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


params := inferenceendpoints.ListInferenceEndpointJobsRequestParams{
  Page: util.ToPointer(int64(1)),
  PageSize: util.ToPointer(int64(1)),
}

response, err := client.InferenceEndpoints.ListInferenceEndpointJobs(context.Background(), "acme-corp", "transcribe", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateInferenceEndpointJob

Creates a new inference endpoint job.

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs`

**Parameters**

| Name                          | Type                          | Required | Description                                                                                                                                                                                                                                         |
| :---------------------------- | :---------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                           | Context                       | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName              | string                        | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| inferenceEndpointName         | string                        | ✅       | The inference endpoint name.                                                                                                                                                                                                                        |
| inferenceEndpointJobPrototype | InferenceEndpointJobPrototype | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`InferenceEndpointJob`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


request := inferenceendpoints.InferenceEndpointJobPrototype{
  Input: []byte{},
  Metadata: []byte{},
  Webhook: util.ToPointer("webhook"),
  WebhookUrl: util.ToPointer("https://webhook.example.com/events"),
}

response, err := client.InferenceEndpoints.CreateInferenceEndpointJob(context.Background(), "acme-corp", "transcribe", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetInferenceEndpointJob

Gets an inference endpoint job.

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs/{inference_endpoint_job_id}`

**Parameters**

| Name                   | Type    | Required | Description                                                                                                                                                                                                                                         |
| :--------------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                    | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName       | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| inferenceEndpointName  | string  | ✅       | The inference endpoint name.                                                                                                                                                                                                                        |
| inferenceEndpointJobId | string  | ✅       | The inference endpoint job identifier.                                                                                                                                                                                                              |

**Return Type**

`InferenceEndpointJob`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.InferenceEndpoints.GetInferenceEndpointJob(context.Background(), "acme-corp", "transcribe", "2fc459a1-1c09-4a34-ade7-54d03fc51d6a")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteInferenceEndpointJob

Cancels an inference endpoint job.

- HTTP Method: `DELETE`
- Endpoint: `/organizations/{organization_name}/inference-endpoints/{inference_endpoint_name}/jobs/{inference_endpoint_job_id}`

**Parameters**

| Name                   | Type    | Required | Description                                                                                                                                                                                                                                         |
| :--------------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                    | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName       | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| inferenceEndpointName  | string  | ✅       | The inference endpoint name.                                                                                                                                                                                                                        |
| inferenceEndpointJobId | string  | ✅       | The inference endpoint job identifier.                                                                                                                                                                                                              |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.InferenceEndpoints.DeleteInferenceEndpointJob(context.Background(), "acme-corp", "transcribe", "2fc459a1-1c09-4a34-ade7-54d03fc51d6a")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
