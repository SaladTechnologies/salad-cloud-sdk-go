package shared

import "encoding/json"

type ContainerGroupProbeTcp struct {
	Port *int64 `json:"port,omitempty" required:"true" min:"0" max:"65536"`
}

func (c *ContainerGroupProbeTcp) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupProbeTcp) SetPort(port int64) {
	c.Port = &port
}

func (c ContainerGroupProbeTcp) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupProbeTcp to string"
	}
	return string(jsonData)
}
