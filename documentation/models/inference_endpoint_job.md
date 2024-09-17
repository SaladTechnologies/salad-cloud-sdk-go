# InferenceEndpointJob

Represents a inference endpoint job

**Properties**

| Name                  | Type                                           | Required | Description                            |
| :-------------------- | :--------------------------------------------- | :------- | :------------------------------------- |
| Id                    | string                                         | ✅       |                                        |
| Input                 | any                                            | ✅       | The job input. May be any valid JSON.  |
| InferenceEndpointName | string                                         | ✅       | The inference endpoint name            |
| Status                | inferenceendpoints.InferenceEndpointJobStatus  | ✅       |                                        |
| Events                | []inferenceendpoints.InferenceEndpointJobEvent | ✅       |                                        |
| OrganizationName      | string                                         | ✅       | The organization name                  |
| CreateTime            | string                                         | ✅       |                                        |
| UpdateTime            | string                                         | ✅       |                                        |
| Metadata              | any                                            | ❌       |                                        |
| Webhook               | string                                         | ❌       |                                        |
| Output                | any                                            | ❌       | The job output. May be any valid JSON. |

# InferenceEndpointJobStatus

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| pending   | string | ✅       | "pending"   |
| running   | string | ✅       | "running"   |
| succeeded | string | ✅       | "succeeded" |
| cancelled | string | ✅       | "cancelled" |
| failed    | string | ✅       | "failed"    |
