package containergroups

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents a container
type CreateContainer struct {
	Image *string `json:"image,omitempty" required:"true" maxLength:"1024" minLength:"1"`
	// Represents a container resource requirements
	Resources *shared.ContainerResourceRequirements `json:"resources,omitempty" required:"true"`
	// Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image.
	Command                []string                               `json:"command,omitempty" maxItems:"100"`
	Priority               *shared.ContainerGroupPriority         `json:"priority,omitempty"`
	EnvironmentVariables   map[string]*string                     `json:"environment_variables,omitempty"`
	Logging                *CreateContainerLogging                `json:"logging,omitempty"`
	RegistryAuthentication *CreateContainerRegistryAuthentication `json:"registry_authentication,omitempty"`
}

func (c *CreateContainer) SetImage(image string) {
	c.Image = &image
}

func (c *CreateContainer) GetImage() *string {
	if c == nil {
		return nil
	}
	return c.Image
}

func (c *CreateContainer) SetResources(resources shared.ContainerResourceRequirements) {
	c.Resources = &resources
}

func (c *CreateContainer) GetResources() *shared.ContainerResourceRequirements {
	if c == nil {
		return nil
	}
	return c.Resources
}

func (c *CreateContainer) SetCommand(command []string) {
	c.Command = command
}

func (c *CreateContainer) GetCommand() []string {
	if c == nil {
		return nil
	}
	return c.Command
}

func (c *CreateContainer) SetPriority(priority shared.ContainerGroupPriority) {
	c.Priority = &priority
}

func (c *CreateContainer) GetPriority() *shared.ContainerGroupPriority {
	if c == nil {
		return nil
	}
	return c.Priority
}

