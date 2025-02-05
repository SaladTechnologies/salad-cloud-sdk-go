package containergroups

import (
	"encoding/json"
)

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
	touched map[string]bool
}

func (c *ContainerGroupInstance) GetInstanceId() *string {
	if c == nil {
		return nil
	}
	return c.InstanceId
}

func (c *ContainerGroupInstance) SetInstanceId(instanceId string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["InstanceId"] = true
	c.InstanceId = &instanceId
}

func (c *ContainerGroupInstance) SetInstanceIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["InstanceId"] = true
	c.InstanceId = nil
}

func (c *ContainerGroupInstance) GetMachineId() *string {
	if c == nil {
		return nil
	}
	return c.MachineId
}

func (c *ContainerGroupInstance) SetMachineId(machineId string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["MachineId"] = true
	c.MachineId = &machineId
}

func (c *ContainerGroupInstance) SetMachineIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["MachineId"] = true
	c.MachineId = nil
}

func (c *ContainerGroupInstance) GetState() *State {
	if c == nil {
		return nil
	}
	return c.State
}

func (c *ContainerGroupInstance) SetState(state State) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["State"] = true
	c.State = &state
}

func (c *ContainerGroupInstance) SetStateNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["State"] = true
	c.State = nil
}

func (c *ContainerGroupInstance) GetUpdateTime() *string {
	if c == nil {
		return nil
	}
	return c.UpdateTime
}

func (c *ContainerGroupInstance) SetUpdateTime(updateTime string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["UpdateTime"] = true
	c.UpdateTime = &updateTime
}

func (c *ContainerGroupInstance) SetUpdateTimeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["UpdateTime"] = true
	c.UpdateTime = nil
}

func (c *ContainerGroupInstance) GetVersion() *int64 {
	if c == nil {
		return nil
	}
	return c.Version
}

func (c *ContainerGroupInstance) SetVersion(version int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Version"] = true
	c.Version = &version
}

func (c *ContainerGroupInstance) SetVersionNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Version"] = true
	c.Version = nil
}

func (c *ContainerGroupInstance) GetReady() *bool {
	if c == nil {
		return nil
	}
	return c.Ready
}

func (c *ContainerGroupInstance) SetReady(ready bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Ready"] = true
	c.Ready = &ready
}

func (c *ContainerGroupInstance) SetReadyNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Ready"] = true
	c.Ready = nil
}

func (c *ContainerGroupInstance) GetStarted() *bool {
	if c == nil {
		return nil
	}
	return c.Started
}

func (c *ContainerGroupInstance) SetStarted(started bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Started"] = true
	c.Started = &started
}

func (c *ContainerGroupInstance) SetStartedNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Started"] = true
	c.Started = nil
}

func (c ContainerGroupInstance) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["InstanceId"] && c.InstanceId == nil {
		data["instance_id"] = nil
	} else if c.InstanceId != nil {
		data["instance_id"] = c.InstanceId
	}

	if c.touched["MachineId"] && c.MachineId == nil {
		data["machine_id"] = nil
	} else if c.MachineId != nil {
		data["machine_id"] = c.MachineId
	}

	if c.touched["State"] && c.State == nil {
		data["state"] = nil
	} else if c.State != nil {
		data["state"] = c.State
	}

	if c.touched["UpdateTime"] && c.UpdateTime == nil {
		data["update_time"] = nil
	} else if c.UpdateTime != nil {
		data["update_time"] = c.UpdateTime
	}

	if c.touched["Version"] && c.Version == nil {
		data["version"] = nil
	} else if c.Version != nil {
		data["version"] = c.Version
	}

	if c.touched["Ready"] && c.Ready == nil {
		data["ready"] = nil
	} else if c.Ready != nil {
		data["ready"] = c.Ready
	}

	if c.touched["Started"] && c.Started == nil {
		data["started"] = nil
	} else if c.Started != nil {
		data["started"] = c.Started
	}

	return json.Marshal(data)
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
