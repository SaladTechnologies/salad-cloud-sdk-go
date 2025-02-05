package organizationdata

import (
	"encoding/json"
)

// Represents a list of GPU classes
type GpuClassesList struct {
	// The list of GPU classes
	Items   []GpuClass `json:"items,omitempty" required:"true" maxItems:"100"`
	touched map[string]bool
}

func (g *GpuClassesList) GetItems() []GpuClass {
	if g == nil {
		return nil
	}
	return g.Items
}

func (g *GpuClassesList) SetItems(items []GpuClass) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Items"] = true
	g.Items = items
}

func (g *GpuClassesList) SetItemsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Items"] = true
	g.Items = nil
}

func (g GpuClassesList) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Items"] && g.Items == nil {
		data["items"] = nil
	} else if g.Items != nil {
		data["items"] = g.Items
	}

	return json.Marshal(data)
}

func (g GpuClassesList) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GpuClassesList to string"
	}
	return string(jsonData)
}
