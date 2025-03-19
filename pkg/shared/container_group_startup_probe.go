package shared

import "encoding/json"

// Defines a probe that checks if a container application has started successfully. Startup probes help prevent applications from being prematurely marked as unhealthy during initialization. The probe can use HTTP requests, TCP connections, gRPC calls, or shell commands to determine startup status.
type ContainerGroupStartupProbe struct {
	// Defines the exec action for a probe in a container group. This is used to execute a command inside a container for health checks.
	Exec *ContainerGroupProbeExec `json:"exec,omitempty"`
	// Number of times the probe must fail before considering the container not started
	FailureThreshold *int64 `json:"failure_threshold,omitempty" required:"true" min:"1" max:"20"`
	// Configuration for gRPC-based health probes in container groups, used to determine container health status.
	Grpc *ContainerGroupGRpcProbe `json:"grpc,omitempty"`
	// Defines HTTP probe configuration for container health checks within a container group.
	Http *ContainerGroupHttpProbeConfiguration `json:"http,omitempty"`
	// Number of seconds to wait after container startup before the first probe is executed
	InitialDelaySeconds *int64 `json:"initial_delay_seconds,omitempty" required:"true" min:"0" max:"1200"`
	// Configuration for a TCP probe used to check container health via network connectivity.
	Tcp *ContainerGroupTcpProbe `json:"tcp,omitempty"`
	// How frequently (in seconds) to perform the probe
	PeriodSeconds *int64 `json:"period_seconds,omitempty" required:"true" min:"1" max:"120"`
	// Minimum consecutive successes required for the probe to be considered successful
	SuccessThreshold *int64 `json:"success_threshold,omitempty" required:"true" min:"1" max:"10"`
	// Maximum time (in seconds) to wait for a probe response before considering it failed
	TimeoutSeconds *int64 `json:"timeout_seconds,omitempty" required:"true" min:"1" max:"60"`
}

func (c *ContainerGroupStartupProbe) GetExec() *ContainerGroupProbeExec {
	if c == nil {
		return nil
	}
	return c.Exec
}

func (c *ContainerGroupStartupProbe) SetExec(exec ContainerGroupProbeExec) {
	c.Exec = &exec
}

func (c *ContainerGroupStartupProbe) GetFailureThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.FailureThreshold
}

func (c *ContainerGroupStartupProbe) SetFailureThreshold(failureThreshold int64) {
	c.FailureThreshold = &failureThreshold
}

func (c *ContainerGroupStartupProbe) GetGrpc() *ContainerGroupGRpcProbe {
	if c == nil {
		return nil
	}
	return c.Grpc
}

func (c *ContainerGroupStartupProbe) SetGrpc(grpc ContainerGroupGRpcProbe) {
	c.Grpc = &grpc
}

func (c *ContainerGroupStartupProbe) GetHttp() *ContainerGroupHttpProbeConfiguration {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerGroupStartupProbe) SetHttp(http ContainerGroupHttpProbeConfiguration) {
	c.Http = &http
}

func (c *ContainerGroupStartupProbe) GetInitialDelaySeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.InitialDelaySeconds
}

func (c *ContainerGroupStartupProbe) SetInitialDelaySeconds(initialDelaySeconds int64) {
	c.InitialDelaySeconds = &initialDelaySeconds
}

func (c *ContainerGroupStartupProbe) GetTcp() *ContainerGroupTcpProbe {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerGroupStartupProbe) SetTcp(tcp ContainerGroupTcpProbe) {
	c.Tcp = &tcp
}

func (c *ContainerGroupStartupProbe) GetPeriodSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.PeriodSeconds
}

func (c *ContainerGroupStartupProbe) SetPeriodSeconds(periodSeconds int64) {
	c.PeriodSeconds = &periodSeconds
}

func (c *ContainerGroupStartupProbe) GetSuccessThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.SuccessThreshold
}

func (c *ContainerGroupStartupProbe) SetSuccessThreshold(successThreshold int64) {
	c.SuccessThreshold = &successThreshold
}

func (c *ContainerGroupStartupProbe) GetTimeoutSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.TimeoutSeconds
}

func (c *ContainerGroupStartupProbe) SetTimeoutSeconds(timeoutSeconds int64) {
	c.TimeoutSeconds = &timeoutSeconds
}

func (c ContainerGroupStartupProbe) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupStartupProbe to string"
	}
	return string(jsonData)
}
