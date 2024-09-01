package containergroups

// Represents a list of container group instances
type ContainerGroupInstances struct {
	Instances []ContainerGroupInstance `json:"instances,omitempty" required:"true" maxItems:"1000"`
}

func (c *ContainerGroupInstances) SetInstances(instances []ContainerGroupInstance) {
	c.Instances = instances
}

func (c *ContainerGroupInstances) GetInstances() []ContainerGroupInstance {
	if c == nil {
		return nil
	}
	return c.Instances
}
