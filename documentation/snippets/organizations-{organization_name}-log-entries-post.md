```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/logs"
)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)

logEntryQuerySortOrder := logs.LOG_ENTRY_QUERY_SORT_ORDER_DESC

request := logs.LogEntryQuery{
  EndTime: util.ToPointer("EndTime"),
  PageSize: util.ToPointer(int64(123)),
  Query: util.ToPointer("Query"),
  SortOrder: &logEntryQuerySortOrder,
  StartTime: util.ToPointer("StartTime"),
}

response, err := client.Logs.QueryLogEntries(context.Background(), "organizationName", request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
