# QuotasService

A list of all methods in the `QuotasService` service. Click on the method name to view detailed information about that method.

| Methods                 | Description                  |
| :---------------------- | :--------------------------- |
| [GetQuotas](#getquotas) | Gets the organization quotas |

## GetQuotas

Gets the organization quotas

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/quotas`

**Parameters**

| Name             | Type    | Required | Description                  |
| :--------------- | :------ | :------- | :--------------------------- |
| ctx              | Context | ✅       | Default go language context  |
| organizationName | string  | ✅       | The unique organization name |

**Return Type**

`Quotas`

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

response, err := client.Quotas.GetQuotas(context.Background(), "organizationName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```
