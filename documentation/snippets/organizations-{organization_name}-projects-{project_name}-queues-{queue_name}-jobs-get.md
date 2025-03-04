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


params := queues.ListQueueJobsRequestParams{

}

response, err := client.Queues.ListQueueJobs(context.Background(), "organizationName", "projectName", "queueName", params)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
