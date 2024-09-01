package shared

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
)

// Represents a container
type Container struct {
	Image *string `json:"image,omitempty" required:"true" maxLength:"1024" minLength:"1"`
	// Represents a container resource requirements
	Resources            *ContainerResourceRequirements `json:"resources,omitempty" required:"true"`
	Command              []string                       `json:"command,omitempty" required:"true" maxItems:"100"`
	Priority             *ContainerGroupPriority        `json:"priority,omitempty"`
	Size                 *int64                         `json:"size,omitempty"`
	Hash                 *string                        `json:"hash,omitempty"`
	EnvironmentVariables map[string]*string             `json:"environment_variables,omitempty"`
	Logging              *ContainerLogging              `json:"logging,omitempty"`
}

func (c *Container) SetImage(image string) {
	c.Image = &image
}

func (c *Container) GetImage() *string {
	if c == nil {
		return nil
	}
	return c.Image
}

func (c *Container) SetResources(resources ContainerResourceRequirements) {
	c.Resources = &resources
}

func (c *Container) GetResources() *ContainerResourceRequirements {
	if c == nil {
		return nil
	}
	return c.Resources
}

func (c *Container) SetCommand(command []string) {
	c.Command = command
}

func (c *Container) GetCommand() []string {
	if c == nil {
		return nil
	}
	return c.Command
}

func (c *Container) SetPriority(priority ContainerGroupPriority) {
	c.Priority = &priority
}

func (c *Container) GetPriority() *ContainerGroupPriority {
	if c == nil {
		return nil
	}
	return c.Priority
}

func (c *Container) SetSize(size int64) {
	c.Size = &size
}

func (c *Container) GetSize() *int64 {
	if c == nil {
		return nil
	}
	return c.Size
}

func (c *Container) SetHash(hash string) {
	c.Hash = &hash
}

