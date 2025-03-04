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
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/containergroups"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


containerResourceRequirements := shared.ContainerResourceRequirements{
  Cpu: util.ToPointer(int64(123)),
  Memory: util.ToPointer(int64(123)),
  GpuClasses: []string{},
  StorageAmount: util.ToPointer(util.Nullable[int64]{ Value: int64(123) }),
}

containerGroupPriority := shared.CONTAINER_GROUP_PRIORITY_HIGH


loggingAxiom2 := containergroups.LoggingAxiom2{
  Host: util.ToPointer("Host"),
  ApiToken: util.ToPointer("ApiToken"),
  Dataset: util.ToPointer("Dataset"),
}


datadogTags2 := containergroups.DatadogTags2{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

loggingDatadog2 := containergroups.LoggingDatadog2{
  Host: util.ToPointer("Host"),
  ApiKey: util.ToPointer("ApiKey"),
  Tags: []containergroups.DatadogTags2{datadogTags2},
}


loggingNewRelic2 := containergroups.LoggingNewRelic2{
  Host: util.ToPointer("Host"),
  IngestionKey: util.ToPointer("IngestionKey"),
}


loggingSplunk2 := containergroups.LoggingSplunk2{
  Host: util.ToPointer("Host"),
  Token: util.ToPointer("Token"),
}


loggingTcp2 := containergroups.LoggingTcp2{
  Host: util.ToPointer("Host"),
  Port: util.ToPointer(int64(123)),
}

httpFormat2 := containergroups.HTTP_FORMAT2_JSON


httpHeaders3 := containergroups.HttpHeaders3{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

httpCompression2 := containergroups.HTTP_COMPRESSION2_NONE

loggingHttp2 := containergroups.LoggingHttp2{
  Host: util.ToPointer("Host"),
  Port: util.ToPointer(int64(123)),
  User: util.ToPointer(util.Nullable[string]{ Value: "User" }),
  Password: util.ToPointer(util.Nullable[string]{ Value: "Password" }),
  Path: util.ToPointer(util.Nullable[string]{ Value: "Path" }),
  Format: &httpFormat2,
  Headers: []containergroups.HttpHeaders3{httpHeaders3},
  Compression: &httpCompression2,
}

createContainerLogging := containergroups.CreateContainerLogging{
  Axiom: &loggingAxiom2,
  Datadog: &loggingDatadog2,
  NewRelic: &loggingNewRelic2,
  Splunk: &loggingSplunk2,
  Tcp: &loggingTcp2,
  Http: &loggingHttp2,
}


registryAuthenticationBasic1 := containergroups.RegistryAuthenticationBasic1{
  Username: util.ToPointer("Username"),
  Password: util.ToPointer("Password"),
}


registryAuthenticationGcpGcr1 := containergroups.RegistryAuthenticationGcpGcr1{
  ServiceKey: util.ToPointer("ServiceKey"),
}


registryAuthenticationAwsEcr1 := containergroups.RegistryAuthenticationAwsEcr1{
  AccessKeyId: util.ToPointer("AccessKeyId"),
  SecretAccessKey: util.ToPointer("SecretAccessKey"),
}


registryAuthenticationDockerHub1 := containergroups.RegistryAuthenticationDockerHub1{
  Username: util.ToPointer("Username"),
  PersonalAccessToken: util.ToPointer("PersonalAccessToken"),
}


registryAuthenticationGcpGar1 := containergroups.RegistryAuthenticationGcpGar1{
  ServiceKey: util.ToPointer("ServiceKey"),
}

createContainerRegistryAuthentication := containergroups.CreateContainerRegistryAuthentication{
  Basic: &registryAuthenticationBasic1,
  GcpGcr: &registryAuthenticationGcpGcr1,
  AwsEcr: &registryAuthenticationAwsEcr1,
  DockerHub: &registryAuthenticationDockerHub1,
  GcpGar: &registryAuthenticationGcpGar1,
}

createContainer := containergroups.CreateContainer{
  Image: util.ToPointer("Image"),
  Resources: &containerResourceRequirements,
  Command: []string{},
  Priority: &containerGroupPriority,
  EnvironmentVariables: map[string]string{},
  Logging: &createContainerLogging,
  RegistryAuthentication: &createContainerRegistryAuthentication,
  ImageCaching: util.ToPointer(true),
}

containerRestartPolicy := shared.CONTAINER_RESTART_POLICY_ALWAYS

countryCode := shared.COUNTRY_CODE_AF

containerNetworkingProtocol := shared.CONTAINER_NETWORKING_PROTOCOL_HTTP

createContainerGroupNetworkingLoadBalancer := containergroups.CREATE_CONTAINER_GROUP_NETWORKING_LOAD_BALANCER_ROUND_ROBIN

createContainerGroupNetworking := containergroups.CreateContainerGroupNetworking{
  Protocol: &containerNetworkingProtocol,
  Port: util.ToPointer(int64(123)),
  Auth: util.ToPointer(true),
  LoadBalancer: &createContainerGroupNetworkingLoadBalancer,
  SingleConnectionLimit: util.ToPointer(true),
  ClientRequestTimeout: util.ToPointer(int64(123)),
  ServerResponseTimeout: util.ToPointer(int64(123)),
}


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{
  Port: util.ToPointer(int64(123)),
}

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  Scheme: &containerProbeHttpScheme,
  Headers: []shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2},
}


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{
  Service: util.ToPointer("Service"),
  Port: util.ToPointer(int64(123)),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}

