# CreateContainerGroup

Represents a request to create a container group

**Properties**

| Name            | Type                                           | Required | Description                                                                                     |
| :-------------- | :--------------------------------------------- | :------- | :---------------------------------------------------------------------------------------------- |
| Name            | string                                         | ✅       |                                                                                                 |
| Container       | containergroups.CreateContainer                | ✅       | Represents a container                                                                          |
| AutostartPolicy | bool                                           | ✅       |                                                                                                 |
| RestartPolicy   | shared.ContainerRestartPolicy                  | ✅       |                                                                                                 |
| Replicas        | int64                                          | ✅       |                                                                                                 |
| DisplayName     | string                                         | ❌       |                                                                                                 |
| CountryCodes    | []shared.CountryCode                           | ❌       | List of countries nodes must be located in. Remove this field to permit nodes from any country. |
| Networking      | containergroups.CreateContainerGroupNetworking | ❌       | Represents container group networking parameters                                                |
| LivenessProbe   | shared.ContainerGroupLivenessProbe             | ❌       | Represents the container group liveness probe                                                   |
| ReadinessProbe  | shared.ContainerGroupReadinessProbe            | ❌       | Represents the container group readiness probe                                                  |
| StartupProbe    | shared.ContainerGroupStartupProbe              | ❌       | Represents the container group startup probe                                                    |
| QueueConnection | shared.ContainerGroupQueueConnection           | ❌       | Represents container group queue connection                                                     |
| QueueAutoscaler | shared.QueueAutoscaler                         | ❌       | Represents the autoscaling rules for a queue                                                    |
