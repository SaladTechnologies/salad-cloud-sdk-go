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
