package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Specifies the resource requirements for a container.
type ContainerResourceRequirements struct {
	// The number of CPU cores required by the container. Must be between 1 and 16.
	Cpu *int64 `json:"cpu,omitempty" required:"true" min:"1" max:"16"`
	// The amount of memory (in MB) required by the container. Must be between 1024 MB and 61440 MB.
	Memory *int64 `json:"memory,omitempty" required:"true" min:"1024" max:"61440"`
	// A list of GPU class UUIDs required by the container. Can be null if no GPU is required.
	GpuClasses *util.Nullable[[]string] `json:"gpu_classes,omitempty" required:"true" maxItems:"100"`
	// The amount of storage (in bytes) required by the container. Must be between 1 GB (1073741824 bytes) and 50 GB (53687091200 bytes).
	StorageAmount *int64 `json:"storage_amount,omitempty" min:"1073741824" max:"53687091200"`
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

func (c *ContainerResourceRequirements) GetMemory() *int64 {
	if c == nil {
		return nil
	}
	return c.Memory
}

func (c *ContainerResourceRequirements) SetMemory(memory int64) {
	c.Memory = &memory
}

func (c *ContainerResourceRequirements) GetGpuClasses() *util.Nullable[[]string] {
	if c == nil {
		return nil
	}
	return c.GpuClasses
}

func (c *ContainerResourceRequirements) SetGpuClasses(gpuClasses util.Nullable[[]string]) {
	c.GpuClasses = &gpuClasses
}

func (c *ContainerResourceRequirements) SetGpuClassesNull() {
	c.GpuClasses = &util.Nullable[[]string]{IsNull: true}
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
