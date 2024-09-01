package shared

type ContainerGroupProbeHttp struct {
	Path    *string                           `json:"path,omitempty" required:"true"`
	Port    *int64                            `json:"port,omitempty" required:"true" min:"0" max:"65536"`
	Scheme  *ContainerProbeHttpScheme         `json:"scheme,omitempty"`
	Headers []ContainerGroupProbeHttpHeaders2 `json:"headers,omitempty"`
}

func (c *ContainerGroupProbeHttp) SetPath(path string) {
	c.Path = &path
}

func (c *ContainerGroupProbeHttp) GetPath() *string {
	if c == nil {
		return nil
	}
	return c.Path
}

func (c *ContainerGroupProbeHttp) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerGroupProbeHttp) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupProbeHttp) SetScheme(scheme ContainerProbeHttpScheme) {
	c.Scheme = &scheme
}

func (c *ContainerGroupProbeHttp) GetScheme() *ContainerProbeHttpScheme {
	if c == nil {
		return nil
	}
	return c.Scheme
}

func (c *ContainerGroupProbeHttp) SetHeaders(headers []ContainerGroupProbeHttpHeaders2) {
	c.Headers = headers
}

func (c *ContainerGroupProbeHttp) GetHeaders() []ContainerGroupProbeHttpHeaders2 {
	if c == nil {
		return nil
	}
	return c.Headers
}
