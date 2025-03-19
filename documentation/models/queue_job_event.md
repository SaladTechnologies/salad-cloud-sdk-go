# QueueJobEvent

Represents an event for queue job

**Properties**

| Name   | Type          | Required | Description                                    |
| :----- | :------------ | :------- | :--------------------------------------------- |
| Action | queues.Action | ✅       | The action that was taken on the queue job     |
| Time   | string        | ✅       | The time the action was taken on the queue job |

# Action

The action that was taken on the queue job

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| created   | string | ✅       | "created"   |
| started   | string | ✅       | "started"   |
| succeeded | string | ✅       | "succeeded" |
| cancelled | string | ✅       | "cancelled" |
| failed    | string | ✅       | "failed"    |
