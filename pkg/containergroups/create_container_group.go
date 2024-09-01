package containergroups

import (
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
}

func (c *CreateContainerGroup) SetName(name string) {
	c.Name = &name
}

func (c *CreateContainerGroup) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateContainerGroup) SetDisplayName(displayName string) {
	c.DisplayName = &displayName
}

func (c *CreateContainerGroup) GetDisplayName() *string {
	if c == nil {
		return nil
	}
	return c.DisplayName
}

func (c *CreateContainerGroup) SetContainer(container CreateContainer) {
	c.Container = &container
}

func (c *CreateContainerGroup) GetContainer() *CreateContainer {
	if c == nil {
		return nil
	}
	return c.Container
}

func (c *CreateContainerGroup) SetAutostartPolicy(autostartPolicy bool) {
	c.AutostartPolicy = &autostartPolicy
}

func (c *CreateContainerGroup) GetAutostartPolicy() *bool {
	if c == nil {
		return nil
	}
	return c.AutostartPolicy
}

func (c *CreateContainerGroup) SetRestartPolicy(restartPolicy shared.ContainerRestartPolicy) {
	c.RestartPolicy = &restartPolicy
}

func (c *CreateContainerGroup) GetRestartPolicy() *shared.ContainerRestartPolicy {
	if c == nil {
		return nil
	}
	return c.RestartPolicy
}

func (c *CreateContainerGroup) SetReplicas(replicas int64) {
	c.Replicas = &replicas
}

func (c *CreateContainerGroup) GetReplicas() *int64 {
	if c == nil {
		return nil
	}
	return c.Replicas
}

func (c *CreateContainerGroup) SetCountryCodes(countryCodes []shared.CountryCode) {
	c.CountryCodes = countryCodes
}

func (c *CreateContainerGroup) GetCountryCodes() []shared.CountryCode {
	if c == nil {
		return nil
	}
	return c.CountryCodes
}

func (c *CreateContainerGroup) SetNetworking(networking CreateContainerGroupNetworking) {
	c.Networking = &networking
}

func (c *CreateContainerGroup) GetNetworking() *CreateContainerGroupNetworking {
	if c == nil {
		return nil
	}
	return c.Networking
}

func (c *CreateContainerGroup) SetLivenessProbe(livenessProbe shared.ContainerGroupLivenessProbe) {
	c.LivenessProbe = &livenessProbe
}

func (c *CreateContainerGroup) GetLivenessProbe() *shared.ContainerGroupLivenessProbe {
	if c == nil {
		return nil
	}
	return c.LivenessProbe
}

func (c *CreateContainerGroup) SetReadinessProbe(readinessProbe shared.ContainerGroupReadinessProbe) {
	c.ReadinessProbe = &readinessProbe
}

func (c *CreateContainerGroup) GetReadinessProbe() *shared.ContainerGroupReadinessProbe {
	if c == nil {
		return nil
	}
	return c.ReadinessProbe
}

func (c *CreateContainerGroup) SetStartupProbe(startupProbe shared.ContainerGroupStartupProbe) {
	c.StartupProbe = &startupProbe
}

func (c *CreateContainerGroup) GetStartupProbe() *shared.ContainerGroupStartupProbe {
	if c == nil {
		return nil
	}
	return c.StartupProbe
}

func (c *CreateContainerGroup) SetQueueConnection(queueConnection shared.ContainerGroupQueueConnection) {
	c.QueueConnection = &queueConnection
}

func (c *CreateContainerGroup) GetQueueConnection() *shared.ContainerGroupQueueConnection {
	if c == nil {
		return nil
	}
	return c.QueueConnection
}
