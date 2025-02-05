package organizationdata

import (
	"encoding/json"
)

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
	touched      map[string]bool
}

func (g *GpuClass) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GpuClass) SetId(id string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = &id
}

func (g *GpuClass) SetIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = nil
}

func (g *GpuClass) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GpuClass) SetName(name string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = &name
}

func (g *GpuClass) SetNameNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = nil
}

func (g *GpuClass) GetPrices() []GpuClassPrice {
	if g == nil {
		return nil
	}
	return g.Prices
}

func (g *GpuClass) SetPrices(prices []GpuClassPrice) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Prices"] = true
	g.Prices = prices
}

func (g *GpuClass) SetPricesNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Prices"] = true
	g.Prices = nil
}

func (g *GpuClass) GetIsHighDemand() *bool {
	if g == nil {
		return nil
	}
	return g.IsHighDemand
}

func (g *GpuClass) SetIsHighDemand(isHighDemand bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["IsHighDemand"] = true
	g.IsHighDemand = &isHighDemand
}

func (g *GpuClass) SetIsHighDemandNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["IsHighDemand"] = true
	g.IsHighDemand = nil
}

func (g GpuClass) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Id"] && g.Id == nil {
		data["id"] = nil
	} else if g.Id != nil {
		data["id"] = g.Id
	}

	if g.touched["Name"] && g.Name == nil {
		data["name"] = nil
	} else if g.Name != nil {
		data["name"] = g.Name
	}

	if g.touched["Prices"] && g.Prices == nil {
		data["prices"] = nil
	} else if g.Prices != nil {
		data["prices"] = g.Prices
	}

	if g.touched["IsHighDemand"] && g.IsHighDemand == nil {
		data["is_high_demand"] = nil
	} else if g.IsHighDemand != nil {
		data["is_high_demand"] = g.IsHighDemand
	}

	return json.Marshal(data)
}

func (g GpuClass) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GpuClass to string"
	}
	return string(jsonData)
}
