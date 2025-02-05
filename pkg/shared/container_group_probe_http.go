package shared

import (
	"encoding/json"
)

type ContainerGroupProbeHttp struct {
	Path    *string                           `json:"path,omitempty" required:"true"`
	Port    *int64                            `json:"port,omitempty" required:"true" min:"0" max:"65536"`
	Scheme  *ContainerProbeHttpScheme         `json:"scheme,omitempty"`
	Headers []ContainerGroupProbeHttpHeaders2 `json:"headers,omitempty"`
	touched map[string]bool
}

func (c *ContainerGroupProbeHttp) GetPath() *string {
	if c == nil {
		return nil
	}
	return c.Path
}

func (c *ContainerGroupProbeHttp) SetPath(path string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Path"] = true
	c.Path = &path
}

func (c *ContainerGroupProbeHttp) SetPathNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Path"] = true
	c.Path = nil
}

func (c *ContainerGroupProbeHttp) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupProbeHttp) SetPort(port int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = &port
}

func (c *ContainerGroupProbeHttp) SetPortNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = nil
}

func (c *ContainerGroupProbeHttp) GetScheme() *ContainerProbeHttpScheme {
	if c == nil {
		return nil
	}
	return c.Scheme
}

func (c *ContainerGroupProbeHttp) SetScheme(scheme ContainerProbeHttpScheme) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Scheme"] = true
	c.Scheme = &scheme
}

func (c *ContainerGroupProbeHttp) SetSchemeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Scheme"] = true
	c.Scheme = nil
}

func (c *ContainerGroupProbeHttp) GetHeaders() []ContainerGroupProbeHttpHeaders2 {
	if c == nil {
		return nil
	}
	return c.Headers
}

func (c *ContainerGroupProbeHttp) SetHeaders(headers []ContainerGroupProbeHttpHeaders2) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Headers"] = true
	c.Headers = headers
}

func (c *ContainerGroupProbeHttp) SetHeadersNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Headers"] = true
	c.Headers = nil
}

func (c ContainerGroupProbeHttp) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Path"] && c.Path == nil {
		data["path"] = nil
	} else if c.Path != nil {
		data["path"] = c.Path
	}

	if c.touched["Port"] && c.Port == nil {
		data["port"] = nil
	} else if c.Port != nil {
		data["port"] = c.Port
	}

	if c.touched["Scheme"] && c.Scheme == nil {
		data["scheme"] = nil
	} else if c.Scheme != nil {
		data["scheme"] = c.Scheme
	}

	if c.touched["Headers"] && c.Headers == nil {
		data["headers"] = nil
	} else if c.Headers != nil {
		data["headers"] = c.Headers
	}

	return json.Marshal(data)
}

func (c ContainerGroupProbeHttp) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupProbeHttp to string"
	}
	return string(jsonData)
}
