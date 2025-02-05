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

fmt.Println(response)
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
containerResourceRequirements.SetGpuClasses([]string{})
containerResourceRequirements.SetStorageAmount(int64(123))

containerGroupPriority := shared.CONTAINER_GROUP_PRIORITY_HIGH


loggingAxiom2 := containergroups.LoggingAxiom2{}
loggingAxiom2.SetHost("Host")
loggingAxiom2.SetApiToken("ApiToken")
loggingAxiom2.SetDataset("Dataset")


datadogTags2 := containergroups.DatadogTags2{}
datadogTags2.SetName("Name")
datadogTags2.SetValue("Value")

loggingDatadog2 := containergroups.LoggingDatadog2{}
loggingDatadog2.SetHost("Host")
loggingDatadog2.SetApiKey("ApiKey")
loggingDatadog2.SetTags([]containergroups.DatadogTags2{datadogTags2})


loggingNewRelic2 := containergroups.LoggingNewRelic2{}
loggingNewRelic2.SetHost("Host")
loggingNewRelic2.SetIngestionKey("IngestionKey")


loggingSplunk2 := containergroups.LoggingSplunk2{}
loggingSplunk2.SetHost("Host")
loggingSplunk2.SetToken("Token")


loggingTcp2 := containergroups.LoggingTcp2{}
loggingTcp2.SetHost("Host")
loggingTcp2.SetPort(int64(123))

httpFormat2 := containergroups.HTTP_FORMAT2_JSON


httpHeaders3 := containergroups.HttpHeaders3{}
httpHeaders3.SetName("Name")
httpHeaders3.SetValue("Value")

httpCompression2 := containergroups.HTTP_COMPRESSION2_NONE

loggingHttp2 := containergroups.LoggingHttp2{}
loggingHttp2.SetHost("Host")
loggingHttp2.SetPort(int64(123))
loggingHttp2.SetUser("User")
loggingHttp2.SetPassword("Password")
loggingHttp2.SetPath("Path")
loggingHttp2.SetFormat(httpFormat2)
loggingHttp2.SetHeaders([]containergroups.HttpHeaders3{httpHeaders3})
loggingHttp2.SetCompression(httpCompression2)

createContainerLogging := containergroups.CreateContainerLogging{}
createContainerLogging.SetAxiom(loggingAxiom2)
createContainerLogging.SetDatadog(loggingDatadog2)
createContainerLogging.SetNewRelic(loggingNewRelic2)
createContainerLogging.SetSplunk(loggingSplunk2)
createContainerLogging.SetTcp(loggingTcp2)
createContainerLogging.SetHttp(loggingHttp2)


registryAuthenticationBasic1 := containergroups.RegistryAuthenticationBasic1{}
registryAuthenticationBasic1.SetUsername("Username")
registryAuthenticationBasic1.SetPassword("Password")


registryAuthenticationGcpGcr1 := containergroups.RegistryAuthenticationGcpGcr1{}
registryAuthenticationGcpGcr1.SetServiceKey("ServiceKey")


registryAuthenticationAwsEcr1 := containergroups.RegistryAuthenticationAwsEcr1{}
registryAuthenticationAwsEcr1.SetAccessKeyId("AccessKeyId")
registryAuthenticationAwsEcr1.SetSecretAccessKey("SecretAccessKey")


registryAuthenticationDockerHub1 := containergroups.RegistryAuthenticationDockerHub1{}
registryAuthenticationDockerHub1.SetUsername("Username")
registryAuthenticationDockerHub1.SetPersonalAccessToken("PersonalAccessToken")


registryAuthenticationGcpGar1 := containergroups.RegistryAuthenticationGcpGar1{}
registryAuthenticationGcpGar1.SetServiceKey("ServiceKey")

createContainerRegistryAuthentication := containergroups.CreateContainerRegistryAuthentication{}
createContainerRegistryAuthentication.SetBasic(registryAuthenticationBasic1)
createContainerRegistryAuthentication.SetGcpGcr(registryAuthenticationGcpGcr1)
createContainerRegistryAuthentication.SetAwsEcr(registryAuthenticationAwsEcr1)
createContainerRegistryAuthentication.SetDockerHub(registryAuthenticationDockerHub1)
createContainerRegistryAuthentication.SetGcpGar(registryAuthenticationGcpGar1)

createContainer := containergroups.CreateContainer{}
createContainer.SetImage("Image")
createContainer.SetResources(containerResourceRequirements)
createContainer.SetCommand([]string{})
createContainer.SetPriority(containerGroupPriority)
createContainer.SetEnvironmentVariables(map[string]string{})
createContainer.SetLogging(createContainerLogging)
createContainer.SetRegistryAuthentication(createContainerRegistryAuthentication)

