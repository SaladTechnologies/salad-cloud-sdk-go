package containergroups

import "encoding/json"

// Specifies the resource requirements for creating a container.
type CreateContainerResourceRequirements struct {
	// The number of CPU cores required by the container. Must be between 1 and 16.
	Cpu *int64 `json:"cpu,omitempty" required:"true" min:"1" max:"16"`
	// The amount of memory (in MB) required by the container. Must be between 1024 MB and 61440 MB.
	Memory *int64 `json:"memory,omitempty" required:"true" min:"1024" max:"61440"`
	// A list of GPU class UUIDs required by the container. Can be null if no GPU is required.
	GpuClasses []string `json:"gpu_classes,omitempty" maxItems:"100"`
	// The amount of storage (in bytes) required by the container. Must be between 1 GB (1073741824 bytes) and 250 GB (268435456000 bytes).
	StorageAmount *int64 `json:"storage_amount,omitempty" min:"1073741824" max:"268435456000"`
	// The size of the shared memory (/dev/shm) in MB. If not specified, defaults to 64MB.
	ShmSize *int64 `json:"shm_size,omitempty" min:"64" max:"2147483647"`
}

func (c *CreateContainerResourceRequirements) GetCpu() *int64 {
	if c == nil {
		return nil
	}
	return c.Cpu
}

func (c *CreateContainerResourceRequirements) SetCpu(cpu int64) {
	c.Cpu = &cpu
}

func (c *CreateContainerResourceRequirements) GetMemory() *int64 {
	if c == nil {
		return nil
	}
	return c.Memory
}

func (c *CreateContainerResourceRequirements) SetMemory(memory int64) {
	c.Memory = &memory
}

func (c *CreateContainerResourceRequirements) GetGpuClasses() []string {
	if c == nil {
		return nil
	}
	return c.GpuClasses
}

func (c *CreateContainerResourceRequirements) SetGpuClasses(gpuClasses []string) {
	c.GpuClasses = gpuClasses
}

func (c *CreateContainerResourceRequirements) GetStorageAmount() *int64 {
	if c == nil {
		return nil
	}
	return c.StorageAmount
}

func (c *CreateContainerResourceRequirements) SetStorageAmount(storageAmount int64) {
	c.StorageAmount = &storageAmount
}

func (c *CreateContainerResourceRequirements) GetShmSize() *int64 {
	if c == nil {
		return nil
	}
	return c.ShmSize
}

func (c *CreateContainerResourceRequirements) SetShmSize(shmSize int64) {
	c.ShmSize = &shmSize
}

func (c CreateContainerResourceRequirements) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateContainerResourceRequirements to string"
	}
	return string(jsonData)
}
