package containergroups

import "encoding/json"

// A collection of container group instances returned as part of a paginated response or batch operation result.
type ContainerGroupInstanceCollection struct {
	// An array of container group instances, each representing a deployed container group with its current state and configuration information.
	Instances []ContainerGroupInstance `json:"instances,omitempty" required:"true" maxItems:"1000"`
}

func (c *ContainerGroupInstanceCollection) GetInstances() []ContainerGroupInstance {
	if c == nil {
		return nil
	}
	return c.Instances
}

func (c *ContainerGroupInstanceCollection) SetInstances(instances []ContainerGroupInstance) {
	c.Instances = instances
}

func (c ContainerGroupInstanceCollection) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupInstanceCollection to string"
	}
	return string(jsonData)
}
