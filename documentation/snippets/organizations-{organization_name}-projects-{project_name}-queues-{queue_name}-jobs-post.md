```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/queues"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


request := queues.CreateQueueJob{}
request.SetInput(any)

response, err := client.Queues.CreateQueueJob(context.Background(), "organizationName", "projectName", "queueName", request)
if err != nil {
  panic(err)
}

fmt.Print(response)

```
