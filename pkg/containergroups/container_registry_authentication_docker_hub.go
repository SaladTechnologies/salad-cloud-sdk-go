package containergroups

import "encoding/json"

// Authentication details for Docker Hub registry
type ContainerRegistryAuthenticationDockerHub struct {
	// Docker Hub username
	Username *string `json:"username,omitempty" required:"true" maxLength:"10000" minLength:"1" pattern:"^.*$"`
	// Docker Hub personal access token (PAT)
	PersonalAccessToken *string `json:"personal_access_token,omitempty" required:"true" maxLength:"10000" minLength:"1" pattern:"^.*$"`
}

func (c *ContainerRegistryAuthenticationDockerHub) GetUsername() *string {
	if c == nil {
		return nil
	}
	return c.Username
}

func (c *ContainerRegistryAuthenticationDockerHub) SetUsername(username string) {
	c.Username = &username
}

func (c *ContainerRegistryAuthenticationDockerHub) GetPersonalAccessToken() *string {
	if c == nil {
		return nil
	}
	return c.PersonalAccessToken
}

func (c *ContainerRegistryAuthenticationDockerHub) SetPersonalAccessToken(personalAccessToken string) {
	c.PersonalAccessToken = &personalAccessToken
}

func (c ContainerRegistryAuthenticationDockerHub) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerRegistryAuthenticationDockerHub to string"
	}
	return string(jsonData)
}
