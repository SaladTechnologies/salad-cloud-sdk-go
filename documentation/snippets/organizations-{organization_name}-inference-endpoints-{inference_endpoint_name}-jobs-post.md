```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)


request := inferenceendpoints.InferenceEndpointJobPrototype{
  Input: []byte{},
  Metadata: []byte{},
  Webhook: util.ToPointer("webhook"),
  WebhookUrl: util.ToPointer("https://webhook.example.com/events"),
}

response, err := client.InferenceEndpoints.CreateInferenceEndpointJob(context.Background(), "acme-corp", "transcribe", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
