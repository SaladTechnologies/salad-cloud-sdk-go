package containergroups

import "encoding/json"

// Authentication details for Google Artifact Registry (GAR)
type ContainerRegistryAuthenticationGcpGar struct {
	// GCP service account key in JSON format for GAR authentication
	ServiceKey *string `json:"service_key,omitempty" required:"true" maxLength:"10000" minLength:"1" pattern:"^.*$"`
}

func (c *ContainerRegistryAuthenticationGcpGar) GetServiceKey() *string {
	if c == nil {
		return nil
	}
	return c.ServiceKey
}

func (c *ContainerRegistryAuthenticationGcpGar) SetServiceKey(serviceKey string) {
	c.ServiceKey = &serviceKey
}

func (c ContainerRegistryAuthenticationGcpGar) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerRegistryAuthenticationGcpGar to string"
	}
	return string(jsonData)
}
