# InferenceEndpointJob

Represents a inference endpoint job

**Properties**

| Name                  | Type                                                                              | Required | Description                                    |
| :-------------------- | :-------------------------------------------------------------------------------- | :------- | :--------------------------------------------- |
| CreateTime            | string                                                                            | ✅       | The time the job was created.                  |
| Events                | [][inferenceendpoints.InferenceEndpointJobEvent](inference_endpoint_job_event.md) | ✅       | The list of events.                            |
| Id                    | string                                                                            | ✅       | The inference endpoint job identifier.         |
| InferenceEndpointName | string                                                                            | ✅       | The inference endpoint name.                   |
| Input                 | any                                                                               | ✅       | The job input. May be any valid JSON.          |
| OrganizationName      | string                                                                            | ✅       | The organization name.                         |
| Status                | [inferenceendpoints.Status](status.md)                                            | ✅       | The current status.                            |
| UpdateTime            | string                                                                            | ✅       | The time the job was last updated.             |
| Metadata              | any                                                                               | ❌       | The job metadata. May be any valid JSON.       |
| Output                | any                                                                               | ❌       | The job output. May be any valid JSON.         |
| Webhook               | string                                                                            | ❌       | The webhook URL called when the job completes. |
| WebhookUrl            | string                                                                            | ❌       | The webhook URL called when the job completes. |
