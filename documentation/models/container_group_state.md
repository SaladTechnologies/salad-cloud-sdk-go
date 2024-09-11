# ContainerGroupState

Represents a container group state

**Properties**

| Name                 | Type                                       | Required | Description                                        |
| :------------------- | :----------------------------------------- | :------- | :------------------------------------------------- |
| Status               | `shared.ContainerGroupStatus`              | ✅       |                                                    |
| StartTime            | `string`                                   | ✅       |                                                    |
| FinishTime           | `string`                                   | ✅       |                                                    |
| InstanceStatusCounts | `shared.ContainerGroupInstanceStatusCount` | ✅       | Represents a container group instance status count |
| Description          | `string`                                   | ❌       |                                                    |
