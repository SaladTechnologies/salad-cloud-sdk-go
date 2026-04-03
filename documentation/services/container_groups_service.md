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
| [UpdateContainerGroupInstance](#updatecontainergroupinstance)         | Updates a container group instance                                                                                          |
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

`ContainerGroupCollection`

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

response, err := client.ContainerGroups.ListContainerGroups(context.Background(), "acme-corp", "dev-env")
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

| Name                          | Type                          | Required | Description                                                                                                                                                                                                                                         |
| :---------------------------- | :---------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                           | Context                       | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName              | string                        | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName                   | string                        | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupCreationRequest | ContainerGroupCreationRequest | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`shared.ContainerGroup`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/containergroups"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


axiomLoggingConfiguration := shared.AxiomLoggingConfiguration{
  ApiToken: util.ToPointer("api_token"),
  Dataset: util.ToPointer("dataset"),
  Host: util.ToPointer("host"),
}


datadogTagForContainerLogging := shared.DatadogTagForContainerLogging{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

datadogLoggingConfiguration := shared.DatadogLoggingConfiguration{
  ApiKey: util.ToPointer("api_key"),
  Host: util.ToPointer("host"),
  Tags: []shared.DatadogTagForContainerLogging{datadogTagForContainerLogging},
}

containerLoggingHttpCompression := shared.CONTAINER_LOGGING_HTTP_COMPRESSION_NONE

containerLoggingHttpFormat := shared.CONTAINER_LOGGING_HTTP_FORMAT_JSON


containerLoggingHttpHeader := shared.ContainerLoggingHttpHeader{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

containerLoggingConfigurationHttp2 := containergroups.ContainerLoggingConfigurationHttp2{
  Compression: &containerLoggingHttpCompression,
  Format: &containerLoggingHttpFormat,
  Headers: []shared.ContainerLoggingHttpHeader{containerLoggingHttpHeader},
  Host: util.ToPointer("host"),
  Password: util.ToPointer(util.Nullable[string]{ Value: "password" }),
  Path: util.ToPointer(util.Nullable[string]{ Value: "path" }),
  Port: util.ToPointer(int64(42056)),
  User: util.ToPointer(util.Nullable[string]{ Value: "user" }),
}


newRelicLoggingConfiguration := shared.NewRelicLoggingConfiguration{
  Host: util.ToPointer("host"),
  IngestionKey: util.ToPointer("ingestion_key"),
}


containerLoggingSplunkConfiguration := shared.ContainerLoggingSplunkConfiguration{
  Host: util.ToPointer("host"),
  Token: util.ToPointer("token"),
}


tcpLoggingConfiguration := shared.TcpLoggingConfiguration{
  Host: util.ToPointer("host"),
  Port: util.ToPointer(int64(44671)),
}

containerConfigurationLogging := containergroups.ContainerConfigurationLogging{
  Axiom: &axiomLoggingConfiguration,
  Datadog: &datadogLoggingConfiguration,
  Http: &containerLoggingConfigurationHttp2,
  NewRelic: &newRelicLoggingConfiguration,
  Splunk: &containerLoggingSplunkConfiguration,
  Tcp: &tcpLoggingConfiguration,
}

containerGroupPriority := shared.CONTAINER_GROUP_PRIORITY_HIGH


containerRegistryAuthenticationAwsEcr := containergroups.ContainerRegistryAuthenticationAwsEcr{
  AccessKeyId: util.ToPointer("access_key_id"),
  SecretAccessKey: util.ToPointer("secret_access_key"),
}


containerRegistryAuthenticationBasic := containergroups.ContainerRegistryAuthenticationBasic{
  Password: util.ToPointer("password"),
  Username: util.ToPointer("username"),
}


containerRegistryAuthenticationDockerHub := containergroups.ContainerRegistryAuthenticationDockerHub{
  PersonalAccessToken: util.ToPointer("personal_access_token"),
  Username: util.ToPointer("username"),
}


containerRegistryAuthenticationGcpGar := containergroups.ContainerRegistryAuthenticationGcpGar{
  ServiceKey: util.ToPointer("service_key"),
}


containerRegistryAuthenticationGcpGcr := containergroups.ContainerRegistryAuthenticationGcpGcr{
  ServiceKey: util.ToPointer("service_key"),
}

containerRegistryAuthentication := containergroups.ContainerRegistryAuthentication{
  AwsEcr: &containerRegistryAuthenticationAwsEcr,
  Basic: &containerRegistryAuthenticationBasic,
  DockerHub: &containerRegistryAuthenticationDockerHub,
  GcpGar: &containerRegistryAuthenticationGcpGar,
  GcpGcr: &containerRegistryAuthenticationGcpGcr,
}


createContainerResourceRequirements := containergroups.CreateContainerResourceRequirements{
  Cpu: util.ToPointer(int64(827)),
  GpuClasses: []string{},
  Memory: util.ToPointer(int64(734164836)),
  ShmSize: util.ToPointer(int64(64)),
  StorageAmount: util.ToPointer(int64(761306530849177.9)),
}

containerConfiguration := containergroups.ContainerConfiguration{
  Command: []string{},
  EnvironmentVariables: map[string]string{},
  Image: util.ToPointer("acme/:latest"),
  ImageCaching: util.ToPointer(true),
  Logging: &containerConfigurationLogging,
  Priority: &containerGroupPriority,
  RegistryAuthentication: &containerRegistryAuthentication,
  Resources: &createContainerResourceRequirements,
}

countryCode := shared.COUNTRY_CODE_AF


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}


containerGroupGRpcProbe := shared.ContainerGroupGRpcProbe{
  Port: util.ToPointer(int64(37648)),
  Service: util.ToPointer("service"),
}


containerGroupProbeHttpHeader := shared.ContainerGroupProbeHttpHeader{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

httpScheme := shared.HTTP_SCHEME_HTTP

containerGroupHttpProbeConfiguration := shared.ContainerGroupHttpProbeConfiguration{
  Headers: []shared.ContainerGroupProbeHttpHeader{containerGroupProbeHttpHeader},
  Path: util.ToPointer("path"),
  Port: util.ToPointer(int64(29069)),
  Scheme: &httpScheme,
}


containerGroupTcpProbe := shared.ContainerGroupTcpProbe{
  Port: util.ToPointer(int64(13817)),
}

containerGroupLivenessProbe := shared.ContainerGroupLivenessProbe{
  Exec: &containerGroupProbeExec,
  FailureThreshold: util.ToPointer(int64(3)),
  Grpc: &containerGroupGRpcProbe,
  Http: &containerGroupHttpProbeConfiguration,
  InitialDelaySeconds: util.ToPointer(int64(670)),
  PeriodSeconds: util.ToPointer(int64(10)),
  SuccessThreshold: util.ToPointer(int64(1)),
  Tcp: &containerGroupTcpProbe,
  TimeoutSeconds: util.ToPointer(int64(30)),
}

theContainerGroupNetworkingLoadBalancer := shared.THE_CONTAINER_GROUP_NETWORKING_LOAD_BALANCER_ROUND_ROBIN

containerNetworkingProtocol := shared.CONTAINER_NETWORKING_PROTOCOL_HTTP

createContainerGroupNetworking := containergroups.CreateContainerGroupNetworking{
  Auth: util.ToPointer(true),
  ClientRequestTimeout: util.ToPointer(int64(100000)),
  LoadBalancer: &theContainerGroupNetworkingLoadBalancer,
  Port: util.ToPointer(int64(60000)),
  Protocol: &containerNetworkingProtocol,
  ServerResponseTimeout: util.ToPointer(int64(100000)),
  SingleConnectionLimit: util.ToPointer(true),
}


queueBasedAutoscalerConfiguration := shared.QueueBasedAutoscalerConfiguration{
  DesiredQueueLength: util.ToPointer(int64(53)),
  MaxDownscalePerMinute: util.ToPointer(int64(59)),
  MaxReplicas: util.ToPointer(int64(321)),
  MaxUpscalePerMinute: util.ToPointer(int64(100)),
  MinReplicas: util.ToPointer(int64(54)),
  PollingPeriod: util.ToPointer(int64(140)),
}


containerGroupQueueConnection := shared.ContainerGroupQueueConnection{
  Path: util.ToPointer("path"),
  Port: util.ToPointer(int64(47568)),
  QueueName: util.ToPointer("z1h-3z01x9"),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}


containerGroupGRpcProbe := shared.ContainerGroupGRpcProbe{
  Port: util.ToPointer(int64(37648)),
  Service: util.ToPointer("service"),
}


containerGroupProbeHttpHeader := shared.ContainerGroupProbeHttpHeader{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

httpScheme := shared.HTTP_SCHEME_HTTP

containerGroupHttpProbeConfiguration := shared.ContainerGroupHttpProbeConfiguration{
  Headers: []shared.ContainerGroupProbeHttpHeader{containerGroupProbeHttpHeader},
  Path: util.ToPointer("path"),
  Port: util.ToPointer(int64(29069)),
  Scheme: &httpScheme,
}


containerGroupTcpProbe := shared.ContainerGroupTcpProbe{
  Port: util.ToPointer(int64(13817)),
}

containerGroupReadinessProbe := shared.ContainerGroupReadinessProbe{
  Exec: &containerGroupProbeExec,
  FailureThreshold: util.ToPointer(int64(3)),
  Grpc: &containerGroupGRpcProbe,
  Http: &containerGroupHttpProbeConfiguration,
  InitialDelaySeconds: util.ToPointer(int64(262)),
  PeriodSeconds: util.ToPointer(int64(1)),
  SuccessThreshold: util.ToPointer(int64(1)),
  Tcp: &containerGroupTcpProbe,
  TimeoutSeconds: util.ToPointer(int64(1)),
}

containerRestartPolicy := shared.CONTAINER_RESTART_POLICY_ALWAYS


containerGroupScalingAction := shared.ContainerGroupScalingAction{
  Replicas: util.ToPointer(int64(461)),
  Schedule: util.ToPointer("7kwC/T8C   da       x6Ci   bM-rgGYn     bDY6,vT"),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}


containerGroupGRpcProbe := shared.ContainerGroupGRpcProbe{
  Port: util.ToPointer(int64(37648)),
  Service: util.ToPointer("service"),
}


containerGroupProbeHttpHeader := shared.ContainerGroupProbeHttpHeader{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

httpScheme := shared.HTTP_SCHEME_HTTP

containerGroupHttpProbeConfiguration := shared.ContainerGroupHttpProbeConfiguration{
  Headers: []shared.ContainerGroupProbeHttpHeader{containerGroupProbeHttpHeader},
  Path: util.ToPointer("path"),
  Port: util.ToPointer(int64(29069)),
  Scheme: &httpScheme,
}


containerGroupTcpProbe := shared.ContainerGroupTcpProbe{
  Port: util.ToPointer(int64(13817)),
}

containerGroupStartupProbe := shared.ContainerGroupStartupProbe{
  Exec: &containerGroupProbeExec,
  FailureThreshold: util.ToPointer(int64(15)),
  Grpc: &containerGroupGRpcProbe,
  Http: &containerGroupHttpProbeConfiguration,
  InitialDelaySeconds: util.ToPointer(int64(503)),
  PeriodSeconds: util.ToPointer(int64(3)),
  SuccessThreshold: util.ToPointer(int64(2)),
  Tcp: &containerGroupTcpProbe,
  TimeoutSeconds: util.ToPointer(int64(10)),
}

request := containergroups.ContainerGroupCreationRequest{
  AutostartPolicy: util.ToPointer(true),
  Container: &containerConfiguration,
  CountryCodes: []shared.CountryCode{countryCode},
  DisplayName: util.ToPointer("KMg0KyVwpb"),
  LivenessProbe: &containerGroupLivenessProbe,
  Name: util.ToPointer("name"),
  Networking: &createContainerGroupNetworking,
  QueueAutoscaler: &queueBasedAutoscalerConfiguration,
  QueueConnection: &containerGroupQueueConnection,
  ReadinessProbe: &containerGroupReadinessProbe,
  Replicas: util.ToPointer(int64(77)),
  RestartPolicy: &containerRestartPolicy,
  ScalingActions: []shared.ContainerGroupScalingAction{containerGroupScalingAction},
  ScheduledScalingEnabled: util.ToPointer(true),
  StartupProbe: &containerGroupStartupProbe,
}

response, err := client.ContainerGroups.CreateContainerGroup(context.Background(), "acme-corp", "dev-env", request)
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
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.ContainerGroups.GetContainerGroup(context.Background(), "acme-corp", "dev-env", "mandlebrot")
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

| Name                | Type                | Required | Description                                                                                                                                                                                                                                         |
| :------------------ | :------------------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                 | Context             | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName    | string              | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName         | string              | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName  | string              | ✅       | The unique container group name                                                                                                                                                                                                                     |
| containerGroupPatch | ContainerGroupPatch | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`shared.ContainerGroup`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/containergroups"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


axiomLoggingConfiguration := shared.AxiomLoggingConfiguration{
  ApiToken: util.ToPointer("api_token"),
  Dataset: util.ToPointer("dataset"),
  Host: util.ToPointer("host"),
}


datadogTagForContainerLogging := shared.DatadogTagForContainerLogging{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

datadogLoggingConfiguration := shared.DatadogLoggingConfiguration{
  ApiKey: util.ToPointer("api_key"),
  Host: util.ToPointer("host"),
  Tags: []shared.DatadogTagForContainerLogging{datadogTagForContainerLogging},
}

containerLoggingHttpCompression := shared.CONTAINER_LOGGING_HTTP_COMPRESSION_NONE

containerLoggingHttpFormat := shared.CONTAINER_LOGGING_HTTP_FORMAT_JSON


containerLoggingHttpHeader := shared.ContainerLoggingHttpHeader{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

containerLoggingConfigurationHttp1 := shared.ContainerLoggingConfigurationHttp1{
  Compression: &containerLoggingHttpCompression,
  Format: &containerLoggingHttpFormat,
  Headers: []shared.ContainerLoggingHttpHeader{containerLoggingHttpHeader},
  Host: util.ToPointer("host"),
  Password: util.ToPointer(util.Nullable[string]{ Value: "password" }),
  Path: util.ToPointer(util.Nullable[string]{ Value: "path" }),
  Port: util.ToPointer(int64(55354)),
  User: util.ToPointer(util.Nullable[string]{ Value: "user" }),
}


newRelicLoggingConfiguration := shared.NewRelicLoggingConfiguration{
  Host: util.ToPointer("host"),
  IngestionKey: util.ToPointer("ingestion_key"),
}


containerLoggingSplunkConfiguration := shared.ContainerLoggingSplunkConfiguration{
  Host: util.ToPointer("host"),
  Token: util.ToPointer("token"),
}


tcpLoggingConfiguration := shared.TcpLoggingConfiguration{
  Host: util.ToPointer("host"),
  Port: util.ToPointer(int64(44671)),
}

updateContainerLogging := containergroups.UpdateContainerLogging{
  Axiom: &axiomLoggingConfiguration,
  Datadog: &datadogLoggingConfiguration,
  Http: &containerLoggingConfigurationHttp1,
  NewRelic: &newRelicLoggingConfiguration,
  Splunk: &containerLoggingSplunkConfiguration,
  Tcp: &tcpLoggingConfiguration,
}

containerGroupPriority := shared.CONTAINER_GROUP_PRIORITY_HIGH


containerRegistryAuthenticationAwsEcr := containergroups.ContainerRegistryAuthenticationAwsEcr{
  AccessKeyId: util.ToPointer("access_key_id"),
  SecretAccessKey: util.ToPointer("secret_access_key"),
}


containerRegistryAuthenticationBasic := containergroups.ContainerRegistryAuthenticationBasic{
  Password: util.ToPointer("password"),
  Username: util.ToPointer("username"),
}


containerRegistryAuthenticationDockerHub := containergroups.ContainerRegistryAuthenticationDockerHub{
  PersonalAccessToken: util.ToPointer("personal_access_token"),
  Username: util.ToPointer("username"),
}


containerRegistryAuthenticationGcpGar := containergroups.ContainerRegistryAuthenticationGcpGar{
  ServiceKey: util.ToPointer("service_key"),
}


containerRegistryAuthenticationGcpGcr := containergroups.ContainerRegistryAuthenticationGcpGcr{
  ServiceKey: util.ToPointer("service_key"),
}

containerRegistryAuthentication := containergroups.ContainerRegistryAuthentication{
  AwsEcr: &containerRegistryAuthenticationAwsEcr,
  Basic: &containerRegistryAuthenticationBasic,
  DockerHub: &containerRegistryAuthenticationDockerHub,
  GcpGar: &containerRegistryAuthenticationGcpGar,
  GcpGcr: &containerRegistryAuthenticationGcpGcr,
}


containerResourceUpdateSchema := containergroups.ContainerResourceUpdateSchema{
  Cpu: util.ToPointer(util.Nullable[int64]{ Value: int64(191) }),
  GpuClasses: []string{},
  Memory: util.ToPointer(util.Nullable[int64]{ Value: int64(909130690) }),
  ShmSize: util.ToPointer(util.Nullable[int64]{ Value: int64(64) }),
  StorageAmount: util.ToPointer(util.Nullable[int64]{ Value: int64(699558298076245) }),
}

updateContainer := containergroups.UpdateContainer{
  Command: []string{},
  EnvironmentVariables: map[string]string{},
  Image: util.ToPointer(util.Nullable[string]{ Value: "image" }),
  ImageCaching: util.ToPointer(true),
  Logging: &updateContainerLogging,
  Priority: &containerGroupPriority,
  RegistryAuthentication: &containerRegistryAuthentication,
  Resources: &containerResourceUpdateSchema,
}

countryCode := shared.COUNTRY_CODE_AF


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}


containerGroupGRpcProbe := shared.ContainerGroupGRpcProbe{
  Port: util.ToPointer(int64(37648)),
  Service: util.ToPointer("service"),
}


containerGroupProbeHttpHeader := shared.ContainerGroupProbeHttpHeader{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

httpScheme := shared.HTTP_SCHEME_HTTP

containerGroupHttpProbeConfiguration := shared.ContainerGroupHttpProbeConfiguration{
  Headers: []shared.ContainerGroupProbeHttpHeader{containerGroupProbeHttpHeader},
  Path: util.ToPointer("path"),
  Port: util.ToPointer(int64(29069)),
  Scheme: &httpScheme,
}


containerGroupTcpProbe := shared.ContainerGroupTcpProbe{
  Port: util.ToPointer(int64(13817)),
}

containerGroupLivenessProbe := shared.ContainerGroupLivenessProbe{
  Exec: &containerGroupProbeExec,
  FailureThreshold: util.ToPointer(int64(3)),
  Grpc: &containerGroupGRpcProbe,
  Http: &containerGroupHttpProbeConfiguration,
  InitialDelaySeconds: util.ToPointer(int64(670)),
  PeriodSeconds: util.ToPointer(int64(10)),
  SuccessThreshold: util.ToPointer(int64(1)),
  Tcp: &containerGroupTcpProbe,
  TimeoutSeconds: util.ToPointer(int64(30)),
}


updateContainerGroupNetworking := containergroups.UpdateContainerGroupNetworking{
  Port: util.ToPointer(util.Nullable[int64]{ Value: int64(17025) }),
}


queueBasedAutoscalerConfiguration := shared.QueueBasedAutoscalerConfiguration{
  DesiredQueueLength: util.ToPointer(int64(53)),
  MaxDownscalePerMinute: util.ToPointer(int64(59)),
  MaxReplicas: util.ToPointer(int64(321)),
  MaxUpscalePerMinute: util.ToPointer(int64(100)),
  MinReplicas: util.ToPointer(int64(54)),
  PollingPeriod: util.ToPointer(int64(140)),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}


containerGroupGRpcProbe := shared.ContainerGroupGRpcProbe{
  Port: util.ToPointer(int64(37648)),
  Service: util.ToPointer("service"),
}


containerGroupProbeHttpHeader := shared.ContainerGroupProbeHttpHeader{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

httpScheme := shared.HTTP_SCHEME_HTTP

containerGroupHttpProbeConfiguration := shared.ContainerGroupHttpProbeConfiguration{
  Headers: []shared.ContainerGroupProbeHttpHeader{containerGroupProbeHttpHeader},
  Path: util.ToPointer("path"),
  Port: util.ToPointer(int64(29069)),
  Scheme: &httpScheme,
}


containerGroupTcpProbe := shared.ContainerGroupTcpProbe{
  Port: util.ToPointer(int64(13817)),
}

containerGroupReadinessProbe := shared.ContainerGroupReadinessProbe{
  Exec: &containerGroupProbeExec,
  FailureThreshold: util.ToPointer(int64(3)),
  Grpc: &containerGroupGRpcProbe,
  Http: &containerGroupHttpProbeConfiguration,
  InitialDelaySeconds: util.ToPointer(int64(262)),
  PeriodSeconds: util.ToPointer(int64(1)),
  SuccessThreshold: util.ToPointer(int64(1)),
  Tcp: &containerGroupTcpProbe,
  TimeoutSeconds: util.ToPointer(int64(1)),
}


containerGroupScalingAction := shared.ContainerGroupScalingAction{
  Replicas: util.ToPointer(int64(461)),
  Schedule: util.ToPointer("7kwC/T8C   da       x6Ci   bM-rgGYn     bDY6,vT"),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}


containerGroupGRpcProbe := shared.ContainerGroupGRpcProbe{
  Port: util.ToPointer(int64(37648)),
  Service: util.ToPointer("service"),
}


containerGroupProbeHttpHeader := shared.ContainerGroupProbeHttpHeader{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

httpScheme := shared.HTTP_SCHEME_HTTP

containerGroupHttpProbeConfiguration := shared.ContainerGroupHttpProbeConfiguration{
  Headers: []shared.ContainerGroupProbeHttpHeader{containerGroupProbeHttpHeader},
  Path: util.ToPointer("path"),
  Port: util.ToPointer(int64(29069)),
  Scheme: &httpScheme,
}


containerGroupTcpProbe := shared.ContainerGroupTcpProbe{
  Port: util.ToPointer(int64(13817)),
}

containerGroupStartupProbe := shared.ContainerGroupStartupProbe{
  Exec: &containerGroupProbeExec,
  FailureThreshold: util.ToPointer(int64(15)),
  Grpc: &containerGroupGRpcProbe,
  Http: &containerGroupHttpProbeConfiguration,
  InitialDelaySeconds: util.ToPointer(int64(503)),
  PeriodSeconds: util.ToPointer(int64(3)),
  SuccessThreshold: util.ToPointer(int64(2)),
  Tcp: &containerGroupTcpProbe,
  TimeoutSeconds: util.ToPointer(int64(10)),
}

request := containergroups.ContainerGroupPatch{
  Container: &updateContainer,
  CountryCodes: []shared.CountryCode{countryCode},
  DisplayName: util.ToPointer(util.Nullable[string]{ Value: "ykW8yoj HBQ" }),
  LivenessProbe: &containerGroupLivenessProbe,
  Networking: &updateContainerGroupNetworking,
  QueueAutoscaler: &queueBasedAutoscalerConfiguration,
  ReadinessProbe: &containerGroupReadinessProbe,
  Replicas: util.ToPointer(util.Nullable[int64]{ Value: int64(102) }),
  ScalingActions: []shared.ContainerGroupScalingAction{containerGroupScalingAction},
  ScheduledScalingEnabled: util.ToPointer(true),
  StartupProbe: &containerGroupStartupProbe,
}

response, err := client.ContainerGroups.UpdateContainerGroup(context.Background(), "acme-corp", "dev-env", "mandlebrot", request)
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
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.ContainerGroups.DeleteContainerGroup(context.Background(), "acme-corp", "dev-env", "mandlebrot")
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
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.ContainerGroups.StartContainerGroup(context.Background(), "acme-corp", "dev-env", "mandlebrot")
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
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.ContainerGroups.StopContainerGroup(context.Background(), "acme-corp", "dev-env", "mandlebrot")
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

`ContainerGroupInstanceCollection`

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

response, err := client.ContainerGroups.ListContainerGroupInstances(context.Background(), "acme-corp", "dev-env", "mandlebrot")
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
| containerGroupInstanceId | string  | ✅       | The unique container group instance identifier                                                                                                                                                                                                      |

**Return Type**

`ContainerGroupInstance`

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

response, err := client.ContainerGroups.GetContainerGroupInstance(context.Background(), "acme-corp", "dev-env", "mandlebrot", "db3a4591-efc3-46c0-b06a-3d820c0ec100")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## UpdateContainerGroupInstance

Updates a container group instance

- HTTP Method: `PATCH`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}`

**Parameters**

| Name                        | Type                        | Required | Description                                                                                                                                                                                                                                         |
| :-------------------------- | :-------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                         | Context                     | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName            | string                      | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName                 | string                      | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName          | string                      | ✅       | The unique container group name                                                                                                                                                                                                                     |
| containerGroupInstanceId    | string                      | ✅       | The unique container group instance identifier                                                                                                                                                                                                      |
| containerGroupInstancePatch | ContainerGroupInstancePatch | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`ContainerGroupInstance`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/containergroups"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


request := containergroups.ContainerGroupInstancePatch{
  DeletionCost: util.ToPointer(util.Nullable[int64]{ Value: int64(19725) }),
}

response, err := client.ContainerGroups.UpdateContainerGroupInstance(context.Background(), "acme-corp", "dev-env", "mandlebrot", "db3a4591-efc3-46c0-b06a-3d820c0ec100", request)
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
| containerGroupInstanceId | string  | ✅       | The unique container group instance identifier                                                                                                                                                                                                      |

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

response, err := client.ContainerGroups.ReallocateContainerGroupInstance(context.Background(), "acme-corp", "dev-env", "mandlebrot", "db3a4591-efc3-46c0-b06a-3d820c0ec100")
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
| containerGroupInstanceId | string  | ✅       | The unique container group instance identifier                                                                                                                                                                                                      |

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

response, err := client.ContainerGroups.RecreateContainerGroupInstance(context.Background(), "acme-corp", "dev-env", "mandlebrot", "db3a4591-efc3-46c0-b06a-3d820c0ec100")
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
| containerGroupInstanceId | string  | ✅       | The unique container group instance identifier                                                                                                                                                                                                      |

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

response, err := client.ContainerGroups.RestartContainerGroupInstance(context.Background(), "acme-corp", "dev-env", "mandlebrot", "db3a4591-efc3-46c0-b06a-3d820c0ec100")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
