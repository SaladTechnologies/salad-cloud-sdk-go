# OrganizationDataService

A list of all methods in the `OrganizationDataService` service. Click on the method name to view detailed information about that method.

| Methods                                   | Description                                          |
| :---------------------------------------- | :--------------------------------------------------- |
| [ListGpuClasses](#listgpuclasses)         | List the GPU Classes                                 |
| [GetCpuAvailability](#getcpuavailability) | Gets the CPU availability for the given organization |
| [GetGpuAvailability](#getgpuavailability) | Gets the GPU availability for the given organization |

## ListGpuClasses

List the GPU Classes

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/gpu-classes`

**Parameters**

| Name             | Type    | Required | Description                                                                                                                                                                                                                                         |
| :--------------- | :------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx              | Context | ✅       | Default go language context                                                                                                                                                                                                                         |
| organizationName | string  | ✅       | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. |

**Return Type**

`GpuClassesList`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"

)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

response, err := client.OrganizationData.ListGpuClasses(context.Background(), "acme-corp")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

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
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/organizationdata"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

countryCode := shared.COUNTRY_CODE_AF

request := organizationdata.CpuAvailabilityPrototype{
  Cpu: util.ToPointer(util.Nullable[int64]{ Value: int64(4) }),
  Memory: util.ToPointer(util.Nullable[int64]{ Value: int64(8192) }),
  StorageAmount: util.ToPointer(util.Nullable[int64]{ Value: int64(1000000000) }),
  CountryCodes: []shared.CountryCode{countryCode},
}

response, err := client.OrganizationData.GetCpuAvailability(context.Background(), "acme-corp", request)
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
  "github.com/saladtechnologies/salad-cloud-sdk-go/pkg/organizationdata"
)

config := saladcloudsdkconfig.NewConfig()
config.SetApiKey("API_KEY")
client := saladcloudsdk.NewSaladCloudSdk(config)

countryCode := shared.COUNTRY_CODE_AF

request := organizationdata.GpuAvailabilityPrototype{
  GpuClasses: []string{},
  Cpu: util.ToPointer(util.Nullable[int64]{ Value: int64(4) }),
  Memory: util.ToPointer(util.Nullable[int64]{ Value: int64(8192) }),
  StorageAmount: util.ToPointer(util.Nullable[int64]{ Value: int64(1000000000) }),
  CountryCodes: []shared.CountryCode{countryCode},
}

response, err := client.OrganizationData.GetGpuAvailability(context.Background(), "acme-corp", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
