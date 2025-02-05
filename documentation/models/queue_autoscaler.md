# QueueAutoscaler

Represents the autoscaling rules for a queue

**Properties**

| Name                  | Type  | Required | Description                                                     |
| :-------------------- | :---- | :------- | :-------------------------------------------------------------- |
| MinReplicas           | int64 | ✅       | The minimum number of instances the container can scale down to |
| MaxReplicas           | int64 | ✅       | The maximum number of instances the container can scale up to   |
| DesiredQueueLength    | int64 | ✅       |                                                                 |
| PollingPeriod         | int64 | ❌       | The period (in seconds) in which the queue checks the formula   |
| MaxUpscalePerMinute   | int64 | ❌       | The maximum number of instances that can be added per minute    |
| MaxDownscalePerMinute | int64 | ❌       | The maximum number of instances that can be removed per minute  |
