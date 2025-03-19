package shared

import "encoding/json"

// Defines a liveness probe for container groups that determines when to restart a container if it becomes unhealthy
type ContainerGroupLivenessProbe struct {
	// Defines the exec action for a probe in a container group. This is used to execute a command inside a container for health checks.
	Exec *ContainerGroupProbeExec `json:"exec,omitempty"`
	// Number of consecutive failures required to consider the probe as failed
	FailureThreshold *int64 `json:"failure_threshold,omitempty" required:"true" min:"1" max:"20"`
	// Configuration for gRPC-based health probes in container groups, used to determine container health status.
	Grpc *ContainerGroupGRpcProbe `json:"grpc,omitempty"`
	// Defines HTTP probe configuration for container health checks within a container group.
	Http *ContainerGroupHttpProbeConfiguration `json:"http,omitempty"`
	// Number of seconds to wait after container start before initiating liveness probes
	InitialDelaySeconds *int64 `json:"initial_delay_seconds,omitempty" required:"true" min:"0" max:"1200"`
	// Frequency in seconds at which the probe should be executed
	PeriodSeconds *int64 `json:"period_seconds,omitempty" required:"true" min:"1" max:"120"`
	// Number of consecutive successes required to consider the probe successful
	SuccessThreshold *int64 `json:"success_threshold,omitempty" required:"true" min:"1" max:"10"`
	// Configuration for a TCP probe used to check container health via network connectivity.
	Tcp *ContainerGroupTcpProbe `json:"tcp,omitempty"`
	// Number of seconds after which the probe times out if no response is received
	TimeoutSeconds *int64 `json:"timeout_seconds,omitempty" required:"true" min:"1" max:"60"`
}

func (c *ContainerGroupLivenessProbe) GetExec() *ContainerGroupProbeExec {
	if c == nil {
		return nil
	}
	return c.Exec
}

func (c *ContainerGroupLivenessProbe) SetExec(exec ContainerGroupProbeExec) {
	c.Exec = &exec
}

func (c *ContainerGroupLivenessProbe) GetFailureThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.FailureThreshold
}

func (c *ContainerGroupLivenessProbe) SetFailureThreshold(failureThreshold int64) {
	c.FailureThreshold = &failureThreshold
}

func (c *ContainerGroupLivenessProbe) GetGrpc() *ContainerGroupGRpcProbe {
	if c == nil {
		return nil
	}
	return c.Grpc
}

func (c *ContainerGroupLivenessProbe) SetGrpc(grpc ContainerGroupGRpcProbe) {
	c.Grpc = &grpc
}

func (c *ContainerGroupLivenessProbe) GetHttp() *ContainerGroupHttpProbeConfiguration {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerGroupLivenessProbe) SetHttp(http ContainerGroupHttpProbeConfiguration) {
	c.Http = &http
}

func (c *ContainerGroupLivenessProbe) GetInitialDelaySeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.InitialDelaySeconds
}

func (c *ContainerGroupLivenessProbe) SetInitialDelaySeconds(initialDelaySeconds int64) {
	c.InitialDelaySeconds = &initialDelaySeconds
}

func (c *ContainerGroupLivenessProbe) GetPeriodSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.PeriodSeconds
}

func (c *ContainerGroupLivenessProbe) SetPeriodSeconds(periodSeconds int64) {
	c.PeriodSeconds = &periodSeconds
}

func (c *ContainerGroupLivenessProbe) GetSuccessThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.SuccessThreshold
}

func (c *ContainerGroupLivenessProbe) SetSuccessThreshold(successThreshold int64) {
	c.SuccessThreshold = &successThreshold
}

func (c *ContainerGroupLivenessProbe) GetTcp() *ContainerGroupTcpProbe {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerGroupLivenessProbe) SetTcp(tcp ContainerGroupTcpProbe) {
	c.Tcp = &tcp
}

func (c *ContainerGroupLivenessProbe) GetTimeoutSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.TimeoutSeconds
}

func (c *ContainerGroupLivenessProbe) SetTimeoutSeconds(timeoutSeconds int64) {
	c.TimeoutSeconds = &timeoutSeconds
}

func (c ContainerGroupLivenessProbe) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupLivenessProbe to string"
	}
	return string(jsonData)
}
