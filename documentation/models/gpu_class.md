# GpuClass

Represents a GPU Class

**Properties**

| Name         | Type                             | Required | Description                                          |
| :----------- | :------------------------------- | :------- | :--------------------------------------------------- |
| Id           | string                           | ✅       | The unique identifier                                |
| Name         | string                           | ✅       | The GPU class name                                   |
| Prices       | []organizationdata.GpuClassPrice | ✅       | The list of prices for each container group priority |
| IsHighDemand | bool                             | ❌       | Whether the GPU class is in high demand              |
| GpuClassType | organizationdata.GpuClassType    | ❌       | The type of GPU class                                |
| MinVcpu      | int64                            | ❌       | The minimum vCPU count                               |
| MaxVcpu      | int64                            | ❌       | The maximum vCPU count                               |
| MinRam       | int64                            | ❌       | The minimum RAM amount in GB                         |
| MaxRam       | int64                            | ❌       | The maximum RAM amount in GB                         |
| MinStorage   | int64                            | ❌       | The minimum storage amount in GB                     |
| MaxStorage   | int64                            | ❌       | The maximum storage amount in GB                     |

# GpuClassType

The type of GPU class

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| community | string | ✅       | "community" |
| secure    | string | ✅       | "secure"    |
