# UpdateContainerGroup

Represents a request to update a container group

**Properties**

| Name            | Type                                           | Required | Description                                                                                     |
| :-------------- | :--------------------------------------------- | :------- | :---------------------------------------------------------------------------------------------- |
| DisplayName     | string                                         | ❌       |                                                                                                 |
| Container       | containergroups.UpdateContainer                | ❌       | Represents an update container object                                                           |
| Replicas        | int64                                          | ❌       |                                                                                                 |
| CountryCodes    | []shared.CountryCode                           | ❌       | List of countries nodes must be located in. Remove this field to permit nodes from any country. |
| Networking      | containergroups.UpdateContainerGroupNetworking | ❌       | Represents update container group networking parameters                                         |
| LivenessProbe   | shared.ContainerGroupLivenessProbe             | ❌       | Represents the container group liveness probe                                                   |
| ReadinessProbe  | shared.ContainerGroupReadinessProbe            | ❌       | Represents the container group readiness probe                                                  |
| StartupProbe    | shared.ContainerGroupStartupProbe              | ❌       | Represents the container group startup probe                                                    |
| QueueAutoscaler | shared.QueueAutoscaler                         | ❌       | Represents the autoscaling rules for a queue                                                    |
