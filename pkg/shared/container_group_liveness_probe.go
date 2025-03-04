package shared

import "encoding/json"

// Represents the container group liveness probe
type ContainerGroupLivenessProbe struct {
	Tcp                 *ContainerGroupProbeTcp  `json:"tcp,omitempty"`
	Http                *ContainerGroupProbeHttp `json:"http,omitempty"`
	Grpc                *ContainerGroupProbeGrpc `json:"grpc,omitempty"`
	Exec                *ContainerGroupProbeExec `json:"exec,omitempty"`
	InitialDelaySeconds *int64                   `json:"initial_delay_seconds,omitempty" required:"true" min:"0"`
	PeriodSeconds       *int64                   `json:"period_seconds,omitempty" required:"true" min:"1"`
	TimeoutSeconds      *int64                   `json:"timeout_seconds,omitempty" required:"true" min:"1"`
	SuccessThreshold    *int64                   `json:"success_threshold,omitempty" required:"true" min:"1"`
	FailureThreshold    *int64                   `json:"failure_threshold,omitempty" required:"true" min:"1"`
}

func (c *ContainerGroupLivenessProbe) GetTcp() *ContainerGroupProbeTcp {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerGroupLivenessProbe) SetTcp(tcp ContainerGroupProbeTcp) {
	c.Tcp = &tcp
}

func (c *ContainerGroupLivenessProbe) GetHttp() *ContainerGroupProbeHttp {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerGroupLivenessProbe) SetHttp(http ContainerGroupProbeHttp) {
	c.Http = &http
}

func (c *ContainerGroupLivenessProbe) GetGrpc() *ContainerGroupProbeGrpc {
	if c == nil {
		return nil
	}
	return c.Grpc
}

func (c *ContainerGroupLivenessProbe) SetGrpc(grpc ContainerGroupProbeGrpc) {
	c.Grpc = &grpc
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

func (c *ContainerGroupLivenessProbe) GetTimeoutSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.TimeoutSeconds
}

func (c *ContainerGroupLivenessProbe) SetTimeoutSeconds(timeoutSeconds int64) {
	c.TimeoutSeconds = &timeoutSeconds
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

func (c *ContainerGroupLivenessProbe) GetFailureThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.FailureThreshold
}

func (c *ContainerGroupLivenessProbe) SetFailureThreshold(failureThreshold int64) {
	c.FailureThreshold = &failureThreshold
}

func (c ContainerGroupLivenessProbe) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupLivenessProbe to string"
	}
	return string(jsonData)
}