containerRestartPolicy := shared.CONTAINER_RESTART_POLICY_ALWAYS

countryCode := shared.COUNTRY_CODE_AF

containerNetworkingProtocol := shared.CONTAINER_NETWORKING_PROTOCOL_HTTP

createContainerGroupNetworkingLoadBalancer := containergroups.CREATE_CONTAINER_GROUP_NETWORKING_LOAD_BALANCER_ROUND_ROBIN

createContainerGroupNetworking := containergroups.CreateContainerGroupNetworking{}
createContainerGroupNetworking.SetProtocol(containerNetworkingProtocol)
createContainerGroupNetworking.SetPort(int64(123))
createContainerGroupNetworking.SetAuth(true)
createContainerGroupNetworking.SetLoadBalancer(createContainerGroupNetworkingLoadBalancer)
createContainerGroupNetworking.SetSingleConnectionLimit(true)
createContainerGroupNetworking.SetClientRequestTimeout(int64(123))
createContainerGroupNetworking.SetServerResponseTimeout(int64(123))


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{}
containerGroupProbeTcp.SetPort(int64(123))

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{}
containerGroupProbeHttpHeaders2.SetName("Name")
containerGroupProbeHttpHeaders2.SetValue("Value")

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{}
containerGroupProbeHttp.SetPath("Path")
containerGroupProbeHttp.SetPort(int64(123))
containerGroupProbeHttp.SetScheme(containerProbeHttpScheme)
containerGroupProbeHttp.SetHeaders([]shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2})


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{}
containerGroupProbeGrpc.SetService("Service")
containerGroupProbeGrpc.SetPort(int64(123))


containerGroupProbeExec := shared.ContainerGroupProbeExec{}
containerGroupProbeExec.SetCommand([]string{})

containerGroupLivenessProbe := shared.ContainerGroupLivenessProbe{}
containerGroupLivenessProbe.SetTcp(containerGroupProbeTcp)
containerGroupLivenessProbe.SetHttp(containerGroupProbeHttp)
containerGroupLivenessProbe.SetGrpc(containerGroupProbeGrpc)
containerGroupLivenessProbe.SetExec(containerGroupProbeExec)
containerGroupLivenessProbe.SetInitialDelaySeconds(int64(123))
containerGroupLivenessProbe.SetPeriodSeconds(int64(123))
containerGroupLivenessProbe.SetTimeoutSeconds(int64(123))
containerGroupLivenessProbe.SetSuccessThreshold(int64(123))
containerGroupLivenessProbe.SetFailureThreshold(int64(123))


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{}
containerGroupProbeTcp.SetPort(int64(123))

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{}
containerGroupProbeHttpHeaders2.SetName("Name")
containerGroupProbeHttpHeaders2.SetValue("Value")

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{}
containerGroupProbeHttp.SetPath("Path")
containerGroupProbeHttp.SetPort(int64(123))
containerGroupProbeHttp.SetScheme(containerProbeHttpScheme)
containerGroupProbeHttp.SetHeaders([]shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2})


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{}
containerGroupProbeGrpc.SetService("Service")
containerGroupProbeGrpc.SetPort(int64(123))


containerGroupProbeExec := shared.ContainerGroupProbeExec{}
containerGroupProbeExec.SetCommand([]string{})

containerGroupReadinessProbe := shared.ContainerGroupReadinessProbe{}
containerGroupReadinessProbe.SetTcp(containerGroupProbeTcp)
containerGroupReadinessProbe.SetHttp(containerGroupProbeHttp)
containerGroupReadinessProbe.SetGrpc(containerGroupProbeGrpc)
containerGroupReadinessProbe.SetExec(containerGroupProbeExec)
containerGroupReadinessProbe.SetInitialDelaySeconds(int64(123))
containerGroupReadinessProbe.SetPeriodSeconds(int64(123))
containerGroupReadinessProbe.SetTimeoutSeconds(int64(123))
containerGroupReadinessProbe.SetSuccessThreshold(int64(123))
containerGroupReadinessProbe.SetFailureThreshold(int64(123))


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{}
containerGroupProbeTcp.SetPort(int64(123))

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{}
containerGroupProbeHttpHeaders2.SetName("Name")
containerGroupProbeHttpHeaders2.SetValue("Value")

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{}
containerGroupProbeHttp.SetPath("Path")
containerGroupProbeHttp.SetPort(int64(123))
containerGroupProbeHttp.SetScheme(containerProbeHttpScheme)
containerGroupProbeHttp.SetHeaders([]shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2})


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{}
containerGroupProbeGrpc.SetService("Service")
containerGroupProbeGrpc.SetPort(int64(123))


