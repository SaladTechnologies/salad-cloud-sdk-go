package shared

// Represents a container resource requirements
type ContainerResourceRequirements struct {
	Cpu           *int64   `json:"cpu,omitempty" required:"true" min:"1" max:"16"`
	Memory        *int64   `json:"memory,omitempty" required:"true" min:"1024" max:"30720"`
	GpuClasses    []string `json:"gpu_classes,omitempty"`
	StorageAmount *int64   `json:"storage_amount,omitempty" min:"1073741824" max:"53687091200"`
}

func (c *ContainerResourceRequirements) SetCpu(cpu int64) {
	c.Cpu = &cpu
}

func (c *ContainerResourceRequirements) GetCpu() *int64 {
	if c == nil {
		return nil
	}
	return c.Cpu
}

func (c *ContainerResourceRequirements) SetMemory(memory int64) {
	c.Memory = &memory
}

func (c *ContainerResourceRequirements) GetMemory() *int64 {
	if c == nil {
		return nil
	}
	return c.Memory
}

func (c *ContainerResourceRequirements) SetGpuClasses(gpuClasses []string) {
	c.GpuClasses = gpuClasses
}

func (c *ContainerResourceRequirements) GetGpuClasses() []string {
	if c == nil {
		return nil
	}
	return c.GpuClasses
}

func (c *ContainerResourceRequirements) SetStorageAmount(storageAmount int64) {
	c.StorageAmount = &storageAmount
}

func (c *ContainerResourceRequirements) GetStorageAmount() *int64 {
	if c == nil {
		return nil
	}
	return c.StorageAmount
}
