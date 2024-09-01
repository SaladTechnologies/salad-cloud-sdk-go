package shared

// Represents container group networking parameters
type ContainerGroupNetworking struct {
	Protocol *ContainerNetworkingProtocol `json:"protocol,omitempty" required:"true"`
	Port     *int64                       `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	Auth     *bool                        `json:"auth,omitempty" required:"true"`
	Dns      *string                      `json:"dns,omitempty" required:"true"`
}

func (c *ContainerGroupNetworking) SetProtocol(protocol ContainerNetworkingProtocol) {
	c.Protocol = &protocol
}

func (c *ContainerGroupNetworking) GetProtocol() *ContainerNetworkingProtocol {
	if c == nil {
		return nil
	}
	return c.Protocol
}

func (c *ContainerGroupNetworking) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerGroupNetworking) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupNetworking) SetAuth(auth bool) {
	c.Auth = &auth
}

func (c *ContainerGroupNetworking) GetAuth() *bool {
	if c == nil {
		return nil
	}
	return c.Auth
}

func (c *ContainerGroupNetworking) SetDns(dns string) {
	c.Dns = &dns
}

func (c *ContainerGroupNetworking) GetDns() *string {
	if c == nil {
		return nil
	}
	return c.Dns
}
