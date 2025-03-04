package organizationdata

import "encoding/json"

// Represents a list of GPU classes
type GpuClassesList struct {
	// The list of GPU classes
	Items []GpuClass `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (g *GpuClassesList) GetItems() []GpuClass {
	if g == nil {
		return nil
	}
	return g.Items
}

func (g *GpuClassesList) SetItems(items []GpuClass) {
	g.Items = items
}

func (g GpuClassesList) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GpuClassesList to string"
	}
	return string(jsonData)
}
