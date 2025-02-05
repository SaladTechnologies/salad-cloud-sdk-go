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


params := inferenceendpoints.GetInferenceEndpointJobsRequestParams{}


response, err := client.InferenceEndpoints.GetInferenceEndpointJobs(context.Background(), "organizationName", "inferenceEndpointName", params)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
