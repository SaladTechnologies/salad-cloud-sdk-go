package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents a request to create a container group
type CreateContainerGroup struct {
	Name        *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	DisplayName *string `json:"display_name,omitempty" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// Represents a container
	Container       *CreateContainer               `json:"container,omitempty" required:"true"`
	AutostartPolicy *bool                          `json:"autostart_policy,omitempty" required:"true"`
	RestartPolicy   *shared.ContainerRestartPolicy `json:"restart_policy,omitempty" required:"true"`
	Replicas        *int64                         `json:"replicas,omitempty" required:"true" min:"0" max:"250"`
	// List of countries nodes must be located in. Remove this field to permit nodes from any country.
	CountryCodes []shared.CountryCode `json:"country_codes,omitempty" minItems:"1" maxItems:"500"`
	// Represents container group networking parameters
	Networking *CreateContainerGroupNetworking `json:"networking,omitempty"`
	// Represents the container group liveness probe
	LivenessProbe *shared.ContainerGroupLivenessProbe `json:"liveness_probe,omitempty"`
	// Represents the container group readiness probe
	ReadinessProbe *shared.ContainerGroupReadinessProbe `json:"readiness_probe,omitempty"`
	// Represents the container group startup probe
	StartupProbe *shared.ContainerGroupStartupProbe `json:"startup_probe,omitempty"`
	// Represents container group queue connection
	QueueConnection *shared.ContainerGroupQueueConnection `json:"queue_connection,omitempty"`
	// Represents the autoscaling rules for a queue
	QueueAutoscaler *shared.QueueAutoscaler `json:"queue_autoscaler,omitempty"`
	touched         map[string]bool
}

func (c *CreateContainerGroup) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateContainerGroup) SetName(name string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = &name
}

func (c *CreateContainerGroup) SetNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = nil
}

func (c *CreateContainerGroup) GetDisplayName() *string {
	if c == nil {
		return nil
	}
	return c.DisplayName
}

func (c *CreateContainerGroup) SetDisplayName(displayName string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["DisplayName"] = true
	c.DisplayName = &displayName
}

func (c *CreateContainerGroup) SetDisplayNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["DisplayName"] = true
	c.DisplayName = nil
}

func (c *CreateContainerGroup) GetContainer() *CreateContainer {
	if c == nil {
		return nil
	}
	return c.Container
}

func (c *CreateContainerGroup) SetContainer(container CreateContainer) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Container"] = true
	c.Container = &container
}

func (c *CreateContainerGroup) SetContainerNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Container"] = true
	c.Container = nil
}

func (c *CreateContainerGroup) GetAutostartPolicy() *bool {
	if c == nil {
		return nil
	}
	return c.AutostartPolicy
}

func (c *CreateContainerGroup) SetAutostartPolicy(autostartPolicy bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AutostartPolicy"] = true
	c.AutostartPolicy = &autostartPolicy
}

func (c *CreateContainerGroup) SetAutostartPolicyNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AutostartPolicy"] = true
	c.AutostartPolicy = nil
}

func (c *CreateContainerGroup) GetRestartPolicy() *shared.ContainerRestartPolicy {
	if c == nil {
		return nil
	}
	return c.RestartPolicy
}

func (c *CreateContainerGroup) SetRestartPolicy(restartPolicy shared.ContainerRestartPolicy) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["RestartPolicy"] = true
	c.RestartPolicy = &restartPolicy
}

func (c *CreateContainerGroup) SetRestartPolicyNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["RestartPolicy"] = true
	c.RestartPolicy = nil
}

func (c *CreateContainerGroup) GetReplicas() *int64 {
	if c == nil {
		return nil
	}
	return c.Replicas
}

func (c *CreateContainerGroup) SetReplicas(replicas int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Replicas"] = true
	c.Replicas = &replicas
}

func (c *CreateContainerGroup) SetReplicasNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Replicas"] = true
	c.Replicas = nil
}

func (c *CreateContainerGroup) GetCountryCodes() []shared.CountryCode {
	if c == nil {
		return nil
	}
	return c.CountryCodes
}

func (c *CreateContainerGroup) SetCountryCodes(countryCodes []shared.CountryCode) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CountryCodes"] = true
	c.CountryCodes = countryCodes
}

func (c *CreateContainerGroup) SetCountryCodesNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CountryCodes"] = true
	c.CountryCodes = nil
}

func (c *CreateContainerGroup) GetNetworking() *CreateContainerGroupNetworking {
	if c == nil {
		return nil
	}
	return c.Networking
}

func (c *CreateContainerGroup) SetNetworking(networking CreateContainerGroupNetworking) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Networking"] = true
	c.Networking = &networking
}

func (c *CreateContainerGroup) SetNetworkingNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Networking"] = true
	c.Networking = nil
}

func (c *CreateContainerGroup) GetLivenessProbe() *shared.ContainerGroupLivenessProbe {
	if c == nil {
		return nil
	}
	return c.LivenessProbe
}

func (c *CreateContainerGroup) SetLivenessProbe(livenessProbe shared.ContainerGroupLivenessProbe) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["LivenessProbe"] = true
	c.LivenessProbe = &livenessProbe
}

func (c *CreateContainerGroup) SetLivenessProbeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["LivenessProbe"] = true
	c.LivenessProbe = nil
}

func (c *CreateContainerGroup) GetReadinessProbe() *shared.ContainerGroupReadinessProbe {
	if c == nil {
		return nil
	}
	return c.ReadinessProbe
}

func (c *CreateContainerGroup) SetReadinessProbe(readinessProbe shared.ContainerGroupReadinessProbe) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ReadinessProbe"] = true
	c.ReadinessProbe = &readinessProbe
}

func (c *CreateContainerGroup) SetReadinessProbeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ReadinessProbe"] = true
	c.ReadinessProbe = nil
}

func (c *CreateContainerGroup) GetStartupProbe() *shared.ContainerGroupStartupProbe {
	if c == nil {
		return nil
	}
	return c.StartupProbe
}

func (c *CreateContainerGroup) SetStartupProbe(startupProbe shared.ContainerGroupStartupProbe) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartupProbe"] = true
	c.StartupProbe = &startupProbe
}

func (c *CreateContainerGroup) SetStartupProbeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["StartupProbe"] = true
	c.StartupProbe = nil
}

func (c *CreateContainerGroup) GetQueueConnection() *shared.ContainerGroupQueueConnection {
	if c == nil {
		return nil
	}
	return c.QueueConnection
}

func (c *CreateContainerGroup) SetQueueConnection(queueConnection shared.ContainerGroupQueueConnection) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueueConnection"] = true
	c.QueueConnection = &queueConnection
}

func (c *CreateContainerGroup) SetQueueConnectionNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueueConnection"] = true
	c.QueueConnection = nil
}

func (c *CreateContainerGroup) GetQueueAutoscaler() *shared.QueueAutoscaler {
	if c == nil {
		return nil
	}
	return c.QueueAutoscaler
}

func (c *CreateContainerGroup) SetQueueAutoscaler(queueAutoscaler shared.QueueAutoscaler) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueueAutoscaler"] = true
	c.QueueAutoscaler = &queueAutoscaler
}

func (c *CreateContainerGroup) SetQueueAutoscalerNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueueAutoscaler"] = true
	c.QueueAutoscaler = nil
}

func (c CreateContainerGroup) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

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

	if c.touched["QueueAutoscaler"] && c.QueueAutoscaler == nil {
		data["queue_autoscaler"] = nil
	} else if c.QueueAutoscaler != nil {
		data["queue_autoscaler"] = c.QueueAutoscaler
	}

	return json.Marshal(data)
}

func (c CreateContainerGroup) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateContainerGroup to string"
	}
	return string(jsonData)
}
