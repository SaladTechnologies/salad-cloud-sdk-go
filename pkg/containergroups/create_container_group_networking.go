package containergroups

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents container group networking parameters
type CreateContainerGroupNetworking struct {
	Protocol *shared.ContainerNetworkingProtocol `json:"protocol,omitempty" required:"true"`
	Port     *int64                              `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	Auth     *bool                               `json:"auth,omitempty" required:"true"`
}

func (c *CreateContainerGroupNetworking) SetProtocol(protocol shared.ContainerNetworkingProtocol) {
	c.Protocol = &protocol
}

func (c *CreateContainerGroupNetworking) GetProtocol() *shared.ContainerNetworkingProtocol {
	if c == nil {
		return nil
	}
	return c.Protocol
}

func (c *CreateContainerGroupNetworking) SetPort(port int64) {
	c.Port = &port
}

func (c *CreateContainerGroupNetworking) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *CreateContainerGroupNetworking) SetAuth(auth bool) {
	c.Auth = &auth
}

func (c *CreateContainerGroupNetworking) GetAuth() *bool {
	if c == nil {
		return nil
	}
	return c.Auth
}
