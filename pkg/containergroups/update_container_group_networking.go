package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents update container group networking parameters
type UpdateContainerGroupNetworking struct {
	// The port number to expose on the container group
	Port *util.Nullable[int64] `json:"port,omitempty" min:"1" max:"65535"`
}

func (u *UpdateContainerGroupNetworking) GetPort() *util.Nullable[int64] {
	if u == nil {
		return nil
	}
	return u.Port
}

func (u *UpdateContainerGroupNetworking) SetPort(port util.Nullable[int64]) {
	u.Port = &port
}

func (u *UpdateContainerGroupNetworking) SetPortNull() {
	u.Port = &util.Nullable[int64]{IsNull: true}
}

func (u UpdateContainerGroupNetworking) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateContainerGroupNetworking to string"
	}
	return string(jsonData)
}
