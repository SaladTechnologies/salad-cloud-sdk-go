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
