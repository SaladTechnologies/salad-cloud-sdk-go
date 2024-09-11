# InferenceEndpointJobEvent

Represents an event for inference endpoint job

**Properties**

| Name   | Type                                                 | Required | Description |
| :----- | :--------------------------------------------------- | :------- | :---------- |
| Action | `inferenceendpoints.InferenceEndpointJobEventAction` | ✅       |             |
| Time   | `string`                                             | ✅       |             |

# InferenceEndpointJobEventAction

**Properties**

| Name      | Type     | Required | Description |
| :-------- | :------- | :------- | :---------- |
| created   | `string` | ✅       | "created"   |
| started   | `string` | ✅       | "started"   |
| succeeded | `string` | ✅       | "succeeded" |
| cancelled | `string` | ✅       | "cancelled" |
| failed    | `string` | ✅       | "failed"    |
