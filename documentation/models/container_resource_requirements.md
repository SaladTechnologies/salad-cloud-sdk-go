# ContainerResourceRequirements

Specifies the resource requirements for a container.

**Properties**

| Name          | Type     | Required | Description                                                                                                                        |
| :------------ | :------- | :------- | :--------------------------------------------------------------------------------------------------------------------------------- |
| Cpu           | int64    | ✅       | The number of CPU cores required by the container. Must be between 1 and 16.                                                       |
| Memory        | int64    | ✅       | The amount of memory (in MB) required by the container. Must be between 1024 MB and 61440 MB.                                      |
| GpuClasses    | []string | ✅       | A list of GPU class UUIDs required by the container. Can be null if no GPU is required.                                            |
| StorageAmount | int64    | ❌       | The amount of storage (in bytes) required by the container. Must be between 1 GB (1073741824 bytes) and 50 GB (53687091200 bytes). |
