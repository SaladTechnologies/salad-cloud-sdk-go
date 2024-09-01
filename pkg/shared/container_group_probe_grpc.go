package shared

type ContainerGroupProbeGrpc struct {
	Service *string `json:"service,omitempty" required:"true"`
	Port    *int64  `json:"port,omitempty" required:"true" min:"0" max:"65536"`
}

func (c *ContainerGroupProbeGrpc) SetService(service string) {
	c.Service = &service
}

func (c *ContainerGroupProbeGrpc) GetService() *string {
	if c == nil {
		return nil
	}
	return c.Service
}

func (c *ContainerGroupProbeGrpc) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerGroupProbeGrpc) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}
