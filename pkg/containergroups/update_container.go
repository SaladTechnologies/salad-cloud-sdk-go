package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents an update container object
type UpdateContainer struct {
	// Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image.
	Command *util.Nullable[[]string] `json:"command,omitempty" maxItems:"100"`
	// Environment variables to set in the container.
	EnvironmentVariables map[string]*string `json:"environment_variables,omitempty"`
	// The container image to use.
	Image *util.Nullable[string] `json:"image,omitempty" maxLength:"1024" minLength:"1" pattern:"^.*$"`
	// The container image caching.
	ImageCaching *bool `json:"image_caching,omitempty"`
	// Configuration options for directing container logs to a logging provider. This schema enables you to specify a single logging destination for container output, supporting monitoring, debugging, and analytics use cases. Each provider has its own configuration parameters defined in the referenced schemas. Only one logging provider can be selected at a time.
	Logging *util.Nullable[UpdateContainerLogging] `json:"logging,omitempty"`
	// Specifies the priority level for container group execution, which determines resource allocation and scheduling precedence.
	Priority *util.Nullable[shared.ContainerGroupPriority] `json:"priority,omitempty"`
	// Authentication configuration for various container registry types, including AWS ECR, Docker Hub, GCP GAR, GCP GCR, and basic authentication.
	RegistryAuthentication *ContainerRegistryAuthentication `json:"registry_authentication,omitempty"`
	// Defines the resource specifications that can be modified for a container group, including CPU, memory, GPU classes, and storage allocations.
	Resources *util.Nullable[ContainerResourceUpdateSchema] `json:"resources,omitempty"`
}

func (u *UpdateContainer) GetCommand() *util.Nullable[[]string] {
	if u == nil {
		return nil
	}
	return u.Command
}

func (u *UpdateContainer) SetCommand(command util.Nullable[[]string]) {
	u.Command = &command
}

func (u *UpdateContainer) SetCommandNull() {
	u.Command = &util.Nullable[[]string]{IsNull: true}
}

func (u *UpdateContainer) GetEnvironmentVariables() map[string]*string {
	if u == nil {
		return nil
	}
	return u.EnvironmentVariables
}

func (u *UpdateContainer) SetEnvironmentVariables(environmentVariables map[string]*string) {
	u.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (u *UpdateContainer) GetImage() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.Image
}

func (u *UpdateContainer) SetImage(image util.Nullable[string]) {
	u.Image = &image
}

func (u *UpdateContainer) SetImageNull() {
	u.Image = &util.Nullable[string]{IsNull: true}
}

func (u *UpdateContainer) GetImageCaching() *bool {
	if u == nil {
		return nil
	}
	return u.ImageCaching
}

func (u *UpdateContainer) SetImageCaching(imageCaching bool) {
	u.ImageCaching = &imageCaching
}

func (u *UpdateContainer) GetLogging() *util.Nullable[UpdateContainerLogging] {
	if u == nil {
		return nil
	}
	return u.Logging
}

func (u *UpdateContainer) SetLogging(logging util.Nullable[UpdateContainerLogging]) {
	u.Logging = &logging
}

func (u *UpdateContainer) SetLoggingNull() {
	u.Logging = &util.Nullable[UpdateContainerLogging]{IsNull: true}
}

func (u *UpdateContainer) GetPriority() *util.Nullable[shared.ContainerGroupPriority] {
	if u == nil {
		return nil
	}
	return u.Priority
}

func (u *UpdateContainer) SetPriority(priority util.Nullable[shared.ContainerGroupPriority]) {
	u.Priority = &priority
}

func (u *UpdateContainer) SetPriorityNull() {
	u.Priority = &util.Nullable[shared.ContainerGroupPriority]{IsNull: true}
}

func (u *UpdateContainer) GetRegistryAuthentication() *ContainerRegistryAuthentication {
	if u == nil {
		return nil
	}
	return u.RegistryAuthentication
}

func (u *UpdateContainer) SetRegistryAuthentication(registryAuthentication ContainerRegistryAuthentication) {
	u.RegistryAuthentication = &registryAuthentication
}

func (u *UpdateContainer) GetResources() *util.Nullable[ContainerResourceUpdateSchema] {
	if u == nil {
		return nil
	}
	return u.Resources
}

func (u *UpdateContainer) SetResources(resources util.Nullable[ContainerResourceUpdateSchema]) {
	u.Resources = &resources
}

func (u *UpdateContainer) SetResourcesNull() {
	u.Resources = &util.Nullable[ContainerResourceUpdateSchema]{IsNull: true}
}

func (u UpdateContainer) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateContainer to string"
	}
	return string(jsonData)
}
