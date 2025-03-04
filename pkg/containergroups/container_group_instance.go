package containergroups

import "encoding/json"

// Represents the details of a single container group instance
type ContainerGroupInstance struct {
	// The unique instance ID
	InstanceId *string `json:"instance_id,omitempty" required:"true"`
	// The machine ID
	MachineId *string `json:"machine_id,omitempty" required:"true"`
	// The state of the container group instance
	State *State `json:"state,omitempty" required:"true"`
	// The UTC date & time when the workload on this machine transitioned to the current state
	UpdateTime *string `json:"update_time,omitempty" required:"true"`
	// The version of the running container group
	Version *int64 `json:"version,omitempty" required:"true" min:"1"`
	// Specifies whether the container group instance is currently passing its readiness check. If no readiness probe is defined, is true once fully started.
	Ready *bool `json:"ready,omitempty"`
	// Specifies whether the container group instance passed its startup probe. Is always true when no startup probe is defined.
	Started *bool `json:"started,omitempty"`
}

func (c *ContainerGroupInstance) GetInstanceId() *string {
	if c == nil {
		return nil
	}
	return c.InstanceId
}

func (c *ContainerGroupInstance) SetInstanceId(instanceId string) {
	c.InstanceId = &instanceId
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

func (c *ContainerGroupInstance) GetState() *State {
	if c == nil {
		return nil
	}
	return c.State
}

func (c *ContainerGroupInstance) SetState(state State) {
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

func (c ContainerGroupInstance) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupInstance to string"
	}
	return string(jsonData)
}

// The state of the container group instance
type State string

const (
	STATE_ALLOCATING  State = "allocating"
	STATE_DOWNLOADING State = "downloading"
	STATE_CREATING    State = "creating"
	STATE_RUNNING     State = "running"
	STATE_STOPPING    State = "stopping"
)
