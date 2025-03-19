package shared

import "encoding/json"

// Configuration for gRPC-based health probes in container groups, used to determine container health status.
type ContainerGroupGRpcProbe struct {
	// The port number on which the gRPC health check service is exposed.
	Port *int64 `json:"port,omitempty" required:"true" min:"0" max:"65536"`
	// The name of the gRPC service that implements the health check protocol.
	Service *string `json:"service,omitempty" required:"true" maxLength:"1024" pattern:"^.*$"`
}

func (c *ContainerGroupGRpcProbe) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupGRpcProbe) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerGroupGRpcProbe) GetService() *string {
	if c == nil {
		return nil
	}
	return c.Service
}

func (c *ContainerGroupGRpcProbe) SetService(service string) {
	c.Service = &service
}

func (c ContainerGroupGRpcProbe) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupGRpcProbe to string"
	}
	return string(jsonData)
}
