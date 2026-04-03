package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/unmarshal"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a request to update a container group
type ContainerGroupPatch struct {
	// Represents an update container object
	Container *util.Nullable[UpdateContainer] `json:"container,omitempty"`
	// List of countries nodes must be located in. Remove this field to permit nodes from any country.
	CountryCodes *util.Nullable[[]shared.CountryCode] `json:"country_codes,omitempty" minItems:"1" maxItems:"500"`
	// The display name for the container group. If null is provided, the display name will be set to the container group name.
	DisplayName *util.Nullable[string] `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// Defines a liveness probe for container groups that determines when to restart a container if it becomes unhealthy
	LivenessProbe *util.Nullable[shared.ContainerGroupLivenessProbe] `json:"liveness_probe,omitempty"`
	// Represents update container group networking parameters
	Networking *UpdateContainerGroupNetworking `json:"networking,omitempty"`
	// Defines configuration for automatically scaling container instances based on queue length. The autoscaler monitors a queue and adjusts the number of running replicas to maintain the desired queue length.
	QueueAutoscaler *shared.QueueBasedAutoscalerConfiguration `json:"queue_autoscaler,omitempty"`
	// Defines how to check if a container is ready to serve traffic. The readiness probe determines whether the container's application is ready to accept traffic. If the readiness probe fails, the container is considered not ready and traffic will not be sent to it.
	ReadinessProbe *util.Nullable[shared.ContainerGroupReadinessProbe] `json:"readiness_probe,omitempty"`
	// The desired number of instances for your container group deployment.
	Replicas *util.Nullable[int64] `json:"replicas,omitempty" min:"0" max:"500"`
	// List of scaling actions configurations
	ScalingActions []shared.ContainerGroupScalingAction `json:"scaling-actions,omitempty" maxItems:"100"`
	// Indicates if scheduled scaling is enabled
	ScheduledScalingEnabled *bool `json:"scheduled-scaling-enabled,omitempty"`
	// Defines a probe that checks if a container application has started successfully. Startup probes help prevent applications from being prematurely marked as unhealthy during initialization. The probe can use HTTP requests, TCP connections, gRPC calls, or shell commands to determine startup status.
	StartupProbe *util.Nullable[shared.ContainerGroupStartupProbe] `json:"startup_probe,omitempty"`
}

func (c *ContainerGroupPatch) GetContainer() *util.Nullable[UpdateContainer] {
	if c == nil {
		return nil
	}
	return c.Container
}

func (c *ContainerGroupPatch) SetContainer(container util.Nullable[UpdateContainer]) {
	c.Container = &container
}

func (c *ContainerGroupPatch) SetContainerNull() {
	c.Container = &util.Nullable[UpdateContainer]{IsNull: true}
}

func (c *ContainerGroupPatch) GetCountryCodes() *util.Nullable[[]shared.CountryCode] {
	if c == nil {
		return nil
	}
	return c.CountryCodes
}

func (c *ContainerGroupPatch) SetCountryCodes(countryCodes util.Nullable[[]shared.CountryCode]) {
	c.CountryCodes = &countryCodes
}

func (c *ContainerGroupPatch) SetCountryCodesNull() {
	c.CountryCodes = &util.Nullable[[]shared.CountryCode]{IsNull: true}
}

func (c *ContainerGroupPatch) GetDisplayName() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.DisplayName
}

func (c *ContainerGroupPatch) SetDisplayName(displayName util.Nullable[string]) {
	c.DisplayName = &displayName
}

func (c *ContainerGroupPatch) SetDisplayNameNull() {
	c.DisplayName = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerGroupPatch) GetLivenessProbe() *util.Nullable[shared.ContainerGroupLivenessProbe] {
	if c == nil {
		return nil
	}
	return c.LivenessProbe
}

func (c *ContainerGroupPatch) SetLivenessProbe(livenessProbe util.Nullable[shared.ContainerGroupLivenessProbe]) {
	c.LivenessProbe = &livenessProbe
}

func (c *ContainerGroupPatch) SetLivenessProbeNull() {
	c.LivenessProbe = &util.Nullable[shared.ContainerGroupLivenessProbe]{IsNull: true}
}

func (c *ContainerGroupPatch) GetNetworking() *UpdateContainerGroupNetworking {
	if c == nil {
		return nil
	}
	return c.Networking
}

func (c *ContainerGroupPatch) SetNetworking(networking UpdateContainerGroupNetworking) {
	c.Networking = &networking
}

func (c *ContainerGroupPatch) GetQueueAutoscaler() *shared.QueueBasedAutoscalerConfiguration {
	if c == nil {
		return nil
	}
	return c.QueueAutoscaler
}

func (c *ContainerGroupPatch) SetQueueAutoscaler(queueAutoscaler shared.QueueBasedAutoscalerConfiguration) {
	c.QueueAutoscaler = &queueAutoscaler
}

func (c *ContainerGroupPatch) GetReadinessProbe() *util.Nullable[shared.ContainerGroupReadinessProbe] {
	if c == nil {
		return nil
	}
	return c.ReadinessProbe
}

func (c *ContainerGroupPatch) SetReadinessProbe(readinessProbe util.Nullable[shared.ContainerGroupReadinessProbe]) {
	c.ReadinessProbe = &readinessProbe
}

func (c *ContainerGroupPatch) SetReadinessProbeNull() {
	c.ReadinessProbe = &util.Nullable[shared.ContainerGroupReadinessProbe]{IsNull: true}
}

func (c *ContainerGroupPatch) GetReplicas() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.Replicas
}

func (c *ContainerGroupPatch) SetReplicas(replicas util.Nullable[int64]) {
	c.Replicas = &replicas
}

func (c *ContainerGroupPatch) SetReplicasNull() {
	c.Replicas = &util.Nullable[int64]{IsNull: true}
}

func (c *ContainerGroupPatch) GetScalingActions() []shared.ContainerGroupScalingAction {
	if c == nil {
		return nil
	}
	return c.ScalingActions
}

func (c *ContainerGroupPatch) SetScalingActions(scalingActions []shared.ContainerGroupScalingAction) {
	c.ScalingActions = scalingActions
}

func (c *ContainerGroupPatch) GetScheduledScalingEnabled() *bool {
	if c == nil {
		return nil
	}
	return c.ScheduledScalingEnabled
}

func (c *ContainerGroupPatch) SetScheduledScalingEnabled(scheduledScalingEnabled bool) {
	c.ScheduledScalingEnabled = &scheduledScalingEnabled
}

func (c *ContainerGroupPatch) GetStartupProbe() *util.Nullable[shared.ContainerGroupStartupProbe] {
	if c == nil {
		return nil
	}
	return c.StartupProbe
}

func (c *ContainerGroupPatch) SetStartupProbe(startupProbe util.Nullable[shared.ContainerGroupStartupProbe]) {
	c.StartupProbe = &startupProbe
}

func (c *ContainerGroupPatch) SetStartupProbeNull() {
	c.StartupProbe = &util.Nullable[shared.ContainerGroupStartupProbe]{IsNull: true}
}

func (c ContainerGroupPatch) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupPatch to string"
	}
	return string(jsonData)
}

func (c *ContainerGroupPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
