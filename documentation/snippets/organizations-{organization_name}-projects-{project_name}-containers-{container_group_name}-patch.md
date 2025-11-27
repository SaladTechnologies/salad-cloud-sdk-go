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
  Host: util.ToPointer("host"),
  ApiToken: util.ToPointer("api_token"),
  Dataset: util.ToPointer("dataset"),
}


datadogTagForContainerLogging := shared.DatadogTagForContainerLogging{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

datadogLoggingConfiguration := shared.DatadogLoggingConfiguration{
  Host: util.ToPointer("host"),
  ApiKey: util.ToPointer("api_key"),
  Tags: []shared.DatadogTagForContainerLogging{datadogTagForContainerLogging},
}

containerLoggingHttpFormat := shared.CONTAINER_LOGGING_HTTP_FORMAT_JSON


containerLoggingHttpHeader := shared.ContainerLoggingHttpHeader{
  Name: util.ToPointer("name"),
  Value: util.ToPointer("value"),
}

containerLoggingHttpCompression := shared.CONTAINER_LOGGING_HTTP_COMPRESSION_NONE

containerLoggingConfigurationHttp1 := shared.ContainerLoggingConfigurationHttp1{
  Host: util.ToPointer("host"),
  Port: util.ToPointer(int64(55354)),
  User: util.ToPointer(util.Nullable[string]{ Value: "user" }),
  Password: util.ToPointer(util.Nullable[string]{ Value: "password" }),
  Path: util.ToPointer(util.Nullable[string]{ Value: "path" }),
  Format: &containerLoggingHttpFormat,
  Headers: []shared.ContainerLoggingHttpHeader{containerLoggingHttpHeader},
  Compression: &containerLoggingHttpCompression,
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
  Username: util.ToPointer("username"),
  Password: util.ToPointer("password"),
}


containerRegistryAuthenticationDockerHub := containergroups.ContainerRegistryAuthenticationDockerHub{
  Username: util.ToPointer("username"),
  PersonalAccessToken: util.ToPointer("personal_access_token"),
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
  Cpu: util.ToPointer(util.Nullable[int64]{ Value: int64(1013) }),
  Memory: util.ToPointer(util.Nullable[int64]{ Value: int64(352043675) }),
  GpuClasses: []string{},
  StorageAmount: util.ToPointer(util.Nullable[int64]{ Value: int64(1032076497908566.1) }),
  ShmSize: util.ToPointer(util.Nullable[int64]{ Value: int64(64) }),
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


updateContainerGroupNetworking := containergroups.UpdateContainerGroupNetworking{
  Port: util.ToPointer(util.Nullable[int64]{ Value: int64(13142) }),
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
  InitialDelaySeconds: util.ToPointer(int64(1106)),
  Tcp: &containerGroupTcpProbe,
  PeriodSeconds: util.ToPointer(int64(3)),
  SuccessThreshold: util.ToPointer(int64(2)),
  TimeoutSeconds: util.ToPointer(int64(10)),
}


queueBasedAutoscalerConfiguration := shared.QueueBasedAutoscalerConfiguration{
  DesiredQueueLength: util.ToPointer(int64(53)),
  MaxReplicas: util.ToPointer(int64(291)),
  MaxDownscalePerMinute: util.ToPointer(int64(65)),
  MaxUpscalePerMinute: util.ToPointer(int64(100)),
  MinReplicas: util.ToPointer(int64(54)),
  PollingPeriod: util.ToPointer(int64(140)),
}

request := containergroups.ContainerGroupPatch{
  DisplayName: util.ToPointer(util.Nullable[string]{ Value: "ZJjdnvu" }),
  Container: &updateContainer,
  Replicas: util.ToPointer(util.Nullable[int64]{ Value: int64(56) }),
  CountryCodes: []shared.CountryCode{countryCode},
  Networking: &updateContainerGroupNetworking,
  LivenessProbe: &containerGroupLivenessProbe,
  ReadinessProbe: &containerGroupReadinessProbe,
  StartupProbe: &containerGroupStartupProbe,
  QueueAutoscaler: &queueBasedAutoscalerConfiguration,
}

response, err := client.ContainerGroups.UpdateContainerGroup(context.Background(), "acme-corp", "dev-env", "mandlebrot", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
