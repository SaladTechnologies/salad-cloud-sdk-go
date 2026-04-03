package organizations

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/unmarshal"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

type GpuAvailabilityPrototype struct {
	// A list of country codes where the resources are available
	CountryCodes []shared.CountryCode `json:"country_codes,omitempty"`
	// The number of available CPU cores
	Cpu *util.Nullable[int64] `json:"cpu,omitempty"`
	// A list of available GPU class names
	GpuClasses []string `json:"gpu_classes,omitempty" required:"true" minItems:"1"`
	// The amount of available memory in MB
	Memory *util.Nullable[int64] `json:"memory,omitempty"`
	// The amount of available storage in bytes
	StorageAmount *util.Nullable[int64] `json:"storage_amount,omitempty"`
}

func (g *GpuAvailabilityPrototype) GetCountryCodes() []shared.CountryCode {
	if g == nil {
		return nil
	}
	return g.CountryCodes
}

func (g *GpuAvailabilityPrototype) SetCountryCodes(countryCodes []shared.CountryCode) {
	g.CountryCodes = countryCodes
}

func (g *GpuAvailabilityPrototype) GetCpu() *util.Nullable[int64] {
	if g == nil {
		return nil
	}
	return g.Cpu
}

func (g *GpuAvailabilityPrototype) SetCpu(cpu util.Nullable[int64]) {
	g.Cpu = &cpu
}

func (g *GpuAvailabilityPrototype) SetCpuNull() {
	g.Cpu = &util.Nullable[int64]{IsNull: true}
}

func (g *GpuAvailabilityPrototype) GetGpuClasses() []string {
	if g == nil {
		return nil
	}
	return g.GpuClasses
}

func (g *GpuAvailabilityPrototype) SetGpuClasses(gpuClasses []string) {
	g.GpuClasses = gpuClasses
}

func (g *GpuAvailabilityPrototype) GetMemory() *util.Nullable[int64] {
	if g == nil {
		return nil
	}
	return g.Memory
}

func (g *GpuAvailabilityPrototype) SetMemory(memory util.Nullable[int64]) {
	g.Memory = &memory
}

func (g *GpuAvailabilityPrototype) SetMemoryNull() {
	g.Memory = &util.Nullable[int64]{IsNull: true}
}

func (g *GpuAvailabilityPrototype) GetStorageAmount() *util.Nullable[int64] {
	if g == nil {
		return nil
	}
	return g.StorageAmount
}

func (g *GpuAvailabilityPrototype) SetStorageAmount(storageAmount util.Nullable[int64]) {
	g.StorageAmount = &storageAmount
}

func (g *GpuAvailabilityPrototype) SetStorageAmountNull() {
	g.StorageAmount = &util.Nullable[int64]{IsNull: true}
}

func (g GpuAvailabilityPrototype) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GpuAvailabilityPrototype to string"
	}
	return string(jsonData)
}

func (g *GpuAvailabilityPrototype) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, g)
}