func (c *Container) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *Container) SetEnvironmentVariables(environmentVariables map[string]*string) {
	c.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (c *Container) GetEnvironmentVariables() map[string]*string {
	if c == nil {
		return nil
	}
	return c.EnvironmentVariables
}

func (c *Container) SetLogging(logging ContainerLogging) {
	c.Logging = &logging
}

func (c *Container) GetLogging() *ContainerLogging {
	if c == nil {
		return nil
	}
	return c.Logging
}

type ContainerLogging struct {
	Axiom    *LoggingAxiom1    `json:"axiom,omitempty"`
	Datadog  *LoggingDatadog1  `json:"datadog,omitempty"`
	NewRelic *LoggingNewRelic1 `json:"new_relic,omitempty"`
	Splunk   *LoggingSplunk1   `json:"splunk,omitempty"`
	Tcp      *LoggingTcp1      `json:"tcp,omitempty"`
	Http     *LoggingHttp1     `json:"http,omitempty"`
}

func (c *ContainerLogging) SetAxiom(axiom LoggingAxiom1) {
	c.Axiom = &axiom
}

func (c *ContainerLogging) GetAxiom() *LoggingAxiom1 {
	if c == nil {
		return nil
	}
	return c.Axiom
}

func (c *ContainerLogging) SetDatadog(datadog LoggingDatadog1) {
	c.Datadog = &datadog
}

func (c *ContainerLogging) GetDatadog() *LoggingDatadog1 {
	if c == nil {
		return nil
	}
	return c.Datadog
}

func (c *ContainerLogging) SetNewRelic(newRelic LoggingNewRelic1) {
	c.NewRelic = &newRelic
}

func (c *ContainerLogging) GetNewRelic() *LoggingNewRelic1 {
	if c == nil {
		return nil
	}
	return c.NewRelic
}

func (c *ContainerLogging) SetSplunk(splunk LoggingSplunk1) {
	c.Splunk = &splunk
}

func (c *ContainerLogging) GetSplunk() *LoggingSplunk1 {
	if c == nil {
		return nil
	}
	return c.Splunk
}

func (c *ContainerLogging) SetTcp(tcp LoggingTcp1) {
	c.Tcp = &tcp
}

func (c *ContainerLogging) GetTcp() *LoggingTcp1 {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerLogging) SetHttp(http LoggingHttp1) {
	c.Http = &http
}

func (c *ContainerLogging) GetHttp() *LoggingHttp1 {
	if c == nil {
		return nil
	}
	return c.Http
}

type LoggingAxiom1 struct {
	Host     *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiToken *string `json:"api_token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Dataset  *string `json:"dataset,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingAxiom1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingAxiom1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingAxiom1) SetApiToken(apiToken string) {
	l.ApiToken = &apiToken
}

func (l *LoggingAxiom1) GetApiToken() *string {
	if l == nil {
		return nil
	}
	return l.ApiToken
}

func (l *LoggingAxiom1) SetDataset(dataset string) {
	l.Dataset = &dataset
}

func (l *LoggingAxiom1) GetDataset() *string {
	if l == nil {
		return nil
	}
	return l.Dataset
}

type LoggingDatadog1 struct {
	Host   *string        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiKey *string        `json:"api_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Tags   []DatadogTags1 `json:"tags,omitempty" maxItems:"1000"`
}

func (l *LoggingDatadog1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingDatadog1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingDatadog1) SetApiKey(apiKey string) {
	l.ApiKey = &apiKey
}

func (l *LoggingDatadog1) GetApiKey() *string {
	if l == nil {
		return nil
	}
	return l.ApiKey
}

func (l *LoggingDatadog1) SetTags(tags []DatadogTags1) {
	l.Tags = tags
}

func (l *LoggingDatadog1) GetTags() []DatadogTags1 {
	if l == nil {
		return nil
	}
	return l.Tags
}

type DatadogTags1 struct {
	Name  *string `json:"name,omitempty" required:"true"`
	Value *string `json:"value,omitempty" required:"true"`
}

func (d *DatadogTags1) SetName(name string) {
	d.Name = &name
}

func (d *DatadogTags1) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DatadogTags1) SetValue(value string) {
	d.Value = &value
}

func (d *DatadogTags1) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

type LoggingNewRelic1 struct {
	Host         *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	IngestionKey *string `json:"ingestion_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingNewRelic1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingNewRelic1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingNewRelic1) SetIngestionKey(ingestionKey string) {
	l.IngestionKey = &ingestionKey
}

func (l *LoggingNewRelic1) GetIngestionKey() *string {
	if l == nil {
		return nil
	}
	return l.IngestionKey
}

type LoggingSplunk1 struct {
	Host  *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Token *string `json:"token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingSplunk1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingSplunk1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingSplunk1) SetToken(token string) {
	l.Token = &token
}

func (l *LoggingSplunk1) GetToken() *string {
	if l == nil {
		return nil
	}
	return l.Token
}

type LoggingTcp1 struct {
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
}

func (l *LoggingTcp1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingTcp1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingTcp1) SetPort(port int64) {
	l.Port = &port
}

func (l *LoggingTcp1) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

type LoggingHttp1 struct {
	Host        *string           `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port        *int64            `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	User        *string           `json:"user,omitempty"`
	Password    *string           `json:"password,omitempty"`
	Path        *string           `json:"path,omitempty"`
	Format      *HttpFormat1      `json:"format,omitempty" required:"true"`
	Headers     []HttpHeaders1    `json:"headers,omitempty" maxItems:"1000"`
	Compression *HttpCompression1 `json:"compression,omitempty" required:"true"`
}

func (l *LoggingHttp1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingHttp1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingHttp1) SetPort(port int64) {
	l.Port = &port
}

func (l *LoggingHttp1) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingHttp1) SetUser(user string) {
	l.User = &user
}

func (l *LoggingHttp1) GetUser() *string {
	if l == nil {
		return nil
	}
	return l.User
}

func (l *LoggingHttp1) SetPassword(password string) {
	l.Password = &password
}

func (l *LoggingHttp1) GetPassword() *string {
	if l == nil {
		return nil
	}
	return l.Password
}

func (l *LoggingHttp1) SetPath(path string) {
	l.Path = &path
}

func (l *LoggingHttp1) GetPath() *string {
	if l == nil {
		return nil
	}
	return l.Path
}

func (l *LoggingHttp1) SetFormat(format HttpFormat1) {
	l.Format = &format
}

func (l *LoggingHttp1) GetFormat() *HttpFormat1 {
	if l == nil {
		return nil
	}
	return l.Format
}

func (l *LoggingHttp1) SetHeaders(headers []HttpHeaders1) {
	l.Headers = headers
}

func (l *LoggingHttp1) GetHeaders() []HttpHeaders1 {
	if l == nil {
		return nil
	}
	return l.Headers
}

func (l *LoggingHttp1) SetCompression(compression HttpCompression1) {
	l.Compression = &compression
}

func (l *LoggingHttp1) GetCompression() *HttpCompression1 {
	if l == nil {
		return nil
	}
	return l.Compression
}

type HttpFormat1 string

const (
	HTTP_FORMAT1_JSON       HttpFormat1 = "json"
	HTTP_FORMAT1_JSON_LINES HttpFormat1 = "json_lines"
)

type HttpHeaders1 struct {
	Name  *string `json:"name,omitempty" required:"true"`
	Value *string `json:"value,omitempty" required:"true"`
}

func (h *HttpHeaders1) SetName(name string) {
	h.Name = &name
}

func (h *HttpHeaders1) GetName() *string {
	if h == nil {
		return nil
	}
	return h.Name
}

func (h *HttpHeaders1) SetValue(value string) {
	h.Value = &value
}

func (h *HttpHeaders1) GetValue() *string {
	if h == nil {
		return nil
	}
	return h.Value
}

type HttpCompression1 string

const (
	HTTP_COMPRESSION1_NONE HttpCompression1 = "none"
	HTTP_COMPRESSION1_GZIP HttpCompression1 = "gzip"
)
