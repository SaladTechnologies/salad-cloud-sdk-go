package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a container group state
type ContainerGroupState struct {
	Status      *ContainerGroupStatus  `json:"status,omitempty" required:"true"`
	Description *util.Nullable[string] `json:"description,omitempty"`
	StartTime   *string                `json:"start_time,omitempty" required:"true"`
	FinishTime  *string                `json:"finish_time,omitempty" required:"true"`
	// Represents a container group instance status count
	InstanceStatusCounts *ContainerGroupInstanceStatusCount `json:"instance_status_counts,omitempty" required:"true"`
}

func (c *ContainerGroupState) GetStatus() *ContainerGroupStatus {
	if c == nil {
		return nil
	}
	return c.Status
}

func (c *ContainerGroupState) SetStatus(status ContainerGroupStatus) {
	c.Status = &status
}

func (c *ContainerGroupState) GetDescription() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Description
}

func (c *ContainerGroupState) SetDescription(description util.Nullable[string]) {
	c.Description = &description
}

func (c *ContainerGroupState) SetDescriptionNull() {
	c.Description = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerGroupState) GetStartTime() *string {
	if c == nil {
		return nil
	}
	return c.StartTime
}

func (c *ContainerGroupState) SetStartTime(startTime string) {
	c.StartTime = &startTime
}

func (c *ContainerGroupState) GetFinishTime() *string {
	if c == nil {
		return nil
	}
	return c.FinishTime
}

func (c *ContainerGroupState) SetFinishTime(finishTime string) {
	c.FinishTime = &finishTime
}

func (c *ContainerGroupState) GetInstanceStatusCounts() *ContainerGroupInstanceStatusCount {
	if c == nil {
		return nil
	}
	return c.InstanceStatusCounts
}

func (c *ContainerGroupState) SetInstanceStatusCounts(instanceStatusCounts ContainerGroupInstanceStatusCount) {
	c.InstanceStatusCounts = &instanceStatusCounts
}

func (c ContainerGroupState) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupState to string"
	}
	return string(jsonData)
}
