package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents a list of container groups
type ContainerGroupList struct {
	Items   []shared.ContainerGroup `json:"items,omitempty" required:"true" maxItems:"100"`
	touched map[string]bool
}

func (c *ContainerGroupList) GetItems() []shared.ContainerGroup {
	if c == nil {
		return nil
	}
	return c.Items
}

func (c *ContainerGroupList) SetItems(items []shared.ContainerGroup) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Items"] = true
	c.Items = items
}

func (c *ContainerGroupList) SetItemsNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Items"] = true
	c.Items = nil
}

func (c ContainerGroupList) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Items"] && c.Items == nil {
		data["items"] = nil
	} else if c.Items != nil {
		data["items"] = c.Items
	}

	return json.Marshal(data)
}

func (c ContainerGroupList) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupList to string"
	}
	return string(jsonData)
}
