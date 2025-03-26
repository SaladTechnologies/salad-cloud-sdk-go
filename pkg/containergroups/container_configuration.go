package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Configuration for creating a container within a container group. Defines the container image, resource requirements, environment variables, and other settings needed to deploy and run the container.
type ContainerConfiguration struct {
	// Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image. Each element in the array represents a command segment or argument.
	Command *util.Nullable[[]string] `json:"command,omitempty" maxItems:"100"`
	// Key-value pairs of environment variables to set within the container. These variables will be available to processes running inside the container.
	EnvironmentVariables map[string]*string `json:"environment_variables,omitempty"`
	// The container image.
	Image *string `json:"image,omitempty" required:"true" maxLength:"2048" minLength:"1" pattern:"^.*$"`
	// The container image caching.
	ImageCaching *bool `json:"image_caching,omitempty"`
	// Configuration options for directing container logs to a logging provider. This schema enables you to specify a single logging destination for container output, supporting monitoring, debugging, and analytics use cases. Each provider has its own configuration parameters defined in the referenced schemas. Only one logging provider can be selected at a time.
	Logging *ContainerConfigurationLogging `json:"logging,omitempty"`
	// Specifies the priority level for container group execution, which determines resource allocation and scheduling precedence.
	Priority *util.Nullable[shared.ContainerGroupPriority] `json:"priority,omitempty"`
	// Authentication configuration for various container registry types, including AWS ECR, Docker Hub, GCP GAR, GCP GCR, and basic authentication.
	RegistryAuthentication *ContainerRegistryAuthentication `json:"registry_authentication,omitempty"`
	// Specifies the resource requirements for a container.
	Resources *shared.ContainerResourceRequirements `json:"resources,omitempty" required:"true"`
}

func (c *ContainerConfiguration) GetCommand() *util.Nullable[[]string] {
	if c == nil {
		return nil
	}
	return c.Command
}

func (c *ContainerConfiguration) SetCommand(command util.Nullable[[]string]) {
	c.Command = &command
}

func (c *ContainerConfiguration) SetCommandNull() {
	c.Command = &util.Nullable[[]string]{IsNull: true}
}

func (c *ContainerConfiguration) GetEnvironmentVariables() map[string]*string {
	if c == nil {
		return nil
	}
	return c.EnvironmentVariables
}

func (c *ContainerConfiguration) SetEnvironmentVariables(environmentVariables map[string]*string) {
	c.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (c *ContainerConfiguration) GetImage() *string {
	if c == nil {
		return nil
	}
	return c.Image
}

func (c *ContainerConfiguration) SetImage(image string) {
	c.Image = &image
}

func (c *ContainerConfiguration) GetImageCaching() *bool {
	if c == nil {
		return nil
	}
	return c.ImageCaching
}

func (c *ContainerConfiguration) SetImageCaching(imageCaching bool) {
	c.ImageCaching = &imageCaching
}

func (c *ContainerConfiguration) GetLogging() *ContainerConfigurationLogging {
	if c == nil {
		return nil
	}
	return c.Logging
}

func (c *ContainerConfiguration) SetLogging(logging ContainerConfigurationLogging) {
	c.Logging = &logging
}

func (c *ContainerConfiguration) GetPriority() *util.Nullable[shared.ContainerGroupPriority] {
	if c == nil {
		return nil
	}
	return c.Priority
}

func (c *ContainerConfiguration) SetPriority(priority util.Nullable[shared.ContainerGroupPriority]) {
	c.Priority = &priority
}

func (c *ContainerConfiguration) SetPriorityNull() {
	c.Priority = &util.Nullable[shared.ContainerGroupPriority]{IsNull: true}
}

func (c *ContainerConfiguration) GetRegistryAuthentication() *ContainerRegistryAuthentication {
	if c == nil {
		return nil
	}
	return c.RegistryAuthentication
}

func (c *ContainerConfiguration) SetRegistryAuthentication(registryAuthentication ContainerRegistryAuthentication) {
	c.RegistryAuthentication = &registryAuthentication
}

func (c *ContainerConfiguration) GetResources() *shared.ContainerResourceRequirements {
	if c == nil {
		return nil
	}
	return c.Resources
}

func (c *ContainerConfiguration) SetResources(resources shared.ContainerResourceRequirements) {
	c.Resources = &resources
}

func (c ContainerConfiguration) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerConfiguration to string"
	}
	return string(jsonData)
}
