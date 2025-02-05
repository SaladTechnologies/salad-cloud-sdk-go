package shared

import (
	"encoding/json"
)

type ContainerGroupProbeGrpc struct {
	Service *string `json:"service,omitempty" required:"true"`
	Port    *int64  `json:"port,omitempty" required:"true" min:"0" max:"65536"`
	touched map[string]bool
}

func (c *ContainerGroupProbeGrpc) GetService() *string {
	if c == nil {
		return nil
	}
	return c.Service
}

func (c *ContainerGroupProbeGrpc) SetService(service string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Service"] = true
	c.Service = &service
}

func (c *ContainerGroupProbeGrpc) SetServiceNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Service"] = true
	c.Service = nil
}

func (c *ContainerGroupProbeGrpc) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupProbeGrpc) SetPort(port int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = &port
}

func (c *ContainerGroupProbeGrpc) SetPortNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = nil
}

func (c ContainerGroupProbeGrpc) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Service"] && c.Service == nil {
		data["service"] = nil
	} else if c.Service != nil {
		data["service"] = c.Service
	}

	if c.touched["Port"] && c.Port == nil {
		data["port"] = nil
	} else if c.Port != nil {
		data["port"] = c.Port
	}

	return json.Marshal(data)
}

func (c ContainerGroupProbeGrpc) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupProbeGrpc to string"
	}
	return string(jsonData)
}
