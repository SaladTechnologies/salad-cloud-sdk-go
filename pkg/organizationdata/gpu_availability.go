package organizationdata

import "encoding/json"

type GpuAvailability struct {
	// The number of available GPU batches
	AvailableGpuBatch *int64 `json:"available_gpu_batch,omitempty"`
	// The number of available low-end GPUs
	AvailableGpuLow *int64 `json:"available_gpu_low,omitempty"`
	// The number of available medium-end GPUs
	AvailableGpuMedium *int64 `json:"available_gpu_medium,omitempty"`
	// The number of available high-end GPUs
	AvailableGpuHigh *int64 `json:"available_gpu_high,omitempty"`
	// The number of on-call GPUs available
	OnCallGpu *int64 `json:"on_call_gpu,omitempty"`
}

func (g *GpuAvailability) GetAvailableGpuBatch() *int64 {
	if g == nil {
		return nil
	}
	return g.AvailableGpuBatch
}

func (g *GpuAvailability) SetAvailableGpuBatch(availableGpuBatch int64) {
	g.AvailableGpuBatch = &availableGpuBatch
}

func (g *GpuAvailability) GetAvailableGpuLow() *int64 {
	if g == nil {
		return nil
	}
	return g.AvailableGpuLow
}

func (g *GpuAvailability) SetAvailableGpuLow(availableGpuLow int64) {
	g.AvailableGpuLow = &availableGpuLow
}

func (g *GpuAvailability) GetAvailableGpuMedium() *int64 {
	if g == nil {
		return nil
	}
	return g.AvailableGpuMedium
}

func (g *GpuAvailability) SetAvailableGpuMedium(availableGpuMedium int64) {
	g.AvailableGpuMedium = &availableGpuMedium
}

func (g *GpuAvailability) GetAvailableGpuHigh() *int64 {
	if g == nil {
		return nil
	}
	return g.AvailableGpuHigh
}

func (g *GpuAvailability) SetAvailableGpuHigh(availableGpuHigh int64) {
	g.AvailableGpuHigh = &availableGpuHigh
}

func (g *GpuAvailability) GetOnCallGpu() *int64 {
	if g == nil {
		return nil
	}
	return g.OnCallGpu
}

func (g *GpuAvailability) SetOnCallGpu(onCallGpu int64) {
	g.OnCallGpu = &onCallGpu
}

func (g GpuAvailability) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GpuAvailability to string"
	}
	return string(jsonData)
}
