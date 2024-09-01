package shared

type ContainerGroupProbeTcp struct {
	Port *int64 `json:"port,omitempty" required:"true" min:"0" max:"65536"`
}

func (c *ContainerGroupProbeTcp) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerGroupProbeTcp) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}
