package containergroups

import "encoding/json"

// A Container Group Instance represents a running instance of a container group on a specific machine. It provides information about the execution state, readiness, and version of the deployed container group.
type ContainerGroupInstance struct {
	// The container group instance identifier.
	Id *string `json:"id,omitempty" required:"true"`
	// The container group machine identifier.
	MachineId *string `json:"machine_id,omitempty" required:"true"`
	// The state of the container group instance
	State *TheContainerGroupInstanceState `json:"state,omitempty" required:"true"`
	// The UTC timestamp when the container group instance last changed its state. This helps track the lifecycle and state transitions of the instance.
	UpdateTime *string `json:"update_time,omitempty" required:"true"`
	// The version of the container group definition currently running on this instance. Used to track deployment and update progress across the container group fleet.
	Version *int64 `json:"version,omitempty" required:"true" min:"1" max:"2147483647"`
	// Indicates whether the container group instance is currently passing its readiness checks and is able to receive traffic or perform its intended function. If no readiness probe is defined, this will be true once the instance is fully started.
	Ready *bool `json:"ready,omitempty"`
	// Indicates whether the container group instance has successfully completed its startup sequence and passed any configured startup probes. This will always be true when no startup probe is defined for the container group.
	Started *bool `json:"started,omitempty"`
	// The cost of deleting the container group instance
	DeletionCost *int64 `json:"deletion_cost,omitempty" min:"0" max:"100000"`
}

func (c *ContainerGroupInstance) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *ContainerGroupInstance) SetId(id string) {
	c.Id = &id
}

func (c *ContainerGroupInstance) GetMachineId() *string {
	if c == nil {
		return nil
	}
	return c.MachineId
}

func (c *ContainerGroupInstance) SetMachineId(machineId string) {
	c.MachineId = &machineId
}

func (c *ContainerGroupInstance) GetState() *TheContainerGroupInstanceState {
	if c == nil {
		return nil
	}
	return c.State
}

func (c *ContainerGroupInstance) SetState(state TheContainerGroupInstanceState) {
	c.State = &state
}

func (c *ContainerGroupInstance) GetUpdateTime() *string {
	if c == nil {
		return nil
	}
	return c.UpdateTime
}

func (c *ContainerGroupInstance) SetUpdateTime(updateTime string) {
	c.UpdateTime = &updateTime
}

func (c *ContainerGroupInstance) GetVersion() *int64 {
	if c == nil {
		return nil
	}
	return c.Version
}

func (c *ContainerGroupInstance) SetVersion(version int64) {
	c.Version = &version
}

func (c *ContainerGroupInstance) GetReady() *bool {
	if c == nil {
		return nil
	}
	return c.Ready
}

func (c *ContainerGroupInstance) SetReady(ready bool) {
	c.Ready = &ready
}

func (c *ContainerGroupInstance) GetStarted() *bool {
	if c == nil {
		return nil
	}
	return c.Started
}

func (c *ContainerGroupInstance) SetStarted(started bool) {
	c.Started = &started
}

func (c *ContainerGroupInstance) GetDeletionCost() *int64 {
	if c == nil {
		return nil
	}
	return c.DeletionCost
}

func (c *ContainerGroupInstance) SetDeletionCost(deletionCost int64) {
	c.DeletionCost = &deletionCost
}

func (c ContainerGroupInstance) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupInstance to string"
	}
	return string(jsonData)
}
