package containergroups

import (
	"encoding/json"
)

// Represents a list of container group instances
type ContainerGroupInstances struct {
	Instances []ContainerGroupInstance `json:"instances,omitempty" required:"true" maxItems:"1000"`
	touched   map[string]bool
}

func (c *ContainerGroupInstances) GetInstances() []ContainerGroupInstance {
	if c == nil {
		return nil
	}
	return c.Instances
}

func (c *ContainerGroupInstances) SetInstances(instances []ContainerGroupInstance) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Instances"] = true
	c.Instances = instances
}

func (c *ContainerGroupInstances) SetInstancesNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Instances"] = true
	c.Instances = nil
}

func (c ContainerGroupInstances) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Instances"] && c.Instances == nil {
		data["instances"] = nil
	} else if c.Instances != nil {
		data["instances"] = c.Instances
	}

	return json.Marshal(data)
}

func (c ContainerGroupInstances) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupInstances to string"
	}
	return string(jsonData)
}
