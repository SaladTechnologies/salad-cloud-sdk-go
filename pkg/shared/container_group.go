package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// A container group definition that represents a scalable set of identical containers running as a distributed service
type ContainerGroup struct {
	// Defines whether containers in this group should automatically start when deployed (true) or require manual starting (false)
	AutostartPolicy *bool `json:"autostart_policy,omitempty" required:"true"`
	// Represents a container with its configuration and resource requirements.
	Container *Container `json:"container,omitempty" required:"true"`
	// List of country codes where container instances are permitted to run. When not specified or empty, containers may run in any available region.
	CountryCodes []CountryCode `json:"country_codes,omitempty" required:"true" maxItems:"500"`
	// ISO 8601 timestamp when this container group was initially created
	CreateTime *string `json:"create_time,omitempty" required:"true"`
	// Represents the operational state of a container group during its lifecycle, including timing information, status, and instance distribution metrics. This state captures the current execution status, start and finish times, and provides visibility into the operational health across instances.
	CurrentState *ContainerGroupState `json:"current_state,omitempty" required:"true"`
	// The display-friendly name of the resource.
	DisplayName *string `json:"display_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The container group identifier.
	Id *string `json:"id,omitempty" required:"true"`
	// Defines a liveness probe for container groups that determines when to restart a container if it becomes unhealthy
	LivenessProbe *util.Nullable[ContainerGroupLivenessProbe] `json:"liveness_probe,omitempty"`
	// The container group name.
	Name *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// Network configuration for container groups that defines connectivity, routing, and access control settings
	Networking *ContainerGroupNetworkingConfiguration `json:"networking,omitempty"`
	// The organization name.
	OrganizationName *string `json:"organization_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// Indicates whether a configuration change has been requested but not yet applied to all containers in the group
	PendingChange *bool `json:"pending_change,omitempty" required:"true"`
	// Specifies the priority level for container group execution, which determines resource allocation and scheduling precedence.
	Priority *util.Nullable[ContainerGroupPriority] `json:"priority,omitempty" required:"true"`
	// The project name.
	ProjectName *string `json:"project_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// Defines configuration for automatically scaling container instances based on queue length. The autoscaler monitors a queue and adjusts the number of running replicas to maintain the desired queue length.
	QueueAutoscaler *QueueBasedAutoscalerConfiguration `json:"queue_autoscaler,omitempty"`
	// Configuration for connecting a container group to a message queue system, enabling asynchronous communication between services.
	QueueConnection *ContainerGroupQueueConnection `json:"queue_connection,omitempty"`
	// Defines how to check if a container is ready to serve traffic. The readiness probe determines whether the container's application is ready to accept traffic. If the readiness probe fails, the container is considered not ready and traffic will not be sent to it.
	ReadinessProbe *util.Nullable[ContainerGroupReadinessProbe] `json:"readiness_probe,omitempty"`
	// The container group replicas.
	Replicas *int64 `json:"replicas,omitempty" required:"true" min:"0" max:"500"`
	// Specifies the policy for restarting containers when they exit or fail.
	RestartPolicy *ContainerRestartPolicy `json:"restart_policy,omitempty" required:"true"`
	// Defines a probe that checks if a container application has started successfully. Startup probes help prevent applications from being prematurely marked as unhealthy during initialization. The probe can use HTTP requests, TCP connections, gRPC calls, or shell commands to determine startup status.
	StartupProbe *util.Nullable[ContainerGroupStartupProbe] `json:"startup_probe,omitempty"`
	// ISO 8601 timestamp when this container group was last updated
	UpdateTime *string `json:"update_time,omitempty" required:"true"`
	// Incremental version number that increases with each configuration change to the container group
	Version *int64 `json:"version,omitempty" required:"true" min:"1" max:"2147483647"`
}

func (c *ContainerGroup) GetAutostartPolicy() *bool {
	if c == nil {
		return nil
	}
	return c.AutostartPolicy
}

func (c *ContainerGroup) SetAutostartPolicy(autostartPolicy bool) {
	c.AutostartPolicy = &autostartPolicy
}

func (c *ContainerGroup) GetContainer() *Container {
	if c == nil {
		return nil
	}
	return c.Container
}

func (c *ContainerGroup) SetContainer(container Container) {
	c.Container = &container
}

func (c *ContainerGroup) GetCountryCodes() []CountryCode {
	if c == nil {
		return nil
	}
	return c.CountryCodes
}

func (c *ContainerGroup) SetCountryCodes(countryCodes []CountryCode) {
	c.CountryCodes = countryCodes
}

func (c *ContainerGroup) GetCreateTime() *string {
	if c == nil {
		return nil
	}
	return c.CreateTime
}

func (c *ContainerGroup) SetCreateTime(createTime string) {
	c.CreateTime = &createTime
}

func (c *ContainerGroup) GetCurrentState() *ContainerGroupState {
	if c == nil {
		return nil
	}
	return c.CurrentState
}

func (c *ContainerGroup) SetCurrentState(currentState ContainerGroupState) {
	c.CurrentState = &currentState
}

func (c *ContainerGroup) GetDisplayName() *string {
	if c == nil {
		return nil
	}
	return c.DisplayName
}

func (c *ContainerGroup) SetDisplayName(displayName string) {
	c.DisplayName = &displayName
}

func (c *ContainerGroup) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *ContainerGroup) SetId(id string) {
	c.Id = &id
}

func (c *ContainerGroup) GetLivenessProbe() *util.Nullable[ContainerGroupLivenessProbe] {
	if c == nil {
		return nil
	}
	return c.LivenessProbe
}

func (c *ContainerGroup) SetLivenessProbe(livenessProbe util.Nullable[ContainerGroupLivenessProbe]) {
	c.LivenessProbe = &livenessProbe
}

func (c *ContainerGroup) SetLivenessProbeNull() {
	c.LivenessProbe = &util.Nullable[ContainerGroupLivenessProbe]{IsNull: true}
}

func (c *ContainerGroup) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *ContainerGroup) SetName(name string) {
	c.Name = &name
}

func (c *ContainerGroup) GetNetworking() *ContainerGroupNetworkingConfiguration {
	if c == nil {
		return nil
	}
	return c.Networking
}

func (c *ContainerGroup) SetNetworking(networking ContainerGroupNetworkingConfiguration) {
	c.Networking = &networking
}

func (c *ContainerGroup) GetOrganizationName() *string {
	if c == nil {
		return nil
	}
	return c.OrganizationName
}

func (c *ContainerGroup) SetOrganizationName(organizationName string) {
	c.OrganizationName = &organizationName
}

func (c *ContainerGroup) GetPendingChange() *bool {
	if c == nil {
		return nil
	}
	return c.PendingChange
}

func (c *ContainerGroup) SetPendingChange(pendingChange bool) {
	c.PendingChange = &pendingChange
}

func (c *ContainerGroup) GetPriority() *util.Nullable[ContainerGroupPriority] {
	if c == nil {
		return nil
	}
	return c.Priority
}

func (c *ContainerGroup) SetPriority(priority util.Nullable[ContainerGroupPriority]) {
	c.Priority = &priority
}

func (c *ContainerGroup) SetPriorityNull() {
	c.Priority = &util.Nullable[ContainerGroupPriority]{IsNull: true}
}

func (c *ContainerGroup) GetProjectName() *string {
	if c == nil {
		return nil
	}
	return c.ProjectName
}

func (c *ContainerGroup) SetProjectName(projectName string) {
	c.ProjectName = &projectName
}

func (c *ContainerGroup) GetQueueAutoscaler() *QueueBasedAutoscalerConfiguration {
	if c == nil {
		return nil
	}
	return c.QueueAutoscaler
}

func (c *ContainerGroup) SetQueueAutoscaler(queueAutoscaler QueueBasedAutoscalerConfiguration) {
	c.QueueAutoscaler = &queueAutoscaler
}

func (c *ContainerGroup) GetQueueConnection() *ContainerGroupQueueConnection {
	if c == nil {
		return nil
	}
	return c.QueueConnection
}

func (c *ContainerGroup) SetQueueConnection(queueConnection ContainerGroupQueueConnection) {
	c.QueueConnection = &queueConnection
}

func (c *ContainerGroup) GetReadinessProbe() *util.Nullable[ContainerGroupReadinessProbe] {
	if c == nil {
		return nil
	}
	return c.ReadinessProbe
}

func (c *ContainerGroup) SetReadinessProbe(readinessProbe util.Nullable[ContainerGroupReadinessProbe]) {
	c.ReadinessProbe = &readinessProbe
}

func (c *ContainerGroup) SetReadinessProbeNull() {
	c.ReadinessProbe = &util.Nullable[ContainerGroupReadinessProbe]{IsNull: true}
}

func (c *ContainerGroup) GetReplicas() *int64 {
	if c == nil {
		return nil
	}
	return c.Replicas
}

func (c *ContainerGroup) SetReplicas(replicas int64) {
	c.Replicas = &replicas
}

func (c *ContainerGroup) GetRestartPolicy() *ContainerRestartPolicy {
	if c == nil {
		return nil
	}
	return c.RestartPolicy
}

func (c *ContainerGroup) SetRestartPolicy(restartPolicy ContainerRestartPolicy) {
	c.RestartPolicy = &restartPolicy
}

func (c *ContainerGroup) GetStartupProbe() *util.Nullable[ContainerGroupStartupProbe] {
	if c == nil {
		return nil
	}
	return c.StartupProbe
}

func (c *ContainerGroup) SetStartupProbe(startupProbe util.Nullable[ContainerGroupStartupProbe]) {
	c.StartupProbe = &startupProbe
}

func (c *ContainerGroup) SetStartupProbeNull() {
	c.StartupProbe = &util.Nullable[ContainerGroupStartupProbe]{IsNull: true}
}

func (c *ContainerGroup) GetUpdateTime() *string {
	if c == nil {
		return nil
	}
	return c.UpdateTime
}

func (c *ContainerGroup) SetUpdateTime(updateTime string) {
	c.UpdateTime = &updateTime
}

func (c *ContainerGroup) GetVersion() *int64 {
	if c == nil {
		return nil
	}
	return c.Version
}

func (c *ContainerGroup) SetVersion(version int64) {
	c.Version = &version
}

func (c ContainerGroup) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroup to string"
	}
	return string(jsonData)
}
