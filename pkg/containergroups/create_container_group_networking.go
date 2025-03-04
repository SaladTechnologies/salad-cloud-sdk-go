package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents container group networking parameters
type CreateContainerGroupNetworking struct {
	Protocol              *shared.ContainerNetworkingProtocol         `json:"protocol,omitempty" required:"true"`
	Port                  *int64                                      `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	Auth                  *bool                                       `json:"auth,omitempty" required:"true"`
	LoadBalancer          *CreateContainerGroupNetworkingLoadBalancer `json:"load_balancer,omitempty"`
	SingleConnectionLimit *bool                                       `json:"single_connection_limit,omitempty"`
	ClientRequestTimeout  *int64                                      `json:"client_request_timeout,omitempty" min:"1" max:"100000"`
	ServerResponseTimeout *int64                                      `json:"server_response_timeout,omitempty" min:"1" max:"100000"`
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

func (c *CreateContainerGroupNetworking) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *CreateContainerGroupNetworking) SetPort(port int64) {
	c.Port = &port
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

func (c *CreateContainerGroupNetworking) GetLoadBalancer() *CreateContainerGroupNetworkingLoadBalancer {
	if c == nil {
		return nil
	}
	return c.LoadBalancer
}

func (c *CreateContainerGroupNetworking) SetLoadBalancer(loadBalancer CreateContainerGroupNetworkingLoadBalancer) {
	c.LoadBalancer = &loadBalancer
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

func (c *CreateContainerGroupNetworking) GetClientRequestTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ClientRequestTimeout
}

func (c *CreateContainerGroupNetworking) SetClientRequestTimeout(clientRequestTimeout int64) {
	c.ClientRequestTimeout = &clientRequestTimeout
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

func (c CreateContainerGroupNetworking) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateContainerGroupNetworking to string"
	}
	return string(jsonData)
}

type CreateContainerGroupNetworkingLoadBalancer string

const (
	CREATE_CONTAINER_GROUP_NETWORKING_LOAD_BALANCER_ROUND_ROBIN                 CreateContainerGroupNetworkingLoadBalancer = "round_robin"
	CREATE_CONTAINER_GROUP_NETWORKING_LOAD_BALANCER_LEAST_NUMBER_OF_CONNECTIONS CreateContainerGroupNetworkingLoadBalancer = "least_number_of_connections"
)
