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

| Name             | Type    | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |

**Return Type**

`QueueCollection`

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

response, err := client.Queues.ListQueues(context.Background(), "acme-corp", "dev-env")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateQueue

Creates a new queue in the given project.

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues`

**Parameters**

| Name             | Type           | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context        | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string         | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | string         | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queuePrototype   | QueuePrototype | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`Queue`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


request := queues.QueuePrototype{
  Description: util.ToPointer("description"),
  DisplayName: util.ToPointer("tLWSUinMUjM"),
  Name: util.ToPointer("name"),
}

response, err := client.Queues.CreateQueue(context.Background(), "acme-corp", "dev-env", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetQueue

Gets an existing queue in the given project.

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}`

**Parameters**

| Name             | Type    | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | string  | ✅       | The queue name.                                                                                                                                                                                                                                     |

**Return Type**

`Queue`

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

response, err := client.Queues.GetQueue(context.Background(), "acme-corp", "dev-env", "fifo-queue")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## UpdateQueue

Updates an existing queue in the given project.

- HTTP Method: `PATCH`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}`

**Parameters**

| Name             | Type       | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :--------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context    | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string     | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | string     | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | string     | ✅       | The queue name.                                                                                                                                                                                                                                     |
| queuePatch       | QueuePatch | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`Queue`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


request := queues.QueuePatch{
  Description: util.ToPointer(util.Nullable[string]{ Value: "description" }),
  DisplayName: util.ToPointer(util.Nullable[string]{ Value: "hB13" }),
}

response, err := client.Queues.UpdateQueue(context.Background(), "acme-corp", "dev-env", "fifo-queue", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteQueue

Deletes an existing queue in the given project.

- HTTP Method: `DELETE`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}`

**Parameters**

| Name             | Type    | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | string  | ✅       | The queue name.                                                                                                                                                                                                                                     |

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

response, err := client.Queues.DeleteQueue(context.Background(), "acme-corp", "dev-env", "fifo-queue")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListQueueJobs

Gets the list of jobs in a queue

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs`

**Parameters**

| Name             | Type                       | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context                    | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string                     | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | string                     | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | string                     | ✅       | The queue name.                                                                                                                                                                                                                                     |
| params           | ListQueueJobsRequestParams | ✅       | Additional request parameters                                                                                                                                                                                                                       |

**Return Type**

`QueueJobCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


params := queues.ListQueueJobsRequestParams{
  Page: util.ToPointer(int64(1)),
  PageSize: util.ToPointer(int64(1)),
}

response, err := client.Queues.ListQueueJobs(context.Background(), "acme-corp", "dev-env", "fifo-queue", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateQueueJob

Creates a new job

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs`

**Parameters**

| Name              | Type              | Required | Description                                                                                                                                                                                                                                         |
| :---------------- | :---------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx               | Context           | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName  | string            | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName       | string            | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName         | string            | ✅       | The queue name.                                                                                                                                                                                                                                     |
| queueJobPrototype | QueueJobPrototype | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`QueueJob`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


request := queues.QueueJobPrototype{
  Input: []byte{},
  Metadata: []byte{},
  Webhook: util.ToPointer("webhook"),
}

response, err := client.Queues.CreateQueueJob(context.Background(), "acme-corp", "dev-env", "fifo-queue", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetQueueJob

Gets a job in a queue

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id}`

**Parameters**

| Name             | Type    | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | string  | ✅       | The queue name.                                                                                                                                                                                                                                     |
| queueJobId       | string  | ✅       | The job identifier. This is automatically generated and assigned when the job is created.                                                                                                                                                           |

**Return Type**

`QueueJob`

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

response, err := client.Queues.GetQueueJob(context.Background(), "acme-corp", "dev-env", "fifo-queue", "7dcd6922-50e9-4d56-89b5-91cde26f0211")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteQueueJob

Cancels a job in a queue

- HTTP Method: `DELETE`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id}`

**Parameters**

| Name             | Type    | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| queueName        | string  | ✅       | The queue name.                                                                                                                                                                                                                                     |
| queueJobId       | string  | ✅       | The job identifier. This is automatically generated and assigned when the job is created.                                                                                                                                                           |

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

response, err := client.Queues.DeleteQueueJob(context.Background(), "acme-corp", "dev-env", "fifo-queue", "7dcd6922-50e9-4d56-89b5-91cde26f0211")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