containerGroupLivenessProbe := shared.ContainerGroupLivenessProbe{
  Tcp: &containerGroupProbeTcp,
  Http: &containerGroupProbeHttp,
  Grpc: &containerGroupProbeGrpc,
  Exec: &containerGroupProbeExec,
  InitialDelaySeconds: util.ToPointer(int64(123)),
  PeriodSeconds: util.ToPointer(int64(123)),
  TimeoutSeconds: util.ToPointer(int64(123)),
  SuccessThreshold: util.ToPointer(int64(123)),
  FailureThreshold: util.ToPointer(int64(123)),
}


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{
  Port: util.ToPointer(int64(123)),
}

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  Scheme: &containerProbeHttpScheme,
  Headers: []shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2},
}


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{
  Service: util.ToPointer("Service"),
  Port: util.ToPointer(int64(123)),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}

containerGroupReadinessProbe := shared.ContainerGroupReadinessProbe{
  Tcp: &containerGroupProbeTcp,
  Http: &containerGroupProbeHttp,
  Grpc: &containerGroupProbeGrpc,
  Exec: &containerGroupProbeExec,
  InitialDelaySeconds: util.ToPointer(int64(123)),
  PeriodSeconds: util.ToPointer(int64(123)),
  TimeoutSeconds: util.ToPointer(int64(123)),
  SuccessThreshold: util.ToPointer(int64(123)),
  FailureThreshold: util.ToPointer(int64(123)),
}


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{
  Port: util.ToPointer(int64(123)),
}

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  Scheme: &containerProbeHttpScheme,
  Headers: []shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2},
}


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{
  Service: util.ToPointer("Service"),
  Port: util.ToPointer(int64(123)),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}

containerGroupStartupProbe := shared.ContainerGroupStartupProbe{
  Tcp: &containerGroupProbeTcp,
  Http: &containerGroupProbeHttp,
  Grpc: &containerGroupProbeGrpc,
  Exec: &containerGroupProbeExec,
  InitialDelaySeconds: util.ToPointer(int64(123)),
  PeriodSeconds: util.ToPointer(int64(123)),
  TimeoutSeconds: util.ToPointer(int64(123)),
  SuccessThreshold: util.ToPointer(int64(123)),
  FailureThreshold: util.ToPointer(int64(123)),
}


containerGroupQueueConnection := shared.ContainerGroupQueueConnection{
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  QueueName: util.ToPointer("QueueName"),
}


queueAutoscaler := shared.QueueAutoscaler{
  MinReplicas: util.ToPointer(int64(123)),
  MaxReplicas: util.ToPointer(int64(123)),
  DesiredQueueLength: util.ToPointer(int64(123)),
  PollingPeriod: util.ToPointer(int64(123)),
  MaxUpscalePerMinute: util.ToPointer(int64(123)),
  MaxDownscalePerMinute: util.ToPointer(int64(123)),
}

request := containergroups.CreateContainerGroup{
  Name: util.ToPointer("Name"),
  DisplayName: util.ToPointer(util.Nullable[string]{ Value: "DisplayName" }),
  Container: &createContainer,
  AutostartPolicy: util.ToPointer(true),
  RestartPolicy: &containerRestartPolicy,
  Replicas: util.ToPointer(int64(123)),
  CountryCodes: []shared.CountryCode{countryCode},
  Networking: &createContainerGroupNetworking,
  LivenessProbe: &containerGroupLivenessProbe,
  ReadinessProbe: &containerGroupReadinessProbe,
  StartupProbe: &containerGroupStartupProbe,
  QueueConnection: &containerGroupQueueConnection,
  QueueAutoscaler: &queueAutoscaler,
}

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
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/containergroups"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


