package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a container group
type ContainerGroup struct {
	Id          *string `json:"id,omitempty" required:"true"`
	Name        *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	DisplayName *string `json:"display_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// Represents a container
	Container       *Container              `json:"container,omitempty" required:"true"`
	AutostartPolicy *bool                   `json:"autostart_policy,omitempty" required:"true"`
	RestartPolicy   *ContainerRestartPolicy `json:"restart_policy,omitempty" required:"true"`
	Replicas        *int64                  `json:"replicas,omitempty" required:"true" min:"0" max:"100"`
	// Represents a container group state
	CurrentState *ContainerGroupState `json:"current_state,omitempty" required:"true"`
	// List of countries nodes must be located in. Remove this field to permit nodes from any country.
	CountryCodes []CountryCode `json:"country_codes,omitempty" minItems:"1" maxItems:"500"`
	// Represents container group networking parameters
	Networking *util.Nullable[ContainerGroupNetworking] `json:"networking,omitempty"`
	// Represents the container group liveness probe
	LivenessProbe *util.Nullable[ContainerGroupLivenessProbe] `json:"liveness_probe,omitempty"`
	// Represents the container group readiness probe
	ReadinessProbe *util.Nullable[ContainerGroupReadinessProbe] `json:"readiness_probe,omitempty"`
	// Represents the container group startup probe
	StartupProbe *util.Nullable[ContainerGroupStartupProbe] `json:"startup_probe,omitempty"`
	// Represents container group queue connection
	QueueConnection *util.Nullable[ContainerGroupQueueConnection] `json:"queue_connection,omitempty"`
	CreateTime      *string                                       `json:"create_time,omitempty" required:"true"`
	UpdateTime      *string                                       `json:"update_time,omitempty" required:"true"`
	PendingChange   *bool                                         `json:"pending_change,omitempty" required:"true"`
	Version         *int64                                        `json:"version,omitempty" required:"true" min:"1"`
	// Represents the autoscaling rules for a queue
	QueueAutoscaler *util.Nullable[QueueAutoscaler] `json:"queue_autoscaler,omitempty"`
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

func (c *ContainerGroup) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *ContainerGroup) SetName(name string) {
	c.Name = &name
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

func (c *ContainerGroup) GetContainer() *Container {
	if c == nil {
		return nil
	}
	return c.Container
}

func (c *ContainerGroup) SetContainer(container Container) {
	c.Container = &container
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

func (c *ContainerGroup) GetRestartPolicy() *ContainerRestartPolicy {
	if c == nil {
		return nil
	}
	return c.RestartPolicy
}

func (c *ContainerGroup) SetRestartPolicy(restartPolicy ContainerRestartPolicy) {
	c.RestartPolicy = &restartPolicy
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

func (c *ContainerGroup) GetCurrentState() *ContainerGroupState {
	if c == nil {
		return nil
	}
	return c.CurrentState
}

func (c *ContainerGroup) SetCurrentState(currentState ContainerGroupState) {
	c.CurrentState = &currentState
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

func (c *ContainerGroup) GetNetworking() *util.Nullable[ContainerGroupNetworking] {
	if c == nil {
		return nil
	}
	return c.Networking
}

func (c *ContainerGroup) SetNetworking(networking util.Nullable[ContainerGroupNetworking]) {
	c.Networking = &networking
}

func (c *ContainerGroup) SetNetworkingNull() {
	c.Networking = &util.Nullable[ContainerGroupNetworking]{IsNull: true}
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

func (c *ContainerGroup) GetQueueConnection() *util.Nullable[ContainerGroupQueueConnection] {
	if c == nil {
		return nil
	}
	return c.QueueConnection
}

func (c *ContainerGroup) SetQueueConnection(queueConnection util.Nullable[ContainerGroupQueueConnection]) {
	c.QueueConnection = &queueConnection
}

func (c *ContainerGroup) SetQueueConnectionNull() {
	c.QueueConnection = &util.Nullable[ContainerGroupQueueConnection]{IsNull: true}
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

func (c *ContainerGroup) GetUpdateTime() *string {
	if c == nil {
		return nil
	}
	return c.UpdateTime
}

func (c *ContainerGroup) SetUpdateTime(updateTime string) {
	c.UpdateTime = &updateTime
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

func (c *ContainerGroup) GetVersion() *int64 {
	if c == nil {
		return nil
	}
	return c.Version
}

func (c *ContainerGroup) SetVersion(version int64) {
	c.Version = &version
}

func (c *ContainerGroup) GetQueueAutoscaler() *util.Nullable[QueueAutoscaler] {
	if c == nil {
		return nil
	}
	return c.QueueAutoscaler
}

func (c *ContainerGroup) SetQueueAutoscaler(queueAutoscaler util.Nullable[QueueAutoscaler]) {
	c.QueueAutoscaler = &queueAutoscaler
}

func (c *ContainerGroup) SetQueueAutoscalerNull() {
	c.QueueAutoscaler = &util.Nullable[QueueAutoscaler]{IsNull: true}
}

func (c ContainerGroup) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroup to string"
	}
	return string(jsonData)
}
