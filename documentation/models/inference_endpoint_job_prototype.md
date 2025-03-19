# InferenceEndpointJobPrototype

Represents a request to create a inference endpoint job

**Properties**

| Name       | Type   | Required | Description                                              |
| :--------- | :----- | :------- | :------------------------------------------------------- |
| Input      | any    | ✅       | The job input. May be any valid JSON.                    |
| Metadata   | any    | ❌       | The job metadata. May be any valid JSON.                 |
| Webhook    | string | ❌       | The webhook URL to which the job results will be POSTed. |
| WebhookUrl | string | ❌       | The webhook URL to which the job results will be POSTed. |
