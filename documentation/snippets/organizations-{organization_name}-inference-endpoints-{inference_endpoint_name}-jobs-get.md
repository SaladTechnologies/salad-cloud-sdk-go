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


params := inferenceendpoints.ListInferenceEndpointJobsRequestParams{
  Page: util.ToPointer(int64(1)),
  PageSize: util.ToPointer(int64(1)),
}

response, err := client.InferenceEndpoints.ListInferenceEndpointJobs(context.Background(), "acme-corp", "transcribe", params)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
