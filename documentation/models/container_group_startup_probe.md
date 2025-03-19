# ContainerGroupStartupProbe

Defines a probe that checks if a container application has started successfully. Startup probes help prevent applications from being prematurely marked as unhealthy during initialization. The probe can use HTTP requests, TCP connections, gRPC calls, or shell commands to determine startup status.

**Properties**

| Name                | Type                                        | Required | Description                                                                                                                       |
| :------------------ | :------------------------------------------ | :------- | :-------------------------------------------------------------------------------------------------------------------------------- |
| FailureThreshold    | int64                                       | ✅       | Number of times the probe must fail before considering the container not started                                                  |
| InitialDelaySeconds | int64                                       | ✅       | Number of seconds to wait after container startup before the first probe is executed                                              |
| PeriodSeconds       | int64                                       | ✅       | How frequently (in seconds) to perform the probe                                                                                  |
| SuccessThreshold    | int64                                       | ✅       | Minimum consecutive successes required for the probe to be considered successful                                                  |
| TimeoutSeconds      | int64                                       | ✅       | Maximum time (in seconds) to wait for a probe response before considering it failed                                               |
| Exec                | shared.ContainerGroupProbeExec              | ❌       | Defines the exec action for a probe in a container group. This is used to execute a command inside a container for health checks. |
| Grpc                | shared.ContainerGroupGRpcProbe              | ❌       | Configuration for gRPC-based health probes in container groups, used to determine container health status.                        |
| Http                | shared.ContainerGroupHttpProbeConfiguration | ❌       | Defines HTTP probe configuration for container health checks within a container group.                                            |
| Tcp                 | shared.ContainerGroupTcpProbe               | ❌       | Configuration for a TCP probe used to check container health via network connectivity.                                            |
