# InferenceEndpointJob

Represents a inference endpoint job

**Properties**

| Name                  | Type                                           | Required | Description                                    |
| :-------------------- | :--------------------------------------------- | :------- | :--------------------------------------------- |
| Id                    | string                                         | ✅       | The unique identifier.                         |
| Input                 | any                                            | ✅       | The job input. May be any valid JSON.          |
| InferenceEndpointName | string                                         | ✅       | The inference endpoint name.                   |
| Status                | inferenceendpoints.InferenceEndpointJobStatus  | ✅       | The current status.                            |
| Events                | []inferenceendpoints.InferenceEndpointJobEvent | ✅       | The list of events.                            |
| OrganizationName      | string                                         | ✅       | The organization name.                         |
| CreateTime            | string                                         | ✅       | The time the job was created.                  |
| UpdateTime            | string                                         | ✅       | The time the job was last updated.             |
| Metadata              | any                                            | ❌       | The job metadata. May be any valid JSON.       |
| Webhook               | string                                         | ❌       | The webhook URL called when the job completes. |
| Output                | any                                            | ❌       | The job output. May be any valid JSON.         |

# InferenceEndpointJobStatus

The current status.

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| pending   | string | ✅       | "pending"   |
| running   | string | ✅       | "running"   |
| succeeded | string | ✅       | "succeeded" |
| cancelled | string | ✅       | "cancelled" |
| failed    | string | ✅       | "failed"    |
