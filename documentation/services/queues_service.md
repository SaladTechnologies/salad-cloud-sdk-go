# QueuesService

A list of all methods in the `QueuesService` service. Click on the method name to view detailed information about that method.

| Methods                           | Description                    |
| :-------------------------------- | :----------------------------- |
| [ListQueues](#listqueues)         | Gets the list of queues        |
| [CreateQueue](#createqueue)       | Creates a new queue            |
| [GetQueue](#getqueue)             | Gets a queue                   |
| [UpdateQueue](#updatequeue)       | Updates a queue                |
| [DeleteQueue](#deletequeue)       | Deletes a queue                |
| [ListQueueJobs](#listqueuejobs)   | Retrieves a list of queue jobs |
| [CreateQueueJob](#createqueuejob) | Creates a new job              |
| [GetQueueJob](#getqueuejob)       | Retrieves a job in a queue     |
| [DeleteQueueJob](#deletequeuejob) | Deletes a queue job            |

## ListQueues

Gets the list of queues

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues`

**Parameters**

| Name             | Type    | Required | Description                  |
| :--------------- | :------ | :------- | :--------------------------- |
| ctx              | Context | ✅       | Default go language context  |
| organizationName | string  | ✅       | The unique organization name |
| projectName      | string  | ✅       | The unique project name      |

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

Creates a new queue

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues`

**Parameters**

| Name             | Type        | Required | Description                  |
| :--------------- | :---------- | :------- | :--------------------------- |
| ctx              | Context     | ✅       | Default go language context  |
| organizationName | string      | ✅       | The unique organization name |
| projectName      | string      | ✅       | The unique project name      |
| createQueue      | CreateQueue | ✅       |                              |

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

Gets a queue

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}`

**Parameters**

| Name             | Type    | Required | Description                  |
| :--------------- | :------ | :------- | :--------------------------- |
| ctx              | Context | ✅       | Default go language context  |
| organizationName | string  | ✅       | The unique organization name |
| projectName      | string  | ✅       | The unique project name      |
| queueName        | string  | ✅       | The unique queue name        |

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

Updates a queue

- HTTP Method: `PATCH`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}`

**Parameters**

| Name             | Type        | Required | Description                  |
| :--------------- | :---------- | :------- | :--------------------------- |
| ctx              | Context     | ✅       | Default go language context  |
| organizationName | string      | ✅       | The unique organization name |
| projectName      | string      | ✅       | The unique project name      |
| queueName        | string      | ✅       | The unique queue name        |
| updateQueue      | UpdateQueue | ✅       |                              |

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

Deletes a queue

- HTTP Method: `DELETE`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}`

**Parameters**

| Name             | Type    | Required | Description                  |
| :--------------- | :------ | :------- | :--------------------------- |
| ctx              | Context | ✅       | Default go language context  |
| organizationName | string  | ✅       | The unique organization name |
| projectName      | string  | ✅       | The unique project name      |
| queueName        | string  | ✅       | The unique queue name        |

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

Retrieves a list of queue jobs

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs`

**Parameters**

| Name             | Type                       | Required | Description                   |
| :--------------- | :------------------------- | :------- | :---------------------------- |
| ctx              | Context                    | ✅       | Default go language context   |
| organizationName | string                     | ✅       | The unique organization name  |
| projectName      | string                     | ✅       | The unique project name       |
| queueName        | string                     | ✅       | The unique queue name         |
| params           | ListQueueJobsRequestParams | ✅       | Additional request parameters |

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

| Name             | Type           | Required | Description                  |
| :--------------- | :------------- | :------- | :--------------------------- |
| ctx              | Context        | ✅       | Default go language context  |
| organizationName | string         | ✅       | The unique organization name |
| projectName      | string         | ✅       | The unique project name      |
| queueName        | string         | ✅       | The unique queue name        |
| createQueueJob   | CreateQueueJob | ✅       |                              |

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

Retrieves a job in a queue

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id}`

**Parameters**

| Name             | Type    | Required | Description                  |
| :--------------- | :------ | :------- | :--------------------------- |
| ctx              | Context | ✅       | Default go language context  |
| organizationName | string  | ✅       | The unique organization name |
| projectName      | string  | ✅       | The unique project name      |
| queueName        | string  | ✅       | The unique queue name        |
| queueJobId       | string  | ✅       | The unique job id            |

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

Deletes a queue job

- HTTP Method: `DELETE`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id}`

**Parameters**

| Name             | Type    | Required | Description                  |
| :--------------- | :------ | :------- | :--------------------------- |
| ctx              | Context | ✅       | Default go language context  |
| organizationName | string  | ✅       | The unique organization name |
| projectName      | string  | ✅       | The unique project name      |
| queueName        | string  | ✅       | The unique queue name        |
| queueJobId       | string  | ✅       | The unique job id            |

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
