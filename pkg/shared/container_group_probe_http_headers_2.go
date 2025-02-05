package shared

import (
	"encoding/json"
)

type ContainerGroupProbeHttpHeaders2 struct {
	Name    *string `json:"name,omitempty" required:"true"`
	Value   *string `json:"value,omitempty" required:"true"`
	touched map[string]bool
}

func (c *ContainerGroupProbeHttpHeaders2) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *ContainerGroupProbeHttpHeaders2) SetName(name string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = &name
}

func (c *ContainerGroupProbeHttpHeaders2) SetNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = nil
}

func (c *ContainerGroupProbeHttpHeaders2) GetValue() *string {
	if c == nil {
		return nil
	}
	return c.Value
}

func (c *ContainerGroupProbeHttpHeaders2) SetValue(value string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Value"] = true
	c.Value = &value
}

func (c *ContainerGroupProbeHttpHeaders2) SetValueNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Value"] = true
	c.Value = nil
}

func (c ContainerGroupProbeHttpHeaders2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Name"] && c.Name == nil {
		data["name"] = nil
	} else if c.Name != nil {
		data["name"] = c.Name
	}

	if c.touched["Value"] && c.Value == nil {
		data["value"] = nil
	} else if c.Value != nil {
		data["value"] = c.Value
	}

	return json.Marshal(data)
}

func (c ContainerGroupProbeHttpHeaders2) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupProbeHttpHeaders2 to string"
	}
	return string(jsonData)
}
