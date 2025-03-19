package shared

import "encoding/json"

// Defines how to check if a container is ready to serve traffic. The readiness probe determines whether the container's application is ready to accept traffic. If the readiness probe fails, the container is considered not ready and traffic will not be sent to it.
type ContainerGroupReadinessProbe struct {
	// Defines the exec action for a probe in a container group. This is used to execute a command inside a container for health checks.
	Exec *ContainerGroupProbeExec `json:"exec,omitempty"`
	// The number of consecutive failures required to consider the probe failed. After this many consecutive failures, the container is marked as not ready.
	FailureThreshold *int64 `json:"failure_threshold,omitempty" required:"true" min:"1" max:"20"`
	// Configuration for gRPC-based health probes in container groups, used to determine container health status.
	Grpc *ContainerGroupGRpcProbe `json:"grpc,omitempty"`
	// Defines HTTP probe configuration for container health checks within a container group.
	Http *ContainerGroupHttpProbeConfiguration `json:"http,omitempty"`
	// The time in seconds to wait after the container starts before initiating the first probe. This allows time for the application to initialize before being tested.
	InitialDelaySeconds *int64 `json:"initial_delay_seconds,omitempty" required:"true" min:"0" max:"1200"`
	// How frequently (in seconds) the probe should be executed during the container's lifetime. Specifies the interval between consecutive probe executions.
	PeriodSeconds *int64 `json:"period_seconds,omitempty" required:"true" min:"1" max:"120"`
	// The minimum consecutive successes required to consider the probe successful after it has failed. Defines how many successful probe results are needed to transition from failure to success.
	SuccessThreshold *int64 `json:"success_threshold,omitempty" required:"true" min:"1" max:"10"`
	// Configuration for a TCP probe used to check container health via network connectivity.
	Tcp *ContainerGroupTcpProbe `json:"tcp,omitempty"`
	// The maximum time in seconds that the probe has to complete. If the probe doesn't return a result before the timeout, it's considered failed.
	TimeoutSeconds *int64 `json:"timeout_seconds,omitempty" required:"true" min:"1" max:"60"`
}

func (c *ContainerGroupReadinessProbe) GetExec() *ContainerGroupProbeExec {
	if c == nil {
		return nil
	}
	return c.Exec
}

func (c *ContainerGroupReadinessProbe) SetExec(exec ContainerGroupProbeExec) {
	c.Exec = &exec
}

func (c *ContainerGroupReadinessProbe) GetFailureThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.FailureThreshold
}

func (c *ContainerGroupReadinessProbe) SetFailureThreshold(failureThreshold int64) {
	c.FailureThreshold = &failureThreshold
}

func (c *ContainerGroupReadinessProbe) GetGrpc() *ContainerGroupGRpcProbe {
	if c == nil {
		return nil
	}
	return c.Grpc
}

func (c *ContainerGroupReadinessProbe) SetGrpc(grpc ContainerGroupGRpcProbe) {
	c.Grpc = &grpc
}

func (c *ContainerGroupReadinessProbe) GetHttp() *ContainerGroupHttpProbeConfiguration {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerGroupReadinessProbe) SetHttp(http ContainerGroupHttpProbeConfiguration) {
	c.Http = &http
}

func (c *ContainerGroupReadinessProbe) GetInitialDelaySeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.InitialDelaySeconds
}

func (c *ContainerGroupReadinessProbe) SetInitialDelaySeconds(initialDelaySeconds int64) {
	c.InitialDelaySeconds = &initialDelaySeconds
}

func (c *ContainerGroupReadinessProbe) GetPeriodSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.PeriodSeconds
}

func (c *ContainerGroupReadinessProbe) SetPeriodSeconds(periodSeconds int64) {
	c.PeriodSeconds = &periodSeconds
}

func (c *ContainerGroupReadinessProbe) GetSuccessThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.SuccessThreshold
}

func (c *ContainerGroupReadinessProbe) SetSuccessThreshold(successThreshold int64) {
	c.SuccessThreshold = &successThreshold
}

func (c *ContainerGroupReadinessProbe) GetTcp() *ContainerGroupTcpProbe {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerGroupReadinessProbe) SetTcp(tcp ContainerGroupTcpProbe) {
	c.Tcp = &tcp
}

func (c *ContainerGroupReadinessProbe) GetTimeoutSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.TimeoutSeconds
}

func (c *ContainerGroupReadinessProbe) SetTimeoutSeconds(timeoutSeconds int64) {
	c.TimeoutSeconds = &timeoutSeconds
}

func (c ContainerGroupReadinessProbe) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupReadinessProbe to string"
	}
	return string(jsonData)
}