resources := containergroups.Resources{
  Cpu: util.ToPointer(util.Nullable[int64]{ Value: int64(123) }),
  Memory: util.ToPointer(util.Nullable[int64]{ Value: int64(123) }),
  GpuClasses: []string{},
  StorageAmount: util.ToPointer(util.Nullable[int64]{ Value: int64(123) }),
}

containerGroupPriority := shared.CONTAINER_GROUP_PRIORITY_HIGH


loggingAxiom3 := containergroups.LoggingAxiom3{
  Host: util.ToPointer("Host"),
  ApiToken: util.ToPointer("ApiToken"),
  Dataset: util.ToPointer("Dataset"),
}


datadogTags3 := containergroups.DatadogTags3{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

loggingDatadog3 := containergroups.LoggingDatadog3{
  Host: util.ToPointer("Host"),
  ApiKey: util.ToPointer("ApiKey"),
  Tags: []containergroups.DatadogTags3{datadogTags3},
}


loggingNewRelic3 := containergroups.LoggingNewRelic3{
  Host: util.ToPointer("Host"),
  IngestionKey: util.ToPointer("IngestionKey"),
}


loggingSplunk3 := containergroups.LoggingSplunk3{
  Host: util.ToPointer("Host"),
  Token: util.ToPointer("Token"),
}


loggingTcp3 := containergroups.LoggingTcp3{
  Host: util.ToPointer("Host"),
  Port: util.ToPointer(int64(123)),
}

httpFormat3 := containergroups.HTTP_FORMAT3_JSON


httpHeaders4 := containergroups.HttpHeaders4{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

httpCompression3 := containergroups.HTTP_COMPRESSION3_NONE

loggingHttp3 := containergroups.LoggingHttp3{
  Host: util.ToPointer("Host"),
  Port: util.ToPointer(int64(123)),
  User: util.ToPointer(util.Nullable[string]{ Value: "User" }),
  Password: util.ToPointer(util.Nullable[string]{ Value: "Password" }),
  Path: util.ToPointer(util.Nullable[string]{ Value: "Path" }),
  Format: &httpFormat3,
  Headers: []containergroups.HttpHeaders4{httpHeaders4},
  Compression: &httpCompression3,
}

updateContainerLogging := containergroups.UpdateContainerLogging{
  Axiom: &loggingAxiom3,
  Datadog: &loggingDatadog3,
  NewRelic: &loggingNewRelic3,
  Splunk: &loggingSplunk3,
  Tcp: &loggingTcp3,
  Http: &loggingHttp3,
}


registryAuthenticationBasic2 := containergroups.RegistryAuthenticationBasic2{
  Username: util.ToPointer("Username"),
  Password: util.ToPointer("Password"),
}


registryAuthenticationGcpGcr2 := containergroups.RegistryAuthenticationGcpGcr2{
  ServiceKey: util.ToPointer("ServiceKey"),
}


registryAuthenticationAwsEcr2 := containergroups.RegistryAuthenticationAwsEcr2{
  AccessKeyId: util.ToPointer("AccessKeyId"),
  SecretAccessKey: util.ToPointer("SecretAccessKey"),
}


registryAuthenticationDockerHub2 := containergroups.RegistryAuthenticationDockerHub2{
  Username: util.ToPointer("Username"),
  PersonalAccessToken: util.ToPointer("PersonalAccessToken"),
}


registryAuthenticationGcpGar2 := containergroups.RegistryAuthenticationGcpGar2{
  ServiceKey: util.ToPointer("ServiceKey"),
}

updateContainerRegistryAuthentication := containergroups.UpdateContainerRegistryAuthentication{
  Basic: &registryAuthenticationBasic2,
  GcpGcr: &registryAuthenticationGcpGcr2,
  AwsEcr: &registryAuthenticationAwsEcr2,
  DockerHub: &registryAuthenticationDockerHub2,
  GcpGar: &registryAuthenticationGcpGar2,
}

updateContainer := containergroups.UpdateContainer{
  Image: util.ToPointer(util.Nullable[string]{ Value: "Image" }),
  Resources: &resources,
  Command: []string{},
  Priority: &containerGroupPriority,
  EnvironmentVariables: map[string]string{},
  Logging: &updateContainerLogging,
  RegistryAuthentication: &updateContainerRegistryAuthentication,
  ImageCaching: util.ToPointer(true),
}

countryCode := shared.COUNTRY_CODE_AF


updateContainerGroupNetworking := containergroups.UpdateContainerGroupNetworking{
  Port: util.ToPointer(util.Nullable[int64]{ Value: int64(123) }),
}


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{
  Port: util.ToPointer(int64(123)),
}

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  Scheme: &containerProbeHttpScheme,
  Headers: []shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2},
}


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{
  Service: util.ToPointer("Service"),
  Port: util.ToPointer(int64(123)),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}

