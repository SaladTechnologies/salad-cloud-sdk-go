package organizationdata

// Represents a list of GPU classes
type GpuClassesList struct {
	// The list of GPU classes
	Items []GpuClass `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (g *GpuClassesList) SetItems(items []GpuClass) {
	g.Items = items
}

func (g *GpuClassesList) GetItems() []GpuClass {
	if g == nil {
		return nil
	}
	return g.Items
}
