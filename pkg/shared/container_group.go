package shared

import (
	"encoding/json"
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
	Networking *ContainerGroupNetworking `json:"networking,omitempty"`
	// Represents the container group liveness probe
	LivenessProbe *ContainerGroupLivenessProbe `json:"liveness_probe,omitempty"`
	// Represents the container group readiness probe
	ReadinessProbe *ContainerGroupReadinessProbe `json:"readiness_probe,omitempty"`
	// Represents the container group startup probe
	StartupProbe *ContainerGroupStartupProbe `json:"startup_probe,omitempty"`
	// Represents container group queue connection
	QueueConnection *ContainerGroupQueueConnection `json:"queue_connection,omitempty"`
	CreateTime      *string                        `json:"create_time,omitempty" required:"true"`
	UpdateTime      *string                        `json:"update_time,omitempty" required:"true"`
	PendingChange   *bool                          `json:"pending_change,omitempty" required:"true"`
	Version         *int64                         `json:"version,omitempty" required:"true" min:"1"`
	// Represents the autoscaling rules for a queue
	QueueAutoscaler *QueueAutoscaler `json:"queue_autoscaler,omitempty"`
	touched         map[string]bool
}

func (c *ContainerGroup) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *ContainerGroup) SetId(id string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Id"] = true
	c.Id = &id
}

func (c *ContainerGroup) SetIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Id"] = true
	c.Id = nil
}

func (c *ContainerGroup) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *ContainerGroup) SetName(name string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = &name
}

func (c *ContainerGroup) SetNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = nil
}

func (c *ContainerGroup) GetDisplayName() *string {
	if c == nil {
		return nil
	}
	return c.DisplayName
}

func (c *ContainerGroup) SetDisplayName(displayName string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["DisplayName"] = true
	c.DisplayName = &displayName
}

func (c *ContainerGroup) SetDisplayNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["DisplayName"] = true
	c.DisplayName = nil
}

func (c *ContainerGroup) GetContainer() *Container {
	if c == nil {
		return nil
	}
	return c.Container
}

func (c *ContainerGroup) SetContainer(container Container) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Container"] = true
	c.Container = &container
}

func (c *ContainerGroup) SetContainerNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Container"] = true
	c.Container = nil
}

func (c *ContainerGroup) GetAutostartPolicy() *bool {
	if c == nil {
		return nil
	}
	return c.AutostartPolicy
}

func (c *ContainerGroup) SetAutostartPolicy(autostartPolicy bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AutostartPolicy"] = true
	c.AutostartPolicy = &autostartPolicy
}

func (c *ContainerGroup) SetAutostartPolicyNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AutostartPolicy"] = true
	c.AutostartPolicy = nil
}

func (c *ContainerGroup) GetRestartPolicy() *ContainerRestartPolicy {
	if c == nil {
		return nil
	}
	return c.RestartPolicy
}

func (c *ContainerGroup) SetRestartPolicy(restartPolicy ContainerRestartPolicy) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["RestartPolicy"] = true
	c.RestartPolicy = &restartPolicy
}

func (c *ContainerGroup) SetRestartPolicyNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["RestartPolicy"] = true
	c.RestartPolicy = nil
}

func (c *ContainerGroup) GetReplicas() *int64 {
	if c == nil {
		return nil
	}
	return c.Replicas
}

func (c *ContainerGroup) SetReplicas(replicas int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Replicas"] = true
	c.Replicas = &replicas
}

func (c *ContainerGroup) SetReplicasNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Replicas"] = true
	c.Replicas = nil
}

func (c *ContainerGroup) GetCurrentState() *ContainerGroupState {
	if c == nil {
		return nil
	}
	return c.CurrentState
}

func (c *ContainerGroup) SetCurrentState(currentState ContainerGroupState) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CurrentState"] = true
	c.CurrentState = &currentState
}

func (c *ContainerGroup) SetCurrentStateNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CurrentState"] = true
	c.CurrentState = nil
}

func (c *ContainerGroup) GetCountryCodes() []CountryCode {
	if c == nil {
		return nil
	}
	return c.CountryCodes
}

