package shared

import "encoding/json"

// Represents an HTTP header used for container logging configuration.
type ContainerLoggingHttpHeader struct {
	// The name of the HTTP header
	Name *string `json:"name,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// The value of the HTTP header
	Value *string `json:"value,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
}

func (c *ContainerLoggingHttpHeader) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *ContainerLoggingHttpHeader) SetName(name string) {
	c.Name = &name
}

func (c *ContainerLoggingHttpHeader) GetValue() *string {
	if c == nil {
		return nil
	}
	return c.Value
}

func (c *ContainerLoggingHttpHeader) SetValue(value string) {
	c.Value = &value
}

func (c ContainerLoggingHttpHeader) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerLoggingHttpHeader to string"
	}
	return string(jsonData)
}
