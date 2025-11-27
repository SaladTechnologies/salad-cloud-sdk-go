```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.Queues.DeleteQueueJob(context.Background(), "acme-corp", "dev-env", "fifo-queue", "7dcd6922-50e9-4d56-89b5-91cde26f0211")
if err != nil {
  panic(err)
}

fmt.Println(response)

```
