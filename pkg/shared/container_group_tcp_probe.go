package shared

import "encoding/json"

// Configuration for a TCP probe used to check container health via network connectivity.
type ContainerGroupTcpProbe struct {
	// The TCP port number that the probe should connect to. Must be a valid port number between 0 and 65535.
	Port *int64 `json:"port,omitempty" required:"true" min:"0" max:"65535"`
}

func (c *ContainerGroupTcpProbe) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupTcpProbe) SetPort(port int64) {
	c.Port = &port
}

func (c ContainerGroupTcpProbe) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupTcpProbe to string"
	}
	return string(jsonData)
}
