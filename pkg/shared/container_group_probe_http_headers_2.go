package shared

type ContainerGroupProbeHttpHeaders2 struct {
	Name  *string `json:"name,omitempty" required:"true"`
	Value *string `json:"value,omitempty" required:"true"`
}

func (c *ContainerGroupProbeHttpHeaders2) SetName(name string) {
	c.Name = &name
}

func (c *ContainerGroupProbeHttpHeaders2) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *ContainerGroupProbeHttpHeaders2) SetValue(value string) {
	c.Value = &value
}

func (c *ContainerGroupProbeHttpHeaders2) GetValue() *string {
	if c == nil {
		return nil
	}
	return c.Value
}
