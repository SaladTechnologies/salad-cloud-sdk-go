# LogsService

A list of all methods in the `LogsService` service. Click on the method name to view detailed information about that method.

| Methods                             | Description                                                                                                               |
| :---------------------------------- | :------------------------------------------------------------------------------------------------------------------------ |
| [QueryLogEntries](#querylogentries) | Retrieve a collection of _log entries_ for the _organization_ identified by `{organization_name}` matching the log query. |

## QueryLogEntries

Retrieve a collection of _log entries_ for the _organization_ identified by `{organization_name}` matching the log query.

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/log-entries`

**Parameters**

| Name             | Type          | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context       | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string        | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| logEntryQuery    | LogEntryQuery | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`LogEntryCollection`

**Example Usage Code Snippet**

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
