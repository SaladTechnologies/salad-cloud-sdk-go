package organizationdata

// Represents a GPU Class
type GpuClass struct {
	// The unique identifier
	Id *string `json:"id,omitempty" required:"true"`
	// The GPU class name
	Name *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2"`
	// The list of prices for each container group priority
	Prices []GpuClassPrice `json:"prices,omitempty" required:"true" minItems:"1" maxItems:"100"`
	// Whether the GPU class is in high demand
	IsHighDemand *bool `json:"is_high_demand,omitempty"`
}

func (g *GpuClass) SetId(id string) {
	g.Id = &id
}

func (g *GpuClass) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GpuClass) SetName(name string) {
	g.Name = &name
}

func (g *GpuClass) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GpuClass) SetPrices(prices []GpuClassPrice) {
	g.Prices = prices
}

func (g *GpuClass) GetPrices() []GpuClassPrice {
	if g == nil {
		return nil
	}
	return g.Prices
}

func (g *GpuClass) SetIsHighDemand(isHighDemand bool) {
	g.IsHighDemand = &isHighDemand
}

func (g *GpuClass) GetIsHighDemand() *bool {
	if g == nil {
		return nil
	}
	return g.IsHighDemand
}
