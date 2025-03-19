package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a request to create a container group, which manages a collection of container instances with shared configuration and scaling policies
type ContainerGroupCreationRequest struct {
	// Determines whether the container group should start automatically when created (true) or remain stopped until manually started (false)
	AutostartPolicy *bool `json:"autostart_policy,omitempty" required:"true"`
	// Configuration for creating a container within a container group. Defines the container image, resource requirements, environment variables, and other settings needed to deploy and run the container.
	Container *ContainerConfiguration `json:"container,omitempty" required:"true"`
	// List of countries nodes must be located in. Remove this field to permit nodes from any country.
	CountryCodes []shared.CountryCode `json:"country_codes,omitempty" minItems:"1" maxItems:"500"`
	// Human-readable name for the container group that can include spaces and special characters, used for display purposes
	DisplayName *string `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// Defines a liveness probe for container groups that determines when to restart a container if it becomes unhealthy
	LivenessProbe *util.Nullable[shared.ContainerGroupLivenessProbe] `json:"liveness_probe,omitempty"`
	// Unique identifier for the container group that must follow DNS naming conventions (lowercase alphanumeric with hyphens)
	Name *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// Network configuration for container groups specifying connectivity parameters, including authentication, protocol, and timeout settings
	Networking *CreateContainerGroupNetworking `json:"networking,omitempty"`
	// Defines configuration for automatically scaling container instances based on queue length. The autoscaler monitors a queue and adjusts the number of running replicas to maintain the desired queue length.
	QueueAutoscaler *shared.QueueBasedAutoscalerConfiguration `json:"queue_autoscaler,omitempty"`
	// Configuration for connecting a container group to a message queue system, enabling asynchronous communication between services.
	QueueConnection *shared.ContainerGroupQueueConnection `json:"queue_connection,omitempty"`
	// Defines how to check if a container is ready to serve traffic. The readiness probe determines whether the container's application is ready to accept traffic. If the readiness probe fails, the container is considered not ready and traffic will not be sent to it.
	ReadinessProbe *util.Nullable[shared.ContainerGroupReadinessProbe] `json:"readiness_probe,omitempty"`
	// Number of container instances to deploy and maintain for this container group
	Replicas *int64 `json:"replicas,omitempty" required:"true" min:"0" max:"500"`
	// Specifies the policy for restarting containers when they exit or fail.
	RestartPolicy *shared.ContainerRestartPolicy `json:"restart_policy,omitempty" required:"true"`
	// Defines a probe that checks if a container application has started successfully. Startup probes help prevent applications from being prematurely marked as unhealthy during initialization. The probe can use HTTP requests, TCP connections, gRPC calls, or shell commands to determine startup status.
	StartupProbe *util.Nullable[shared.ContainerGroupStartupProbe] `json:"startup_probe,omitempty"`
}

func (c *ContainerGroupCreationRequest) GetAutostartPolicy() *bool {
	if c == nil {
		return nil
	}
	return c.AutostartPolicy
}

func (c *ContainerGroupCreationRequest) SetAutostartPolicy(autostartPolicy bool) {
	c.AutostartPolicy = &autostartPolicy
}

func (c *ContainerGroupCreationRequest) GetContainer() *ContainerConfiguration {
	if c == nil {
		return nil
	}
	return c.Container
}

func (c *ContainerGroupCreationRequest) SetContainer(container ContainerConfiguration) {
	c.Container = &container
}

func (c *ContainerGroupCreationRequest) GetCountryCodes() []shared.CountryCode {
	if c == nil {
		return nil
	}
	return c.CountryCodes
}

func (c *ContainerGroupCreationRequest) SetCountryCodes(countryCodes []shared.CountryCode) {
	c.CountryCodes = countryCodes
}

func (c *ContainerGroupCreationRequest) GetDisplayName() *string {
	if c == nil {
		return nil
	}
	return c.DisplayName
}

func (c *ContainerGroupCreationRequest) SetDisplayName(displayName string) {
	c.DisplayName = &displayName
}

func (c *ContainerGroupCreationRequest) GetLivenessProbe() *util.Nullable[shared.ContainerGroupLivenessProbe] {
	if c == nil {
		return nil
	}
	return c.LivenessProbe
}

func (c *ContainerGroupCreationRequest) SetLivenessProbe(livenessProbe util.Nullable[shared.ContainerGroupLivenessProbe]) {
	c.LivenessProbe = &livenessProbe
}

func (c *ContainerGroupCreationRequest) SetLivenessProbeNull() {
	c.LivenessProbe = &util.Nullable[shared.ContainerGroupLivenessProbe]{IsNull: true}
}

func (c *ContainerGroupCreationRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *ContainerGroupCreationRequest) SetName(name string) {
	c.Name = &name
}

func (c *ContainerGroupCreationRequest) GetNetworking() *CreateContainerGroupNetworking {
	if c == nil {
		return nil
	}
	return c.Networking
}

func (c *ContainerGroupCreationRequest) SetNetworking(networking CreateContainerGroupNetworking) {
	c.Networking = &networking
}

func (c *ContainerGroupCreationRequest) GetQueueAutoscaler() *shared.QueueBasedAutoscalerConfiguration {
	if c == nil {
		return nil
	}
	return c.QueueAutoscaler
}

func (c *ContainerGroupCreationRequest) SetQueueAutoscaler(queueAutoscaler shared.QueueBasedAutoscalerConfiguration) {
	c.QueueAutoscaler = &queueAutoscaler
}

func (c *ContainerGroupCreationRequest) GetQueueConnection() *shared.ContainerGroupQueueConnection {
	if c == nil {
		return nil
	}
	return c.QueueConnection
}

func (c *ContainerGroupCreationRequest) SetQueueConnection(queueConnection shared.ContainerGroupQueueConnection) {
	c.QueueConnection = &queueConnection
}

func (c *ContainerGroupCreationRequest) GetReadinessProbe() *util.Nullable[shared.ContainerGroupReadinessProbe] {
	if c == nil {
		return nil
	}
	return c.ReadinessProbe
}

func (c *ContainerGroupCreationRequest) SetReadinessProbe(readinessProbe util.Nullable[shared.ContainerGroupReadinessProbe]) {
	c.ReadinessProbe = &readinessProbe
}

func (c *ContainerGroupCreationRequest) SetReadinessProbeNull() {
	c.ReadinessProbe = &util.Nullable[shared.ContainerGroupReadinessProbe]{IsNull: true}
}

func (c *ContainerGroupCreationRequest) GetReplicas() *int64 {
	if c == nil {
		return nil
	}
	return c.Replicas
}

func (c *ContainerGroupCreationRequest) SetReplicas(replicas int64) {
	c.Replicas = &replicas
}

func (c *ContainerGroupCreationRequest) GetRestartPolicy() *shared.ContainerRestartPolicy {
	if c == nil {
		return nil
	}
	return c.RestartPolicy
}

func (c *ContainerGroupCreationRequest) SetRestartPolicy(restartPolicy shared.ContainerRestartPolicy) {
	c.RestartPolicy = &restartPolicy
}

func (c *ContainerGroupCreationRequest) GetStartupProbe() *util.Nullable[shared.ContainerGroupStartupProbe] {
	if c == nil {
		return nil
	}
	return c.StartupProbe
}

func (c *ContainerGroupCreationRequest) SetStartupProbe(startupProbe util.Nullable[shared.ContainerGroupStartupProbe]) {
	c.StartupProbe = &startupProbe
}

func (c *ContainerGroupCreationRequest) SetStartupProbeNull() {
	c.StartupProbe = &util.Nullable[shared.ContainerGroupStartupProbe]{IsNull: true}
}

func (c ContainerGroupCreationRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupCreationRequest to string"
	}
	return string(jsonData)
}
