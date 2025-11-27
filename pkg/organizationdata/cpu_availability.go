package organizationdata

import "encoding/json"

type CpuAvailability struct {
	// The number of available CPU cores
	AvailableCpuBatch *int64 `json:"available_cpu_batch,omitempty"`
	// The amount of on-call CPU
	OnCallCpu *int64 `json:"on_call_cpu,omitempty"`
}

func (c *CpuAvailability) GetAvailableCpuBatch() *int64 {
	if c == nil {
		return nil
	}
	return c.AvailableCpuBatch
}

func (c *CpuAvailability) SetAvailableCpuBatch(availableCpuBatch int64) {
	c.AvailableCpuBatch = &availableCpuBatch
}

func (c *CpuAvailability) GetOnCallCpu() *int64 {
	if c == nil {
		return nil
	}
	return c.OnCallCpu
}

func (c *CpuAvailability) SetOnCallCpu(onCallCpu int64) {
	c.OnCallCpu = &onCallCpu
}

func (c CpuAvailability) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CpuAvailability to string"
	}
	return string(jsonData)
}
