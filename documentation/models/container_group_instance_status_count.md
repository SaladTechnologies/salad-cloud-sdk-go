# ContainerGroupInstanceStatusCount

A summary of container group instances categorized by their current lifecycle status

**Properties**

| Name            | Type  | Required | Description                                                                    |
| :-------------- | :---- | :------- | :----------------------------------------------------------------------------- |
| AllocatingCount | int64 | ✅       | The number of container instances that are currently being allocated resources |
| CreatingCount   | int64 | ✅       | The number of container instances that are in the process of being created     |
| RunningCount    | int64 | ✅       | The number of container instances that are currently running and operational   |
| StoppingCount   | int64 | ✅       | The number of container instances that are in the process of stopping          |
