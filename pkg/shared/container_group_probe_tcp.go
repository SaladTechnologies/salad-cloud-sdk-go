package shared

import (
	"encoding/json"
)

type ContainerGroupProbeTcp struct {
	Port    *int64 `json:"port,omitempty" required:"true" min:"0" max:"65536"`
	touched map[string]bool
}

func (c *ContainerGroupProbeTcp) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupProbeTcp) SetPort(port int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = &port
}

func (c *ContainerGroupProbeTcp) SetPortNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = nil
}

func (c ContainerGroupProbeTcp) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Port"] && c.Port == nil {
		data["port"] = nil
	} else if c.Port != nil {
		data["port"] = c.Port
	}

	return json.Marshal(data)
}

func (c ContainerGroupProbeTcp) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupProbeTcp to string"
	}
	return string(jsonData)
}
