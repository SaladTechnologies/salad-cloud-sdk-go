package containergroups

import "encoding/json"

// Authentication details for AWS Elastic Container Registry (ECR)
type ContainerRegistryAuthenticationAwsEcr struct {
	// AWS access key ID used for ECR authentication
	AccessKeyId *string `json:"access_key_id,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// AWS secret access key used for ECR authentication
	SecretAccessKey *string `json:"secret_access_key,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
}

func (c *ContainerRegistryAuthenticationAwsEcr) GetAccessKeyId() *string {
	if c == nil {
		return nil
	}
	return c.AccessKeyId
}

func (c *ContainerRegistryAuthenticationAwsEcr) SetAccessKeyId(accessKeyId string) {
	c.AccessKeyId = &accessKeyId
}

func (c *ContainerRegistryAuthenticationAwsEcr) GetSecretAccessKey() *string {
	if c == nil {
		return nil
	}
	return c.SecretAccessKey
}

func (c *ContainerRegistryAuthenticationAwsEcr) SetSecretAccessKey(secretAccessKey string) {
	c.SecretAccessKey = &secretAccessKey
}

func (c ContainerRegistryAuthenticationAwsEcr) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerRegistryAuthenticationAwsEcr to string"
	}
	return string(jsonData)
}
