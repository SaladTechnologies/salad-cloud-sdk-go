package containergroups

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents an update container object
type UpdateContainer struct {
	Image     *string    `json:"image,omitempty" maxLength:"1024" minLength:"1"`
	Resources *Resources `json:"resources,omitempty"`
	// Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image.
	Command                []string                               `json:"command,omitempty" maxItems:"100"`
	Priority               *shared.ContainerGroupPriority         `json:"priority,omitempty"`
	EnvironmentVariables   map[string]*string                     `json:"environment_variables,omitempty"`
	Logging                *UpdateContainerLogging                `json:"logging,omitempty"`
	RegistryAuthentication *UpdateContainerRegistryAuthentication `json:"registry_authentication,omitempty"`
}

func (u *UpdateContainer) SetImage(image string) {
	u.Image = &image
}

func (u *UpdateContainer) GetImage() *string {
	if u == nil {
		return nil
	}
	return u.Image
}

func (u *UpdateContainer) SetResources(resources Resources) {
	u.Resources = &resources
}

func (u *UpdateContainer) GetResources() *Resources {
	if u == nil {
		return nil
	}
	return u.Resources
}

func (u *UpdateContainer) SetCommand(command []string) {
	u.Command = command
}

func (u *UpdateContainer) GetCommand() []string {
	if u == nil {
		return nil
	}
	return u.Command
}

func (u *UpdateContainer) SetPriority(priority shared.ContainerGroupPriority) {
	u.Priority = &priority
}

func (u *UpdateContainer) GetPriority() *shared.ContainerGroupPriority {
	if u == nil {
		return nil
	}
	return u.Priority
}

