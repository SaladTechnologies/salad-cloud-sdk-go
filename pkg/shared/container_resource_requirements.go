package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a container resource requirements
type ContainerResourceRequirements struct {
	Cpu           *int64                   `json:"cpu,omitempty" required:"true" min:"1" max:"16"`
	Memory        *int64                   `json:"memory,omitempty" required:"true" min:"1024" max:"61440"`
	GpuClasses    *util.Nullable[[]string] `json:"gpu_classes,omitempty"`
	StorageAmount *util.Nullable[int64]    `json:"storage_amount,omitempty" min:"1073741824" max:"53687091200"`
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

func (c *ContainerResourceRequirements) GetStorageAmount() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.StorageAmount
}

func (c *ContainerResourceRequirements) SetStorageAmount(storageAmount util.Nullable[int64]) {
	c.StorageAmount = &storageAmount
}

func (c *ContainerResourceRequirements) SetStorageAmountNull() {
	c.StorageAmount = &util.Nullable[int64]{IsNull: true}
}

func (c ContainerResourceRequirements) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerResourceRequirements to string"
	}
	return string(jsonData)
}
