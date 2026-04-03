# GpuClass

Represents a GPU Class

**Properties**

| Name         | Type                                                   | Required | Description                                          |
| :----------- | :----------------------------------------------------- | :------- | :--------------------------------------------------- |
| Id           | string                                                 | ✅       | The unique identifier                                |
| Name         | string                                                 | ✅       | The GPU class name                                   |
| Prices       | [][organizationdata.GpuClassPrice](gpu_class_price.md) | ✅       | The list of prices for each container group priority |
| GpuClassType | organizationdata.GpuClassType                          | ❌       | The type of GPU class                                |
| GpuCount     | int64                                                  | ❌       | The number of GPUs in the cluster                    |
| IsHighDemand | bool                                                   | ❌       | Whether the GPU class is in high demand              |
| MaxRam       | int64                                                  | ❌       | The maximum RAM amount in MB                         |
| MaxStorage   | int64                                                  | ❌       | The maximum storage amount in bytes                  |
| MaxVcpu      | int64                                                  | ❌       | The maximum vCPU count                               |
| MinRam       | int64                                                  | ❌       | The minimum RAM amount in MB                         |
| MinStorage   | int64                                                  | ❌       | The minimum storage amount in bytes                  |
| MinVcpu      | int64                                                  | ❌       | The minimum vCPU count                               |

# GpuClassType

The type of GPU class

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| Community | string | ✅       | "community" |
| Secure    | string | ✅       | "secure"    |
