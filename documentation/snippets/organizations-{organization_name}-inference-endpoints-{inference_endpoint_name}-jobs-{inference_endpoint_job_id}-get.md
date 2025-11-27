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

response, err := client.InferenceEndpoints.GetInferenceEndpointJob(context.Background(), "acme-corp", "transcribe", "2fc459a1-1c09-4a34-ade7-54d03fc51d6a")
if err != nil {
  panic(err)
}

fmt.Println(response)

```
