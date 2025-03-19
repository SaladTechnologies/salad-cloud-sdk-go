package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// A paginated collection of container groups that provides a structured way to access multiple container group resources in a single response.
type ContainerGroupCollection struct {
	// An array containing container group objects. Each object represents a discrete container group with its own properties, configuration, and status.
	Items []shared.ContainerGroup `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (c *ContainerGroupCollection) GetItems() []shared.ContainerGroup {
	if c == nil {
		return nil
	}
	return c.Items
}

func (c *ContainerGroupCollection) SetItems(items []shared.ContainerGroup) {
	c.Items = items
}

func (c ContainerGroupCollection) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupCollection to string"
	}
	return string(jsonData)
}
