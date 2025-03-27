package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.
type ContainerLoggingConfigurationHttp1 struct {
	// The hostname or IP address of the HTTP logging endpoint
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// The port number of the HTTP logging endpoint (1-65535)
	Port *int64 `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	// Optional username for HTTP authentication
	User *util.Nullable[string] `json:"user,omitempty" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// Optional password for HTTP authentication
	Password *util.Nullable[string] `json:"password,omitempty" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// Optional URL path for the HTTP endpoint
	Path *util.Nullable[string] `json:"path,omitempty" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// The format in which logs will be delivered
	Format *ContainerLoggingHttpFormat `json:"format,omitempty" required:"true"`
	// Optional HTTP headers to include in log transmission requests
	Headers     *util.Nullable[[]ContainerLoggingHttpHeader] `json:"headers,omitempty" required:"true" maxItems:"1000"`
	Compression any                                          `json:"compression,omitempty" required:"true"`
}

func (c *ContainerLoggingConfigurationHttp1) GetHost() *string {
	if c == nil {
		return nil
	}
	return c.Host
}

func (c *ContainerLoggingConfigurationHttp1) SetHost(host string) {
	c.Host = &host
}

func (c *ContainerLoggingConfigurationHttp1) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerLoggingConfigurationHttp1) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerLoggingConfigurationHttp1) GetUser() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.User
}

func (c *ContainerLoggingConfigurationHttp1) SetUser(user util.Nullable[string]) {
	c.User = &user
}

func (c *ContainerLoggingConfigurationHttp1) SetUserNull() {
	c.User = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerLoggingConfigurationHttp1) GetPassword() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Password
}

func (c *ContainerLoggingConfigurationHttp1) SetPassword(password util.Nullable[string]) {
	c.Password = &password
}

func (c *ContainerLoggingConfigurationHttp1) SetPasswordNull() {
	c.Password = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerLoggingConfigurationHttp1) GetPath() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Path
}

func (c *ContainerLoggingConfigurationHttp1) SetPath(path util.Nullable[string]) {
	c.Path = &path
}

func (c *ContainerLoggingConfigurationHttp1) SetPathNull() {
	c.Path = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerLoggingConfigurationHttp1) GetFormat() *ContainerLoggingHttpFormat {
	if c == nil {
		return nil
	}
	return c.Format
}

func (c *ContainerLoggingConfigurationHttp1) SetFormat(format ContainerLoggingHttpFormat) {
	c.Format = &format
}

func (c *ContainerLoggingConfigurationHttp1) GetHeaders() *util.Nullable[[]ContainerLoggingHttpHeader] {
	if c == nil {
		return nil
	}
	return c.Headers
}

func (c *ContainerLoggingConfigurationHttp1) SetHeaders(headers util.Nullable[[]ContainerLoggingHttpHeader]) {
	c.Headers = &headers
}

func (c *ContainerLoggingConfigurationHttp1) SetHeadersNull() {
	c.Headers = &util.Nullable[[]ContainerLoggingHttpHeader]{IsNull: true}
}

func (c *ContainerLoggingConfigurationHttp1) GetCompression() any {
	if c == nil {
		return nil
	}
	return c.Compression
}

func (c *ContainerLoggingConfigurationHttp1) SetCompression(compression any) {
	c.Compression = compression
}

func (c ContainerLoggingConfigurationHttp1) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerLoggingConfigurationHttp1 to string"
	}
	return string(jsonData)
}
