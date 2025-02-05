package containergroups

import (
	"encoding/json"
)

// Represents update container group networking parameters
type UpdateContainerGroupNetworking struct {
	Port    *int64 `json:"port,omitempty" min:"1" max:"65535"`
	touched map[string]bool
}

func (u *UpdateContainerGroupNetworking) GetPort() *int64 {
	if u == nil {
		return nil
	}
	return u.Port
}

func (u *UpdateContainerGroupNetworking) SetPort(port int64) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Port"] = true
	u.Port = &port
}

func (u *UpdateContainerGroupNetworking) SetPortNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Port"] = true
	u.Port = nil
}

func (u UpdateContainerGroupNetworking) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["Port"] && u.Port == nil {
		data["port"] = nil
	} else if u.Port != nil {
		data["port"] = u.Port
	}

	return json.Marshal(data)
}

func (u UpdateContainerGroupNetworking) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateContainerGroupNetworking to string"
	}
	return string(jsonData)
}
