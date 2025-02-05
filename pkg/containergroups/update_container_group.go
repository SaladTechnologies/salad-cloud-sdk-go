package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents a request to update a container group
type UpdateContainerGroup struct {
	DisplayName *string `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// Represents an update container object
	Container *UpdateContainer `json:"container,omitempty"`
	Replicas  *int64           `json:"replicas,omitempty" min:"0" max:"250"`
	// List of countries nodes must be located in. Remove this field to permit nodes from any country.
	CountryCodes []shared.CountryCode `json:"country_codes,omitempty" minItems:"1" maxItems:"500"`
	// Represents update container group networking parameters
	Networking *UpdateContainerGroupNetworking `json:"networking,omitempty"`
	// Represents the container group liveness probe
	LivenessProbe *shared.ContainerGroupLivenessProbe `json:"liveness_probe,omitempty"`
	// Represents the container group readiness probe
	ReadinessProbe *shared.ContainerGroupReadinessProbe `json:"readiness_probe,omitempty"`
	// Represents the container group startup probe
	StartupProbe *shared.ContainerGroupStartupProbe `json:"startup_probe,omitempty"`
	// Represents the autoscaling rules for a queue
	QueueAutoscaler *shared.QueueAutoscaler `json:"queue_autoscaler,omitempty"`
	touched         map[string]bool
}

func (u *UpdateContainerGroup) GetDisplayName() *string {
	if u == nil {
		return nil
	}
	return u.DisplayName
}

func (u *UpdateContainerGroup) SetDisplayName(displayName string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["DisplayName"] = true
	u.DisplayName = &displayName
}

func (u *UpdateContainerGroup) SetDisplayNameNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["DisplayName"] = true
	u.DisplayName = nil
}

func (u *UpdateContainerGroup) GetContainer() *UpdateContainer {
	if u == nil {
		return nil
	}
	return u.Container
}

func (u *UpdateContainerGroup) SetContainer(container UpdateContainer) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Container"] = true
	u.Container = &container
}

func (u *UpdateContainerGroup) SetContainerNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Container"] = true
	u.Container = nil
}

func (u *UpdateContainerGroup) GetReplicas() *int64 {
	if u == nil {
		return nil
	}
	return u.Replicas
}

func (u *UpdateContainerGroup) SetReplicas(replicas int64) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Replicas"] = true
	u.Replicas = &replicas
}

func (u *UpdateContainerGroup) SetReplicasNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Replicas"] = true
	u.Replicas = nil
}

func (u *UpdateContainerGroup) GetCountryCodes() []shared.CountryCode {
	if u == nil {
		return nil
	}
	return u.CountryCodes
}

func (u *UpdateContainerGroup) SetCountryCodes(countryCodes []shared.CountryCode) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["CountryCodes"] = true
	u.CountryCodes = countryCodes
}

func (u *UpdateContainerGroup) SetCountryCodesNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["CountryCodes"] = true
	u.CountryCodes = nil
}

func (u *UpdateContainerGroup) GetNetworking() *UpdateContainerGroupNetworking {
	if u == nil {
		return nil
	}
	return u.Networking
}

func (u *UpdateContainerGroup) SetNetworking(networking UpdateContainerGroupNetworking) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Networking"] = true
	u.Networking = &networking
}

func (u *UpdateContainerGroup) SetNetworkingNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Networking"] = true
	u.Networking = nil
}

func (u *UpdateContainerGroup) GetLivenessProbe() *shared.ContainerGroupLivenessProbe {
	if u == nil {
		return nil
	}
	return u.LivenessProbe
}

func (u *UpdateContainerGroup) SetLivenessProbe(livenessProbe shared.ContainerGroupLivenessProbe) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["LivenessProbe"] = true
	u.LivenessProbe = &livenessProbe
}

func (u *UpdateContainerGroup) SetLivenessProbeNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["LivenessProbe"] = true
	u.LivenessProbe = nil
}

func (u *UpdateContainerGroup) GetReadinessProbe() *shared.ContainerGroupReadinessProbe {
	if u == nil {
		return nil
	}
	return u.ReadinessProbe
}

func (u *UpdateContainerGroup) SetReadinessProbe(readinessProbe shared.ContainerGroupReadinessProbe) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["ReadinessProbe"] = true
	u.ReadinessProbe = &readinessProbe
}

func (u *UpdateContainerGroup) SetReadinessProbeNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["ReadinessProbe"] = true
	u.ReadinessProbe = nil
}

func (u *UpdateContainerGroup) GetStartupProbe() *shared.ContainerGroupStartupProbe {
	if u == nil {
		return nil
	}
	return u.StartupProbe
}

func (u *UpdateContainerGroup) SetStartupProbe(startupProbe shared.ContainerGroupStartupProbe) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["StartupProbe"] = true
	u.StartupProbe = &startupProbe
}

func (u *UpdateContainerGroup) SetStartupProbeNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["StartupProbe"] = true
	u.StartupProbe = nil
}

func (u *UpdateContainerGroup) GetQueueAutoscaler() *shared.QueueAutoscaler {
	if u == nil {
		return nil
	}
	return u.QueueAutoscaler
}

func (u *UpdateContainerGroup) SetQueueAutoscaler(queueAutoscaler shared.QueueAutoscaler) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["QueueAutoscaler"] = true
	u.QueueAutoscaler = &queueAutoscaler
}

func (u *UpdateContainerGroup) SetQueueAutoscalerNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["QueueAutoscaler"] = true
	u.QueueAutoscaler = nil
}

func (u UpdateContainerGroup) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["DisplayName"] && u.DisplayName == nil {
		data["display_name"] = nil
	} else if u.DisplayName != nil {
		data["display_name"] = u.DisplayName
	}

	if u.touched["Container"] && u.Container == nil {
		data["container"] = nil
	} else if u.Container != nil {
		data["container"] = u.Container
	}

	if u.touched["Replicas"] && u.Replicas == nil {
		data["replicas"] = nil
	} else if u.Replicas != nil {
		data["replicas"] = u.Replicas
	}

	if u.touched["CountryCodes"] && u.CountryCodes == nil {
		data["country_codes"] = nil
	} else if u.CountryCodes != nil {
		data["country_codes"] = u.CountryCodes
	}

	if u.touched["Networking"] && u.Networking == nil {
		data["networking"] = nil
	} else if u.Networking != nil {
		data["networking"] = u.Networking
	}

	if u.touched["LivenessProbe"] && u.LivenessProbe == nil {
		data["liveness_probe"] = nil
	} else if u.LivenessProbe != nil {
		data["liveness_probe"] = u.LivenessProbe
	}

	if u.touched["ReadinessProbe"] && u.ReadinessProbe == nil {
		data["readiness_probe"] = nil
	} else if u.ReadinessProbe != nil {
		data["readiness_probe"] = u.ReadinessProbe
	}

	if u.touched["StartupProbe"] && u.StartupProbe == nil {
		data["startup_probe"] = nil
	} else if u.StartupProbe != nil {
		data["startup_probe"] = u.StartupProbe
	}

	if u.touched["QueueAutoscaler"] && u.QueueAutoscaler == nil {
		data["queue_autoscaler"] = nil
	} else if u.QueueAutoscaler != nil {
		data["queue_autoscaler"] = u.QueueAutoscaler
	}

	return json.Marshal(data)
}

func (u UpdateContainerGroup) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateContainerGroup to string"
	}
	return string(jsonData)
}
