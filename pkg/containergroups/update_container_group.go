package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a request to update a container group
type UpdateContainerGroup struct {
	DisplayName *util.Nullable[string] `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// Represents an update container object
	Container *util.Nullable[UpdateContainer] `json:"container,omitempty"`
	Replicas  *util.Nullable[int64]           `json:"replicas,omitempty" min:"0" max:"500"`
	// List of countries nodes must be located in. Remove this field to permit nodes from any country.
	CountryCodes *util.Nullable[[]shared.CountryCode] `json:"country_codes,omitempty" minItems:"1" maxItems:"500"`
	// Represents update container group networking parameters
	Networking *UpdateContainerGroupNetworking `json:"networking,omitempty"`
	// Represents the container group liveness probe
	LivenessProbe *util.Nullable[shared.ContainerGroupLivenessProbe] `json:"liveness_probe,omitempty"`
	// Represents the container group readiness probe
	ReadinessProbe *util.Nullable[shared.ContainerGroupReadinessProbe] `json:"readiness_probe,omitempty"`
	// Represents the container group startup probe
	StartupProbe *util.Nullable[shared.ContainerGroupStartupProbe] `json:"startup_probe,omitempty"`
	// Represents the autoscaling rules for a queue
	QueueAutoscaler *util.Nullable[shared.QueueAutoscaler] `json:"queue_autoscaler,omitempty"`
}

func (u *UpdateContainerGroup) GetDisplayName() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.DisplayName
}

func (u *UpdateContainerGroup) SetDisplayName(displayName util.Nullable[string]) {
	u.DisplayName = &displayName
}

func (u *UpdateContainerGroup) SetDisplayNameNull() {
	u.DisplayName = &util.Nullable[string]{IsNull: true}
}

func (u *UpdateContainerGroup) GetContainer() *util.Nullable[UpdateContainer] {
	if u == nil {
		return nil
	}
	return u.Container
}

func (u *UpdateContainerGroup) SetContainer(container util.Nullable[UpdateContainer]) {
	u.Container = &container
}

func (u *UpdateContainerGroup) SetContainerNull() {
	u.Container = &util.Nullable[UpdateContainer]{IsNull: true}
}

func (u *UpdateContainerGroup) GetReplicas() *util.Nullable[int64] {
	if u == nil {
		return nil
	}
	return u.Replicas
}

func (u *UpdateContainerGroup) SetReplicas(replicas util.Nullable[int64]) {
	u.Replicas = &replicas
}

func (u *UpdateContainerGroup) SetReplicasNull() {
	u.Replicas = &util.Nullable[int64]{IsNull: true}
}

func (u *UpdateContainerGroup) GetCountryCodes() *util.Nullable[[]shared.CountryCode] {
	if u == nil {
		return nil
	}
	return u.CountryCodes
}

func (u *UpdateContainerGroup) SetCountryCodes(countryCodes util.Nullable[[]shared.CountryCode]) {
	u.CountryCodes = &countryCodes
}

func (u *UpdateContainerGroup) SetCountryCodesNull() {
	u.CountryCodes = &util.Nullable[[]shared.CountryCode]{IsNull: true}
}

func (u *UpdateContainerGroup) GetNetworking() *UpdateContainerGroupNetworking {
	if u == nil {
		return nil
	}
	return u.Networking
}

func (u *UpdateContainerGroup) SetNetworking(networking UpdateContainerGroupNetworking) {
	u.Networking = &networking
}

func (u *UpdateContainerGroup) GetLivenessProbe() *util.Nullable[shared.ContainerGroupLivenessProbe] {
	if u == nil {
		return nil
	}
	return u.LivenessProbe
}

func (u *UpdateContainerGroup) SetLivenessProbe(livenessProbe util.Nullable[shared.ContainerGroupLivenessProbe]) {
	u.LivenessProbe = &livenessProbe
}

func (u *UpdateContainerGroup) SetLivenessProbeNull() {
	u.LivenessProbe = &util.Nullable[shared.ContainerGroupLivenessProbe]{IsNull: true}
}

func (u *UpdateContainerGroup) GetReadinessProbe() *util.Nullable[shared.ContainerGroupReadinessProbe] {
	if u == nil {
		return nil
	}
	return u.ReadinessProbe
}

func (u *UpdateContainerGroup) SetReadinessProbe(readinessProbe util.Nullable[shared.ContainerGroupReadinessProbe]) {
	u.ReadinessProbe = &readinessProbe
}

func (u *UpdateContainerGroup) SetReadinessProbeNull() {
	u.ReadinessProbe = &util.Nullable[shared.ContainerGroupReadinessProbe]{IsNull: true}
}

func (u *UpdateContainerGroup) GetStartupProbe() *util.Nullable[shared.ContainerGroupStartupProbe] {
	if u == nil {
		return nil
	}
	return u.StartupProbe
}

func (u *UpdateContainerGroup) SetStartupProbe(startupProbe util.Nullable[shared.ContainerGroupStartupProbe]) {
	u.StartupProbe = &startupProbe
}

func (u *UpdateContainerGroup) SetStartupProbeNull() {
	u.StartupProbe = &util.Nullable[shared.ContainerGroupStartupProbe]{IsNull: true}
}

func (u *UpdateContainerGroup) GetQueueAutoscaler() *util.Nullable[shared.QueueAutoscaler] {
	if u == nil {
		return nil
	}
	return u.QueueAutoscaler
}

func (u *UpdateContainerGroup) SetQueueAutoscaler(queueAutoscaler util.Nullable[shared.QueueAutoscaler]) {
	u.QueueAutoscaler = &queueAutoscaler
}

func (u *UpdateContainerGroup) SetQueueAutoscalerNull() {
	u.QueueAutoscaler = &util.Nullable[shared.QueueAutoscaler]{IsNull: true}
}

func (u UpdateContainerGroup) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateContainerGroup to string"
	}
	return string(jsonData)
}
