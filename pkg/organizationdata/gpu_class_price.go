package organizationdata

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents the price of a GPU class for a given container group priority
type GpuClassPrice struct {
	Priority *shared.ContainerGroupPriority `json:"priority,omitempty" required:"true"`
	// The price
	Price   *string `json:"price,omitempty" required:"true" maxLength:"20" minLength:"1"`
	touched map[string]bool
}

func (g *GpuClassPrice) GetPriority() *shared.ContainerGroupPriority {
	if g == nil {
		return nil
	}
	return g.Priority
}

func (g *GpuClassPrice) SetPriority(priority shared.ContainerGroupPriority) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Priority"] = true
	g.Priority = &priority
}

func (g *GpuClassPrice) SetPriorityNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Priority"] = true
	g.Priority = nil
}

func (g *GpuClassPrice) GetPrice() *string {
	if g == nil {
		return nil
	}
	return g.Price
}

func (g *GpuClassPrice) SetPrice(price string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Price"] = true
	g.Price = &price
}

func (g *GpuClassPrice) SetPriceNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Price"] = true
	g.Price = nil
}

func (g GpuClassPrice) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Priority"] && g.Priority == nil {
		data["priority"] = nil
	} else if g.Priority != nil {
		data["priority"] = g.Priority
	}

	if g.touched["Price"] && g.Price == nil {
		data["price"] = nil
	} else if g.Price != nil {
		data["price"] = g.Price
	}

	return json.Marshal(data)
}

func (g GpuClassPrice) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GpuClassPrice to string"
	}
	return string(jsonData)
}