containerGroupProbeExec := shared.ContainerGroupProbeExec{}
containerGroupProbeExec.SetCommand([]string{})

containerGroupStartupProbe := shared.ContainerGroupStartupProbe{}
containerGroupStartupProbe.SetTcp(containerGroupProbeTcp)
containerGroupStartupProbe.SetHttp(containerGroupProbeHttp)
containerGroupStartupProbe.SetGrpc(containerGroupProbeGrpc)
containerGroupStartupProbe.SetExec(containerGroupProbeExec)
containerGroupStartupProbe.SetInitialDelaySeconds(int64(123))
containerGroupStartupProbe.SetPeriodSeconds(int64(123))
containerGroupStartupProbe.SetTimeoutSeconds(int64(123))
containerGroupStartupProbe.SetSuccessThreshold(int64(123))
containerGroupStartupProbe.SetFailureThreshold(int64(123))


containerGroupQueueConnection := shared.ContainerGroupQueueConnection{}
containerGroupQueueConnection.SetPath("Path")
containerGroupQueueConnection.SetPort(int64(123))
containerGroupQueueConnection.SetQueueName("QueueName")


queueAutoscaler := shared.QueueAutoscaler{}
queueAutoscaler.SetMinReplicas(int64(123))
queueAutoscaler.SetMaxReplicas(int64(123))
queueAutoscaler.SetDesiredQueueLength(int64(123))
queueAutoscaler.SetPollingPeriod(int64(123))
queueAutoscaler.SetMaxUpscalePerMinute(int64(123))
queueAutoscaler.SetMaxDownscalePerMinute(int64(123))

request := containergroups.CreateContainerGroup{}
request.SetName("Name")
request.SetDisplayName("DisplayName")
request.SetContainer(createContainer)
request.SetAutostartPolicy(true)
request.SetRestartPolicy(containerRestartPolicy)
request.SetReplicas(int64(123))
request.SetCountryCodes([]shared.CountryCode{countryCode})
request.SetNetworking(createContainerGroupNetworking)
request.SetLivenessProbe(containerGroupLivenessProbe)
request.SetReadinessProbe(containerGroupReadinessProbe)
request.SetStartupProbe(containerGroupStartupProbe)
request.SetQueueConnection(containerGroupQueueConnection)
request.SetQueueAutoscaler(queueAutoscaler)

