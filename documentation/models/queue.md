# Queue

Represents a queue.

**Properties**

| Name               | Type                                          | Required | Description                                                                                                                                                |
| :----------------- | :-------------------------------------------- | :------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ContainerGroups    | [][shared.ContainerGroup](container_group.md) | ✅       | The container groups that are part of this queue. Each container group represents a scalable set of identical containers running as a distributed service. |
| CreateTime         | string                                        | ✅       | The date and time the queue was created.                                                                                                                   |
| DisplayName        | string                                        | ✅       | The display name. This may be used as a more human-readable name.                                                                                          |
| Id                 | string                                        | ✅       | The queue identifier. This is automatically generated and assigned when the queue is created.                                                              |
| Name               | string                                        | ✅       | The queue name. This must be unique within the project.                                                                                                    |
| UpdateTime         | string                                        | ✅       | The date and time the queue was last updated.                                                                                                              |
| CurrentQueueLength | int64                                         | ❌       | The current length of the queue                                                                                                                            |
| Description        | string                                        | ❌       | The description. This may be used as a space for notes or other information about the queue.                                                               |
