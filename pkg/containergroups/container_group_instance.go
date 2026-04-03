package containergroups

import "encoding/json"

// A Container Group Instance represents a running instance of a container group on a specific machine. It provides information about the execution state, readiness, and version of the deployed container group.
type ContainerGroupInstance struct {
	// The percentage of CPU used by this container group instance. This is updated every minute.
	CpuPercent *float64 `json:"cpu_percent,omitempty" min:"0"`
	// The total CPU usage in seconds for this container group instance. This is updated every minute.
	CpuUsage *int64 `json:"cpu_usage,omitempty" min:"0"`
	// The total CPU usage in seconds for this container group instance since it was started. This is updated every minute.
	CpuUsageTotal *int64 `json:"cpu_usage_total,omitempty" min:"0"`
	// The cost of deleting the container group instance
	DeletionCost *int64 `json:"deletion_cost,omitempty" min:"0" max:"100000"`
	// The container group instance identifier.
	Id *string `json:"id,omitempty" required:"true"`
	// The container group machine identifier.
	MachineId *string `json:"machine_id,omitempty" required:"true"`
	// The memory usage in MB for this container group instance. This is updated every minute.
	MemoryUsageMb *float64 `json:"memory_usage_mb,omitempty" min:"0"`
	// The percentage of memory used by this container group instance. This is updated every minute.
	MemoryUsagePercent *float64 `json:"memory_usage_percent,omitempty" min:"0"`
	// The progress percentage of pulling the container image. This is only relevant when the instance state is 'downloading'.
	PullingProgress *float64 `json:"pulling_progress,omitempty" min:"0" max:"100"`
	// Indicates whether the container group instance is currently passing its readiness checks and is able to receive traffic or perform its intended function. If no readiness probe is defined, this will be true once the instance is fully started.
	Ready *bool `json:"ready,omitempty"`
	// The SSH host key fingerprint of the container group instance
	SshHostKeyFingerprint *string `json:"ssh_host_key_fingerprint,omitempty" maxLength:"256" minLength:"1"`
	// The SSH IP address of the container group instance
	SshIp *string `json:"ssh_ip,omitempty"`
	// The SSH port of the container group instance
	SshPort *int64 `json:"ssh_port,omitempty" min:"1" max:"65535"`
	// Indicates whether the container group instance has successfully completed its startup sequence and passed any configured startup probes. This will always be true when no startup probe is defined for the container group.
	Started *bool `json:"started,omitempty"`
	// The state of the container group instance
	State *TheContainerGroupInstanceState `json:"state,omitempty" required:"true"`
	// The UTC timestamp when the container group instance last changed its state. This helps track the lifecycle and state transitions of the instance.
	UpdateTime *string `json:"update_time,omitempty" required:"true"`
	// The version of the container group definition currently running on this instance. Used to track deployment and update progress across the container group fleet.
	Version *int64 `json:"version,omitempty" required:"true" min:"1" max:"2147483647"`
}

func (c *ContainerGroupInstance) GetCpuPercent() *float64 {
	if c == nil {
		return nil
	}
	return c.CpuPercent
}

func (c *ContainerGroupInstance) SetCpuPercent(cpuPercent float64) {
	c.CpuPercent = &cpuPercent
}

func (c *ContainerGroupInstance) GetCpuUsage() *int64 {
	if c == nil {
		return nil
	}
	return c.CpuUsage
}

func (c *ContainerGroupInstance) SetCpuUsage(cpuUsage int64) {
	c.CpuUsage = &cpuUsage
}

func (c *ContainerGroupInstance) GetCpuUsageTotal() *int64 {
	if c == nil {
		return nil
	}
	return c.CpuUsageTotal
}

func (c *ContainerGroupInstance) SetCpuUsageTotal(cpuUsageTotal int64) {
	c.CpuUsageTotal = &cpuUsageTotal
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

func (c *ContainerGroupInstance) GetMemoryUsageMb() *float64 {
	if c == nil {
		return nil
	}
	return c.MemoryUsageMb
}

func (c *ContainerGroupInstance) SetMemoryUsageMb(memoryUsageMb float64) {
	c.MemoryUsageMb = &memoryUsageMb
}

func (c *ContainerGroupInstance) GetMemoryUsagePercent() *float64 {
	if c == nil {
		return nil
	}
	return c.MemoryUsagePercent
}

func (c *ContainerGroupInstance) SetMemoryUsagePercent(memoryUsagePercent float64) {
	c.MemoryUsagePercent = &memoryUsagePercent
}

func (c *ContainerGroupInstance) GetPullingProgress() *float64 {
	if c == nil {
		return nil
	}
	return c.PullingProgress
}

func (c *ContainerGroupInstance) SetPullingProgress(pullingProgress float64) {
	c.PullingProgress = &pullingProgress
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

func (c *ContainerGroupInstance) GetSshHostKeyFingerprint() *string {
	if c == nil {
		return nil
	}
	return c.SshHostKeyFingerprint
}

func (c *ContainerGroupInstance) SetSshHostKeyFingerprint(sshHostKeyFingerprint string) {
	c.SshHostKeyFingerprint = &sshHostKeyFingerprint
}

func (c *ContainerGroupInstance) GetSshIp() *string {
	if c == nil {
		return nil
	}
	return c.SshIp
}

func (c *ContainerGroupInstance) SetSshIp(sshIp string) {
	c.SshIp = &sshIp
}

func (c *ContainerGroupInstance) GetSshPort() *int64 {
	if c == nil {
		return nil
	}
	return c.SshPort
}

func (c *ContainerGroupInstance) SetSshPort(sshPort int64) {
	c.SshPort = &sshPort
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

func (c ContainerGroupInstance) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupInstance to string"
	}
	return string(jsonData)
}
