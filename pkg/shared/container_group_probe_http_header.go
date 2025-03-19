package shared

import "encoding/json"

type ContainerGroupProbeHttpHeader struct {
	// The name of the HTTP header
	Name *string `json:"name,omitempty" required:"true" maxLength:"256" minLength:"1" pattern:"^.*$"`
	// The value associated with the HTTP header
	Value *string `json:"value,omitempty" required:"true" maxLength:"1024" minLength:"1" pattern:"^.*$"`
}

func (c *ContainerGroupProbeHttpHeader) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *ContainerGroupProbeHttpHeader) SetName(name string) {
	c.Name = &name
}

func (c *ContainerGroupProbeHttpHeader) GetValue() *string {
	if c == nil {
		return nil
	}
	return c.Value
}

func (c *ContainerGroupProbeHttpHeader) SetValue(value string) {
	c.Value = &value
}

func (c ContainerGroupProbeHttpHeader) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupProbeHttpHeader to string"
	}
	return string(jsonData)
}
