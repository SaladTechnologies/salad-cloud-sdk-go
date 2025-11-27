```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/logs"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

logEntryQuerySortOrder := logs.LOG_ENTRY_QUERY_SORT_ORDER_DESC

request := logs.LogEntryQuery{
  EndTime: util.ToPointer("end_time"),
  PageSize: util.ToPointer(int64(1)),
  Query: util.ToPointer("query"),
  SortOrder: &logEntryQuerySortOrder,
  StartTime: util.ToPointer("start_time"),
}

response, err := client.Logs.QueryLogEntries(context.Background(), "acme-corp", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
