package shared

import "encoding/json"

// Network configuration for container groups that defines connectivity, routing, and access control settings
type ContainerGroupNetworkingConfiguration struct {
	// Whether authentication is required for network access to the container group
	Auth *bool `json:"auth,omitempty" required:"true"`
	// The container group networking client request timeout.
	ClientRequestTimeout *int64 `json:"client_request_timeout,omitempty" min:"1" max:"100000"`
	// Domain name or URL endpoint for the container group's network interface
	Dns *string `json:"dns,omitempty" required:"true" maxLength:"253" minLength:"1" pattern:"^([a-z][a-z0-9-]{0,61}[a-z0-9]\.)*[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The container group networking load balancer.
	LoadBalancer *TheContainerGroupNetworkingLoadBalancer `json:"load_balancer,omitempty" required:"true"`
	// The container group networking port.
	Port *int64 `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	// Defines the communication protocol used for network traffic between containers or external systems. Currently supports HTTP protocol for web-based communication.
	Protocol *ContainerNetworkingProtocol `json:"protocol,omitempty" required:"true"`
	// The container group networking server response timeout.
	ServerResponseTimeout *int64 `json:"server_response_timeout,omitempty" min:"1" max:"100000"`
	// The container group networking single connection limit flag.
	SingleConnectionLimit *bool `json:"single_connection_limit,omitempty"`
}

func (c *ContainerGroupNetworkingConfiguration) GetAuth() *bool {
	if c == nil {
		return nil
	}
	return c.Auth
}

func (c *ContainerGroupNetworkingConfiguration) SetAuth(auth bool) {
	c.Auth = &auth
}

func (c *ContainerGroupNetworkingConfiguration) GetClientRequestTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ClientRequestTimeout
}

func (c *ContainerGroupNetworkingConfiguration) SetClientRequestTimeout(clientRequestTimeout int64) {
	c.ClientRequestTimeout = &clientRequestTimeout
}

func (c *ContainerGroupNetworkingConfiguration) GetDns() *string {
	if c == nil {
		return nil
	}
	return c.Dns
}

func (c *ContainerGroupNetworkingConfiguration) SetDns(dns string) {
	c.Dns = &dns
}

func (c *ContainerGroupNetworkingConfiguration) GetLoadBalancer() *TheContainerGroupNetworkingLoadBalancer {
	if c == nil {
		return nil
	}
	return c.LoadBalancer
}

func (c *ContainerGroupNetworkingConfiguration) SetLoadBalancer(loadBalancer TheContainerGroupNetworkingLoadBalancer) {
	c.LoadBalancer = &loadBalancer
}

func (c *ContainerGroupNetworkingConfiguration) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupNetworkingConfiguration) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerGroupNetworkingConfiguration) GetProtocol() *ContainerNetworkingProtocol {
	if c == nil {
		return nil
	}
	return c.Protocol
}

func (c *ContainerGroupNetworkingConfiguration) SetProtocol(protocol ContainerNetworkingProtocol) {
	c.Protocol = &protocol
}

func (c *ContainerGroupNetworkingConfiguration) GetServerResponseTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ServerResponseTimeout
}

func (c *ContainerGroupNetworkingConfiguration) SetServerResponseTimeout(serverResponseTimeout int64) {
	c.ServerResponseTimeout = &serverResponseTimeout
}

func (c *ContainerGroupNetworkingConfiguration) GetSingleConnectionLimit() *bool {
	if c == nil {
		return nil
	}
	return c.SingleConnectionLimit
}

func (c *ContainerGroupNetworkingConfiguration) SetSingleConnectionLimit(singleConnectionLimit bool) {
	c.SingleConnectionLimit = &singleConnectionLimit
}

func (c ContainerGroupNetworkingConfiguration) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupNetworkingConfiguration to string"
	}
	return string(jsonData)
}