func (c *CreateContainer) SetEnvironmentVariables(environmentVariables map[string]*string) {
	c.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (c *CreateContainer) GetEnvironmentVariables() map[string]*string {
	if c == nil {
		return nil
	}
	return c.EnvironmentVariables
}

func (c *CreateContainer) SetLogging(logging CreateContainerLogging) {
	c.Logging = &logging
}

func (c *CreateContainer) GetLogging() *CreateContainerLogging {
	if c == nil {
		return nil
	}
	return c.Logging
}

func (c *CreateContainer) SetRegistryAuthentication(registryAuthentication CreateContainerRegistryAuthentication) {
	c.RegistryAuthentication = &registryAuthentication
}

func (c *CreateContainer) GetRegistryAuthentication() *CreateContainerRegistryAuthentication {
	if c == nil {
		return nil
	}
	return c.RegistryAuthentication
}

type CreateContainerLogging struct {
	Axiom    *LoggingAxiom2    `json:"axiom,omitempty"`
	Datadog  *LoggingDatadog2  `json:"datadog,omitempty"`
	NewRelic *LoggingNewRelic2 `json:"new_relic,omitempty"`
	Splunk   *LoggingSplunk2   `json:"splunk,omitempty"`
	Tcp      *LoggingTcp2      `json:"tcp,omitempty"`
	Http     *LoggingHttp2     `json:"http,omitempty"`
}

func (c *CreateContainerLogging) SetAxiom(axiom LoggingAxiom2) {
	c.Axiom = &axiom
}

func (c *CreateContainerLogging) GetAxiom() *LoggingAxiom2 {
	if c == nil {
		return nil
	}
	return c.Axiom
}

func (c *CreateContainerLogging) SetDatadog(datadog LoggingDatadog2) {
	c.Datadog = &datadog
}

func (c *CreateContainerLogging) GetDatadog() *LoggingDatadog2 {
	if c == nil {
		return nil
	}
	return c.Datadog
}

func (c *CreateContainerLogging) SetNewRelic(newRelic LoggingNewRelic2) {
	c.NewRelic = &newRelic
}

func (c *CreateContainerLogging) GetNewRelic() *LoggingNewRelic2 {
	if c == nil {
		return nil
	}
	return c.NewRelic
}

func (c *CreateContainerLogging) SetSplunk(splunk LoggingSplunk2) {
	c.Splunk = &splunk
}

func (c *CreateContainerLogging) GetSplunk() *LoggingSplunk2 {
	if c == nil {
		return nil
	}
	return c.Splunk
}

func (c *CreateContainerLogging) SetTcp(tcp LoggingTcp2) {
	c.Tcp = &tcp
}

func (c *CreateContainerLogging) GetTcp() *LoggingTcp2 {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *CreateContainerLogging) SetHttp(http LoggingHttp2) {
	c.Http = &http
}

func (c *CreateContainerLogging) GetHttp() *LoggingHttp2 {
	if c == nil {
		return nil
	}
	return c.Http
}

type LoggingAxiom2 struct {
	Host     *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiToken *string `json:"api_token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Dataset  *string `json:"dataset,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingAxiom2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingAxiom2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingAxiom2) SetApiToken(apiToken string) {
	l.ApiToken = &apiToken
}

func (l *LoggingAxiom2) GetApiToken() *string {
	if l == nil {
		return nil
	}
	return l.ApiToken
}

func (l *LoggingAxiom2) SetDataset(dataset string) {
	l.Dataset = &dataset
}

func (l *LoggingAxiom2) GetDataset() *string {
	if l == nil {
		return nil
	}
	return l.Dataset
}

type LoggingDatadog2 struct {
	Host   *string        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiKey *string        `json:"api_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Tags   []DatadogTags2 `json:"tags,omitempty" maxItems:"1000"`
}

func (l *LoggingDatadog2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingDatadog2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingDatadog2) SetApiKey(apiKey string) {
	l.ApiKey = &apiKey
}

func (l *LoggingDatadog2) GetApiKey() *string {
	if l == nil {
		return nil
	}
	return l.ApiKey
}

func (l *LoggingDatadog2) SetTags(tags []DatadogTags2) {
	l.Tags = tags
}

func (l *LoggingDatadog2) GetTags() []DatadogTags2 {
	if l == nil {
		return nil
	}
	return l.Tags
}

type DatadogTags2 struct {
	Name  *string `json:"name,omitempty" required:"true"`
	Value *string `json:"value,omitempty" required:"true"`
}

func (d *DatadogTags2) SetName(name string) {
	d.Name = &name
}

func (d *DatadogTags2) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DatadogTags2) SetValue(value string) {
	d.Value = &value
}

func (d *DatadogTags2) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

type LoggingNewRelic2 struct {
	Host         *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	IngestionKey *string `json:"ingestion_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingNewRelic2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingNewRelic2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingNewRelic2) SetIngestionKey(ingestionKey string) {
	l.IngestionKey = &ingestionKey
}

func (l *LoggingNewRelic2) GetIngestionKey() *string {
	if l == nil {
		return nil
	}
	return l.IngestionKey
}

type LoggingSplunk2 struct {
	Host  *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Token *string `json:"token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingSplunk2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingSplunk2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingSplunk2) SetToken(token string) {
	l.Token = &token
}

func (l *LoggingSplunk2) GetToken() *string {
	if l == nil {
		return nil
	}
	return l.Token
}

type LoggingTcp2 struct {
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
}

func (l *LoggingTcp2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingTcp2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingTcp2) SetPort(port int64) {
	l.Port = &port
}

func (l *LoggingTcp2) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

type LoggingHttp2 struct {
	Host        *string           `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port        *int64            `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	User        *string           `json:"user,omitempty"`
	Password    *string           `json:"password,omitempty"`
	Path        *string           `json:"path,omitempty"`
	Format      *HttpFormat2      `json:"format,omitempty" required:"true"`
	Headers     []HttpHeaders3    `json:"headers,omitempty" maxItems:"1000"`
	Compression *HttpCompression2 `json:"compression,omitempty" required:"true"`
}

func (l *LoggingHttp2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingHttp2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingHttp2) SetPort(port int64) {
	l.Port = &port
}

func (l *LoggingHttp2) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingHttp2) SetUser(user string) {
	l.User = &user
}

func (l *LoggingHttp2) GetUser() *string {
	if l == nil {
		return nil
	}
	return l.User
}

func (l *LoggingHttp2) SetPassword(password string) {
	l.Password = &password
}

func (l *LoggingHttp2) GetPassword() *string {
	if l == nil {
		return nil
	}
	return l.Password
}

func (l *LoggingHttp2) SetPath(path string) {
	l.Path = &path
}

func (l *LoggingHttp2) GetPath() *string {
	if l == nil {
		return nil
	}
	return l.Path
}

func (l *LoggingHttp2) SetFormat(format HttpFormat2) {
	l.Format = &format
}

func (l *LoggingHttp2) GetFormat() *HttpFormat2 {
	if l == nil {
		return nil
	}
	return l.Format
}

func (l *LoggingHttp2) SetHeaders(headers []HttpHeaders3) {
	l.Headers = headers
}

func (l *LoggingHttp2) GetHeaders() []HttpHeaders3 {
	if l == nil {
		return nil
	}
	return l.Headers
}

func (l *LoggingHttp2) SetCompression(compression HttpCompression2) {
	l.Compression = &compression
}

func (l *LoggingHttp2) GetCompression() *HttpCompression2 {
	if l == nil {
		return nil
	}
	return l.Compression
}

type HttpFormat2 string

const (
	HTTP_FORMAT2_JSON       HttpFormat2 = "json"
	HTTP_FORMAT2_JSON_LINES HttpFormat2 = "json_lines"
)

type HttpHeaders3 struct {
	Name  *string `json:"name,omitempty" required:"true"`
	Value *string `json:"value,omitempty" required:"true"`
}

func (h *HttpHeaders3) SetName(name string) {
	h.Name = &name
}

func (h *HttpHeaders3) GetName() *string {
	if h == nil {
		return nil
	}
	return h.Name
}

func (h *HttpHeaders3) SetValue(value string) {
	h.Value = &value
}

func (h *HttpHeaders3) GetValue() *string {
	if h == nil {
		return nil
	}
	return h.Value
}

type HttpCompression2 string

const (
	HTTP_COMPRESSION2_NONE HttpCompression2 = "none"
	HTTP_COMPRESSION2_GZIP HttpCompression2 = "gzip"
)

type CreateContainerRegistryAuthentication struct {
	Basic     *RegistryAuthenticationBasic1     `json:"basic,omitempty"`
	GcpGcr    *RegistryAuthenticationGcpGcr1    `json:"gcp_gcr,omitempty"`
	AwsEcr    *RegistryAuthenticationAwsEcr1    `json:"aws_ecr,omitempty"`
	DockerHub *RegistryAuthenticationDockerHub1 `json:"docker_hub,omitempty"`
	GcpGar    *RegistryAuthenticationGcpGar1    `json:"gcp_gar,omitempty"`
}

func (c *CreateContainerRegistryAuthentication) SetBasic(basic RegistryAuthenticationBasic1) {
	c.Basic = &basic
}

func (c *CreateContainerRegistryAuthentication) GetBasic() *RegistryAuthenticationBasic1 {
	if c == nil {
		return nil
	}
	return c.Basic
}

func (c *CreateContainerRegistryAuthentication) SetGcpGcr(gcpGcr RegistryAuthenticationGcpGcr1) {
	c.GcpGcr = &gcpGcr
}

func (c *CreateContainerRegistryAuthentication) GetGcpGcr() *RegistryAuthenticationGcpGcr1 {
	if c == nil {
		return nil
	}
	return c.GcpGcr
}

func (c *CreateContainerRegistryAuthentication) SetAwsEcr(awsEcr RegistryAuthenticationAwsEcr1) {
	c.AwsEcr = &awsEcr
}

func (c *CreateContainerRegistryAuthentication) GetAwsEcr() *RegistryAuthenticationAwsEcr1 {
	if c == nil {
		return nil
	}
	return c.AwsEcr
}

func (c *CreateContainerRegistryAuthentication) SetDockerHub(dockerHub RegistryAuthenticationDockerHub1) {
	c.DockerHub = &dockerHub
}

func (c *CreateContainerRegistryAuthentication) GetDockerHub() *RegistryAuthenticationDockerHub1 {
	if c == nil {
		return nil
	}
	return c.DockerHub
}

func (c *CreateContainerRegistryAuthentication) SetGcpGar(gcpGar RegistryAuthenticationGcpGar1) {
	c.GcpGar = &gcpGar
}

func (c *CreateContainerRegistryAuthentication) GetGcpGar() *RegistryAuthenticationGcpGar1 {
	if c == nil {
		return nil
	}
	return c.GcpGar
}

type RegistryAuthenticationBasic1 struct {
	Username *string `json:"username,omitempty" required:"true"`
	Password *string `json:"password,omitempty" required:"true"`
}

func (r *RegistryAuthenticationBasic1) SetUsername(username string) {
	r.Username = &username
}

func (r *RegistryAuthenticationBasic1) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationBasic1) SetPassword(password string) {
	r.Password = &password
}

func (r *RegistryAuthenticationBasic1) GetPassword() *string {
	if r == nil {
		return nil
	}
	return r.Password
}

type RegistryAuthenticationGcpGcr1 struct {
	ServiceKey *string `json:"service_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationGcpGcr1) SetServiceKey(serviceKey string) {
	r.ServiceKey = &serviceKey
}

func (r *RegistryAuthenticationGcpGcr1) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}

