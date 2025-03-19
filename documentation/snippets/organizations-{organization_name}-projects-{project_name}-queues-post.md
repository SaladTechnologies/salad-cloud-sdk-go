```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


request := queues.QueuePrototype{
  Name: util.ToPointer("Name"),
  DisplayName: util.ToPointer("DisplayName"),
  Description: util.ToPointer("Description"),
}

response, err := client.Queues.CreateQueue(context.Background(), "organizationName", "projectName", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
