# ContainerGroupsService

A list of all methods in the `ContainerGroupsService` service. Click on the method name to view detailed information about that method.

| Methods                                                               | Description                                                                                                                 |
| :-------------------------------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------- |
| [ListContainerGroups](#listcontainergroups)                           | Gets the list of container groups                                                                                           |
| [CreateContainerGroup](#createcontainergroup)                         | Creates a new container group                                                                                               |
| [GetContainerGroup](#getcontainergroup)                               | Gets a container group                                                                                                      |
| [UpdateContainerGroup](#updatecontainergroup)                         | Updates a container group                                                                                                   |
| [DeleteContainerGroup](#deletecontainergroup)                         | Deletes a container group                                                                                                   |
| [StartContainerGroup](#startcontainergroup)                           | Starts a container group                                                                                                    |
| [StopContainerGroup](#stopcontainergroup)                             | Stops a container group                                                                                                     |
| [ListContainerGroupInstances](#listcontainergroupinstances)           | Gets the list of container group instances                                                                                  |
| [GetContainerGroupInstance](#getcontainergroupinstance)               | Gets a container group instance                                                                                             |
| [ReallocateContainerGroupInstance](#reallocatecontainergroupinstance) | Reallocates a container group instance to run on a different Salad Node                                                     |
| [RecreateContainerGroupInstance](#recreatecontainergroupinstance)     | Stops a container, destroys it, and starts a new one without requiring the image to be downloaded again on a new Salad Node |
| [RestartContainerGroupInstance](#restartcontainergroupinstance)       | Stops a container and restarts it on the same Salad Node                                                                    |

## ListContainerGroups

Gets the list of container groups

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers`

**Parameters**

| Name             | Type    | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName      | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |

**Return Type**

`ContainerGroupList`

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

response, err := client.ContainerGroups.ListContainerGroups(context.Background(), "organizationName", "projectName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateContainerGroup

Creates a new container group

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers`

**Parameters**

| Name                 | Type                 | Required | Description                                                                                                                                                                                                                                         |
| :------------------- | :------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                  | Context              | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName     | string               | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName          | string               | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| createContainerGroup | CreateContainerGroup | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`shared.ContainerGroup`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/containergroups"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


containerResourceRequirements := shared.ContainerResourceRequirements{}
containerResourceRequirements.SetCpu(int64(123))
containerResourceRequirements.SetMemory(int64(123))

createContainer := containergroups.CreateContainer{}
createContainer.SetImage("Image")
createContainer.SetResources(containerResourceRequirements)

containerRestartPolicy := shared.CONTAINER_RESTART_POLICY_ALWAYS

request := containergroups.CreateContainerGroup{}
request.SetName("Name")
request.SetContainer(createContainer)
request.SetAutostartPolicy(true)
request.SetRestartPolicy(containerRestartPolicy)
request.SetReplicas(int64(123))

response, err := client.ContainerGroups.CreateContainerGroup(context.Background(), "organizationName", "projectName", request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## GetContainerGroup

Gets a container group

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}`

**Parameters**

| Name               | Type    | Required | Description                                                                                                                                                                                                                                         |
| :----------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName   | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName        | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName | string  | ✅       | The unique container group name                                                                                                                                                                                                                     |

**Return Type**

`shared.ContainerGroup`

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

response, err := client.ContainerGroups.GetContainerGroup(context.Background(), "organizationName", "projectName", "containerGroupName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## UpdateContainerGroup

Updates a container group

- HTTP Method: `PATCH`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}`

**Parameters**

| Name                 | Type                 | Required | Description                                                                                                                                                                                                                                         |
| :------------------- | :------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                  | Context              | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName     | string               | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName          | string               | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName   | string               | ✅       | The unique container group name                                                                                                                                                                                                                     |
| updateContainerGroup | UpdateContainerGroup | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`shared.ContainerGroup`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/containergroups"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


request := containergroups.UpdateContainerGroup{}

response, err := client.ContainerGroups.UpdateContainerGroup(context.Background(), "organizationName", "projectName", "containerGroupName", request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## DeleteContainerGroup

Deletes a container group

- HTTP Method: `DELETE`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}`

**Parameters**

| Name               | Type    | Required | Description                                                                                                                                                                                                                                         |
| :----------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName   | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName        | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName | string  | ✅       | The unique container group name                                                                                                                                                                                                                     |

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

response, err := client.ContainerGroups.DeleteContainerGroup(context.Background(), "organizationName", "projectName", "containerGroupName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## StartContainerGroup

Starts a container group

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/start`

**Parameters**

| Name               | Type    | Required | Description                                                                                                                                                                                                                                         |
| :----------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName   | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName        | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName | string  | ✅       | The unique container group name                                                                                                                                                                                                                     |

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

response, err := client.ContainerGroups.StartContainerGroup(context.Background(), "organizationName", "projectName", "containerGroupName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## StopContainerGroup

Stops a container group

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/stop`

**Parameters**

| Name               | Type    | Required | Description                                                                                                                                                                                                                                         |
| :----------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName   | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName        | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName | string  | ✅       | The unique container group name                                                                                                                                                                                                                     |

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

response, err := client.ContainerGroups.StopContainerGroup(context.Background(), "organizationName", "projectName", "containerGroupName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## ListContainerGroupInstances

Gets the list of container group instances

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances`

**Parameters**

| Name               | Type    | Required | Description                                                                                                                                                                                                                                         |
| :----------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName   | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName        | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName | string  | ✅       | The unique container group name                                                                                                                                                                                                                     |

**Return Type**

`ContainerGroupInstances`

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

response, err := client.ContainerGroups.ListContainerGroupInstances(context.Background(), "organizationName", "projectName", "containerGroupName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## GetContainerGroupInstance

Gets a container group instance

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}`

**Parameters**

| Name                     | Type    | Required | Description                                                                                                                                                                                                                                         |
| :----------------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                      | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName         | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName              | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName       | string  | ✅       | The unique container group name                                                                                                                                                                                                                     |
| containerGroupInstanceId | string  | ✅       | The unique instance identifier                                                                                                                                                                                                                      |

**Return Type**

`ContainerGroupInstance`

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

response, err := client.ContainerGroups.GetContainerGroupInstance(context.Background(), "organizationName", "projectName", "containerGroupName", "containerGroupInstanceId")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## ReallocateContainerGroupInstance

Reallocates a container group instance to run on a different Salad Node

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/reallocate`

**Parameters**

| Name                     | Type    | Required | Description                                                                                                                                                                                                                                         |
| :----------------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                      | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName         | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName              | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName       | string  | ✅       | The unique container group name                                                                                                                                                                                                                     |
| containerGroupInstanceId | string  | ✅       | The unique instance identifier                                                                                                                                                                                                                      |

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

response, err := client.ContainerGroups.ReallocateContainerGroupInstance(context.Background(), "organizationName", "projectName", "containerGroupName", "containerGroupInstanceId")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## RecreateContainerGroupInstance

Stops a container, destroys it, and starts a new one without requiring the image to be downloaded again on a new Salad Node

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/recreate`

**Parameters**

| Name                     | Type    | Required | Description                                                                                                                                                                                                                                         |
| :----------------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                      | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName         | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName              | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName       | string  | ✅       | The unique container group name                                                                                                                                                                                                                     |
| containerGroupInstanceId | string  | ✅       | The unique instance identifier                                                                                                                                                                                                                      |

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

response, err := client.ContainerGroups.RecreateContainerGroupInstance(context.Background(), "organizationName", "projectName", "containerGroupName", "containerGroupInstanceId")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## RestartContainerGroupInstance

Stops a container and restarts it on the same Salad Node

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/restart`

**Parameters**

| Name                     | Type    | Required | Description                                                                                                                                                                                                                                         |
| :----------------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                      | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName         | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName              | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName       | string  | ✅       | The unique container group name                                                                                                                                                                                                                     |
| containerGroupInstanceId | string  | ✅       | The unique instance identifier                                                                                                                                                                                                                      |

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

response, err := client.ContainerGroups.RestartContainerGroupInstance(context.Background(), "organizationName", "projectName", "containerGroupName", "containerGroupInstanceId")
if err != nil {
  panic(err)
}

fmt.Print(response)
```
