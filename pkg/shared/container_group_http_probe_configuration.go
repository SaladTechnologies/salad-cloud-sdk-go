package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Defines HTTP probe configuration for container health checks within a container group.
type ContainerGroupHttpProbeConfiguration struct {
	// A collection of HTTP header name-value pairs used for configuring requests and responses in container group endpoints. Each header consists of a name and its corresponding value.
	Headers []ContainerGroupProbeHttpHeader `json:"headers,omitempty" required:"true" minItems:"1" maxItems:"50"`
	// The HTTP path that will be probed to check container health.
	Path *string `json:"path,omitempty" required:"true" maxLength:"2048" minLength:"1" pattern:"^.*$"`
	// The TCP port number to which the HTTP request will be sent.
	Port *int64 `json:"port,omitempty" required:"true" min:"0" max:"65536"`
	// The protocol scheme used for HTTP probe requests in container health checks.
	Scheme *util.Nullable[HttpScheme] `json:"scheme,omitempty" required:"true"`
}

func (c *ContainerGroupHttpProbeConfiguration) GetHeaders() []ContainerGroupProbeHttpHeader {
	if c == nil {
		return nil
	}
	return c.Headers
}

func (c *ContainerGroupHttpProbeConfiguration) SetHeaders(headers []ContainerGroupProbeHttpHeader) {
	c.Headers = headers
}

func (c *ContainerGroupHttpProbeConfiguration) GetPath() *string {
	if c == nil {
		return nil
	}
	return c.Path
}

func (c *ContainerGroupHttpProbeConfiguration) SetPath(path string) {
	c.Path = &path
}

func (c *ContainerGroupHttpProbeConfiguration) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupHttpProbeConfiguration) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerGroupHttpProbeConfiguration) GetScheme() *util.Nullable[HttpScheme] {
	if c == nil {
		return nil
	}
	return c.Scheme
}

func (c *ContainerGroupHttpProbeConfiguration) SetScheme(scheme util.Nullable[HttpScheme]) {
	c.Scheme = &scheme
}

func (c *ContainerGroupHttpProbeConfiguration) SetSchemeNull() {
	c.Scheme = &util.Nullable[HttpScheme]{IsNull: true}
}

func (c ContainerGroupHttpProbeConfiguration) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupHttpProbeConfiguration to string"
	}
	return string(jsonData)
}