func (c *ContainerGroup) SetCountryCodes(countryCodes []CountryCode) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CountryCodes"] = true
	c.CountryCodes = countryCodes
}

func (c *ContainerGroup) SetCountryCodesNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CountryCodes"] = true
	c.CountryCodes = nil
}

func (c *ContainerGroup) GetNetworking() *ContainerGroupNetworking {
	if c == nil {
		return nil
	}
	return c.Networking
}

func (c *ContainerGroup) SetNetworking(networking ContainerGroupNetworking) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Networking"] = true
	c.Networking = &networking
}

func (c *ContainerGroup) SetNetworkingNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Networking"] = true
	c.Networking = nil
}

func (c *ContainerGroup) GetLivenessProbe() *ContainerGroupLivenessProbe {
	if c == nil {
		return nil
	}
	return c.LivenessProbe
}

func (c *ContainerGroup) SetLivenessProbe(livenessProbe ContainerGroupLivenessProbe) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["LivenessProbe"] = true
	c.LivenessProbe = &livenessProbe
}

func (c *ContainerGroup) SetLivenessProbeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["LivenessProbe"] = true
	c.LivenessProbe = nil
}

func (c *ContainerGroup) GetReadinessProbe() *ContainerGroupReadinessProbe {
	if c == nil {
		return nil
	}
	return c.ReadinessProbe
}

func (c *ContainerGroup) SetReadinessProbe(readinessProbe ContainerGroupReadinessProbe) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ReadinessProbe"] = true
	c.ReadinessProbe = &readinessProbe
}

func (c *ContainerGroup) SetReadinessProbeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ReadinessProbe"] = true
	c.ReadinessProbe = nil
}

func (c *ContainerGroup) GetStartupProbe() *ContainerGroupStartupProbe {
	if c == nil {
		return nil
	}
	return c.StartupProbe
}

func (c *ContainerGroup) SetStartupProbe(startupProbe ContainerGroupStartupProbe) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartupProbe"] = true
	c.StartupProbe = &startupProbe
}

func (c *ContainerGroup) SetStartupProbeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartupProbe"] = true
	c.StartupProbe = nil
}

func (c *ContainerGroup) GetQueueConnection() *ContainerGroupQueueConnection {
	if c == nil {
		return nil
	}
	return c.QueueConnection
}

func (c *ContainerGroup) SetQueueConnection(queueConnection ContainerGroupQueueConnection) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueueConnection"] = true
	c.QueueConnection = &queueConnection
}

func (c *ContainerGroup) SetQueueConnectionNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueueConnection"] = true
	c.QueueConnection = nil
}

func (c *ContainerGroup) GetCreateTime() *string {
	if c == nil {
		return nil
	}
	return c.CreateTime
}

func (c *ContainerGroup) SetCreateTime(createTime string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CreateTime"] = true
	c.CreateTime = &createTime
}

func (c *ContainerGroup) SetCreateTimeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CreateTime"] = true
	c.CreateTime = nil
}

func (c *ContainerGroup) GetUpdateTime() *string {
	if c == nil {
		return nil
	}
	return c.UpdateTime
}

func (c *ContainerGroup) SetUpdateTime(updateTime string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["UpdateTime"] = true
	c.UpdateTime = &updateTime
}

func (c *ContainerGroup) SetUpdateTimeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["UpdateTime"] = true
	c.UpdateTime = nil
}

func (c *ContainerGroup) GetPendingChange() *bool {
	if c == nil {
		return nil
	}
	return c.PendingChange
}

func (c *ContainerGroup) SetPendingChange(pendingChange bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["PendingChange"] = true
	c.PendingChange = &pendingChange
}

func (c *ContainerGroup) SetPendingChangeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["PendingChange"] = true
	c.PendingChange = nil
}

func (c *ContainerGroup) GetVersion() *int64 {
	if c == nil {
		return nil
	}
	return c.Version
}

func (c *ContainerGroup) SetVersion(version int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Version"] = true
	c.Version = &version
}

func (c *ContainerGroup) SetVersionNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Version"] = true
	c.Version = nil
}

func (c *ContainerGroup) GetQueueAutoscaler() *QueueAutoscaler {
	if c == nil {
		return nil
	}
	return c.QueueAutoscaler
}

