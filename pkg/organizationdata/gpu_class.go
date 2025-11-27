package organizationdata

import "encoding/json"

// Represents a GPU Class
type GpuClass struct {
	// The unique identifier
	Id *string `json:"id,omitempty" required:"true"`
	// The GPU class name
	Name *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[ -~]{2,63}$"`
	// The list of prices for each container group priority
	Prices []GpuClassPrice `json:"prices,omitempty" required:"true" minItems:"1" maxItems:"100"`
	// Whether the GPU class is in high demand
	IsHighDemand *bool `json:"is_high_demand,omitempty"`
	// The type of GPU class
	GpuClassType *GpuClassType `json:"gpu_class_type,omitempty"`
	// The number of GPUs in the cluster
	GpuCount *int64 `json:"gpu_count,omitempty" min:"1" max:"512"`
	// The minimum vCPU count
	MinVcpu *int64 `json:"min_vcpu,omitempty" min:"0"`
	// The maximum vCPU count
	MaxVcpu *int64 `json:"max_vcpu,omitempty" min:"0"`
	// The minimum RAM amount in MB
	MinRam *int64 `json:"min_ram,omitempty" min:"0"`
	// The maximum RAM amount in MB
	MaxRam *int64 `json:"max_ram,omitempty" min:"0"`
	// The minimum storage amount in bytes
	MinStorage *int64 `json:"min_storage,omitempty" min:"0"`
	// The maximum storage amount in bytes
	MaxStorage *int64 `json:"max_storage,omitempty" min:"0"`
}

func (g *GpuClass) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GpuClass) SetId(id string) {
	g.Id = &id
}

func (g *GpuClass) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GpuClass) SetName(name string) {
	g.Name = &name
}

func (g *GpuClass) GetPrices() []GpuClassPrice {
	if g == nil {
		return nil
	}
	return g.Prices
}

func (g *GpuClass) SetPrices(prices []GpuClassPrice) {
	g.Prices = prices
}

func (g *GpuClass) GetIsHighDemand() *bool {
	if g == nil {
		return nil
	}
	return g.IsHighDemand
}

func (g *GpuClass) SetIsHighDemand(isHighDemand bool) {
	g.IsHighDemand = &isHighDemand
}

func (g *GpuClass) GetGpuClassType() *GpuClassType {
	if g == nil {
		return nil
	}
	return g.GpuClassType
}

func (g *GpuClass) SetGpuClassType(gpuClassType GpuClassType) {
	g.GpuClassType = &gpuClassType
}

func (g *GpuClass) GetGpuCount() *int64 {
	if g == nil {
		return nil
	}
	return g.GpuCount
}

func (g *GpuClass) SetGpuCount(gpuCount int64) {
	g.GpuCount = &gpuCount
}

func (g *GpuClass) GetMinVcpu() *int64 {
	if g == nil {
		return nil
	}
	return g.MinVcpu
}

func (g *GpuClass) SetMinVcpu(minVcpu int64) {
	g.MinVcpu = &minVcpu
}

func (g *GpuClass) GetMaxVcpu() *int64 {
	if g == nil {
		return nil
	}
	return g.MaxVcpu
}

func (g *GpuClass) SetMaxVcpu(maxVcpu int64) {
	g.MaxVcpu = &maxVcpu
}

func (g *GpuClass) GetMinRam() *int64 {
	if g == nil {
		return nil
	}
	return g.MinRam
}

func (g *GpuClass) SetMinRam(minRam int64) {
	g.MinRam = &minRam
}

func (g *GpuClass) GetMaxRam() *int64 {
	if g == nil {
		return nil
	}
	return g.MaxRam
}

func (g *GpuClass) SetMaxRam(maxRam int64) {
	g.MaxRam = &maxRam
}

func (g *GpuClass) GetMinStorage() *int64 {
	if g == nil {
		return nil
	}
	return g.MinStorage
}

func (g *GpuClass) SetMinStorage(minStorage int64) {
	g.MinStorage = &minStorage
}

func (g *GpuClass) GetMaxStorage() *int64 {
	if g == nil {
		return nil
	}
	return g.MaxStorage
}

func (g *GpuClass) SetMaxStorage(maxStorage int64) {
	g.MaxStorage = &maxStorage
}

func (g GpuClass) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GpuClass to string"
	}
	return string(jsonData)
}

// The type of GPU class
type GpuClassType string

const (
	GPU_CLASS_TYPE_COMMUNITY GpuClassType = "community"
	GPU_CLASS_TYPE_SECURE    GpuClassType = "secure"
)