type RegistryAuthenticationAwsEcr1 struct {
	AccessKeyId     *string `json:"access_key_id,omitempty" required:"true"`
	SecretAccessKey *string `json:"secret_access_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationAwsEcr1) SetAccessKeyId(accessKeyId string) {
	r.AccessKeyId = &accessKeyId
}

func (r *RegistryAuthenticationAwsEcr1) GetAccessKeyId() *string {
	if r == nil {
		return nil
	}
	return r.AccessKeyId
}

func (r *RegistryAuthenticationAwsEcr1) SetSecretAccessKey(secretAccessKey string) {
	r.SecretAccessKey = &secretAccessKey
}

func (r *RegistryAuthenticationAwsEcr1) GetSecretAccessKey() *string {
	if r == nil {
		return nil
	}
	return r.SecretAccessKey
}

type RegistryAuthenticationDockerHub1 struct {
	Username            *string `json:"username,omitempty" required:"true"`
	PersonalAccessToken *string `json:"personal_access_token,omitempty" required:"true"`
}

func (r *RegistryAuthenticationDockerHub1) SetUsername(username string) {
	r.Username = &username
}

func (r *RegistryAuthenticationDockerHub1) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationDockerHub1) SetPersonalAccessToken(personalAccessToken string) {
	r.PersonalAccessToken = &personalAccessToken
}

func (r *RegistryAuthenticationDockerHub1) GetPersonalAccessToken() *string {
	if r == nil {
		return nil
	}
	return r.PersonalAccessToken
}

type RegistryAuthenticationGcpGar1 struct {
	ServiceKey *string `json:"service_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationGcpGar1) SetServiceKey(serviceKey string) {
	r.ServiceKey = &serviceKey
}

func (r *RegistryAuthenticationGcpGar1) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}