func (c *ContainerGroup) SetQueueAutoscaler(queueAutoscaler QueueAutoscaler) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueueAutoscaler"] = true
	c.QueueAutoscaler = &queueAutoscaler
}

func (c *ContainerGroup) SetQueueAutoscalerNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueueAutoscaler"] = true
	c.QueueAutoscaler = nil
}

func (c ContainerGroup) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Id"] && c.Id == nil {
		data["id"] = nil
	} else if c.Id != nil {
		data["id"] = c.Id
	}

	if c.touched["Name"] && c.Name == nil {
		data["name"] = nil
	} else if c.Name != nil {
		data["name"] = c.Name
	}

	if c.touched["DisplayName"] && c.DisplayName == nil {
		data["display_name"] = nil
	} else if c.DisplayName != nil {
		data["display_name"] = c.DisplayName
	}

	if c.touched["Container"] && c.Container == nil {
		data["container"] = nil
	} else if c.Container != nil {
		data["container"] = c.Container
	}

	if c.touched["AutostartPolicy"] && c.AutostartPolicy == nil {
		data["autostart_policy"] = nil
	} else if c.AutostartPolicy != nil {
		data["autostart_policy"] = c.AutostartPolicy
	}

	if c.touched["RestartPolicy"] && c.RestartPolicy == nil {
		data["restart_policy"] = nil
	} else if c.RestartPolicy != nil {
		data["restart_policy"] = c.RestartPolicy
	}

	if c.touched["Replicas"] && c.Replicas == nil {
		data["replicas"] = nil
	} else if c.Replicas != nil {
		data["replicas"] = c.Replicas
	}

	if c.touched["CurrentState"] && c.CurrentState == nil {
		data["current_state"] = nil
	} else if c.CurrentState != nil {
		data["current_state"] = c.CurrentState
	}

	if c.touched["CountryCodes"] && c.CountryCodes == nil {
		data["country_codes"] = nil
	} else if c.CountryCodes != nil {
		data["country_codes"] = c.CountryCodes
	}

	if c.touched["Networking"] && c.Networking == nil {
		data["networking"] = nil
	} else if c.Networking != nil {
		data["networking"] = c.Networking
	}

	if c.touched["LivenessProbe"] && c.LivenessProbe == nil {
		data["liveness_probe"] = nil
	} else if c.LivenessProbe != nil {
		data["liveness_probe"] = c.LivenessProbe
	}

	if c.touched["ReadinessProbe"] && c.ReadinessProbe == nil {
		data["readiness_probe"] = nil
	} else if c.ReadinessProbe != nil {
		data["readiness_probe"] = c.ReadinessProbe
	}

	if c.touched["StartupProbe"] && c.StartupProbe == nil {
		data["startup_probe"] = nil
	} else if c.StartupProbe != nil {
		data["startup_probe"] = c.StartupProbe
	}

	if c.touched["QueueConnection"] && c.QueueConnection == nil {
		data["queue_connection"] = nil
	} else if c.QueueConnection != nil {
		data["queue_connection"] = c.QueueConnection
	}

	if c.touched["CreateTime"] && c.CreateTime == nil {
		data["create_time"] = nil
	} else if c.CreateTime != nil {
		data["create_time"] = c.CreateTime
	}

	if c.touched["UpdateTime"] && c.UpdateTime == nil {
		data["update_time"] = nil
	} else if c.UpdateTime != nil {
		data["update_time"] = c.UpdateTime
	}

	if c.touched["PendingChange"] && c.PendingChange == nil {
		data["pending_change"] = nil
	} else if c.PendingChange != nil {
		data["pending_change"] = c.PendingChange
	}

	if c.touched["Version"] && c.Version == nil {
		data["version"] = nil
	} else if c.Version != nil {
		data["version"] = c.Version
	}

	if c.touched["QueueAutoscaler"] && c.QueueAutoscaler == nil {
		data["queue_autoscaler"] = nil
	} else if c.QueueAutoscaler != nil {
		data["queue_autoscaler"] = c.QueueAutoscaler
	}

	return json.Marshal(data)
}

func (c ContainerGroup) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroup to string"
	}
	return string(jsonData)
}
