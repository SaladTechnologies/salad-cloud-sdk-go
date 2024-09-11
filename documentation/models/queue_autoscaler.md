# QueueAutoscaler

Represents the autoscaling rules for a queue

**Properties**

| Name                  | Type    | Required | Description |
| :-------------------- | :------ | :------- | :---------- |
| MinReplicas           | `int64` | ✅       |             |
| MaxReplicas           | `int64` | ✅       |             |
| DesiredQueueLength    | `int64` | ✅       |             |
| PollingPeriod         | `int64` | ❌       |             |
| MaxUpscalePerMinute   | `int64` | ❌       |             |
| MaxDownscalePerMinute | `int64` | ❌       |             |
