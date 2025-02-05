package shared

import (
	"encoding/json"
)

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
	touched             map[string]bool
}

func (c *ContainerGroupReadinessProbe) GetTcp() *ContainerGroupProbeTcp {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerGroupReadinessProbe) SetTcp(tcp ContainerGroupProbeTcp) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Tcp"] = true
	c.Tcp = &tcp
}

func (c *ContainerGroupReadinessProbe) SetTcpNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Tcp"] = true
	c.Tcp = nil
}

func (c *ContainerGroupReadinessProbe) GetHttp() *ContainerGroupProbeHttp {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerGroupReadinessProbe) SetHttp(http ContainerGroupProbeHttp) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Http"] = true
	c.Http = &http
}

func (c *ContainerGroupReadinessProbe) SetHttpNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Http"] = true
	c.Http = nil
}

func (c *ContainerGroupReadinessProbe) GetGrpc() *ContainerGroupProbeGrpc {
	if c == nil {
		return nil
	}
	return c.Grpc
}

func (c *ContainerGroupReadinessProbe) SetGrpc(grpc ContainerGroupProbeGrpc) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Grpc"] = true
	c.Grpc = &grpc
}

func (c *ContainerGroupReadinessProbe) SetGrpcNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Grpc"] = true
	c.Grpc = nil
}

func (c *ContainerGroupReadinessProbe) GetExec() *ContainerGroupProbeExec {
	if c == nil {
		return nil
	}
	return c.Exec
}

func (c *ContainerGroupReadinessProbe) SetExec(exec ContainerGroupProbeExec) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Exec"] = true
	c.Exec = &exec
}

func (c *ContainerGroupReadinessProbe) SetExecNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Exec"] = true
	c.Exec = nil
}

func (c *ContainerGroupReadinessProbe) GetInitialDelaySeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.InitialDelaySeconds
}

func (c *ContainerGroupReadinessProbe) SetInitialDelaySeconds(initialDelaySeconds int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["InitialDelaySeconds"] = true
	c.InitialDelaySeconds = &initialDelaySeconds
}

func (c *ContainerGroupReadinessProbe) SetInitialDelaySecondsNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["InitialDelaySeconds"] = true
	c.InitialDelaySeconds = nil
}

func (c *ContainerGroupReadinessProbe) GetPeriodSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.PeriodSeconds
}

func (c *ContainerGroupReadinessProbe) SetPeriodSeconds(periodSeconds int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["PeriodSeconds"] = true
	c.PeriodSeconds = &periodSeconds
}

func (c *ContainerGroupReadinessProbe) SetPeriodSecondsNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["PeriodSeconds"] = true
	c.PeriodSeconds = nil
}

func (c *ContainerGroupReadinessProbe) GetTimeoutSeconds() *int64 {
	if c == nil {
		return nil
	}
	return c.TimeoutSeconds
}

func (c *ContainerGroupReadinessProbe) SetTimeoutSeconds(timeoutSeconds int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["TimeoutSeconds"] = true
	c.TimeoutSeconds = &timeoutSeconds
}

func (c *ContainerGroupReadinessProbe) SetTimeoutSecondsNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["TimeoutSeconds"] = true
	c.TimeoutSeconds = nil
}

func (c *ContainerGroupReadinessProbe) GetSuccessThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.SuccessThreshold
}

func (c *ContainerGroupReadinessProbe) SetSuccessThreshold(successThreshold int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["SuccessThreshold"] = true
	c.SuccessThreshold = &successThreshold
}

func (c *ContainerGroupReadinessProbe) SetSuccessThresholdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["SuccessThreshold"] = true
	c.SuccessThreshold = nil
}

func (c *ContainerGroupReadinessProbe) GetFailureThreshold() *int64 {
	if c == nil {
		return nil
	}
	return c.FailureThreshold
}

func (c *ContainerGroupReadinessProbe) SetFailureThreshold(failureThreshold int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["FailureThreshold"] = true
	c.FailureThreshold = &failureThreshold
}

func (c *ContainerGroupReadinessProbe) SetFailureThresholdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["FailureThreshold"] = true
	c.FailureThreshold = nil
}

func (c ContainerGroupReadinessProbe) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Tcp"] && c.Tcp == nil {
		data["tcp"] = nil
	} else if c.Tcp != nil {
		data["tcp"] = c.Tcp
	}

	if c.touched["Http"] && c.Http == nil {
		data["http"] = nil
	} else if c.Http != nil {
		data["http"] = c.Http
	}

	if c.touched["Grpc"] && c.Grpc == nil {
		data["grpc"] = nil
	} else if c.Grpc != nil {
		data["grpc"] = c.Grpc
	}

	if c.touched["Exec"] && c.Exec == nil {
		data["exec"] = nil
	} else if c.Exec != nil {
		data["exec"] = c.Exec
	}

	if c.touched["InitialDelaySeconds"] && c.InitialDelaySeconds == nil {
		data["initial_delay_seconds"] = nil
	} else if c.InitialDelaySeconds != nil {
		data["initial_delay_seconds"] = c.InitialDelaySeconds
	}

	if c.touched["PeriodSeconds"] && c.PeriodSeconds == nil {
		data["period_seconds"] = nil
	} else if c.PeriodSeconds != nil {
		data["period_seconds"] = c.PeriodSeconds
	}

	if c.touched["TimeoutSeconds"] && c.TimeoutSeconds == nil {
		data["timeout_seconds"] = nil
	} else if c.TimeoutSeconds != nil {
		data["timeout_seconds"] = c.TimeoutSeconds
	}

	if c.touched["SuccessThreshold"] && c.SuccessThreshold == nil {
		data["success_threshold"] = nil
	} else if c.SuccessThreshold != nil {
		data["success_threshold"] = c.SuccessThreshold
	}

	if c.touched["FailureThreshold"] && c.FailureThreshold == nil {
		data["failure_threshold"] = nil
	} else if c.FailureThreshold != nil {
		data["failure_threshold"] = c.FailureThreshold
	}

	return json.Marshal(data)
}

func (c ContainerGroupReadinessProbe) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupReadinessProbe to string"
	}
	return string(jsonData)
}
