# OrganizationDataService

A list of all methods in the `OrganizationDataService` service. Click on the method name to view detailed information about that method.

| Methods                           | Description          |
| :-------------------------------- | :------------------- |
| [ListGpuClasses](#listgpuclasses) | List the GPU Classes |

## ListGpuClasses

List the GPU Classes

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/gpu-classes`

**Parameters**

| Name             | Type    | Required | Description                  |
| :--------------- | :------ | :------- | :--------------------------- |
| ctx              | Context | ✅       | Default go language context  |
| organizationName | string  | ✅       | The unique organization name |

**Return Type**

`GpuClassesList`

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

response, err := client.OrganizationData.ListGpuClasses(context.Background(), "organizationName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```
