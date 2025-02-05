package shared

import (
	"encoding/json"
)

// Represents a container group instance status count
type ContainerGroupInstanceStatusCount struct {
	AllocatingCount *int64 `json:"allocating_count,omitempty" required:"true" min:"0"`
	CreatingCount   *int64 `json:"creating_count,omitempty" required:"true" min:"0"`
	RunningCount    *int64 `json:"running_count,omitempty" required:"true" min:"0"`
	StoppingCount   *int64 `json:"stopping_count,omitempty" required:"true" min:"0"`
	touched         map[string]bool
}

func (c *ContainerGroupInstanceStatusCount) GetAllocatingCount() *int64 {
	if c == nil {
		return nil
	}
	return c.AllocatingCount
}

func (c *ContainerGroupInstanceStatusCount) SetAllocatingCount(allocatingCount int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AllocatingCount"] = true
	c.AllocatingCount = &allocatingCount
}

func (c *ContainerGroupInstanceStatusCount) SetAllocatingCountNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AllocatingCount"] = true
	c.AllocatingCount = nil
}

func (c *ContainerGroupInstanceStatusCount) GetCreatingCount() *int64 {
	if c == nil {
		return nil
	}
	return c.CreatingCount
}

func (c *ContainerGroupInstanceStatusCount) SetCreatingCount(creatingCount int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CreatingCount"] = true
	c.CreatingCount = &creatingCount
}

func (c *ContainerGroupInstanceStatusCount) SetCreatingCountNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CreatingCount"] = true
	c.CreatingCount = nil
}

func (c *ContainerGroupInstanceStatusCount) GetRunningCount() *int64 {
	if c == nil {
		return nil
	}
	return c.RunningCount
}

func (c *ContainerGroupInstanceStatusCount) SetRunningCount(runningCount int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["RunningCount"] = true
	c.RunningCount = &runningCount
}

func (c *ContainerGroupInstanceStatusCount) SetRunningCountNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["RunningCount"] = true
	c.RunningCount = nil
}

func (c *ContainerGroupInstanceStatusCount) GetStoppingCount() *int64 {
	if c == nil {
		return nil
	}
	return c.StoppingCount
}

func (c *ContainerGroupInstanceStatusCount) SetStoppingCount(stoppingCount int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StoppingCount"] = true
	c.StoppingCount = &stoppingCount
}

func (c *ContainerGroupInstanceStatusCount) SetStoppingCountNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StoppingCount"] = true
	c.StoppingCount = nil
}

func (c ContainerGroupInstanceStatusCount) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["AllocatingCount"] && c.AllocatingCount == nil {
		data["allocating_count"] = nil
	} else if c.AllocatingCount != nil {
		data["allocating_count"] = c.AllocatingCount
	}

	if c.touched["CreatingCount"] && c.CreatingCount == nil {
		data["creating_count"] = nil
	} else if c.CreatingCount != nil {
		data["creating_count"] = c.CreatingCount
	}

	if c.touched["RunningCount"] && c.RunningCount == nil {
		data["running_count"] = nil
	} else if c.RunningCount != nil {
		data["running_count"] = c.RunningCount
	}

	if c.touched["StoppingCount"] && c.StoppingCount == nil {
		data["stopping_count"] = nil
	} else if c.StoppingCount != nil {
		data["stopping_count"] = c.StoppingCount
	}

	return json.Marshal(data)
}

func (c ContainerGroupInstanceStatusCount) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupInstanceStatusCount to string"
	}
	return string(jsonData)
}
