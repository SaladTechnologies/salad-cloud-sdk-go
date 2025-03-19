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


axiomLoggingConfiguration := shared.AxiomLoggingConfiguration{
  Host: util.ToPointer("Host"),
  ApiToken: util.ToPointer("ApiToken"),
  Dataset: util.ToPointer("Dataset"),
}


datadogTagForContainerLogging := shared.DatadogTagForContainerLogging{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

datadogLoggingConfiguration := shared.DatadogLoggingConfiguration{
  Host: util.ToPointer("Host"),
  ApiKey: util.ToPointer("ApiKey"),
  Tags: []shared.DatadogTagForContainerLogging{datadogTagForContainerLogging},
}

format := shared.FORMAT_JSON


containerLoggingHttpHeader := shared.ContainerLoggingHttpHeader{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

compression := shared.COMPRESSION_NONE

containerHttpLoggingConfiguration := shared.ContainerHttpLoggingConfiguration{
  Host: util.ToPointer("Host"),
  Port: util.ToPointer(int64(123)),
  User: util.ToPointer(util.Nullable[string]{ Value: "User" }),
  Password: util.ToPointer(util.Nullable[string]{ Value: "Password" }),
  Path: util.ToPointer(util.Nullable[string]{ Value: "Path" }),
  Format: &format,
  Headers: []shared.ContainerLoggingHttpHeader{containerLoggingHttpHeader},
  Compression: &compression,
}


newRelicLoggingConfiguration := shared.NewRelicLoggingConfiguration{
  Host: util.ToPointer("Host"),
  IngestionKey: util.ToPointer("IngestionKey"),
}


containerLoggingSplunkConfiguration := shared.ContainerLoggingSplunkConfiguration{
  Host: util.ToPointer("Host"),
  Token: util.ToPointer("Token"),
}


tcpLoggingConfiguration := shared.TcpLoggingConfiguration{
  Host: util.ToPointer("Host"),
  Port: util.ToPointer(int64(123)),
}

containerLoggingConfiguration := shared.ContainerLoggingConfiguration{
  Axiom: &axiomLoggingConfiguration,
  Datadog: &datadogLoggingConfiguration,
  Http: &containerHttpLoggingConfiguration,
  NewRelic: &newRelicLoggingConfiguration,
  Splunk: &containerLoggingSplunkConfiguration,
  Tcp: &tcpLoggingConfiguration,
}

containerGroupPriority := shared.CONTAINER_GROUP_PRIORITY_HIGH


containerRegistryAuthenticationAwsEcr := containergroups.ContainerRegistryAuthenticationAwsEcr{
  AccessKeyId: util.ToPointer("AccessKeyId"),
  SecretAccessKey: util.ToPointer("SecretAccessKey"),
}


containerRegistryAuthenticationBasic := containergroups.ContainerRegistryAuthenticationBasic{
  Username: util.ToPointer("Username"),
  Password: util.ToPointer("Password"),
}


containerRegistryAuthenticationDockerHub := containergroups.ContainerRegistryAuthenticationDockerHub{
  Username: util.ToPointer("Username"),
  PersonalAccessToken: util.ToPointer("PersonalAccessToken"),
}


containerRegistryAuthenticationGcpGar := containergroups.ContainerRegistryAuthenticationGcpGar{
  ServiceKey: util.ToPointer("ServiceKey"),
}


containerRegistryAuthenticationGcpGcr := containergroups.ContainerRegistryAuthenticationGcpGcr{
  ServiceKey: util.ToPointer("ServiceKey"),
}

containerRegistryAuthentication := containergroups.ContainerRegistryAuthentication{
  AwsEcr: &containerRegistryAuthenticationAwsEcr,
  Basic: &containerRegistryAuthenticationBasic,
  DockerHub: &containerRegistryAuthenticationDockerHub,
  GcpGar: &containerRegistryAuthenticationGcpGar,
  GcpGcr: &containerRegistryAuthenticationGcpGcr,
}


containerResourceRequirements := shared.ContainerResourceRequirements{
  Cpu: util.ToPointer(int64(123)),
  Memory: util.ToPointer(int64(123)),
  GpuClasses: []string{},
  StorageAmount: util.ToPointer(int64(123)),
}

containerConfiguration := containergroups.ContainerConfiguration{
  Command: []string{},
  EnvironmentVariables: map[string]string{},
  Image: util.ToPointer("Image"),
  ImageCaching: util.ToPointer(true),
  Logging: &containerLoggingConfiguration,
  Priority: &containerGroupPriority,
  RegistryAuthentication: &containerRegistryAuthentication,
  Resources: &containerResourceRequirements,
}

countryCode := shared.COUNTRY_CODE_AF


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}


containerGroupGRpcProbe := shared.ContainerGroupGRpcProbe{
  Port: util.ToPointer(int64(123)),
  Service: util.ToPointer("Service"),
}


containerGroupProbeHttpHeader := shared.ContainerGroupProbeHttpHeader{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

httpScheme := shared.HTTP_SCHEME_HTTP

containerGroupHttpProbeConfiguration := shared.ContainerGroupHttpProbeConfiguration{
  Headers: []shared.ContainerGroupProbeHttpHeader{containerGroupProbeHttpHeader},
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  Scheme: &httpScheme,
}


containerGroupTcpProbe := shared.ContainerGroupTcpProbe{
  Port: util.ToPointer(int64(123)),
}

containerGroupLivenessProbe := shared.ContainerGroupLivenessProbe{
  Exec: &containerGroupProbeExec,
  FailureThreshold: util.ToPointer(int64(123)),
  Grpc: &containerGroupGRpcProbe,
  Http: &containerGroupHttpProbeConfiguration,
  InitialDelaySeconds: util.ToPointer(int64(123)),
  PeriodSeconds: util.ToPointer(int64(123)),
  SuccessThreshold: util.ToPointer(int64(123)),
  Tcp: &containerGroupTcpProbe,
  TimeoutSeconds: util.ToPointer(int64(123)),
}

theContainerGroupNetworkingLoadBalancer := shared.THE_CONTAINER_GROUP_NETWORKING_LOAD_BALANCER_ROUND_ROBIN

containerNetworkingProtocol := shared.CONTAINER_NETWORKING_PROTOCOL_HTTP

createContainerGroupNetworking := containergroups.CreateContainerGroupNetworking{
  Auth: util.ToPointer(true),
  ClientRequestTimeout: util.ToPointer(int64(123)),
  LoadBalancer: &theContainerGroupNetworkingLoadBalancer,
  Port: util.ToPointer(int64(123)),
  Protocol: &containerNetworkingProtocol,
  ServerResponseTimeout: util.ToPointer(int64(123)),
  SingleConnectionLimit: util.ToPointer(true),
}


queueBasedAutoscalerConfiguration := shared.QueueBasedAutoscalerConfiguration{
  DesiredQueueLength: util.ToPointer(int64(123)),
  MaxReplicas: util.ToPointer(int64(123)),
  MaxDownscalePerMinute: util.ToPointer(int64(123)),
  MaxUpscalePerMinute: util.ToPointer(int64(123)),
  MinReplicas: util.ToPointer(int64(123)),
  PollingPeriod: util.ToPointer(int64(123)),
}


containerGroupQueueConnection := shared.ContainerGroupQueueConnection{
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  QueueName: util.ToPointer("QueueName"),
}


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}


containerGroupGRpcProbe := shared.ContainerGroupGRpcProbe{
  Port: util.ToPointer(int64(123)),
  Service: util.ToPointer("Service"),
}


containerGroupProbeHttpHeader := shared.ContainerGroupProbeHttpHeader{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

httpScheme := shared.HTTP_SCHEME_HTTP

containerGroupHttpProbeConfiguration := shared.ContainerGroupHttpProbeConfiguration{
  Headers: []shared.ContainerGroupProbeHttpHeader{containerGroupProbeHttpHeader},
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  Scheme: &httpScheme,
}


containerGroupTcpProbe := shared.ContainerGroupTcpProbe{
  Port: util.ToPointer(int64(123)),
}

containerGroupReadinessProbe := shared.ContainerGroupReadinessProbe{
  Exec: &containerGroupProbeExec,
  FailureThreshold: util.ToPointer(int64(123)),
  Grpc: &containerGroupGRpcProbe,
  Http: &containerGroupHttpProbeConfiguration,
  InitialDelaySeconds: util.ToPointer(int64(123)),
  PeriodSeconds: util.ToPointer(int64(123)),
  SuccessThreshold: util.ToPointer(int64(123)),
  Tcp: &containerGroupTcpProbe,
  TimeoutSeconds: util.ToPointer(int64(123)),
}

containerRestartPolicy := shared.CONTAINER_RESTART_POLICY_ALWAYS


containerGroupProbeExec := shared.ContainerGroupProbeExec{
  Command: []string{},
}


containerGroupGRpcProbe := shared.ContainerGroupGRpcProbe{
  Port: util.ToPointer(int64(123)),
  Service: util.ToPointer("Service"),
}


containerGroupProbeHttpHeader := shared.ContainerGroupProbeHttpHeader{
  Name: util.ToPointer("Name"),
  Value: util.ToPointer("Value"),
}

httpScheme := shared.HTTP_SCHEME_HTTP

containerGroupHttpProbeConfiguration := shared.ContainerGroupHttpProbeConfiguration{
  Headers: []shared.ContainerGroupProbeHttpHeader{containerGroupProbeHttpHeader},
  Path: util.ToPointer("Path"),
  Port: util.ToPointer(int64(123)),
  Scheme: &httpScheme,
}


containerGroupTcpProbe := shared.ContainerGroupTcpProbe{
  Port: util.ToPointer(int64(123)),
}

containerGroupStartupProbe := shared.ContainerGroupStartupProbe{
  Exec: &containerGroupProbeExec,
  FailureThreshold: util.ToPointer(int64(123)),
  Grpc: &containerGroupGRpcProbe,
  Http: &containerGroupHttpProbeConfiguration,
  InitialDelaySeconds: util.ToPointer(int64(123)),
  Tcp: &containerGroupTcpProbe,
  PeriodSeconds: util.ToPointer(int64(123)),
  SuccessThreshold: util.ToPointer(int64(123)),
  TimeoutSeconds: util.ToPointer(int64(123)),
}

request := containergroups.ContainerGroupCreationRequest{
  AutostartPolicy: util.ToPointer(true),
  Container: &containerConfiguration,
  CountryCodes: []shared.CountryCode{countryCode},
  DisplayName: util.ToPointer("DisplayName"),
  LivenessProbe: &containerGroupLivenessProbe,
  Name: util.ToPointer("Name"),
  Networking: &createContainerGroupNetworking,
  QueueAutoscaler: &queueBasedAutoscalerConfiguration,
  QueueConnection: &containerGroupQueueConnection,
  ReadinessProbe: &containerGroupReadinessProbe,
  Replicas: util.ToPointer(int64(123)),
  RestartPolicy: &containerRestartPolicy,
  StartupProbe: &containerGroupStartupProbe,
}

response, err := client.ContainerGroups.CreateContainerGroup(context.Background(), "organizationName", "projectName", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
