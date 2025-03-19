# QueueJob

Represents a queue job

**Properties**

| Name       | Type                   | Required | Description                                      |
| :--------- | :--------------------- | :------- | :----------------------------------------------- |
| Id         | string                 | ✅       | The job identifier                               |
| Input      | any                    | ✅       | The job input. May be any valid JSON.            |
| Status     | queues.QueueJobStatus  | ✅       | The job status                                   |
| Events     | []queues.QueueJobEvent | ✅       | The job events                                   |
| CreateTime | string                 | ✅       | The job creation time                            |
| UpdateTime | string                 | ✅       | The job update time                              |
| Metadata   | any                    | ❌       | Additional metadata for the job                  |
| Webhook    | string                 | ❌       | The webhook URL to notify when the job completes |
| Output     | any                    | ❌       | The job output. May be any valid JSON.           |

# QueueJobStatus

The job status

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| pending   | string | ✅       | "pending"   |
| running   | string | ✅       | "running"   |
| succeeded | string | ✅       | "succeeded" |
| cancelled | string | ✅       | "cancelled" |
| failed    | string | ✅       | "failed"    |
