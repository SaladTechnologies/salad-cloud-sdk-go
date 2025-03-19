# SystemLog

Represents a system log

**Properties**

| Name                  | Type   | Required | Description                                       |
| :-------------------- | :----- | :------- | :------------------------------------------------ |
| EventName             | string | ✅       | The name of the event                             |
| EventTime             | string | ✅       | The UTC date & time when the log item was created |
| ResourceCpu           | int64  | ✅       | The number of CPUs                                |
| ResourceGpuClass      | string | ✅       | The GPU class name                                |
| ResourceMemory        | int64  | ✅       | The memory amount in MB                           |
| ResourceStorageAmount | int64  | ✅       | The storage amount in bytes                       |
| Version               | string | ✅       | The version instance ID                           |
| InstanceId            | string | ❌       | The container group instance identifier.          |
| MachineId             | string | ❌       | The container group machine identifier.           |
