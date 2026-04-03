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
