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
	touched               map[string]bool
}

func (c *CreateContainerGroupNetworking) GetProtocol() *shared.ContainerNetworkingProtocol {
	if c == nil {
		return nil
	}
	return c.Protocol
}

func (c *CreateContainerGroupNetworking) SetProtocol(protocol shared.ContainerNetworkingProtocol) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Protocol"] = true
	c.Protocol = &protocol
}

func (c *CreateContainerGroupNetworking) SetProtocolNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Protocol"] = true
	c.Protocol = nil
}

func (c *CreateContainerGroupNetworking) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *CreateContainerGroupNetworking) SetPort(port int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = &port
}

func (c *CreateContainerGroupNetworking) SetPortNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = nil
}

func (c *CreateContainerGroupNetworking) GetAuth() *bool {
	if c == nil {
		return nil
	}
	return c.Auth
}

func (c *CreateContainerGroupNetworking) SetAuth(auth bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Auth"] = true
	c.Auth = &auth
}

func (c *CreateContainerGroupNetworking) SetAuthNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Auth"] = true
	c.Auth = nil
}

func (c *CreateContainerGroupNetworking) GetLoadBalancer() *CreateContainerGroupNetworkingLoadBalancer {
	if c == nil {
		return nil
	}
	return c.LoadBalancer
}

func (c *CreateContainerGroupNetworking) SetLoadBalancer(loadBalancer CreateContainerGroupNetworkingLoadBalancer) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["LoadBalancer"] = true
	c.LoadBalancer = &loadBalancer
}

func (c *CreateContainerGroupNetworking) SetLoadBalancerNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["LoadBalancer"] = true
	c.LoadBalancer = nil
}

func (c *CreateContainerGroupNetworking) GetSingleConnectionLimit() *bool {
	if c == nil {
		return nil
	}
	return c.SingleConnectionLimit
}

func (c *CreateContainerGroupNetworking) SetSingleConnectionLimit(singleConnectionLimit bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["SingleConnectionLimit"] = true
	c.SingleConnectionLimit = &singleConnectionLimit
}

func (c *CreateContainerGroupNetworking) SetSingleConnectionLimitNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["SingleConnectionLimit"] = true
	c.SingleConnectionLimit = nil
}

func (c *CreateContainerGroupNetworking) GetClientRequestTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ClientRequestTimeout
}

func (c *CreateContainerGroupNetworking) SetClientRequestTimeout(clientRequestTimeout int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ClientRequestTimeout"] = true
	c.ClientRequestTimeout = &clientRequestTimeout
}

func (c *CreateContainerGroupNetworking) SetClientRequestTimeoutNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ClientRequestTimeout"] = true
	c.ClientRequestTimeout = nil
}

func (c *CreateContainerGroupNetworking) GetServerResponseTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ServerResponseTimeout
}

func (c *CreateContainerGroupNetworking) SetServerResponseTimeout(serverResponseTimeout int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ServerResponseTimeout"] = true
	c.ServerResponseTimeout = &serverResponseTimeout
}

func (c *CreateContainerGroupNetworking) SetServerResponseTimeoutNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ServerResponseTimeout"] = true
	c.ServerResponseTimeout = nil
}

func (c CreateContainerGroupNetworking) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Protocol"] && c.Protocol == nil {
		data["protocol"] = nil
	} else if c.Protocol != nil {
		data["protocol"] = c.Protocol
	}

	if c.touched["Port"] && c.Port == nil {
		data["port"] = nil
	} else if c.Port != nil {
		data["port"] = c.Port
	}

	if c.touched["Auth"] && c.Auth == nil {
		data["auth"] = nil
	} else if c.Auth != nil {
		data["auth"] = c.Auth
	}

	if c.touched["LoadBalancer"] && c.LoadBalancer == nil {
		data["load_balancer"] = nil
	} else if c.LoadBalancer != nil {
		data["load_balancer"] = c.LoadBalancer
	}

	if c.touched["SingleConnectionLimit"] && c.SingleConnectionLimit == nil {
		data["single_connection_limit"] = nil
	} else if c.SingleConnectionLimit != nil {
		data["single_connection_limit"] = c.SingleConnectionLimit
	}

	if c.touched["ClientRequestTimeout"] && c.ClientRequestTimeout == nil {
		data["client_request_timeout"] = nil
	} else if c.ClientRequestTimeout != nil {
		data["client_request_timeout"] = c.ClientRequestTimeout
	}

	if c.touched["ServerResponseTimeout"] && c.ServerResponseTimeout == nil {
		data["server_response_timeout"] = nil
	} else if c.ServerResponseTimeout != nil {
		data["server_response_timeout"] = c.ServerResponseTimeout
	}

	return json.Marshal(data)
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
