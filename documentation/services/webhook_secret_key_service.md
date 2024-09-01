# WebhookSecretKeyService

A list of all methods in the `WebhookSecretKeyService` service. Click on the method name to view detailed information about that method.

| Methods                                           | Description                    |
| :------------------------------------------------ | :----------------------------- |
| [GetWebhookSecretKey](#getwebhooksecretkey)       | Gets the webhook secret key    |
| [UpdateWebhookSecretKey](#updatewebhooksecretkey) | Updates the webhook secret key |

## GetWebhookSecretKey

Gets the webhook secret key

- HTTP Method: `GET`
- Endpoint: `/organizations/{organization_name}/webhook-secret-key`

**Parameters**

| Name             | Type    | Required | Description                  |
| :--------------- | :------ | :------- | :--------------------------- |
| ctx              | Context | ✅       | Default go language context  |
| organizationName | string  | ✅       | The unique organization name |

**Return Type**

`WebhookSecretKey`

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

response, err := client.WebhookSecretKey.GetWebhookSecretKey(context.Background(), "organizationName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## UpdateWebhookSecretKey

Updates the webhook secret key

- HTTP Method: `POST`
- Endpoint: `/organizations/{organization_name}/webhook-secret-key`

**Parameters**

| Name             | Type    | Required | Description                  |
| :--------------- | :------ | :------- | :--------------------------- |
| ctx              | Context | ✅       | Default go language context  |
| organizationName | string  | ✅       | The unique organization name |

**Return Type**

`WebhookSecretKey`

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

response, err := client.WebhookSecretKey.UpdateWebhookSecretKey(context.Background(), "organizationName")
if err != nil {
  panic(err)
}

fmt.Print(response)
```
