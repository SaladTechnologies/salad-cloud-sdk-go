package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.
type ContainerHttpLoggingConfiguration struct {
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
	Format *Format `json:"format,omitempty" required:"true"`
	// Optional HTTP headers to include in log transmission requests
	Headers *util.Nullable[[]ContainerLoggingHttpHeader] `json:"headers,omitempty" required:"true" maxItems:"1000"`
	// The compression algorithm to apply to logs before transmission
	Compression *Compression `json:"compression,omitempty" required:"true"`
}

func (c *ContainerHttpLoggingConfiguration) GetHost() *string {
	if c == nil {
		return nil
	}
	return c.Host
}

func (c *ContainerHttpLoggingConfiguration) SetHost(host string) {
	c.Host = &host
}

func (c *ContainerHttpLoggingConfiguration) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerHttpLoggingConfiguration) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerHttpLoggingConfiguration) GetUser() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.User
}

func (c *ContainerHttpLoggingConfiguration) SetUser(user util.Nullable[string]) {
	c.User = &user
}

func (c *ContainerHttpLoggingConfiguration) SetUserNull() {
	c.User = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerHttpLoggingConfiguration) GetPassword() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Password
}

func (c *ContainerHttpLoggingConfiguration) SetPassword(password util.Nullable[string]) {
	c.Password = &password
}

func (c *ContainerHttpLoggingConfiguration) SetPasswordNull() {
	c.Password = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerHttpLoggingConfiguration) GetPath() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Path
}

func (c *ContainerHttpLoggingConfiguration) SetPath(path util.Nullable[string]) {
	c.Path = &path
}

func (c *ContainerHttpLoggingConfiguration) SetPathNull() {
	c.Path = &util.Nullable[string]{IsNull: true}
}

func (c *ContainerHttpLoggingConfiguration) GetFormat() *Format {
	if c == nil {
		return nil
	}
	return c.Format
}

func (c *ContainerHttpLoggingConfiguration) SetFormat(format Format) {
	c.Format = &format
}

func (c *ContainerHttpLoggingConfiguration) GetHeaders() *util.Nullable[[]ContainerLoggingHttpHeader] {
	if c == nil {
		return nil
	}
	return c.Headers
}

func (c *ContainerHttpLoggingConfiguration) SetHeaders(headers util.Nullable[[]ContainerLoggingHttpHeader]) {
	c.Headers = &headers
}

func (c *ContainerHttpLoggingConfiguration) SetHeadersNull() {
	c.Headers = &util.Nullable[[]ContainerLoggingHttpHeader]{IsNull: true}
}

func (c *ContainerHttpLoggingConfiguration) GetCompression() *Compression {
	if c == nil {
		return nil
	}
	return c.Compression
}

func (c *ContainerHttpLoggingConfiguration) SetCompression(compression Compression) {
	c.Compression = &compression
}

func (c ContainerHttpLoggingConfiguration) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerHttpLoggingConfiguration to string"
	}
	return string(jsonData)
}

// The format in which logs will be delivered
type Format string

const (
	FORMAT_JSON       Format = "json"
	FORMAT_JSON_LINES Format = "json_lines"
)

// The compression algorithm to apply to logs before transmission
type Compression string

const (
	COMPRESSION_NONE Compression = "none"
	COMPRESSION_GZIP Compression = "gzip"
)
