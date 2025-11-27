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
