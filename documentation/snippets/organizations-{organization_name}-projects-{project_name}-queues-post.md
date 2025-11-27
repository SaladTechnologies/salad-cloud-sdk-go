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


request := queues.QueuePrototype{
  Name: util.ToPointer("name"),
  DisplayName: util.ToPointer("AV3ysQq"),
  Description: util.ToPointer("description"),
}

response, err := client.Queues.CreateQueue(context.Background(), "acme-corp", "dev-env", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
