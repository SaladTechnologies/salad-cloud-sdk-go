# InferenceEndpointJob

Represents a inference endpoint job

**Properties**

| Name                  | Type                                           | Required | Description                                    |
| :-------------------- | :--------------------------------------------- | :------- | :--------------------------------------------- |
| Id                    | string                                         | ✅       | The inference endpoint job identifier.         |
| InferenceEndpointName | string                                         | ✅       | The inference endpoint name.                   |
| OrganizationName      | string                                         | ✅       | The organization name.                         |
| Input                 | any                                            | ✅       | The job input. May be any valid JSON.          |
| Status                | inferenceendpoints.Status                      | ✅       | The current status.                            |
| Events                | []inferenceendpoints.InferenceEndpointJobEvent | ✅       | The list of events.                            |
| CreateTime            | string                                         | ✅       | The time the job was created.                  |
| UpdateTime            | string                                         | ✅       | The time the job was last updated.             |
| Metadata              | any                                            | ❌       | The job metadata. May be any valid JSON.       |
| Webhook               | string                                         | ❌       | The webhook URL called when the job completes. |
| WebhookUrl            | string                                         | ❌       | The webhook URL called when the job completes. |
| Output                | any                                            | ❌       | The job output. May be any valid JSON.         |
