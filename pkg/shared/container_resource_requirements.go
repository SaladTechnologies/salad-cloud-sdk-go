package shared

import "encoding/json"

// Specifies the resource requirements for a container.
type ContainerResourceRequirements struct {
	// The number of CPU cores required by the container. Must be between 1 and 16.
	Cpu *int64 `json:"cpu,omitempty" required:"true" min:"1" max:"1024"`
	// A list of GPU class UUIDs required by the container. Can be null if no GPU is required.
	GpuClasses []string `json:"gpu_classes,omitempty" required:"true" maxItems:"100"`
	// The amount of memory (in MB) required by the container. Must be between 1024 MB and 61440 MB.
	Memory *int64 `json:"memory,omitempty" required:"true" min:"1024" max:"1073741824"`
	// The size of the shared memory (/dev/shm) in MB. If not specified, defaults to 1024MB.
	ShmSize *int64 `json:"shm_size,omitempty" min:"64" max:"1073741824"`
	// The amount of storage (in bytes) required by the container. Must be between 1 GB (1073741824 bytes) and 250 GB (268435456000 bytes).
	StorageAmount *int64 `json:"storage_amount,omitempty" min:"1073741824" max:"1125899906842624"`
}

func (c *ContainerResourceRequirements) GetCpu() *int64 {
	if c == nil {
		return nil
	}
	return c.Cpu
}

func (c *ContainerResourceRequirements) SetCpu(cpu int64) {
	c.Cpu = &cpu
}

func (c *ContainerResourceRequirements) GetGpuClasses() []string {
	if c == nil {
		return nil
	}
	return c.GpuClasses
}

func (c *ContainerResourceRequirements) SetGpuClasses(gpuClasses []string) {
	c.GpuClasses = gpuClasses
}

func (c *ContainerResourceRequirements) GetMemory() *int64 {
	if c == nil {
		return nil
	}
	return c.Memory
}

func (c *ContainerResourceRequirements) SetMemory(memory int64) {
	c.Memory = &memory
}

func (c *ContainerResourceRequirements) GetShmSize() *int64 {
	if c == nil {
		return nil
	}
	return c.ShmSize
}

func (c *ContainerResourceRequirements) SetShmSize(shmSize int64) {
	c.ShmSize = &shmSize
}

func (c *ContainerResourceRequirements) GetStorageAmount() *int64 {
	if c == nil {
		return nil
	}
	return c.StorageAmount
}

func (c *ContainerResourceRequirements) SetStorageAmount(storageAmount int64) {
	c.StorageAmount = &storageAmount
}

func (c ContainerResourceRequirements) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerResourceRequirements to string"
	}
	return string(jsonData)
}
