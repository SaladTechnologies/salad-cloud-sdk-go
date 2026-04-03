# OrganizationsService

A list of all methods in the `OrganizationsService` service. Click on the method name to view detailed information about that method.

| Methods                                   | Description                                          |
| :---------------------------------------- | :--------------------------------------------------- |
| [GetCpuAvailability](#getcpuavailability) | Gets the CPU availability for the given organization |
| [GetGpuAvailability](#getgpuavailability) | Gets the GPU availability for the given organization |

## GetCpuAvailability

Gets the CPU availability for the given organization

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/availability/sce-cpu-availability`

**Parameters**

| Name                     | Type                     | Required | Description                                                                                                                                                                                                                                         |
| :----------------------- | :----------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                      | Context                  | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName         | string                   | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| cpuAvailabilityPrototype | CpuAvailabilityPrototype | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`CpuAvailability`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/organizations"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

countryCode := shared.COUNTRY_CODE_AF

request := organizations.CpuAvailabilityPrototype{
  CountryCodes: []shared.CountryCode{countryCode},
  Cpu: util.ToPointer(util.Nullable[int64]{ Value: int64(4) }),
  Memory: util.ToPointer(util.Nullable[int64]{ Value: int64(8192) }),
  StorageAmount: util.ToPointer(util.Nullable[int64]{ Value: int64(1000000000) }),
}

response, err := client.Organizations.GetCpuAvailability(context.Background(), "acme-corp", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetGpuAvailability

Gets the GPU availability for the given organization

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/availability/sce-gpu-availability`

**Parameters**

| Name                     | Type                     | Required | Description                                                                                                                                                                                                                                         |
| :----------------------- | :----------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx                      | Context                  | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName         | string                   | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |
| gpuAvailabilityPrototype | GpuAvailabilityPrototype | ✅       |                                                                                                                                                                                                                                                     |

**Return Type**

`GpuAvailability`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/organizations"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

countryCode := shared.COUNTRY_CODE_AF

request := organizations.GpuAvailabilityPrototype{
  CountryCodes: []shared.CountryCode{countryCode},
  Cpu: util.ToPointer(util.Nullable[int64]{ Value: int64(4) }),
  GpuClasses: []string{},
  Memory: util.ToPointer(util.Nullable[int64]{ Value: int64(8192) }),
  StorageAmount: util.ToPointer(util.Nullable[int64]{ Value: int64(1000000000) }),
}

response, err := client.Organizations.GetGpuAvailability(context.Background(), "acme-corp", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
