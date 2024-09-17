# ContainerGroupReadinessProbe

Represents the container group readiness probe

**Properties**

| Name                | Type                           | Required | Description |
| :------------------ | :----------------------------- | :------- | :---------- |
| InitialDelaySeconds | int64                          | ✅       |             |
| PeriodSeconds       | int64                          | ✅       |             |
| TimeoutSeconds      | int64                          | ✅       |             |
| SuccessThreshold    | int64                          | ✅       |             |
| FailureThreshold    | int64                          | ✅       |             |
| Tcp                 | shared.ContainerGroupProbeTcp  | ❌       |             |
| Http                | shared.ContainerGroupProbeHttp | ❌       |             |
| Grpc                | shared.ContainerGroupProbeGrpc | ❌       |             |
| Exec                | shared.ContainerGroupProbeExec | ❌       |             |
