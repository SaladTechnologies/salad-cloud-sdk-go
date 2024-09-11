# QueuesService

A list of all methods in the `QueuesService` service. Click on the method name to view detailed information about that method.

| Methods                           | Description                                     |
| :-------------------------------- | :---------------------------------------------- |
| [ListQueues](#listqueues)         | Gets the list of queues in the given project.   |
| [CreateQueue](#createqueue)       | Creates a new queue in the given project.       |
| [GetQueue](#getqueue)             | Gets an existing queue in the given project.    |
| [UpdateQueue](#updatequeue)       | Updates an existing queue in the given project. |
| [DeleteQueue](#deletequeue)       | Deletes an existing queue in the given project. |
| [ListQueueJobs](#listqueuejobs)   | Gets the list of jobs in a queue                |
| [CreateQueueJob](#createqueuejob) | Creates a new job                               |
| [GetQueueJob](#getqueuejob)       | Gets a job in a queue                           |
| [DeleteQueueJob](#deletequeuejob) | Cancels a job in a queue                        |

## ListQueues

Gets the list of queues in the given project.

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues`

**Parameters**

| Name             | Type      | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :-------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | `Context` | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | `string`  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | `string`  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |

**Return Type**

`QueueList`

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

response, err := client.Queues.ListQueues(context.Background(), "organizationName", "projectName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateQueue

Creates a new queue in the given project.

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues`

**Parameters**

| Name             | Type          | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | `Context`     | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | `string`      | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | `string`      | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| createQueue      | `CreateQueue` | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`Queue`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


request := queues.CreateQueue{}
request.SetName("Name")

response, err := client.Queues.CreateQueue(context.Background(), "organizationName", "projectName", request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## GetQueue

Gets an existing queue in the given project.

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}`

**Parameters**

| Name             | Type      | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :-------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | `Context` | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | `string`  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | `string`  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | `string`  | ✅       | The queue name.                                                                                                                                                                                                                                     |

**Return Type**

`Queue`

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

response, err := client.Queues.GetQueue(context.Background(), "organizationName", "projectName", "queueName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## UpdateQueue

Updates an existing queue in the given project.

- HTTP Method: `PATCH`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}`

**Parameters**

| Name             | Type          | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | `Context`     | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | `string`      | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | `string`      | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | `string`      | ✅       | The queue name.                                                                                                                                                                                                                                     |
| updateQueue      | `UpdateQueue` | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`Queue`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


request := queues.UpdateQueue{}

response, err := client.Queues.UpdateQueue(context.Background(), "organizationName", "projectName", "queueName", request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## DeleteQueue

Deletes an existing queue in the given project.

- HTTP Method: `DELETE`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}`

**Parameters**

| Name             | Type      | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :-------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | `Context` | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | `string`  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | `string`  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | `string`  | ✅       | The queue name.                                                                                                                                                                                                                                     |

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

response, err := client.Queues.DeleteQueue(context.Background(), "organizationName", "projectName", "queueName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## ListQueueJobs

Gets the list of jobs in a queue

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs`

**Parameters**

| Name             | Type                         | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :--------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | `Context`                    | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | `string`                     | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | `string`                     | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | `string`                     | ✅       | The queue name.                                                                                                                                                                                                                                     |
| params           | `ListQueueJobsRequestParams` | ✅       | Additional request parameters                                                                                                                                                                                                                       |

**Return Type**

`QueueJobList`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


params := queues.ListQueueJobsRequestParams{}


response, err := client.Queues.ListQueueJobs(context.Background(), "organizationName", "projectName", "queueName", params)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateQueueJob

Creates a new job

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs`

**Parameters**

| Name             | Type             | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :--------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | `Context`        | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | `string`         | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | `string`         | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | `string`         | ✅       | The queue name.                                                                                                                                                                                                                                     |
| createQueueJob   | `CreateQueueJob` | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`QueueJob`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


request := queues.CreateQueueJob{}
request.SetInput(any)

response, err := client.Queues.CreateQueueJob(context.Background(), "organizationName", "projectName", "queueName", request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## GetQueueJob

Gets a job in a queue

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id}`

**Parameters**

| Name             | Type      | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :-------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | `Context` | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | `string`  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | `string`  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | `string`  | ✅       | The queue name.                                                                                                                                                                                                                                     |
| queueJobId       | `string`  | ✅       | The job identifier. This is automatically generated and assigned when the job is created.                                                                                                                                                           |

**Return Type**

`QueueJob`

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

response, err := client.Queues.GetQueueJob(context.Background(), "organizationName", "projectName", "queueName", "queueJobId")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## DeleteQueueJob

Cancels a job in a queue

- HTTP Method: `DELETE`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id}`

**Parameters**

| Name             | Type      | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :-------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | `Context` | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | `string`  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | `string`  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | `string`  | ✅       | The queue name.                                                                                                                                                                                                                                     |
| queueJobId       | `string`  | ✅       | The job identifier. This is automatically generated and assigned when the job is created.                                                                                                                                                           |

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

response, err := client.Queues.DeleteQueueJob(context.Background(), "organizationName", "projectName", "queueName", "queueJobId")
if err != nil {
  panic(err)
}

fmt.Print(response)
```
