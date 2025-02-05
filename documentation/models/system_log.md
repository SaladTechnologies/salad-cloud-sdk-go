# SystemLog

Represents a system log

**Properties**

| Name                  | Type   | Required | Description                                       |
| :-------------------- | :----- | :------- | :------------------------------------------------ |
| EventName             | string | ✅       | The name of the event                             |
| EventTime             | string | ✅       | The UTC date & time when the log item was created |
| Version               | string | ✅       | The version instance ID                           |
| ResourceCpu           | int64  | ✅       | The number of CPUs                                |
| ResourceMemory        | int64  | ✅       | The memory amount in MB                           |
| ResourceGpuClass      | string | ✅       | The GPU class name                                |
| ResourceStorageAmount | int64  | ✅       | The storage amount in bytes                       |
| InstanceId            | string | ❌       | The unique instance ID                            |
| MachineId             | string | ❌       | The organization-specific machine ID              |
