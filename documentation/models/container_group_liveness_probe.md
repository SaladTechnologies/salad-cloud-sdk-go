# ContainerGroupLivenessProbe

Defines a liveness probe for container groups that determines when to restart a container if it becomes unhealthy

**Properties**

| Name                | Type                                        | Required | Description                                                                                                                       |
| :------------------ | :------------------------------------------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------- |
| FailureThreshold    | int64                                       | ✅       | Number of consecutive failures required to consider the probe as failed                                                           |
| InitialDelaySeconds | int64                                       | ✅       | Number of seconds to wait after container start before initiating liveness probes                                                 |
| PeriodSeconds       | int64                                       | ✅       | Frequency in seconds at which the probe should be executed                                                                        |
| SuccessThreshold    | int64                                       | ✅       | Number of consecutive successes required to consider the probe successful                                                         |
| TimeoutSeconds      | int64                                       | ✅       | Number of seconds after which the probe times out if no response is received                                                      |
| Exec                | shared.ContainerGroupProbeExec              | ❌       | Defines the exec action for a probe in a container group. This is used to execute a command inside a container for health checks. |
| Grpc                | shared.ContainerGroupGRpcProbe              | ❌       | Configuration for gRPC-based health probes in container groups, used to determine container health status.                        |
| Http                | shared.ContainerGroupHttpProbeConfiguration | ❌       | Defines HTTP probe configuration for container health checks within a container group.                                            |
| Tcp                 | shared.ContainerGroupTcpProbe               | ❌       | Configuration for a TCP probe used to check container health via network connectivity.                                            |
