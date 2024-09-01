# ContainerGroupInstance

Represents the details of a single container group instance

**Properties**

| Name       | Type                  | Required | Description                                                                                                                                            |
| :--------- | :-------------------- | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------------- |
| InstanceId | string                | ✅       | The unique instance ID                                                                                                                                 |
| MachineId  | string                | ✅       | The machine ID                                                                                                                                         |
| State      | containergroups.State | ✅       | The state of the container group instance                                                                                                              |
| UpdateTime | string                | ✅       | The UTC date & time when the workload on this machine transitioned to the current state                                                                |
| Version    | int64                 | ✅       | The version of the running container group                                                                                                             |
| Ready      | bool                  | ❌       | Specifies whether the container group instance is currently passing its readiness check. If no readiness probe is defined, is true once fully started. |
| Started    | bool                  | ❌       | Specifies whether the container group instance passed its startup probe. Is always true when no startup probe is defined.                              |

# State

The state of the container group instance

**Properties**

| Name        | Type   | Required | Description   |
| :---------- | :----- | :------- | :------------ |
| allocating  | string | ✅       | "allocating"  |
| downloading | string | ✅       | "downloading" |
| creating    | string | ✅       | "creating"    |
| running     | string | ✅       | "running"     |
| stopping    | string | ✅       | "stopping"    |
