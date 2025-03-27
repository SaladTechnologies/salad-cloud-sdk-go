package containergroups

import "encoding/json"

// Basic username and password authentication for generic container registries
type ContainerRegistryAuthenticationBasic struct {
	// Username for registry authentication
	Username *string `json:"username,omitempty" required:"true" maxLength:"10000" minLength:"1" pattern:"^.*$"`
	// Password for registry authentication
	Password *string `json:"password,omitempty" required:"true" maxLength:"10000" minLength:"1" pattern:"^.*$"`
}

func (c *ContainerRegistryAuthenticationBasic) GetUsername() *string {
	if c == nil {
		return nil
	}
	return c.Username
}

func (c *ContainerRegistryAuthenticationBasic) SetUsername(username string) {
	c.Username = &username
}

func (c *ContainerRegistryAuthenticationBasic) GetPassword() *string {
	if c == nil {
		return nil
	}
	return c.Password
}

func (c *ContainerRegistryAuthenticationBasic) SetPassword(password string) {
	c.Password = &password
}

func (c ContainerRegistryAuthenticationBasic) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerRegistryAuthenticationBasic to string"
	}
	return string(jsonData)
}
