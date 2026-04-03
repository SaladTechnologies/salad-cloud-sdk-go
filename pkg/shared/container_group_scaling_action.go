package shared

import "encoding/json"

// Represents a scaling action configuration for a container group
type ContainerGroupScalingAction struct {
	// The number of replicas to scale to during the scheduled period
	Replicas *int64 `json:"replicas,omitempty" required:"true" min:"0" max:"500"`
	// The cron-style schedule string defining when the scaling should occur
	Schedule *string `json:"schedule,omitempty" required:"true" pattern:"^([0-9A-Za-z*/,-]+)([\t ]+[0-9A-Za-z*/,-]+){4}$"`
}

func (c *ContainerGroupScalingAction) GetReplicas() *int64 {
	if c == nil {
		return nil
	}
	return c.Replicas
}

func (c *ContainerGroupScalingAction) SetReplicas(replicas int64) {
	c.Replicas = &replicas
}

func (c *ContainerGroupScalingAction) GetSchedule() *string {
	if c == nil {
		return nil
	}
	return c.Schedule
}

func (c *ContainerGroupScalingAction) SetSchedule(schedule string) {
	c.Schedule = &schedule
}

func (c ContainerGroupScalingAction) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupScalingAction to string"
	}
	return string(jsonData)
}
