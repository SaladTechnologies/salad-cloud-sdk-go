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

createContainer := containergroups.CreateContainer{}
createContainer.SetImage("Image")
createContainer.SetResources(containerResourceRequirements)

containerRestartPolicy := shared.CONTAINER_RESTART_POLICY_ALWAYS

request := containergroups.CreateContainerGroup{}
request.SetName("Name")
request.SetContainer(createContainer)
request.SetAutostartPolicy(true)
request.SetRestartPolicy(containerRestartPolicy)
request.SetReplicas(int64(123))

response, err := client.ContainerGroups.CreateContainerGroup(context.Background(), "organizationName", "projectName", request)
if err != nil {
  panic(err)
}

fmt.Print(response)

```
