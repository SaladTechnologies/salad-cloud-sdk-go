package containergroups

import (
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
}

func (u *UpdateContainerGroup) SetDisplayName(displayName string) {
	u.DisplayName = &displayName
}

func (u *UpdateContainerGroup) GetDisplayName() *string {
	if u == nil {
		return nil
	}
	return u.DisplayName
}

func (u *UpdateContainerGroup) SetContainer(container UpdateContainer) {
	u.Container = &container
}

func (u *UpdateContainerGroup) GetContainer() *UpdateContainer {
	if u == nil {
		return nil
	}
	return u.Container
}

func (u *UpdateContainerGroup) SetReplicas(replicas int64) {
	u.Replicas = &replicas
}

func (u *UpdateContainerGroup) GetReplicas() *int64 {
	if u == nil {
		return nil
	}
	return u.Replicas
}

func (u *UpdateContainerGroup) SetCountryCodes(countryCodes []shared.CountryCode) {
	u.CountryCodes = countryCodes
}

func (u *UpdateContainerGroup) GetCountryCodes() []shared.CountryCode {
	if u == nil {
		return nil
	}
	return u.CountryCodes
}

func (u *UpdateContainerGroup) SetNetworking(networking UpdateContainerGroupNetworking) {
	u.Networking = &networking
}

func (u *UpdateContainerGroup) GetNetworking() *UpdateContainerGroupNetworking {
	if u == nil {
		return nil
	}
	return u.Networking
}

func (u *UpdateContainerGroup) SetLivenessProbe(livenessProbe shared.ContainerGroupLivenessProbe) {
	u.LivenessProbe = &livenessProbe
}

func (u *UpdateContainerGroup) GetLivenessProbe() *shared.ContainerGroupLivenessProbe {
	if u == nil {
		return nil
	}
	return u.LivenessProbe
}

func (u *UpdateContainerGroup) SetReadinessProbe(readinessProbe shared.ContainerGroupReadinessProbe) {
	u.ReadinessProbe = &readinessProbe
}

func (u *UpdateContainerGroup) GetReadinessProbe() *shared.ContainerGroupReadinessProbe {
	if u == nil {
		return nil
	}
	return u.ReadinessProbe
}

func (u *UpdateContainerGroup) SetStartupProbe(startupProbe shared.ContainerGroupStartupProbe) {
	u.StartupProbe = &startupProbe
}

func (u *UpdateContainerGroup) GetStartupProbe() *shared.ContainerGroupStartupProbe {
	if u == nil {
		return nil
	}
	return u.StartupProbe
}

func (u *UpdateContainerGroup) SetQueueAutoscaler(queueAutoscaler shared.QueueAutoscaler) {
	u.QueueAutoscaler = &queueAutoscaler
}

func (u *UpdateContainerGroup) GetQueueAutoscaler() *shared.QueueAutoscaler {
	if u == nil {
		return nil
	}
	return u.QueueAutoscaler
}
