package containergroups

import "encoding/json"

// Authentication details for Google Container Registry (GCR)
type ContainerRegistryAuthenticationGcpGcr struct {
	// GCP service account key in JSON format for GCR authentication
	ServiceKey *string `json:"service_key,omitempty" required:"true" maxLength:"10000" minLength:"1" pattern:"^.*$"`
}

func (c *ContainerRegistryAuthenticationGcpGcr) GetServiceKey() *string {
	if c == nil {
		return nil
	}
	return c.ServiceKey
}

func (c *ContainerRegistryAuthenticationGcpGcr) SetServiceKey(serviceKey string) {
	c.ServiceKey = &serviceKey
}

func (c ContainerRegistryAuthenticationGcpGcr) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerRegistryAuthenticationGcpGcr to string"
	}
	return string(jsonData)
}
