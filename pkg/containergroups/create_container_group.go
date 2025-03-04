package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a request to create a container group
type CreateContainerGroup struct {
	Name        *string                `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	DisplayName *util.Nullable[string] `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// Represents a container
	Container       *CreateContainer               `json:"container,omitempty" required:"true"`
	AutostartPolicy *bool                          `json:"autostart_policy,omitempty" required:"true"`
	RestartPolicy   *shared.ContainerRestartPolicy `json:"restart_policy,omitempty" required:"true"`
	Replicas        *int64                         `json:"replicas,omitempty" required:"true" min:"0" max:"500"`
	// List of countries nodes must be located in. Remove this field to permit nodes from any country.
	CountryCodes []shared.CountryCode `json:"country_codes,omitempty" minItems:"1" maxItems:"500"`
	// Represents container group networking parameters
	Networking *util.Nullable[CreateContainerGroupNetworking] `json:"networking,omitempty"`
	// Represents the container group liveness probe
	LivenessProbe *util.Nullable[shared.ContainerGroupLivenessProbe] `json:"liveness_probe,omitempty"`
	// Represents the container group readiness probe
	ReadinessProbe *util.Nullable[shared.ContainerGroupReadinessProbe] `json:"readiness_probe,omitempty"`
	// Represents the container group startup probe
	StartupProbe *util.Nullable[shared.ContainerGroupStartupProbe] `json:"startup_probe,omitempty"`
	// Represents container group queue connection
	QueueConnection *util.Nullable[shared.ContainerGroupQueueConnection] `json:"queue_connection,omitempty"`
	// Represents the autoscaling rules for a queue
	QueueAutoscaler *util.Nullable[shared.QueueAutoscaler] `json:"queue_autoscaler,omitempty"`
}

func (c *CreateContainerGroup) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateContainerGroup) SetName(name string) {
	c.Name = &name
}

func (c *CreateContainerGroup) GetDisplayName() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.DisplayName
}

func (c *CreateContainerGroup) SetDisplayName(displayName util.Nullable[string]) {
	c.DisplayName = &displayName
}

func (c *CreateContainerGroup) SetDisplayNameNull() {
	c.DisplayName = &util.Nullable[string]{IsNull: true}
}

func (c *CreateContainerGroup) GetContainer() *CreateContainer {
	if c == nil {
		return nil
	}
	return c.Container
}

func (c *CreateContainerGroup) SetContainer(container CreateContainer) {
	c.Container = &container
}

func (c *CreateContainerGroup) GetAutostartPolicy() *bool {
	if c == nil {
		return nil
	}
	return c.AutostartPolicy
}

func (c *CreateContainerGroup) SetAutostartPolicy(autostartPolicy bool) {
	c.AutostartPolicy = &autostartPolicy
}

func (c *CreateContainerGroup) GetRestartPolicy() *shared.ContainerRestartPolicy {
	if c == nil {
		return nil
	}
	return c.RestartPolicy
}

func (c *CreateContainerGroup) SetRestartPolicy(restartPolicy shared.ContainerRestartPolicy) {
	c.RestartPolicy = &restartPolicy
}

func (c *CreateContainerGroup) GetReplicas() *int64 {
	if c == nil {
		return nil
	}
	return c.Replicas
}

func (c *CreateContainerGroup) SetReplicas(replicas int64) {
	c.Replicas = &replicas
}

func (c *CreateContainerGroup) GetCountryCodes() []shared.CountryCode {
	if c == nil {
		return nil
	}
	return c.CountryCodes
}

func (c *CreateContainerGroup) SetCountryCodes(countryCodes []shared.CountryCode) {
	c.CountryCodes = countryCodes
}

func (c *CreateContainerGroup) GetNetworking() *util.Nullable[CreateContainerGroupNetworking] {
	if c == nil {
		return nil
	}
	return c.Networking
}

func (c *CreateContainerGroup) SetNetworking(networking util.Nullable[CreateContainerGroupNetworking]) {
	c.Networking = &networking
}

func (c *CreateContainerGroup) SetNetworkingNull() {
	c.Networking = &util.Nullable[CreateContainerGroupNetworking]{IsNull: true}
}

func (c *CreateContainerGroup) GetLivenessProbe() *util.Nullable[shared.ContainerGroupLivenessProbe] {
	if c == nil {
		return nil
	}
	return c.LivenessProbe
}

func (c *CreateContainerGroup) SetLivenessProbe(livenessProbe util.Nullable[shared.ContainerGroupLivenessProbe]) {
	c.LivenessProbe = &livenessProbe
}

func (c *CreateContainerGroup) SetLivenessProbeNull() {
	c.LivenessProbe = &util.Nullable[shared.ContainerGroupLivenessProbe]{IsNull: true}
}

func (c *CreateContainerGroup) GetReadinessProbe() *util.Nullable[shared.ContainerGroupReadinessProbe] {
	if c == nil {
		return nil
	}
	return c.ReadinessProbe
}

func (c *CreateContainerGroup) SetReadinessProbe(readinessProbe util.Nullable[shared.ContainerGroupReadinessProbe]) {
	c.ReadinessProbe = &readinessProbe
}

func (c *CreateContainerGroup) SetReadinessProbeNull() {
	c.ReadinessProbe = &util.Nullable[shared.ContainerGroupReadinessProbe]{IsNull: true}
}

func (c *CreateContainerGroup) GetStartupProbe() *util.Nullable[shared.ContainerGroupStartupProbe] {
	if c == nil {
		return nil
	}
	return c.StartupProbe
}

func (c *CreateContainerGroup) SetStartupProbe(startupProbe util.Nullable[shared.ContainerGroupStartupProbe]) {
	c.StartupProbe = &startupProbe
}

func (c *CreateContainerGroup) SetStartupProbeNull() {
	c.StartupProbe = &util.Nullable[shared.ContainerGroupStartupProbe]{IsNull: true}
}

func (c *CreateContainerGroup) GetQueueConnection() *util.Nullable[shared.ContainerGroupQueueConnection] {
	if c == nil {
		return nil
	}
	return c.QueueConnection
}

func (c *CreateContainerGroup) SetQueueConnection(queueConnection util.Nullable[shared.ContainerGroupQueueConnection]) {
	c.QueueConnection = &queueConnection
}

func (c *CreateContainerGroup) SetQueueConnectionNull() {
	c.QueueConnection = &util.Nullable[shared.ContainerGroupQueueConnection]{IsNull: true}
}

func (c *CreateContainerGroup) GetQueueAutoscaler() *util.Nullable[shared.QueueAutoscaler] {
	if c == nil {
		return nil
	}
	return c.QueueAutoscaler
}

func (c *CreateContainerGroup) SetQueueAutoscaler(queueAutoscaler util.Nullable[shared.QueueAutoscaler]) {
	c.QueueAutoscaler = &queueAutoscaler
}

func (c *CreateContainerGroup) SetQueueAutoscalerNull() {
	c.QueueAutoscaler = &util.Nullable[shared.QueueAutoscaler]{IsNull: true}
}

func (c CreateContainerGroup) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateContainerGroup to string"
	}
	return string(jsonData)
}
