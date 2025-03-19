# GpuClassPrice

Represents the price of a GPU class for a given container group priority

**Properties**

| Name     | Type                          | Required | Description                                                                                                                 |
| :------- | :---------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------- |
| Priority | shared.ContainerGroupPriority | ✅       | Specifies the priority level for container group execution, which determines resource allocation and scheduling precedence. |
| Price    | string                        | ✅       | The price                                                                                                                   |