func (u *UpdateContainer) SetEnvironmentVariables(environmentVariables map[string]*string) {
	u.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (u *UpdateContainer) GetEnvironmentVariables() map[string]*string {
	if u == nil {
		return nil
	}
	return u.EnvironmentVariables
}

func (u *UpdateContainer) SetLogging(logging UpdateContainerLogging) {
	u.Logging = &logging
}

func (u *UpdateContainer) GetLogging() *UpdateContainerLogging {
	if u == nil {
		return nil
	}
	return u.Logging
}

func (u *UpdateContainer) SetRegistryAuthentication(registryAuthentication UpdateContainerRegistryAuthentication) {
	u.RegistryAuthentication = &registryAuthentication
}

func (u *UpdateContainer) GetRegistryAuthentication() *UpdateContainerRegistryAuthentication {
	if u == nil {
		return nil
	}
	return u.RegistryAuthentication
}

type Resources struct {
	Cpu           *int64   `json:"cpu,omitempty" min:"1" max:"16"`
	Memory        *int64   `json:"memory,omitempty" min:"1024" max:"30720"`
	GpuClasses    []string `json:"gpu_classes,omitempty"`
	StorageAmount *int64   `json:"storage_amount,omitempty" min:"1073741824" max:"53687091200"`
}

func (r *Resources) SetCpu(cpu int64) {
	r.Cpu = &cpu
}

func (r *Resources) GetCpu() *int64 {
	if r == nil {
		return nil
	}
	return r.Cpu
}

func (r *Resources) SetMemory(memory int64) {
	r.Memory = &memory
}

func (r *Resources) GetMemory() *int64 {
	if r == nil {
		return nil
	}
	return r.Memory
}

func (r *Resources) SetGpuClasses(gpuClasses []string) {
	r.GpuClasses = gpuClasses
}

func (r *Resources) GetGpuClasses() []string {
	if r == nil {
		return nil
	}
	return r.GpuClasses
}

func (r *Resources) SetStorageAmount(storageAmount int64) {
	r.StorageAmount = &storageAmount
}

func (r *Resources) GetStorageAmount() *int64 {
	if r == nil {
		return nil
	}
	return r.StorageAmount
}

type UpdateContainerLogging struct {
	Axiom    *LoggingAxiom3    `json:"axiom,omitempty"`
	Datadog  *LoggingDatadog3  `json:"datadog,omitempty"`
	NewRelic *LoggingNewRelic3 `json:"new_relic,omitempty"`
	Splunk   *LoggingSplunk3   `json:"splunk,omitempty"`
	Tcp      *LoggingTcp3      `json:"tcp,omitempty"`
	Http     *LoggingHttp3     `json:"http,omitempty"`
}

func (u *UpdateContainerLogging) SetAxiom(axiom LoggingAxiom3) {
	u.Axiom = &axiom
}

func (u *UpdateContainerLogging) GetAxiom() *LoggingAxiom3 {
	if u == nil {
		return nil
	}
	return u.Axiom
}

func (u *UpdateContainerLogging) SetDatadog(datadog LoggingDatadog3) {
	u.Datadog = &datadog
}

func (u *UpdateContainerLogging) GetDatadog() *LoggingDatadog3 {
	if u == nil {
		return nil
	}
	return u.Datadog
}

func (u *UpdateContainerLogging) SetNewRelic(newRelic LoggingNewRelic3) {
	u.NewRelic = &newRelic
}

func (u *UpdateContainerLogging) GetNewRelic() *LoggingNewRelic3 {
	if u == nil {
		return nil
	}
	return u.NewRelic
}

func (u *UpdateContainerLogging) SetSplunk(splunk LoggingSplunk3) {
	u.Splunk = &splunk
}

func (u *UpdateContainerLogging) GetSplunk() *LoggingSplunk3 {
	if u == nil {
		return nil
	}
	return u.Splunk
}

func (u *UpdateContainerLogging) SetTcp(tcp LoggingTcp3) {
	u.Tcp = &tcp
}

func (u *UpdateContainerLogging) GetTcp() *LoggingTcp3 {
	if u == nil {
		return nil
	}
	return u.Tcp
}

func (u *UpdateContainerLogging) SetHttp(http LoggingHttp3) {
	u.Http = &http
}

func (u *UpdateContainerLogging) GetHttp() *LoggingHttp3 {
	if u == nil {
		return nil
	}
	return u.Http
}

type LoggingAxiom3 struct {
	Host     *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiToken *string `json:"api_token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Dataset  *string `json:"dataset,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingAxiom3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingAxiom3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingAxiom3) SetApiToken(apiToken string) {
	l.ApiToken = &apiToken
}

func (l *LoggingAxiom3) GetApiToken() *string {
	if l == nil {
		return nil
	}
	return l.ApiToken
}

func (l *LoggingAxiom3) SetDataset(dataset string) {
	l.Dataset = &dataset
}

func (l *LoggingAxiom3) GetDataset() *string {
	if l == nil {
		return nil
	}
	return l.Dataset
}

type LoggingDatadog3 struct {
	Host   *string        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiKey *string        `json:"api_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Tags   []DatadogTags3 `json:"tags,omitempty" maxItems:"1000"`
}

func (l *LoggingDatadog3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingDatadog3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingDatadog3) SetApiKey(apiKey string) {
	l.ApiKey = &apiKey
}

func (l *LoggingDatadog3) GetApiKey() *string {
	if l == nil {
		return nil
	}
	return l.ApiKey
}

func (l *LoggingDatadog3) SetTags(tags []DatadogTags3) {
	l.Tags = tags
}

func (l *LoggingDatadog3) GetTags() []DatadogTags3 {
	if l == nil {
		return nil
	}
	return l.Tags
}

type DatadogTags3 struct {
	Name  *string `json:"name,omitempty" required:"true"`
	Value *string `json:"value,omitempty" required:"true"`
}

func (d *DatadogTags3) SetName(name string) {
	d.Name = &name
}

func (d *DatadogTags3) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DatadogTags3) SetValue(value string) {
	d.Value = &value
}

func (d *DatadogTags3) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

type LoggingNewRelic3 struct {
	Host         *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	IngestionKey *string `json:"ingestion_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingNewRelic3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingNewRelic3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingNewRelic3) SetIngestionKey(ingestionKey string) {
	l.IngestionKey = &ingestionKey
}

func (l *LoggingNewRelic3) GetIngestionKey() *string {
	if l == nil {
		return nil
	}
	return l.IngestionKey
}

type LoggingSplunk3 struct {
	Host  *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Token *string `json:"token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingSplunk3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingSplunk3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingSplunk3) SetToken(token string) {
	l.Token = &token
}

func (l *LoggingSplunk3) GetToken() *string {
	if l == nil {
		return nil
	}
	return l.Token
}

type LoggingTcp3 struct {
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
}

func (l *LoggingTcp3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingTcp3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingTcp3) SetPort(port int64) {
	l.Port = &port
}

func (l *LoggingTcp3) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

type LoggingHttp3 struct {
	Host        *string           `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port        *int64            `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	User        *string           `json:"user,omitempty"`
	Password    *string           `json:"password,omitempty"`
	Path        *string           `json:"path,omitempty"`
	Format      *HttpFormat3      `json:"format,omitempty" required:"true"`
	Headers     []HttpHeaders4    `json:"headers,omitempty" maxItems:"1000"`
	Compression *HttpCompression3 `json:"compression,omitempty" required:"true"`
}

func (l *LoggingHttp3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingHttp3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingHttp3) SetPort(port int64) {
	l.Port = &port
}

func (l *LoggingHttp3) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingHttp3) SetUser(user string) {
	l.User = &user
}

func (l *LoggingHttp3) GetUser() *string {
	if l == nil {
		return nil
	}
	return l.User
}

func (l *LoggingHttp3) SetPassword(password string) {
	l.Password = &password
}

func (l *LoggingHttp3) GetPassword() *string {
	if l == nil {
		return nil
	}
	return l.Password
}

func (l *LoggingHttp3) SetPath(path string) {
	l.Path = &path
}

func (l *LoggingHttp3) GetPath() *string {
	if l == nil {
		return nil
	}
	return l.Path
}

func (l *LoggingHttp3) SetFormat(format HttpFormat3) {
	l.Format = &format
}

func (l *LoggingHttp3) GetFormat() *HttpFormat3 {
	if l == nil {
		return nil
	}
	return l.Format
}

func (l *LoggingHttp3) SetHeaders(headers []HttpHeaders4) {
	l.Headers = headers
}

func (l *LoggingHttp3) GetHeaders() []HttpHeaders4 {
	if l == nil {
		return nil
	}
	return l.Headers
}

func (l *LoggingHttp3) SetCompression(compression HttpCompression3) {
	l.Compression = &compression
}

func (l *LoggingHttp3) GetCompression() *HttpCompression3 {
	if l == nil {
		return nil
	}
	return l.Compression
}

type HttpFormat3 string

const (
	HTTP_FORMAT3_JSON       HttpFormat3 = "json"
	HTTP_FORMAT3_JSON_LINES HttpFormat3 = "json_lines"
)

type HttpHeaders4 struct {
	Name  *string `json:"name,omitempty" required:"true"`
	Value *string `json:"value,omitempty" required:"true"`
}

func (h *HttpHeaders4) SetName(name string) {
	h.Name = &name
}

func (h *HttpHeaders4) GetName() *string {
	if h == nil {
		return nil
	}
	return h.Name
}

func (h *HttpHeaders4) SetValue(value string) {
	h.Value = &value
}

func (h *HttpHeaders4) GetValue() *string {
	if h == nil {
		return nil
	}
	return h.Value
}

type HttpCompression3 string

const (
	HTTP_COMPRESSION3_NONE HttpCompression3 = "none"
	HTTP_COMPRESSION3_GZIP HttpCompression3 = "gzip"
)

type UpdateContainerRegistryAuthentication struct {
	Basic     *RegistryAuthenticationBasic2     `json:"basic,omitempty"`
	GcpGcr    *RegistryAuthenticationGcpGcr2    `json:"gcp_gcr,omitempty"`
	AwsEcr    *RegistryAuthenticationAwsEcr2    `json:"aws_ecr,omitempty"`
	DockerHub *RegistryAuthenticationDockerHub2 `json:"docker_hub,omitempty"`
	GcpGar    *RegistryAuthenticationGcpGar2    `json:"gcp_gar,omitempty"`
}

func (u *UpdateContainerRegistryAuthentication) SetBasic(basic RegistryAuthenticationBasic2) {
	u.Basic = &basic
}

func (u *UpdateContainerRegistryAuthentication) GetBasic() *RegistryAuthenticationBasic2 {
	if u == nil {
		return nil
	}
	return u.Basic
}

func (u *UpdateContainerRegistryAuthentication) SetGcpGcr(gcpGcr RegistryAuthenticationGcpGcr2) {
	u.GcpGcr = &gcpGcr
}

func (u *UpdateContainerRegistryAuthentication) GetGcpGcr() *RegistryAuthenticationGcpGcr2 {
	if u == nil {
		return nil
	}
	return u.GcpGcr
}

func (u *UpdateContainerRegistryAuthentication) SetAwsEcr(awsEcr RegistryAuthenticationAwsEcr2) {
	u.AwsEcr = &awsEcr
}

func (u *UpdateContainerRegistryAuthentication) GetAwsEcr() *RegistryAuthenticationAwsEcr2 {
	if u == nil {
		return nil
	}
	return u.AwsEcr
}

func (u *UpdateContainerRegistryAuthentication) SetDockerHub(dockerHub RegistryAuthenticationDockerHub2) {
	u.DockerHub = &dockerHub
}

func (u *UpdateContainerRegistryAuthentication) GetDockerHub() *RegistryAuthenticationDockerHub2 {
	if u == nil {
		return nil
	}
	return u.DockerHub
}

func (u *UpdateContainerRegistryAuthentication) SetGcpGar(gcpGar RegistryAuthenticationGcpGar2) {
	u.GcpGar = &gcpGar
}

func (u *UpdateContainerRegistryAuthentication) GetGcpGar() *RegistryAuthenticationGcpGar2 {
	if u == nil {
		return nil
	}
	return u.GcpGar
}

type RegistryAuthenticationBasic2 struct {
	Username *string `json:"username,omitempty" required:"true"`
	Password *string `json:"password,omitempty" required:"true"`
}

func (r *RegistryAuthenticationBasic2) SetUsername(username string) {
	r.Username = &username
}

func (r *RegistryAuthenticationBasic2) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationBasic2) SetPassword(password string) {
	r.Password = &password
}

func (r *RegistryAuthenticationBasic2) GetPassword() *string {
	if r == nil {
		return nil
	}
	return r.Password
}

type RegistryAuthenticationGcpGcr2 struct {
	ServiceKey *string `json:"service_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationGcpGcr2) SetServiceKey(serviceKey string) {
	r.ServiceKey = &serviceKey
}

func (r *RegistryAuthenticationGcpGcr2) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}

