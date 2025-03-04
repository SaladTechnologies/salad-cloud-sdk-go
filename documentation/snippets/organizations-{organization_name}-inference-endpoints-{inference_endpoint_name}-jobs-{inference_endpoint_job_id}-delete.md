```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.InferenceEndpoints.CancelInferenceEndpointJob(context.Background(), "organizationName", "inferenceEndpointName", "inferenceEndpointJobId")
if err != nil {
  panic(err)
}

fmt.Println(response)

```