response, err := client.ContainerGroups.CreateContainerGroup(context.Background(), "organizationName", "projectName", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
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

fmt.Println(response)
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


resources := containergroups.Resources{}
resources.SetCpu(int64(123))
resources.SetMemory(int64(123))
resources.SetGpuClasses([]string{})
resources.SetStorageAmount(int64(123))

containerGroupPriority := shared.CONTAINER_GROUP_PRIORITY_HIGH


loggingAxiom3 := containergroups.LoggingAxiom3{}
loggingAxiom3.SetHost("Host")
loggingAxiom3.SetApiToken("ApiToken")
loggingAxiom3.SetDataset("Dataset")


datadogTags3 := containergroups.DatadogTags3{}
datadogTags3.SetName("Name")
datadogTags3.SetValue("Value")

loggingDatadog3 := containergroups.LoggingDatadog3{}
loggingDatadog3.SetHost("Host")
loggingDatadog3.SetApiKey("ApiKey")
loggingDatadog3.SetTags([]containergroups.DatadogTags3{datadogTags3})


loggingNewRelic3 := containergroups.LoggingNewRelic3{}
loggingNewRelic3.SetHost("Host")
loggingNewRelic3.SetIngestionKey("IngestionKey")


loggingSplunk3 := containergroups.LoggingSplunk3{}
loggingSplunk3.SetHost("Host")
loggingSplunk3.SetToken("Token")


loggingTcp3 := containergroups.LoggingTcp3{}
loggingTcp3.SetHost("Host")
loggingTcp3.SetPort(int64(123))

httpFormat3 := containergroups.HTTP_FORMAT3_JSON


httpHeaders4 := containergroups.HttpHeaders4{}
httpHeaders4.SetName("Name")
httpHeaders4.SetValue("Value")

httpCompression3 := containergroups.HTTP_COMPRESSION3_NONE

loggingHttp3 := containergroups.LoggingHttp3{}
loggingHttp3.SetHost("Host")
loggingHttp3.SetPort(int64(123))
loggingHttp3.SetUser("User")
loggingHttp3.SetPassword("Password")
loggingHttp3.SetPath("Path")
loggingHttp3.SetFormat(httpFormat3)
loggingHttp3.SetHeaders([]containergroups.HttpHeaders4{httpHeaders4})
loggingHttp3.SetCompression(httpCompression3)

updateContainerLogging := containergroups.UpdateContainerLogging{}
updateContainerLogging.SetAxiom(loggingAxiom3)
updateContainerLogging.SetDatadog(loggingDatadog3)
updateContainerLogging.SetNewRelic(loggingNewRelic3)
updateContainerLogging.SetSplunk(loggingSplunk3)
updateContainerLogging.SetTcp(loggingTcp3)
updateContainerLogging.SetHttp(loggingHttp3)


registryAuthenticationBasic2 := containergroups.RegistryAuthenticationBasic2{}
registryAuthenticationBasic2.SetUsername("Username")
registryAuthenticationBasic2.SetPassword("Password")


registryAuthenticationGcpGcr2 := containergroups.RegistryAuthenticationGcpGcr2{}
registryAuthenticationGcpGcr2.SetServiceKey("ServiceKey")


registryAuthenticationAwsEcr2 := containergroups.RegistryAuthenticationAwsEcr2{}
registryAuthenticationAwsEcr2.SetAccessKeyId("AccessKeyId")
registryAuthenticationAwsEcr2.SetSecretAccessKey("SecretAccessKey")


registryAuthenticationDockerHub2 := containergroups.RegistryAuthenticationDockerHub2{}
registryAuthenticationDockerHub2.SetUsername("Username")
registryAuthenticationDockerHub2.SetPersonalAccessToken("PersonalAccessToken")


registryAuthenticationGcpGar2 := containergroups.RegistryAuthenticationGcpGar2{}
registryAuthenticationGcpGar2.SetServiceKey("ServiceKey")

updateContainerRegistryAuthentication := containergroups.UpdateContainerRegistryAuthentication{}
updateContainerRegistryAuthentication.SetBasic(registryAuthenticationBasic2)
updateContainerRegistryAuthentication.SetGcpGcr(registryAuthenticationGcpGcr2)
updateContainerRegistryAuthentication.SetAwsEcr(registryAuthenticationAwsEcr2)
updateContainerRegistryAuthentication.SetDockerHub(registryAuthenticationDockerHub2)
updateContainerRegistryAuthentication.SetGcpGar(registryAuthenticationGcpGar2)

updateContainer := containergroups.UpdateContainer{}
updateContainer.SetImage("Image")
updateContainer.SetResources(resources)
updateContainer.SetCommand([]string{})
updateContainer.SetPriority(containerGroupPriority)
updateContainer.SetEnvironmentVariables(map[string]string{})
updateContainer.SetLogging(updateContainerLogging)
updateContainer.SetRegistryAuthentication(updateContainerRegistryAuthentication)

countryCode := shared.COUNTRY_CODE_AF


updateContainerGroupNetworking := containergroups.UpdateContainerGroupNetworking{}
updateContainerGroupNetworking.SetPort(int64(123))


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{}
containerGroupProbeTcp.SetPort(int64(123))

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{}
containerGroupProbeHttpHeaders2.SetName("Name")
containerGroupProbeHttpHeaders2.SetValue("Value")

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{}
containerGroupProbeHttp.SetPath("Path")
containerGroupProbeHttp.SetPort(int64(123))
containerGroupProbeHttp.SetScheme(containerProbeHttpScheme)
containerGroupProbeHttp.SetHeaders([]shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2})


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{}
containerGroupProbeGrpc.SetService("Service")
containerGroupProbeGrpc.SetPort(int64(123))


containerGroupProbeExec := shared.ContainerGroupProbeExec{}
containerGroupProbeExec.SetCommand([]string{})

