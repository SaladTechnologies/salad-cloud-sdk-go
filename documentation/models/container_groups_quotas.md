# ContainerGroupsQuotas

Represents the organization quotas for container groups

**Properties**

| Name                                    | Type  | Required | Description                                                              |
| :-------------------------------------- | :---- | :------- | :----------------------------------------------------------------------- |
| MaxCreatedContainerGroups               | int64 | ❌       | The maximum number of container groups that can be created               |
| ContainerInstanceQuota                  | int64 | ❌       | The maximum number of replicas that can be created for a container group |
| ContainerReplicaQuota                   | int64 | ❌       | The maximum number of replicas that can be created for a container group |
| ContainerReplicasUsed                   | int64 | ❌       | The number of replicas that are currently in use                         |
| MaxContainerGroupReallocationsPerMinute | int64 | ❌       | The maximum number of container group reallocations per minute           |
| MaxContainerGroupRecreatesPerMinute     | int64 | ❌       | The maximum number of container group recreates per minute               |
| MaxContainerGroupRestartsPerMinute      | int64 | ❌       | The maximum number of container group restarts per minute                |
