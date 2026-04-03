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
