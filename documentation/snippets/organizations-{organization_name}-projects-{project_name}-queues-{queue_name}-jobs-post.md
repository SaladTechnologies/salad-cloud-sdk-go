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


request := queues.CreateQueueJob{
  Input: []byte{},
  Metadata: []byte{},
  Webhook: util.ToPointer(util.Nullable[string]{ Value: "Webhook" }),
}

response, err := client.Queues.CreateQueueJob(context.Background(), "organizationName", "projectName", "queueName", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
