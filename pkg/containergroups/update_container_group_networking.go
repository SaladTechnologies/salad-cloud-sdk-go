package containergroups

// Represents update container group networking parameters
type UpdateContainerGroupNetworking struct {
	Port *int64 `json:"port,omitempty" min:"1" max:"65535"`
}

func (u *UpdateContainerGroupNetworking) SetPort(port int64) {
	u.Port = &port
}

func (u *UpdateContainerGroupNetworking) GetPort() *int64 {
	if u == nil {
		return nil
	}
	return u.Port
}
