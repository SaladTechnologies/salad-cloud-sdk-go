package shared

import (
	"encoding/json"
)

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
	touched               map[string]bool
}

func (c *ContainerGroupNetworking) GetProtocol() *ContainerNetworkingProtocol {
	if c == nil {
		return nil
	}
	return c.Protocol
}

func (c *ContainerGroupNetworking) SetProtocol(protocol ContainerNetworkingProtocol) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Protocol"] = true
	c.Protocol = &protocol
}

func (c *ContainerGroupNetworking) SetProtocolNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Protocol"] = true
	c.Protocol = nil
}

func (c *ContainerGroupNetworking) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupNetworking) SetPort(port int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = &port
}

func (c *ContainerGroupNetworking) SetPortNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = nil
}

func (c *ContainerGroupNetworking) GetAuth() *bool {
	if c == nil {
		return nil
	}
	return c.Auth
}

func (c *ContainerGroupNetworking) SetAuth(auth bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Auth"] = true
	c.Auth = &auth
}

func (c *ContainerGroupNetworking) SetAuthNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Auth"] = true
	c.Auth = nil
}

func (c *ContainerGroupNetworking) GetDns() *string {
	if c == nil {
		return nil
	}
	return c.Dns
}

func (c *ContainerGroupNetworking) SetDns(dns string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Dns"] = true
	c.Dns = &dns
}

func (c *ContainerGroupNetworking) SetDnsNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Dns"] = true
	c.Dns = nil
}

func (c *ContainerGroupNetworking) GetLoadBalancer() *ContainerGroupNetworkingLoadBalancer {
	if c == nil {
		return nil
	}
	return c.LoadBalancer
}

func (c *ContainerGroupNetworking) SetLoadBalancer(loadBalancer ContainerGroupNetworkingLoadBalancer) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["LoadBalancer"] = true
	c.LoadBalancer = &loadBalancer
}

func (c *ContainerGroupNetworking) SetLoadBalancerNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["LoadBalancer"] = true
	c.LoadBalancer = nil
}

func (c *ContainerGroupNetworking) GetSingleConnectionLimit() *bool {
	if c == nil {
		return nil
	}
	return c.SingleConnectionLimit
}

func (c *ContainerGroupNetworking) SetSingleConnectionLimit(singleConnectionLimit bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["SingleConnectionLimit"] = true
	c.SingleConnectionLimit = &singleConnectionLimit
}

func (c *ContainerGroupNetworking) SetSingleConnectionLimitNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["SingleConnectionLimit"] = true
	c.SingleConnectionLimit = nil
}

func (c *ContainerGroupNetworking) GetClientRequestTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ClientRequestTimeout
}

func (c *ContainerGroupNetworking) SetClientRequestTimeout(clientRequestTimeout int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ClientRequestTimeout"] = true
	c.ClientRequestTimeout = &clientRequestTimeout
}

func (c *ContainerGroupNetworking) SetClientRequestTimeoutNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ClientRequestTimeout"] = true
	c.ClientRequestTimeout = nil
}

func (c *ContainerGroupNetworking) GetServerResponseTimeout() *int64 {
	if c == nil {
		return nil
	}
	return c.ServerResponseTimeout
}

func (c *ContainerGroupNetworking) SetServerResponseTimeout(serverResponseTimeout int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ServerResponseTimeout"] = true
	c.ServerResponseTimeout = &serverResponseTimeout
}

func (c *ContainerGroupNetworking) SetServerResponseTimeoutNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ServerResponseTimeout"] = true
	c.ServerResponseTimeout = nil
}

func (c ContainerGroupNetworking) MarshalJSON() ([]byte, error) {
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

	if c.touched["Dns"] && c.Dns == nil {
		data["dns"] = nil
	} else if c.Dns != nil {
		data["dns"] = c.Dns
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
