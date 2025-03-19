package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a request to update a container group instance
type ContainerGroupInstancePatch struct {
	// The cost of deleting the container group instance
	DeletionCost *util.Nullable[int64] `json:"deletion_cost,omitempty" min:"0" max:"100000"`
}

func (c *ContainerGroupInstancePatch) GetDeletionCost() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.DeletionCost
}

func (c *ContainerGroupInstancePatch) SetDeletionCost(deletionCost util.Nullable[int64]) {
	c.DeletionCost = &deletionCost
}

func (c *ContainerGroupInstancePatch) SetDeletionCostNull() {
	c.DeletionCost = &util.Nullable[int64]{IsNull: true}
}

func (c ContainerGroupInstancePatch) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupInstancePatch to string"
	}
	return string(jsonData)
}
