# QueueJob

Represents a queue job

**Properties**

| Name       | Type                   | Required | Description                            |
| :--------- | :--------------------- | :------- | :------------------------------------- |
| Id         | string                 | ✅       |                                        |
| Input      | any                    | ✅       | The job input. May be any valid JSON.  |
| Status     | queues.QueueJobStatus  | ✅       |                                        |
| Events     | []queues.QueueJobEvent | ✅       |                                        |
| CreateTime | string                 | ✅       |                                        |
| UpdateTime | string                 | ✅       |                                        |
| Metadata   | any                    | ❌       |                                        |
| Webhook    | string                 | ❌       |                                        |
| Output     | any                    | ❌       | The job output. May be any valid JSON. |

# QueueJobStatus

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| pending   | string | ✅       | "pending"   |
| running   | string | ✅       | "running"   |
| succeeded | string | ✅       | "succeeded" |
| cancelled | string | ✅       | "cancelled" |
| failed    | string | ✅       | "failed"    |
