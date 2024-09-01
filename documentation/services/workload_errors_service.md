# WorkloadErrorsService

A list of all methods in the `WorkloadErrorsService` service. Click on the method name to view detailed information about that method.

| Methods                                 | Description              |
| :-------------------------------------- | :----------------------- |
| [GetWorkloadErrors](#getworkloaderrors) | Gets the workload errors |

## GetWorkloadErrors

Gets the workload errors

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/errors`

**Parameters**

| Name               | Type    | Required | Description                     |
| :----------------- | :------ | :------- | :------------------------------ |
| ctx                | Context | ✅       | Default go language context     |
| organizationName   | string  | ✅       | The unique organization name    |
| projectName        | string  | ✅       | The unique project name         |
| containerGroupName | string  | ✅       | The unique container group name |

**Return Type**

`WorkloadErrorList`

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

response, err := client.WorkloadErrors.GetWorkloadErrors(context.Background(), "organizationName", "projectName", "containerGroupName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```