containerGroupLivenessProbe := shared.ContainerGroupLivenessProbe{
  Tcp: &containerGroupProbeTcp,
  Http: &containerGroupProbeHttp,
  Grpc: &containerGroupProbeGrpc,
  Exec: &containerGroupProbeExec,
  InitialDelaySeconds: util.ToPointer(int64(123)),
  PeriodSeconds: util.ToPointer(int64(123)),
  TimeoutSeconds: util.ToPointer(int64(123)),
  SuccessThreshold: util.ToPointer(int64(123)),
  FailureThreshold: util.ToPointer(int64(123)),
}


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{
  Port: util.ToPointer(int64(123)),
}

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  Scheme: &containerProbeHttpScheme,
  Headers: []shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2},
}


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{
  Service: util.ToPointer("Service"),
  Port: util.ToPointer(int64(123)),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}

containerGroupReadinessProbe := shared.ContainerGroupReadinessProbe{
  Tcp: &containerGroupProbeTcp,
  Http: &containerGroupProbeHttp,
  Grpc: &containerGroupProbeGrpc,
  Exec: &containerGroupProbeExec,
  InitialDelaySeconds: util.ToPointer(int64(123)),
  PeriodSeconds: util.ToPointer(int64(123)),
  TimeoutSeconds: util.ToPointer(int64(123)),
  SuccessThreshold: util.ToPointer(int64(123)),
  FailureThreshold: util.ToPointer(int64(123)),
}


containerGroupProbeTcp := shared.ContainerGroupProbeTcp{
  Port: util.ToPointer(int64(123)),
}

containerProbeHttpScheme := shared.CONTAINER_PROBE_HTTP_SCHEME_HTTP


containerGroupProbeHttpHeaders2 := shared.ContainerGroupProbeHttpHeaders2{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

containerGroupProbeHttp := shared.ContainerGroupProbeHttp{
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  Scheme: &containerProbeHttpScheme,
  Headers: []shared.ContainerGroupProbeHttpHeaders2{containerGroupProbeHttpHeaders2},
}


containerGroupProbeGrpc := shared.ContainerGroupProbeGrpc{
  Service: util.ToPointer("Service"),
  Port: util.ToPointer(int64(123)),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}

containerGroupStartupProbe := shared.ContainerGroupStartupProbe{
  Tcp: &containerGroupProbeTcp,
  Http: &containerGroupProbeHttp,
  Grpc: &containerGroupProbeGrpc,
  Exec: &containerGroupProbeExec,
  InitialDelaySeconds: util.ToPointer(int64(123)),
  PeriodSeconds: util.ToPointer(int64(123)),
  TimeoutSeconds: util.ToPointer(int64(123)),
  SuccessThreshold: util.ToPointer(int64(123)),
  FailureThreshold: util.ToPointer(int64(123)),
}


queueAutoscaler := shared.QueueAutoscaler{
  MinReplicas: util.ToPointer(int64(123)),
  MaxReplicas: util.ToPointer(int64(123)),
  DesiredQueueLength: util.ToPointer(int64(123)),
  PollingPeriod: util.ToPointer(int64(123)),
  MaxUpscalePerMinute: util.ToPointer(int64(123)),
  MaxDownscalePerMinute: util.ToPointer(int64(123)),
}

request := containergroups.UpdateContainerGroup{
  DisplayName: util.ToPointer(util.Nullable[string]{ Value: "DisplayName" }),
  Container: &updateContainer,
  Replicas: util.ToPointer(util.Nullable[int64]{ Value: int64(123) }),
  CountryCodes: []shared.CountryCode{countryCode},
  Networking: &updateContainerGroupNetworking,
  LivenessProbe: &containerGroupLivenessProbe,
  ReadinessProbe: &containerGroupReadinessProbe,
  StartupProbe: &containerGroupStartupProbe,
  QueueAutoscaler: &queueAutoscaler,
}

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
