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


request := queues.CreateQueue{
  Name: util.ToPointer("Name"),
  DisplayName: util.ToPointer(util.Nullable[string]{ Value: "DisplayName" }),
  Description: util.ToPointer(util.Nullable[string]{ Value: "Description" }),
}

response, err := client.Queues.CreateQueue(context.Background(), "organizationName", "projectName", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
