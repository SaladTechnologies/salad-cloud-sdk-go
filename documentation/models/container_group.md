# ContainerGroup

Represents a container group

**Properties**

| Name            | Type                                   | Required | Description                                                                                     |
| :-------------- | :------------------------------------- | :------- | :---------------------------------------------------------------------------------------------- |
| Id              | `string`                               | ✅       |                                                                                                 |
| Name            | `string`                               | ✅       |                                                                                                 |
| DisplayName     | `string`                               | ✅       |                                                                                                 |
| Container       | `shared.Container`                     | ✅       | Represents a container                                                                          |
| AutostartPolicy | `bool`                                 | ✅       |                                                                                                 |
| RestartPolicy   | `shared.ContainerRestartPolicy`        | ✅       |                                                                                                 |
| Replicas        | `int64`                                | ✅       |                                                                                                 |
| CurrentState    | `shared.ContainerGroupState`           | ✅       | Represents a container group state                                                              |
| CreateTime      | `string`                               | ✅       |                                                                                                 |
| UpdateTime      | `string`                               | ✅       |                                                                                                 |
| PendingChange   | `bool`                                 | ✅       |                                                                                                 |
| Version         | `int64`                                | ✅       |                                                                                                 |
| CountryCodes    | `[]shared.CountryCode`                 | ❌       | List of countries nodes must be located in. Remove this field to permit nodes from any country. |
| Networking      | `shared.ContainerGroupNetworking`      | ❌       | Represents container group networking parameters                                                |
| LivenessProbe   | `shared.ContainerGroupLivenessProbe`   | ❌       | Represents the container group liveness probe                                                   |
| ReadinessProbe  | `shared.ContainerGroupReadinessProbe`  | ❌       | Represents the container group readiness probe                                                  |
| StartupProbe    | `shared.ContainerGroupStartupProbe`    | ❌       | Represents the container group startup probe                                                    |
| QueueConnection | `shared.ContainerGroupQueueConnection` | ❌       | Represents container group queue connection                                                     |
