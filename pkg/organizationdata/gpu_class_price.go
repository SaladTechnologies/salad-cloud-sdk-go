package organizationdata

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents the price of a GPU class for a given container group priority
type GpuClassPrice struct {
	// Specifies the priority level for container group execution, which determines resource allocation and scheduling precedence.
	Priority *util.Nullable[shared.ContainerGroupPriority] `json:"priority,omitempty" required:"true"`
	// The price
	Price *string `json:"price,omitempty" required:"true" maxLength:"20" minLength:"1" pattern:"^.*$"`
}

func (g *GpuClassPrice) GetPriority() *util.Nullable[shared.ContainerGroupPriority] {
	if g == nil {
		return nil
	}
	return g.Priority
}

func (g *GpuClassPrice) SetPriority(priority util.Nullable[shared.ContainerGroupPriority]) {
	g.Priority = &priority
}

func (g *GpuClassPrice) SetPriorityNull() {
	g.Priority = &util.Nullable[shared.ContainerGroupPriority]{IsNull: true}
}

func (g *GpuClassPrice) GetPrice() *string {
	if g == nil {
		return nil
	}
	return g.Price
}

func (g *GpuClassPrice) SetPrice(price string) {
	g.Price = &price
}

func (g GpuClassPrice) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GpuClassPrice to string"
	}
	return string(jsonData)
}
