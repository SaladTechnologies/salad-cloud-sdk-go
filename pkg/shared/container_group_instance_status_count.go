package shared

// Represents a container group instance status count
type ContainerGroupInstanceStatusCount struct {
	AllocatingCount *int64 `json:"allocating_count,omitempty" required:"true" min:"0"`
	CreatingCount   *int64 `json:"creating_count,omitempty" required:"true" min:"0"`
	RunningCount    *int64 `json:"running_count,omitempty" required:"true" min:"0"`
	StoppingCount   *int64 `json:"stopping_count,omitempty" required:"true" min:"0"`
}

func (c *ContainerGroupInstanceStatusCount) SetAllocatingCount(allocatingCount int64) {
	c.AllocatingCount = &allocatingCount
}

func (c *ContainerGroupInstanceStatusCount) GetAllocatingCount() *int64 {
	if c == nil {
		return nil
	}
	return c.AllocatingCount
}

func (c *ContainerGroupInstanceStatusCount) SetCreatingCount(creatingCount int64) {
	c.CreatingCount = &creatingCount
}

func (c *ContainerGroupInstanceStatusCount) GetCreatingCount() *int64 {
	if c == nil {
		return nil
	}
	return c.CreatingCount
}

func (c *ContainerGroupInstanceStatusCount) SetRunningCount(runningCount int64) {
	c.RunningCount = &runningCount
}

func (c *ContainerGroupInstanceStatusCount) GetRunningCount() *int64 {
	if c == nil {
		return nil
	}
	return c.RunningCount
}

func (c *ContainerGroupInstanceStatusCount) SetStoppingCount(stoppingCount int64) {
	c.StoppingCount = &stoppingCount
}

func (c *ContainerGroupInstanceStatusCount) GetStoppingCount() *int64 {
	if c == nil {
		return nil
	}
	return c.StoppingCount
}
