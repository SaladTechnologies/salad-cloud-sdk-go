package shared

// Represents the container group readiness probe
type ContainerGroupReadinessProbe struct {
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

func (c *ContainerGroupReadinessProbe) SetTcp(tcp ContainerGroupProbeTcp) {
	c.Tcp = &tcp
}

func (c *ContainerGroupReadinessProbe) GetTcp() *ContainerGroupProbeTcp {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerGroupReadinessProbe) SetHttp(http ContainerGroupProbeHttp) {
	c.Http = &http
}

func (c *ContainerGroupReadinessProbe) GetHttp() *ContainerGroupProbeHttp {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerGroupReadinessProbe) SetGrpc(grpc ContainerGroupProbeGrpc) {
	c.Grpc = &grpc
}

func (c *ContainerGroupReadinessProbe) GetGrpc() *ContainerGroupProbeGrpc {
	if c == nil {
		return nil
	}
	return c.Grpc
}

func (c *ContainerGroupReadinessProbe) SetExec(exec ContainerGroupProbeExec) {
	c.Exec = &exec
}

func (c *ContainerGroupReadinessProbe) GetExec() *ContainerGroupProbeExec {
	if c == nil {
		return nil
	}
	return c.Exec
}

func (c *ContainerGroupReadinessProbe) SetInitialDelaySeconds(initialDelaySeconds int64) {
	c.InitialDelaySeconds = &initialDelaySeconds
}

func (c *ContainerGroupReadinessProbe) GetInitialDelaySeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.InitialDelaySeconds
}

func (c *ContainerGroupReadinessProbe) SetPeriodSeconds(periodSeconds int64) {
	c.PeriodSeconds = &periodSeconds
}

func (c *ContainerGroupReadinessProbe) GetPeriodSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.PeriodSeconds
}

func (c *ContainerGroupReadinessProbe) SetTimeoutSeconds(timeoutSeconds int64) {
	c.TimeoutSeconds = &timeoutSeconds
}

func (c *ContainerGroupReadinessProbe) GetTimeoutSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.TimeoutSeconds
}

func (c *ContainerGroupReadinessProbe) SetSuccessThreshold(successThreshold int64) {
	c.SuccessThreshold = &successThreshold
}

func (c *ContainerGroupReadinessProbe) GetSuccessThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.SuccessThreshold
}

func (c *ContainerGroupReadinessProbe) SetFailureThreshold(failureThreshold int64) {
	c.FailureThreshold = &failureThreshold
}

func (c *ContainerGroupReadinessProbe) GetFailureThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.FailureThreshold
}
