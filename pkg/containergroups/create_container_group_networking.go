package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Network configuration for container groups specifying connectivity parameters, including authentication, protocol, and timeout settings
type CreateContainerGroupNetworking struct {
	// Determines whether authentication is required for network connections to the container group
	Auth *bool `json:"auth,omitempty" required:"true"`
	// The container group networking client request timeout.
	ClientRequestTimeout *int64 `json:"client_request_timeout,omitempty" min:"1" max:"100000"`
	// The container group networking load balancer.
	LoadBalancer *shared.TheContainerGroupNetworkingLoadBalancer `json:"load_balancer,omitempty"`
	// The container group networking port.
	Port *int64 `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	// Defines the communication protocol used for network traffic between containers or external systems. Currently supports HTTP protocol for web-based communication.
	Protocol *shared.ContainerNetworkingProtocol `json:"protocol,omitempty" required:"true"`
	// The container group networking server response timeout.
	ServerResponseTimeout *int64 `json:"server_response_timeout,omitempty" min:"1" max:"100000"`
	// The container group networking single connection limit flag.
	SingleConnectionLimit *bool `json:"single_connection_limit,omitempty"`
}

func (c *CreateContainerGroupNetworking) GetAuth() *bool {
	if c == nil {
		return nil
	}
	return c.Auth
}

func (c *CreateContainerGroupNetworking) SetAuth(auth bool) {
	c.Auth = &auth
}

func (c *CreateContainerGroupNetworking) GetClientRequestTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ClientRequestTimeout
}

func (c *CreateContainerGroupNetworking) SetClientRequestTimeout(clientRequestTimeout int64) {
	c.ClientRequestTimeout = &clientRequestTimeout
}

func (c *CreateContainerGroupNetworking) GetLoadBalancer() *shared.TheContainerGroupNetworkingLoadBalancer {
	if c == nil {
		return nil
	}
	return c.LoadBalancer
}

func (c *CreateContainerGroupNetworking) SetLoadBalancer(loadBalancer shared.TheContainerGroupNetworkingLoadBalancer) {
	c.LoadBalancer = &loadBalancer
}

func (c *CreateContainerGroupNetworking) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *CreateContainerGroupNetworking) SetPort(port int64) {
	c.Port = &port
}

func (c *CreateContainerGroupNetworking) GetProtocol() *shared.ContainerNetworkingProtocol {
	if c == nil {
		return nil
	}
	return c.Protocol
}

func (c *CreateContainerGroupNetworking) SetProtocol(protocol shared.ContainerNetworkingProtocol) {
	c.Protocol = &protocol
}

func (c *CreateContainerGroupNetworking) GetServerResponseTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ServerResponseTimeout
}

func (c *CreateContainerGroupNetworking) SetServerResponseTimeout(serverResponseTimeout int64) {
	c.ServerResponseTimeout = &serverResponseTimeout
}

func (c *CreateContainerGroupNetworking) GetSingleConnectionLimit() *bool {
	if c == nil {
		return nil
	}
	return c.SingleConnectionLimit
}

func (c *CreateContainerGroupNetworking) SetSingleConnectionLimit(singleConnectionLimit bool) {
	c.SingleConnectionLimit = &singleConnectionLimit
}

func (c CreateContainerGroupNetworking) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateContainerGroupNetworking to string"
	}
	return string(jsonData)
}
