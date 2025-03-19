package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Authentication configuration for various container registry types, including AWS ECR, Docker Hub, GCP GAR, GCP GCR, and basic authentication.
type ContainerRegistryAuthentication struct {
	// Authentication details for AWS Elastic Container Registry (ECR)
	AwsEcr *util.Nullable[ContainerRegistryAuthenticationAwsEcr] `json:"aws_ecr,omitempty"`
	// Basic username and password authentication for generic container registries
	Basic *ContainerRegistryAuthenticationBasic `json:"basic,omitempty"`
	// Authentication details for Docker Hub registry
	DockerHub *ContainerRegistryAuthenticationDockerHub `json:"docker_hub,omitempty"`
	// Authentication details for Google Artifact Registry (GAR)
	GcpGar *ContainerRegistryAuthenticationGcpGar `json:"gcp_gar,omitempty"`
	// Authentication details for Google Container Registry (GCR)
	GcpGcr *ContainerRegistryAuthenticationGcpGcr `json:"gcp_gcr,omitempty"`
}

func (c *ContainerRegistryAuthentication) GetAwsEcr() *util.Nullable[ContainerRegistryAuthenticationAwsEcr] {
	if c == nil {
		return nil
	}
	return c.AwsEcr
}

func (c *ContainerRegistryAuthentication) SetAwsEcr(awsEcr util.Nullable[ContainerRegistryAuthenticationAwsEcr]) {
	c.AwsEcr = &awsEcr
}

func (c *ContainerRegistryAuthentication) SetAwsEcrNull() {
	c.AwsEcr = &util.Nullable[ContainerRegistryAuthenticationAwsEcr]{IsNull: true}
}

func (c *ContainerRegistryAuthentication) GetBasic() *ContainerRegistryAuthenticationBasic {
	if c == nil {
		return nil
	}
	return c.Basic
}

func (c *ContainerRegistryAuthentication) SetBasic(basic ContainerRegistryAuthenticationBasic) {
	c.Basic = &basic
}

func (c *ContainerRegistryAuthentication) GetDockerHub() *ContainerRegistryAuthenticationDockerHub {
	if c == nil {
		return nil
	}
	return c.DockerHub
}

func (c *ContainerRegistryAuthentication) SetDockerHub(dockerHub ContainerRegistryAuthenticationDockerHub) {
	c.DockerHub = &dockerHub
}

func (c *ContainerRegistryAuthentication) GetGcpGar() *ContainerRegistryAuthenticationGcpGar {
	if c == nil {
		return nil
	}
	return c.GcpGar
}

func (c *ContainerRegistryAuthentication) SetGcpGar(gcpGar ContainerRegistryAuthenticationGcpGar) {
	c.GcpGar = &gcpGar
}

func (c *ContainerRegistryAuthentication) GetGcpGcr() *ContainerRegistryAuthenticationGcpGcr {
	if c == nil {
		return nil
	}
	return c.GcpGcr
}

func (c *ContainerRegistryAuthentication) SetGcpGcr(gcpGcr ContainerRegistryAuthenticationGcpGcr) {
	c.GcpGcr = &gcpGcr
}

func (c ContainerRegistryAuthentication) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerRegistryAuthentication to string"
	}
	return string(jsonData)
}
