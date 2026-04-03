package organizations

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/unmarshal"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

type CpuAvailabilityPrototype struct {
	// A list of country codes where the resources are available
	CountryCodes []shared.CountryCode `json:"country_codes,omitempty"`
	// The number of available CPU cores
	Cpu *util.Nullable[int64] `json:"cpu,omitempty"`
	// The amount of available memory in MB
	Memory *util.Nullable[int64] `json:"memory,omitempty"`
	// The amount of available storage in bytes
	StorageAmount *util.Nullable[int64] `json:"storage_amount,omitempty"`
}

func (c *CpuAvailabilityPrototype) GetCountryCodes() []shared.CountryCode {
	if c == nil {
		return nil
	}
	return c.CountryCodes
}

func (c *CpuAvailabilityPrototype) SetCountryCodes(countryCodes []shared.CountryCode) {
	c.CountryCodes = countryCodes
}

func (c *CpuAvailabilityPrototype) GetCpu() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.Cpu
}

func (c *CpuAvailabilityPrototype) SetCpu(cpu util.Nullable[int64]) {
	c.Cpu = &cpu
}

func (c *CpuAvailabilityPrototype) SetCpuNull() {
	c.Cpu = &util.Nullable[int64]{IsNull: true}
}

func (c *CpuAvailabilityPrototype) GetMemory() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.Memory
}

func (c *CpuAvailabilityPrototype) SetMemory(memory util.Nullable[int64]) {
	c.Memory = &memory
}

func (c *CpuAvailabilityPrototype) SetMemoryNull() {
	c.Memory = &util.Nullable[int64]{IsNull: true}
}

func (c *CpuAvailabilityPrototype) GetStorageAmount() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.StorageAmount
}

func (c *CpuAvailabilityPrototype) SetStorageAmount(storageAmount util.Nullable[int64]) {
	c.StorageAmount = &storageAmount
}

func (c *CpuAvailabilityPrototype) SetStorageAmountNull() {
	c.StorageAmount = &util.Nullable[int64]{IsNull: true}
}

func (c CpuAvailabilityPrototype) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CpuAvailabilityPrototype to string"
	}
	return string(jsonData)
}

func (c *CpuAvailabilityPrototype) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
