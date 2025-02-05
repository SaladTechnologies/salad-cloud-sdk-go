package shared

import (
	"encoding/json"
)

// Represents a container resource requirements
type ContainerResourceRequirements struct {
	Cpu           *int64   `json:"cpu,omitempty" required:"true" min:"1" max:"16"`
	Memory        *int64   `json:"memory,omitempty" required:"true" min:"1024" max:"61440"`
	GpuClasses    []string `json:"gpu_classes,omitempty"`
	StorageAmount *int64   `json:"storage_amount,omitempty" min:"1073741824" max:"53687091200"`
	touched       map[string]bool
}

func (c *ContainerResourceRequirements) GetCpu() *int64 {
	if c == nil {
		return nil
	}
	return c.Cpu
}

func (c *ContainerResourceRequirements) SetCpu(cpu int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Cpu"] = true
	c.Cpu = &cpu
}

func (c *ContainerResourceRequirements) SetCpuNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Cpu"] = true
	c.Cpu = nil
}

func (c *ContainerResourceRequirements) GetMemory() *int64 {
	if c == nil {
		return nil
	}
	return c.Memory
}

func (c *ContainerResourceRequirements) SetMemory(memory int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Memory"] = true
	c.Memory = &memory
}

func (c *ContainerResourceRequirements) SetMemoryNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Memory"] = true
	c.Memory = nil
}

func (c *ContainerResourceRequirements) GetGpuClasses() []string {
	if c == nil {
		return nil
	}
	return c.GpuClasses
}

func (c *ContainerResourceRequirements) SetGpuClasses(gpuClasses []string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["GpuClasses"] = true
	c.GpuClasses = gpuClasses
}

func (c *ContainerResourceRequirements) SetGpuClassesNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["GpuClasses"] = true
	c.GpuClasses = nil
}

func (c *ContainerResourceRequirements) GetStorageAmount() *int64 {
	if c == nil {
		return nil
	}
	return c.StorageAmount
}

func (c *ContainerResourceRequirements) SetStorageAmount(storageAmount int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StorageAmount"] = true
	c.StorageAmount = &storageAmount
}

func (c *ContainerResourceRequirements) SetStorageAmountNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StorageAmount"] = true
	c.StorageAmount = nil
}

func (c ContainerResourceRequirements) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Cpu"] && c.Cpu == nil {
		data["cpu"] = nil
	} else if c.Cpu != nil {
		data["cpu"] = c.Cpu
	}

	if c.touched["Memory"] && c.Memory == nil {
		data["memory"] = nil
	} else if c.Memory != nil {
		data["memory"] = c.Memory
	}

	if c.touched["GpuClasses"] && c.GpuClasses == nil {
		data["gpu_classes"] = nil
	} else if c.GpuClasses != nil {
		data["gpu_classes"] = c.GpuClasses
	}

	if c.touched["StorageAmount"] && c.StorageAmount == nil {
		data["storage_amount"] = nil
	} else if c.StorageAmount != nil {
		data["storage_amount"] = c.StorageAmount
	}

	return json.Marshal(data)
}

func (c ContainerResourceRequirements) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerResourceRequirements to string"
	}
	return string(jsonData)
}
