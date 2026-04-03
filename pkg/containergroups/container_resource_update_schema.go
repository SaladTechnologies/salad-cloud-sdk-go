package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/unmarshal"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Defines the resource specifications that can be modified for a container group, including CPU, memory, GPU classes, and storage allocations.
type ContainerResourceUpdateSchema struct {
	// The number of CPU cores to allocate to the container (between 1 and 1024).
	Cpu *util.Nullable[int64] `json:"cpu,omitempty" min:"1" max:"1024"`
	// List of GPU class identifiers that the container can use, specified as UUIDs.
	GpuClasses *util.Nullable[[]string] `json:"gpu_classes,omitempty" maxItems:"100"`
	// The amount of memory to allocate to the container in megabytes (between 1024 and 1073741824).
	Memory *util.Nullable[int64] `json:"memory,omitempty" min:"1024" max:"1073741824"`
	// The amount of shared memory to allocate to the container via `/dev/shm` in megabytes (between 64 and 1073741824). If not specified, defaults to 64 MB.
	ShmSize *util.Nullable[int64] `json:"shm_size,omitempty" min:"64" max:"1073741824"`
	// The amount of storage to allocate to the container in bytes (between 1 GB and 1 PB).
	StorageAmount *util.Nullable[int64] `json:"storage_amount,omitempty" min:"1073741824" max:"1125899906842624"`
}

func (c *ContainerResourceUpdateSchema) GetCpu() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.Cpu
}

func (c *ContainerResourceUpdateSchema) SetCpu(cpu util.Nullable[int64]) {
	c.Cpu = &cpu
}

func (c *ContainerResourceUpdateSchema) SetCpuNull() {
	c.Cpu = &util.Nullable[int64]{IsNull: true}
}

func (c *ContainerResourceUpdateSchema) GetGpuClasses() *util.Nullable[[]string] {
	if c == nil {
		return nil
	}
	return c.GpuClasses
}

func (c *ContainerResourceUpdateSchema) SetGpuClasses(gpuClasses util.Nullable[[]string]) {
	c.GpuClasses = &gpuClasses
}

func (c *ContainerResourceUpdateSchema) SetGpuClassesNull() {
	c.GpuClasses = &util.Nullable[[]string]{IsNull: true}
}

func (c *ContainerResourceUpdateSchema) GetMemory() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.Memory
}

func (c *ContainerResourceUpdateSchema) SetMemory(memory util.Nullable[int64]) {
	c.Memory = &memory
}

func (c *ContainerResourceUpdateSchema) SetMemoryNull() {
	c.Memory = &util.Nullable[int64]{IsNull: true}
}

func (c *ContainerResourceUpdateSchema) GetShmSize() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.ShmSize
}

func (c *ContainerResourceUpdateSchema) SetShmSize(shmSize util.Nullable[int64]) {
	c.ShmSize = &shmSize
}

func (c *ContainerResourceUpdateSchema) SetShmSizeNull() {
	c.ShmSize = &util.Nullable[int64]{IsNull: true}
}

func (c *ContainerResourceUpdateSchema) GetStorageAmount() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.StorageAmount
}

func (c *ContainerResourceUpdateSchema) SetStorageAmount(storageAmount util.Nullable[int64]) {
	c.StorageAmount = &storageAmount
}

func (c *ContainerResourceUpdateSchema) SetStorageAmountNull() {
	c.StorageAmount = &util.Nullable[int64]{IsNull: true}
}

func (c ContainerResourceUpdateSchema) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerResourceUpdateSchema to string"
	}
	return string(jsonData)
}

func (c *ContainerResourceUpdateSchema) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
