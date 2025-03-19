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


request := containergroups.ContainerGroupInstancePatch{
  DeletionCost: util.ToPointer(util.Nullable[int64]{ Value: int64(123) }),
}

response, err := client.ContainerGroups.UpdateContainerGroupInstance(context.Background(), "organizationName", "projectName", "containerGroupName", "containerGroupInstanceId", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
