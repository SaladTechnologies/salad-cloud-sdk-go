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
