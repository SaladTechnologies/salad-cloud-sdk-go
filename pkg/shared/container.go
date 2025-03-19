package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a container with its configuration and resource requirements.
type Container struct {
	// List of commands to run inside the container. Each command is a string representing a command-line instruction.
	Command *util.Nullable[[]string] `json:"command,omitempty" required:"true" maxItems:"100"`
	// Environment variables to set in the container.
	EnvironmentVariables map[string]*string `json:"environment_variables,omitempty"`
	// SHA-256 hash (64-character hexadecimal string)
	Hash *string `json:"hash,omitempty" maxLength:"64" minLength:"64" pattern:"^[a-fA-F0-9]{64}$"`
	// The container image.
	Image *string `json:"image,omitempty" required:"true" maxLength:"2048" minLength:"1" pattern:"^.*$"`
	// The container image caching.
	ImageCaching *bool `json:"image_caching,omitempty"`
	// Configuration options for directing container logs to a logging provider. This schema enables you to specify a single logging destination for container output, supporting monitoring, debugging, and analytics use cases. Each provider has its own configuration parameters defined in the referenced schemas. Only one logging provider can be selected at a time.
	Logging *ContainerLoggingConfiguration `json:"logging,omitempty"`
	// Specifies the resource requirements for a container.
	Resources *ContainerResourceRequirements `json:"resources,omitempty" required:"true"`
	// Size of the container in bytes.
	Size *int64 `json:"size,omitempty" min:"0" max:"9223372036854776000"`
}

func (c *Container) GetCommand() *util.Nullable[[]string] {
	if c == nil {
		return nil
	}
	return c.Command
}

func (c *Container) SetCommand(command util.Nullable[[]string]) {
	c.Command = &command
}

func (c *Container) SetCommandNull() {
	c.Command = &util.Nullable[[]string]{IsNull: true}
}

func (c *Container) GetEnvironmentVariables() map[string]*string {
	if c == nil {
		return nil
	}
	return c.EnvironmentVariables
}

func (c *Container) SetEnvironmentVariables(environmentVariables map[string]*string) {
	c.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (c *Container) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *Container) SetHash(hash string) {
	c.Hash = &hash
}

func (c *Container) GetImage() *string {
	if c == nil {
		return nil
	}
	return c.Image
}

func (c *Container) SetImage(image string) {
	c.Image = &image
}

func (c *Container) GetImageCaching() *bool {
	if c == nil {
		return nil
	}
	return c.ImageCaching
}

func (c *Container) SetImageCaching(imageCaching bool) {
	c.ImageCaching = &imageCaching
}

func (c *Container) GetLogging() *ContainerLoggingConfiguration {
	if c == nil {
		return nil
	}
	return c.Logging
}

func (c *Container) SetLogging(logging ContainerLoggingConfiguration) {
	c.Logging = &logging
}

func (c *Container) GetResources() *ContainerResourceRequirements {
	if c == nil {
		return nil
	}
	return c.Resources
}

func (c *Container) SetResources(resources ContainerResourceRequirements) {
	c.Resources = &resources
}

func (c *Container) GetSize() *int64 {
	if c == nil {
		return nil
	}
	return c.Size
}

func (c *Container) SetSize(size int64) {
	c.Size = &size
}

func (c Container) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: Container to string"
	}
	return string(jsonData)
}