type RegistryAuthenticationAwsEcr2 struct {
	AccessKeyId     *string `json:"access_key_id,omitempty" required:"true"`
	SecretAccessKey *string `json:"secret_access_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationAwsEcr2) SetAccessKeyId(accessKeyId string) {
	r.AccessKeyId = &accessKeyId
}

func (r *RegistryAuthenticationAwsEcr2) GetAccessKeyId() *string {
	if r == nil {
		return nil
	}
	return r.AccessKeyId
}

func (r *RegistryAuthenticationAwsEcr2) SetSecretAccessKey(secretAccessKey string) {
	r.SecretAccessKey = &secretAccessKey
}

func (r *RegistryAuthenticationAwsEcr2) GetSecretAccessKey() *string {
	if r == nil {
		return nil
	}
	return r.SecretAccessKey
}

type RegistryAuthenticationDockerHub2 struct {
	Username            *string `json:"username,omitempty" required:"true"`
	PersonalAccessToken *string `json:"personal_access_token,omitempty" required:"true"`
}

func (r *RegistryAuthenticationDockerHub2) SetUsername(username string) {
	r.Username = &username
}

func (r *RegistryAuthenticationDockerHub2) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationDockerHub2) SetPersonalAccessToken(personalAccessToken string) {
	r.PersonalAccessToken = &personalAccessToken
}

func (r *RegistryAuthenticationDockerHub2) GetPersonalAccessToken() *string {
	if r == nil {
		return nil
	}
	return r.PersonalAccessToken
}

type RegistryAuthenticationGcpGar2 struct {
	ServiceKey *string `json:"service_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationGcpGar2) SetServiceKey(serviceKey string) {
	r.ServiceKey = &serviceKey
}

func (r *RegistryAuthenticationGcpGar2) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}
