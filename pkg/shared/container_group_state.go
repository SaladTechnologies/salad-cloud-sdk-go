package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents the operational state of a container group during its lifecycle, including timing information, status, and instance distribution metrics. This state captures the current execution status, start and finish times, and provides visibility into the operational health across instances.
type ContainerGroupState struct {
	// Optional textual description or notes about the current state of the container group
	Description *util.Nullable[string] `json:"description,omitempty" maxLength:"1000" pattern:"^.*$"`
	// Timestamp when the container group execution finished or is expected to finish
	FinishTime *string `json:"finish_time,omitempty" required:"true"`
	// A summary of container group instances categorized by their current lifecycle status
	InstanceStatusCounts *ContainerGroupInstanceStatusCount `json:"instance_status_counts,omitempty" required:"true"`
	// Timestamp when the container group execution started
	StartTime *string `json:"start_time,omitempty" required:"true"`
	// Represents the current operational state of a container group within the Salad platform.
	Status *ContainerGroupStatus `json:"status,omitempty" required:"true"`
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

func (c *ContainerGroupState) GetStartTime() *string {
	if c == nil {
		return nil
	}
	return c.StartTime
}

func (c *ContainerGroupState) SetStartTime(startTime string) {
	c.StartTime = &startTime
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

func (c ContainerGroupState) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupState to string"
	}
	return string(jsonData)
}
