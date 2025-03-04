package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents a list of container groups
type ContainerGroupList struct {
	Items []shared.ContainerGroup `json:"items,omitempty" required:"true" maxItems:"100"`
}

func (c *ContainerGroupList) GetItems() []shared.ContainerGroup {
	if c == nil {
		return nil
	}
	return c.Items
}

func (c *ContainerGroupList) SetItems(items []shared.ContainerGroup) {
	c.Items = items
}

func (c ContainerGroupList) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupList to string"
	}
	return string(jsonData)
}
