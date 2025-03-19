# QueueBasedAutoscalerConfiguration

Defines configuration for automatically scaling container instances based on queue length. The autoscaler monitors a queue and adjusts the number of running replicas to maintain the desired queue length.

**Properties**

| Name                  | Type  | Required | Description                                                                                                           |
| :-------------------- | :---- | :------- | :-------------------------------------------------------------------------------------------------------------------- |
| DesiredQueueLength    | int64 | ✅       | The target number of items in the queue that the autoscaler attempts to maintain by scaling the containers up or down |
| MaxReplicas           | int64 | ✅       | The maximum number of instances the container can scale up to                                                         |
| MinReplicas           | int64 | ✅       | The minimum number of instances the container can scale down to, ensuring baseline availability                       |
| MaxDownscalePerMinute | int64 | ❌       | The maximum number of instances that can be removed per minute to prevent rapid downscaling                           |
| MaxUpscalePerMinute   | int64 | ❌       | The maximum number of instances that can be added per minute to prevent rapid upscaling                               |
| PollingPeriod         | int64 | ❌       | The period (in seconds) in which the autoscaler checks the queue length and applies the scaling formula               |
