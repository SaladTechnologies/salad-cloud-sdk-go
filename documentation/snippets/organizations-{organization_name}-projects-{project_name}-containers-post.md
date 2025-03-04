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
