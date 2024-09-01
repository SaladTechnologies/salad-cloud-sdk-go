package shared

// Represents a container group state
type ContainerGroupState struct {
	Status      *ContainerGroupStatus `json:"status,omitempty" required:"true"`
	Description *string               `json:"description,omitempty"`
	StartTime   *string               `json:"start_time,omitempty" required:"true"`
	FinishTime  *string               `json:"finish_time,omitempty" required:"true"`
	// Represents a container group instance status count
	InstanceStatusCounts *ContainerGroupInstanceStatusCount `json:"instance_status_counts,omitempty" required:"true"`
}

func (c *ContainerGroupState) SetStatus(status ContainerGroupStatus) {
	c.Status = &status
}

func (c *ContainerGroupState) GetStatus() *ContainerGroupStatus {
	if c == nil {
		return nil
	}
	return c.Status
}

func (c *ContainerGroupState) SetDescription(description string) {
	c.Description = &description
}

func (c *ContainerGroupState) GetDescription() *string {
	if c == nil {
		return nil
	}
	return c.Description
}

func (c *ContainerGroupState) SetStartTime(startTime string) {
	c.StartTime = &startTime
}

func (c *ContainerGroupState) GetStartTime() *string {
	if c == nil {
		return nil
	}
	return c.StartTime
}

func (c *ContainerGroupState) SetFinishTime(finishTime string) {
	c.FinishTime = &finishTime
}

func (c *ContainerGroupState) GetFinishTime() *string {
	if c == nil {
		return nil
	}
	return c.FinishTime
}

func (c *ContainerGroupState) SetInstanceStatusCounts(instanceStatusCounts ContainerGroupInstanceStatusCount) {
	c.InstanceStatusCounts = &instanceStatusCounts
}

func (c *ContainerGroupState) GetInstanceStatusCounts() *ContainerGroupInstanceStatusCount {
	if c == nil {
		return nil
	}
	return c.InstanceStatusCounts
}
