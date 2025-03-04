# SystemLogsService

A list of all methods in the `SystemLogsService` service. Click on the method name to view detailed information about that method.

| Methods                         | Description          |
| :------------------------------ | :------------------- |
| [GetSystemLogs](#getsystemlogs) | Gets the System Logs |

## GetSystemLogs

Gets the System Logs

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/system-logs`

**Parameters**

| Name               | Type    | Required | Description                                                                                                                                                                                                                                         |
| :----------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName   | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| projectName        | string  | ✅       | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.                                                                                                                  |
| containerGroupName | string  | ✅       | The unique container group name                                                                                                                                                                                                                     |

**Return Type**

`SystemLogList`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.SystemLogs.GetSystemLogs(context.Background(), "organizationName", "projectName", "containerGroupName")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
