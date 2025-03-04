package shared

import "encoding/json"

// Represents container group networking parameters
type ContainerGroupNetworking struct {
	Protocol              *ContainerNetworkingProtocol          `json:"protocol,omitempty" required:"true"`
	Port                  *int64                                `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	Auth                  *bool                                 `json:"auth,omitempty" required:"true"`
	Dns                   *string                               `json:"dns,omitempty" required:"true"`
	LoadBalancer          *ContainerGroupNetworkingLoadBalancer `json:"load_balancer,omitempty"`
	SingleConnectionLimit *bool                                 `json:"single_connection_limit,omitempty"`
	ClientRequestTimeout  *int64                                `json:"client_request_timeout,omitempty" min:"1" max:"100000"`
	ServerResponseTimeout *int64                                `json:"server_response_timeout,omitempty" min:"1" max:"100000"`
}

func (c *ContainerGroupNetworking) GetProtocol() *ContainerNetworkingProtocol {
	if c == nil {
		return nil
	}
	return c.Protocol
}

func (c *ContainerGroupNetworking) SetProtocol(protocol ContainerNetworkingProtocol) {
	c.Protocol = &protocol
}

func (c *ContainerGroupNetworking) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupNetworking) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerGroupNetworking) GetAuth() *bool {
	if c == nil {
		return nil
	}
	return c.Auth
}

func (c *ContainerGroupNetworking) SetAuth(auth bool) {
	c.Auth = &auth
}

func (c *ContainerGroupNetworking) GetDns() *string {
	if c == nil {
		return nil
	}
	return c.Dns
}

func (c *ContainerGroupNetworking) SetDns(dns string) {
	c.Dns = &dns
}

func (c *ContainerGroupNetworking) GetLoadBalancer() *ContainerGroupNetworkingLoadBalancer {
	if c == nil {
		return nil
	}
	return c.LoadBalancer
}

func (c *ContainerGroupNetworking) SetLoadBalancer(loadBalancer ContainerGroupNetworkingLoadBalancer) {
	c.LoadBalancer = &loadBalancer
}

func (c *ContainerGroupNetworking) GetSingleConnectionLimit() *bool {
	if c == nil {
		return nil
	}
	return c.SingleConnectionLimit
}

func (c *ContainerGroupNetworking) SetSingleConnectionLimit(singleConnectionLimit bool) {
	c.SingleConnectionLimit = &singleConnectionLimit
}

func (c *ContainerGroupNetworking) GetClientRequestTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ClientRequestTimeout
}

func (c *ContainerGroupNetworking) SetClientRequestTimeout(clientRequestTimeout int64) {
	c.ClientRequestTimeout = &clientRequestTimeout
}

func (c *ContainerGroupNetworking) GetServerResponseTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ServerResponseTimeout
}

func (c *ContainerGroupNetworking) SetServerResponseTimeout(serverResponseTimeout int64) {
	c.ServerResponseTimeout = &serverResponseTimeout
}

func (c ContainerGroupNetworking) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupNetworking to string"
	}
	return string(jsonData)
}

type ContainerGroupNetworkingLoadBalancer string

const (
	CONTAINER_GROUP_NETWORKING_LOAD_BALANCER_ROUND_ROBIN                 ContainerGroupNetworkingLoadBalancer = "round_robin"
	CONTAINER_GROUP_NETWORKING_LOAD_BALANCER_LEAST_NUMBER_OF_CONNECTIONS ContainerGroupNetworkingLoadBalancer = "least_number_of_connections"
)