containerGroupLivenessProbe := shared.ContainerGroupLivenessProbe{}
containerGroupLivenessProbe.SetTcp(containerGroupProbeTcp)
containerGroupLivenessProbe.SetHttp(containerGroupProbeHttp)
containerGroupLivenessProbe.SetGrpc(containerGroupProbeGrpc)
containerGroupLivenessProbe.SetExec(containerGroupProbeExec)
containerGroupLivenessProbe.SetInitialDelaySeconds(int64(123))
containerGroupLivenessProbe.SetPeriodSeconds(int64(123))
containerGroupLivenessProbe.SetTimeoutSeconds(int64(123))
containerGroupLivenessProbe.SetSuccessThreshold(int64(123))
containerGroupLivenessProbe.SetFailureThreshold(int64(123))


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{}
containerGroupProbeTcp.SetPort(int64(123))

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{}
containerGroupProbeHttpHeaders2.SetName("Name")
containerGroupProbeHttpHeaders2.SetValue("Value")

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{}
containerGroupProbeHttp.SetPath("Path")
containerGroupProbeHttp.SetPort(int64(123))
containerGroupProbeHttp.SetScheme(containerProbeHttpScheme)
containerGroupProbeHttp.SetHeaders([]shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2})


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{}
containerGroupProbeGrpc.SetService("Service")
containerGroupProbeGrpc.SetPort(int64(123))


containerGroupProbeExec := shared.ContainerGroupProbeExec{}
containerGroupProbeExec.SetCommand([]string{})

containerGroupReadinessProbe := shared.ContainerGroupReadinessProbe{}
containerGroupReadinessProbe.SetTcp(containerGroupProbeTcp)
containerGroupReadinessProbe.SetHttp(containerGroupProbeHttp)
containerGroupReadinessProbe.SetGrpc(containerGroupProbeGrpc)
containerGroupReadinessProbe.SetExec(containerGroupProbeExec)
containerGroupReadinessProbe.SetInitialDelaySeconds(int64(123))
containerGroupReadinessProbe.SetPeriodSeconds(int64(123))
containerGroupReadinessProbe.SetTimeoutSeconds(int64(123))
containerGroupReadinessProbe.SetSuccessThreshold(int64(123))
containerGroupReadinessProbe.SetFailureThreshold(int64(123))


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{}
containerGroupProbeTcp.SetPort(int64(123))

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{}
containerGroupProbeHttpHeaders2.SetName("Name")
containerGroupProbeHttpHeaders2.SetValue("Value")

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{}
containerGroupProbeHttp.SetPath("Path")
containerGroupProbeHttp.SetPort(int64(123))
containerGroupProbeHttp.SetScheme(containerProbeHttpScheme)
containerGroupProbeHttp.SetHeaders([]shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2})


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{}
containerGroupProbeGrpc.SetService("Service")
containerGroupProbeGrpc.SetPort(int64(123))


containerGroupProbeExec := shared.ContainerGroupProbeExec{}
containerGroupProbeExec.SetCommand([]string{})

containerGroupStartupProbe := shared.ContainerGroupStartupProbe{}
containerGroupStartupProbe.SetTcp(containerGroupProbeTcp)
containerGroupStartupProbe.SetHttp(containerGroupProbeHttp)
containerGroupStartupProbe.SetGrpc(containerGroupProbeGrpc)
containerGroupStartupProbe.SetExec(containerGroupProbeExec)
containerGroupStartupProbe.SetInitialDelaySeconds(int64(123))
containerGroupStartupProbe.SetPeriodSeconds(int64(123))
containerGroupStartupProbe.SetTimeoutSeconds(int64(123))
containerGroupStartupProbe.SetSuccessThreshold(int64(123))
containerGroupStartupProbe.SetFailureThreshold(int64(123))


queueAutoscaler := shared.QueueAutoscaler{}
queueAutoscaler.SetMinReplicas(int64(123))
queueAutoscaler.SetMaxReplicas(int64(123))
queueAutoscaler.SetDesiredQueueLength(int64(123))
queueAutoscaler.SetPollingPeriod(int64(123))
queueAutoscaler.SetMaxUpscalePerMinute(int64(123))
queueAutoscaler.SetMaxDownscalePerMinute(int64(123))

request := containergroups.UpdateContainerGroup{}
request.SetDisplayName("DisplayName")
request.SetContainer(updateContainer)
request.SetReplicas(int64(123))
request.SetCountryCodes([]shared.CountryCode{countryCode})
request.SetNetworking(updateContainerGroupNetworking)
request.SetLivenessProbe(containerGroupLivenessProbe)
request.SetReadinessProbe(containerGroupReadinessProbe)
request.SetStartupProbe(containerGroupStartupProbe)
request.SetQueueAutoscaler(queueAutoscaler)

response, err := client.ContainerGroups.UpdateContainerGroup(context.Background(), "organizationName", "projectName", "containerGroupName", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
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

fmt.Println(response)
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

fmt.Println(response)
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

fmt.Println(response)
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

fmt.Println(response)
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

fmt.Println(response)
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

fmt.Println(response)
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

fmt.Println(response)
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

fmt.Println(response)
```
