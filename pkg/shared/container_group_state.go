package shared

import (
	"encoding/json"
)

// Represents a container group state
type ContainerGroupState struct {
	Status      *ContainerGroupStatus `json:"status,omitempty" required:"true"`
	Description *string               `json:"description,omitempty"`
	StartTime   *string               `json:"start_time,omitempty" required:"true"`
	FinishTime  *string               `json:"finish_time,omitempty" required:"true"`
	// Represents a container group instance status count
	InstanceStatusCounts *ContainerGroupInstanceStatusCount `json:"instance_status_counts,omitempty" required:"true"`
	touched              map[string]bool
}

func (c *ContainerGroupState) GetStatus() *ContainerGroupStatus {
	if c == nil {
		return nil
	}
	return c.Status
}

func (c *ContainerGroupState) SetStatus(status ContainerGroupStatus) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Status"] = true
	c.Status = &status
}

func (c *ContainerGroupState) SetStatusNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Status"] = true
	c.Status = nil
}

func (c *ContainerGroupState) GetDescription() *string {
	if c == nil {
		return nil
	}
	return c.Description
}

func (c *ContainerGroupState) SetDescription(description string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Description"] = true
	c.Description = &description
}

func (c *ContainerGroupState) SetDescriptionNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Description"] = true
	c.Description = nil
}

func (c *ContainerGroupState) GetStartTime() *string {
	if c == nil {
		return nil
	}
	return c.StartTime
}

func (c *ContainerGroupState) SetStartTime(startTime string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartTime"] = true
	c.StartTime = &startTime
}

func (c *ContainerGroupState) SetStartTimeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartTime"] = true
	c.StartTime = nil
}

func (c *ContainerGroupState) GetFinishTime() *string {
	if c == nil {
		return nil
	}
	return c.FinishTime
}

func (c *ContainerGroupState) SetFinishTime(finishTime string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["FinishTime"] = true
	c.FinishTime = &finishTime
}

func (c *ContainerGroupState) SetFinishTimeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["FinishTime"] = true
	c.FinishTime = nil
}

func (c *ContainerGroupState) GetInstanceStatusCounts() *ContainerGroupInstanceStatusCount {
	if c == nil {
		return nil
	}
	return c.InstanceStatusCounts
}

func (c *ContainerGroupState) SetInstanceStatusCounts(instanceStatusCounts ContainerGroupInstanceStatusCount) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["InstanceStatusCounts"] = true
	c.InstanceStatusCounts = &instanceStatusCounts
}

func (c *ContainerGroupState) SetInstanceStatusCountsNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["InstanceStatusCounts"] = true
	c.InstanceStatusCounts = nil
}

func (c ContainerGroupState) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Status"] && c.Status == nil {
		data["status"] = nil
	} else if c.Status != nil {
		data["status"] = c.Status
	}

	if c.touched["Description"] && c.Description == nil {
		data["description"] = nil
	} else if c.Description != nil {
		data["description"] = c.Description
	}

	if c.touched["StartTime"] && c.StartTime == nil {
		data["start_time"] = nil
	} else if c.StartTime != nil {
		data["start_time"] = c.StartTime
	}

	if c.touched["FinishTime"] && c.FinishTime == nil {
		data["finish_time"] = nil
	} else if c.FinishTime != nil {
		data["finish_time"] = c.FinishTime
	}

	if c.touched["InstanceStatusCounts"] && c.InstanceStatusCounts == nil {
		data["instance_status_counts"] = nil
	} else if c.InstanceStatusCounts != nil {
		data["instance_status_counts"] = c.InstanceStatusCounts
	}

	return json.Marshal(data)
}

func (c ContainerGroupState) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupState to string"
	}
	return string(jsonData)
}
