package shared

import "encoding/json"

// A summary of container group instances categorized by their current lifecycle status
type ContainerGroupInstanceStatusCount struct {
	// The number of container instances that are currently being allocated resources
	AllocatingCount *int64 `json:"allocating_count,omitempty" required:"true" min:"0" max:"2147483647"`
	// The number of container instances that are in the process of being created
	CreatingCount *int64 `json:"creating_count,omitempty" required:"true" min:"0" max:"2147483647"`
	// The number of container instances that are currently running and operational
	RunningCount *int64 `json:"running_count,omitempty" required:"true" min:"0" max:"2147483647"`
	// The number of container instances that are in the process of stopping
	StoppingCount *int64 `json:"stopping_count,omitempty" required:"true" min:"0" max:"2147483647"`
}

func (c *ContainerGroupInstanceStatusCount) GetAllocatingCount() *int64 {
	if c == nil {
		return nil
	}
	return c.AllocatingCount
}

func (c *ContainerGroupInstanceStatusCount) SetAllocatingCount(allocatingCount int64) {
	c.AllocatingCount = &allocatingCount
}

func (c *ContainerGroupInstanceStatusCount) GetCreatingCount() *int64 {
	if c == nil {
		return nil
	}
	return c.CreatingCount
}

func (c *ContainerGroupInstanceStatusCount) SetCreatingCount(creatingCount int64) {
	c.CreatingCount = &creatingCount
}

func (c *ContainerGroupInstanceStatusCount) GetRunningCount() *int64 {
	if c == nil {
		return nil
	}
	return c.RunningCount
}

func (c *ContainerGroupInstanceStatusCount) SetRunningCount(runningCount int64) {
	c.RunningCount = &runningCount
}

func (c *ContainerGroupInstanceStatusCount) GetStoppingCount() *int64 {
	if c == nil {
		return nil
	}
	return c.StoppingCount
}

func (c *ContainerGroupInstanceStatusCount) SetStoppingCount(stoppingCount int64) {
	c.StoppingCount = &stoppingCount
}

func (c ContainerGroupInstanceStatusCount) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupInstanceStatusCount to string"
	}
	return string(jsonData)
}
