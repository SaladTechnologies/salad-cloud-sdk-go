package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.
type ContainerLoggingConfigurationHttp2 struct {
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
	Format *ContainerHttpLoggingConfigurationFormat2 `json:"format,omitempty" required:"true"`
	// Optional HTTP headers to include in log transmission requests
	Headers []shared.ContainerLoggingHttpHeader `json:"headers,omitempty" maxItems:"1000"`
	// The compression algorithm to apply to logs before transmission
	Compression *ContainerHttpLoggingConfigurationCompression2 `json:"compression,omitempty" required:"true"`
}

func (c *ContainerLoggingConfigurationHttp2) GetHost() *string {
	if c == nil {
		return nil
	}
	return c.Host
}

func (c *ContainerLoggingConfigurationHttp2) SetHost(host string) {
	c.Host = &host
}

func (c *ContainerLoggingConfigurationHttp2) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerLoggingConfigurationHttp2) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerLoggingConfigurationHttp2) GetUser() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.User
}

func (c *ContainerLoggingConfigurationHttp2) SetUser(user util.Nullable[string]) {
	c.User = &user
}

func (c *ContainerLoggingConfigurationHttp2) SetUserNull() {
	c.User = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerLoggingConfigurationHttp2) GetPassword() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Password
}

func (c *ContainerLoggingConfigurationHttp2) SetPassword(password util.Nullable[string]) {
	c.Password = &password
}

func (c *ContainerLoggingConfigurationHttp2) SetPasswordNull() {
	c.Password = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerLoggingConfigurationHttp2) GetPath() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Path
}

func (c *ContainerLoggingConfigurationHttp2) SetPath(path util.Nullable[string]) {
	c.Path = &path
}

func (c *ContainerLoggingConfigurationHttp2) SetPathNull() {
	c.Path = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerLoggingConfigurationHttp2) GetFormat() *ContainerHttpLoggingConfigurationFormat2 {
	if c == nil {
		return nil
	}
	return c.Format
}

func (c *ContainerLoggingConfigurationHttp2) SetFormat(format ContainerHttpLoggingConfigurationFormat2) {
	c.Format = &format
}

func (c *ContainerLoggingConfigurationHttp2) GetHeaders() []shared.ContainerLoggingHttpHeader {
	if c == nil {
		return nil
	}
	return c.Headers
}

func (c *ContainerLoggingConfigurationHttp2) SetHeaders(headers []shared.ContainerLoggingHttpHeader) {
	c.Headers = headers
}

func (c *ContainerLoggingConfigurationHttp2) GetCompression() *ContainerHttpLoggingConfigurationCompression2 {
	if c == nil {
		return nil
	}
	return c.Compression
}

func (c *ContainerLoggingConfigurationHttp2) SetCompression(compression ContainerHttpLoggingConfigurationCompression2) {
	c.Compression = &compression
}

func (c ContainerLoggingConfigurationHttp2) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerLoggingConfigurationHttp2 to string"
	}
	return string(jsonData)
}

// The format in which logs will be delivered
type ContainerHttpLoggingConfigurationFormat2 string

const (
	CONTAINER_HTTP_LOGGING_CONFIGURATION_FORMAT2_JSON       ContainerHttpLoggingConfigurationFormat2 = "json"
	CONTAINER_HTTP_LOGGING_CONFIGURATION_FORMAT2_JSON_LINES ContainerHttpLoggingConfigurationFormat2 = "json_lines"
)

// The compression algorithm to apply to logs before transmission
type ContainerHttpLoggingConfigurationCompression2 string

const (
	CONTAINER_HTTP_LOGGING_CONFIGURATION_COMPRESSION2_NONE ContainerHttpLoggingConfigurationCompression2 = "none"
	CONTAINER_HTTP_LOGGING_CONFIGURATION_COMPRESSION2_GZIP ContainerHttpLoggingConfigurationCompression2 = "gzip"
)
