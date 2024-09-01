```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/inferenceendpoints"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)


request := inferenceendpoints.CreateInferenceEndpointJob{}
request.SetInput(any)

response, err := client.InferenceEndpoints.CreateInferenceEndpointJob(context.Background(), "organizationName", "inferenceEndpointName", request)
if err != nil {
  panic(err)
}

fmt.Print(response)

```
