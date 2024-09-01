package organizationdata

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents the price of a GPU class for a given container group priority
type GpuClassPrice struct {
	Priority *shared.ContainerGroupPriority `json:"priority,omitempty" required:"true"`
	// The price
	Price *string `json:"price,omitempty" required:"true" maxLength:"20" minLength:"1"`
}

func (g *GpuClassPrice) SetPriority(priority shared.ContainerGroupPriority) {
	g.Priority = &priority
}

func (g *GpuClassPrice) GetPriority() *shared.ContainerGroupPriority {
	if g == nil {
		return nil
	}
	return g.Priority
}

func (g *GpuClassPrice) SetPrice(price string) {
	g.Price = &price
}

func (g *GpuClassPrice) GetPrice() *string {
	if g == nil {
		return nil
	}
	return g.Price
}
