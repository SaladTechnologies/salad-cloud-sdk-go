```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.ContainerGroups.ReallocateContainerGroupInstance(context.Background(), "organizationName", "projectName", "containerGroupName", "containerGroupInstanceId")
if err != nil {
  panic(err)
}

fmt.Print(response)

```
