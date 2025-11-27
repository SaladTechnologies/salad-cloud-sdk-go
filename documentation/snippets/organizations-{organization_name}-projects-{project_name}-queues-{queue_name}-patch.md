```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


request := queues.QueuePatch{
  DisplayName: util.ToPointer(util.Nullable[string]{ Value: "5Ecnf" }),
  Description: util.ToPointer(util.Nullable[string]{ Value: "description" }),
}

response, err := client.Queues.UpdateQueue(context.Background(), "acme-corp", "dev-env", "fifo-queue", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
