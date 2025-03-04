package shared

import "encoding/json"

// Represents the container group startup probe
type ContainerGroupStartupProbe struct {
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

func (c *ContainerGroupStartupProbe) GetTcp() *ContainerGroupProbeTcp {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerGroupStartupProbe) SetTcp(tcp ContainerGroupProbeTcp) {
	c.Tcp = &tcp
}

func (c *ContainerGroupStartupProbe) GetHttp() *ContainerGroupProbeHttp {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerGroupStartupProbe) SetHttp(http ContainerGroupProbeHttp) {
	c.Http = &http
}

func (c *ContainerGroupStartupProbe) GetGrpc() *ContainerGroupProbeGrpc {
	if c == nil {
		return nil
	}
	return c.Grpc
}

func (c *ContainerGroupStartupProbe) SetGrpc(grpc ContainerGroupProbeGrpc) {
	c.Grpc = &grpc
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

func (c *ContainerGroupStartupProbe) GetInitialDelaySeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.InitialDelaySeconds
}

func (c *ContainerGroupStartupProbe) SetInitialDelaySeconds(initialDelaySeconds int64) {
	c.InitialDelaySeconds = &initialDelaySeconds
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

func (c *ContainerGroupStartupProbe) GetTimeoutSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.TimeoutSeconds
}

func (c *ContainerGroupStartupProbe) SetTimeoutSeconds(timeoutSeconds int64) {
	c.TimeoutSeconds = &timeoutSeconds
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

func (c *ContainerGroupStartupProbe) GetFailureThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.FailureThreshold
}

func (c *ContainerGroupStartupProbe) SetFailureThreshold(failureThreshold int64) {
	c.FailureThreshold = &failureThreshold
}

func (c ContainerGroupStartupProbe) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupStartupProbe to string"
	}
	return string(jsonData)
}
