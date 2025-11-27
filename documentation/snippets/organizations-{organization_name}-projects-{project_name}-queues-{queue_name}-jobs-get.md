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


params := queues.ListQueueJobsRequestParams{
  Page: util.ToPointer(int64(1)),
  PageSize: util.ToPointer(int64(1)),
}

response, err := client.Queues.ListQueueJobs(context.Background(), "acme-corp", "dev-env", "fifo-queue", params)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
