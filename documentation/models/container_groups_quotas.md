# ContainerGroupsQuotas

**Properties**

| Name                                    | Type    | Required | Description |
| :-------------------------------------- | :------ | :------- | :---------- |
| MaxCreatedContainerGroups               | `int64` | ✅       |             |
| ContainerInstanceQuota                  | `int64` | ✅       |             |
| MaxContainerGroupReallocationsPerMinute | `int64` | ❌       |             |
| MaxContainerGroupRecreatesPerMinute     | `int64` | ❌       |             |
| MaxContainerGroupRestartsPerMinute      | `int64` | ❌       |             |
