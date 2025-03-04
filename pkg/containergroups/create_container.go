package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a container
type CreateContainer struct {
	Image *string `json:"image,omitempty" required:"true" maxLength:"1024" minLength:"1"`
	// Represents a container resource requirements
	Resources *shared.ContainerResourceRequirements `json:"resources,omitempty" required:"true"`
	// Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image.
	Command                *util.Nullable[[]string]                              `json:"command,omitempty" maxItems:"100"`
	Priority               *util.Nullable[shared.ContainerGroupPriority]         `json:"priority,omitempty"`
	EnvironmentVariables   map[string]*string                                    `json:"environment_variables,omitempty"`
	Logging                *util.Nullable[CreateContainerLogging]                `json:"logging,omitempty"`
	RegistryAuthentication *util.Nullable[CreateContainerRegistryAuthentication] `json:"registry_authentication,omitempty"`
	ImageCaching           *bool                                                 `json:"image_caching,omitempty"`
}

func (c *CreateContainer) GetImage() *string {
	if c == nil {
		return nil
	}
	return c.Image
}

func (c *CreateContainer) SetImage(image string) {
	c.Image = &image
}

func (c *CreateContainer) GetResources() *shared.ContainerResourceRequirements {
	if c == nil {
		return nil
	}
	return c.Resources
}

func (c *CreateContainer) SetResources(resources shared.ContainerResourceRequirements) {
	c.Resources = &resources
}

func (c *CreateContainer) GetCommand() *util.Nullable[[]string] {
	if c == nil {
		return nil
	}
	return c.Command
}

func (c *CreateContainer) SetCommand(command util.Nullable[[]string]) {
	c.Command = &command
}

func (c *CreateContainer) SetCommandNull() {
	c.Command = &util.Nullable[[]string]{IsNull: true}
}

func (c *CreateContainer) GetPriority() *util.Nullable[shared.ContainerGroupPriority] {
	if c == nil {
		return nil
	}
	return c.Priority
}

func (c *CreateContainer) SetPriority(priority util.Nullable[shared.ContainerGroupPriority]) {
	c.Priority = &priority
}

func (c *CreateContainer) SetPriorityNull() {
	c.Priority = &util.Nullable[shared.ContainerGroupPriority]{IsNull: true}
}

func (c *CreateContainer) GetEnvironmentVariables() map[string]*string {
	if c == nil {
		return nil
	}
	return c.EnvironmentVariables
}

func (c *CreateContainer) SetEnvironmentVariables(environmentVariables map[string]*string) {
	c.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (c *CreateContainer) GetLogging() *util.Nullable[CreateContainerLogging] {
	if c == nil {
		return nil
	}
	return c.Logging
}

func (c *CreateContainer) SetLogging(logging util.Nullable[CreateContainerLogging]) {
	c.Logging = &logging
}

func (c *CreateContainer) SetLoggingNull() {
	c.Logging = &util.Nullable[CreateContainerLogging]{IsNull: true}
}

func (c *CreateContainer) GetRegistryAuthentication() *util.Nullable[CreateContainerRegistryAuthentication] {
	if c == nil {
		return nil
	}
	return c.RegistryAuthentication
}

func (c *CreateContainer) SetRegistryAuthentication(registryAuthentication util.Nullable[CreateContainerRegistryAuthentication]) {
	c.RegistryAuthentication = &registryAuthentication
}

func (c *CreateContainer) SetRegistryAuthenticationNull() {
	c.RegistryAuthentication = &util.Nullable[CreateContainerRegistryAuthentication]{IsNull: true}
}

func (c *CreateContainer) GetImageCaching() *bool {
	if c == nil {
		return nil
	}
	return c.ImageCaching
}

func (c *CreateContainer) SetImageCaching(imageCaching bool) {
	c.ImageCaching = &imageCaching
}

func (c CreateContainer) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateContainer to string"
	}
	return string(jsonData)
}

type CreateContainerLogging struct {
	Axiom    *util.Nullable[LoggingAxiom2]    `json:"axiom,omitempty"`
	Datadog  *util.Nullable[LoggingDatadog2]  `json:"datadog,omitempty"`
	NewRelic *util.Nullable[LoggingNewRelic2] `json:"new_relic,omitempty"`
	Splunk   *util.Nullable[LoggingSplunk2]   `json:"splunk,omitempty"`
	Tcp      *util.Nullable[LoggingTcp2]      `json:"tcp,omitempty"`
	Http     *util.Nullable[LoggingHttp2]     `json:"http,omitempty"`
}

func (c *CreateContainerLogging) GetAxiom() *util.Nullable[LoggingAxiom2] {
	if c == nil {
		return nil
	}
	return c.Axiom
}

func (c *CreateContainerLogging) SetAxiom(axiom util.Nullable[LoggingAxiom2]) {
	c.Axiom = &axiom
}

func (c *CreateContainerLogging) SetAxiomNull() {
	c.Axiom = &util.Nullable[LoggingAxiom2]{IsNull: true}
}

func (c *CreateContainerLogging) GetDatadog() *util.Nullable[LoggingDatadog2] {
	if c == nil {
		return nil
	}
	return c.Datadog
}

func (c *CreateContainerLogging) SetDatadog(datadog util.Nullable[LoggingDatadog2]) {
	c.Datadog = &datadog
}

func (c *CreateContainerLogging) SetDatadogNull() {
	c.Datadog = &util.Nullable[LoggingDatadog2]{IsNull: true}
}

func (c *CreateContainerLogging) GetNewRelic() *util.Nullable[LoggingNewRelic2] {
	if c == nil {
		return nil
	}
	return c.NewRelic
}

func (c *CreateContainerLogging) SetNewRelic(newRelic util.Nullable[LoggingNewRelic2]) {
	c.NewRelic = &newRelic
}

func (c *CreateContainerLogging) SetNewRelicNull() {
	c.NewRelic = &util.Nullable[LoggingNewRelic2]{IsNull: true}
}

func (c *CreateContainerLogging) GetSplunk() *util.Nullable[LoggingSplunk2] {
	if c == nil {
		return nil
	}
	return c.Splunk
}

func (c *CreateContainerLogging) SetSplunk(splunk util.Nullable[LoggingSplunk2]) {
	c.Splunk = &splunk
}

func (c *CreateContainerLogging) SetSplunkNull() {
	c.Splunk = &util.Nullable[LoggingSplunk2]{IsNull: true}
}

func (c *CreateContainerLogging) GetTcp() *util.Nullable[LoggingTcp2] {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *CreateContainerLogging) SetTcp(tcp util.Nullable[LoggingTcp2]) {
	c.Tcp = &tcp
}

func (c *CreateContainerLogging) SetTcpNull() {
	c.Tcp = &util.Nullable[LoggingTcp2]{IsNull: true}
}

func (c *CreateContainerLogging) GetHttp() *util.Nullable[LoggingHttp2] {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *CreateContainerLogging) SetHttp(http util.Nullable[LoggingHttp2]) {
	c.Http = &http
}

func (c *CreateContainerLogging) SetHttpNull() {
	c.Http = &util.Nullable[LoggingHttp2]{IsNull: true}
}

func (c CreateContainerLogging) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateContainerLogging to string"
	}
	return string(jsonData)
}

type LoggingAxiom2 struct {
	Host     *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiToken *string `json:"api_token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Dataset  *string `json:"dataset,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingAxiom2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingAxiom2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingAxiom2) GetApiToken() *string {
	if l == nil {
		return nil
	}
	return l.ApiToken
}

func (l *LoggingAxiom2) SetApiToken(apiToken string) {
	l.ApiToken = &apiToken
}

func (l *LoggingAxiom2) GetDataset() *string {
	if l == nil {
		return nil
	}
	return l.Dataset
}

func (l *LoggingAxiom2) SetDataset(dataset string) {
	l.Dataset = &dataset
}

func (l LoggingAxiom2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingAxiom2 to string"
	}
	return string(jsonData)
}

type LoggingDatadog2 struct {
	Host   *string                        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiKey *string                        `json:"api_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Tags   *util.Nullable[[]DatadogTags2] `json:"tags,omitempty" maxItems:"1000"`
}

func (l *LoggingDatadog2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingDatadog2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingDatadog2) GetApiKey() *string {
	if l == nil {
		return nil
	}
	return l.ApiKey
}

func (l *LoggingDatadog2) SetApiKey(apiKey string) {
	l.ApiKey = &apiKey
}

func (l *LoggingDatadog2) GetTags() *util.Nullable[[]DatadogTags2] {
	if l == nil {
		return nil
	}
	return l.Tags
}

func (l *LoggingDatadog2) SetTags(tags util.Nullable[[]DatadogTags2]) {
	l.Tags = &tags
}

func (l *LoggingDatadog2) SetTagsNull() {
	l.Tags = &util.Nullable[[]DatadogTags2]{IsNull: true}
}

func (l LoggingDatadog2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingDatadog2 to string"
	}
	return string(jsonData)
}

type DatadogTags2 struct {
	Name  *string `json:"name,omitempty" required:"true"`
	Value *string `json:"value,omitempty" required:"true"`
}

func (d *DatadogTags2) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DatadogTags2) SetName(name string) {
	d.Name = &name
}

func (d *DatadogTags2) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

func (d *DatadogTags2) SetValue(value string) {
	d.Value = &value
}

func (d DatadogTags2) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DatadogTags2 to string"
	}
	return string(jsonData)
}

type LoggingNewRelic2 struct {
	Host         *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	IngestionKey *string `json:"ingestion_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingNewRelic2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingNewRelic2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingNewRelic2) GetIngestionKey() *string {
	if l == nil {
		return nil
	}
	return l.IngestionKey
}

func (l *LoggingNewRelic2) SetIngestionKey(ingestionKey string) {
	l.IngestionKey = &ingestionKey
}

func (l LoggingNewRelic2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingNewRelic2 to string"
	}
	return string(jsonData)
}

type LoggingSplunk2 struct {
	Host  *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Token *string `json:"token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingSplunk2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingSplunk2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingSplunk2) GetToken() *string {
	if l == nil {
		return nil
	}
	return l.Token
}

func (l *LoggingSplunk2) SetToken(token string) {
	l.Token = &token
}

func (l LoggingSplunk2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingSplunk2 to string"
	}
	return string(jsonData)
}

type LoggingTcp2 struct {
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
}

func (l *LoggingTcp2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingTcp2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingTcp2) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingTcp2) SetPort(port int64) {
	l.Port = &port
}

func (l LoggingTcp2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingTcp2 to string"
	}
	return string(jsonData)
}

type LoggingHttp2 struct {
	Host        *string                        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port        *int64                         `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	User        *util.Nullable[string]         `json:"user,omitempty"`
	Password    *util.Nullable[string]         `json:"password,omitempty"`
	Path        *util.Nullable[string]         `json:"path,omitempty"`
	Format      *HttpFormat2                   `json:"format,omitempty" required:"true"`
	Headers     *util.Nullable[[]HttpHeaders3] `json:"headers,omitempty" maxItems:"1000"`
	Compression *HttpCompression2              `json:"compression,omitempty" required:"true"`
}

func (l *LoggingHttp2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingHttp2) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingHttp2) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingHttp2) SetPort(port int64) {
	l.Port = &port
}

func (l *LoggingHttp2) GetUser() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.User
}

func (l *LoggingHttp2) SetUser(user util.Nullable[string]) {
	l.User = &user
}

func (l *LoggingHttp2) SetUserNull() {
	l.User = &util.Nullable[string]{IsNull: true}
}

func (l *LoggingHttp2) GetPassword() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.Password
}

func (l *LoggingHttp2) SetPassword(password util.Nullable[string]) {
	l.Password = &password
}

func (l *LoggingHttp2) SetPasswordNull() {
	l.Password = &util.Nullable[string]{IsNull: true}
}

func (l *LoggingHttp2) GetPath() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.Path
}

func (l *LoggingHttp2) SetPath(path util.Nullable[string]) {
	l.Path = &path
}

func (l *LoggingHttp2) SetPathNull() {
	l.Path = &util.Nullable[string]{IsNull: true}
}

func (l *LoggingHttp2) GetFormat() *HttpFormat2 {
	if l == nil {
		return nil
	}
	return l.Format
}

func (l *LoggingHttp2) SetFormat(format HttpFormat2) {
	l.Format = &format
}

func (l *LoggingHttp2) GetHeaders() *util.Nullable[[]HttpHeaders3] {
	if l == nil {
		return nil
	}
	return l.Headers
}

func (l *LoggingHttp2) SetHeaders(headers util.Nullable[[]HttpHeaders3]) {
	l.Headers = &headers
}

func (l *LoggingHttp2) SetHeadersNull() {
	l.Headers = &util.Nullable[[]HttpHeaders3]{IsNull: true}
}

func (l *LoggingHttp2) GetCompression() *HttpCompression2 {
	if l == nil {
		return nil
	}
	return l.Compression
}

func (l *LoggingHttp2) SetCompression(compression HttpCompression2) {
	l.Compression = &compression
}

func (l LoggingHttp2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingHttp2 to string"
	}
	return string(jsonData)
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

func (h *HttpHeaders3) GetName() *string {
	if h == nil {
		return nil
	}
	return h.Name
}

func (h *HttpHeaders3) SetName(name string) {
	h.Name = &name
}

func (h *HttpHeaders3) GetValue() *string {
	if h == nil {
		return nil
	}
	return h.Value
}

func (h *HttpHeaders3) SetValue(value string) {
	h.Value = &value
}

func (h HttpHeaders3) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: HttpHeaders3 to string"
	}
	return string(jsonData)
}

type HttpCompression2 string

const (
	HTTP_COMPRESSION2_NONE HttpCompression2 = "none"
	HTTP_COMPRESSION2_GZIP HttpCompression2 = "gzip"
)

type CreateContainerRegistryAuthentication struct {
	Basic     *util.Nullable[RegistryAuthenticationBasic1]     `json:"basic,omitempty"`
	GcpGcr    *util.Nullable[RegistryAuthenticationGcpGcr1]    `json:"gcp_gcr,omitempty"`
	AwsEcr    *util.Nullable[RegistryAuthenticationAwsEcr1]    `json:"aws_ecr,omitempty"`
	DockerHub *util.Nullable[RegistryAuthenticationDockerHub1] `json:"docker_hub,omitempty"`
	GcpGar    *util.Nullable[RegistryAuthenticationGcpGar1]    `json:"gcp_gar,omitempty"`
}

func (c *CreateContainerRegistryAuthentication) GetBasic() *util.Nullable[RegistryAuthenticationBasic1] {
	if c == nil {
		return nil
	}
	return c.Basic
}

func (c *CreateContainerRegistryAuthentication) SetBasic(basic util.Nullable[RegistryAuthenticationBasic1]) {
	c.Basic = &basic
}

func (c *CreateContainerRegistryAuthentication) SetBasicNull() {
	c.Basic = &util.Nullable[RegistryAuthenticationBasic1]{IsNull: true}
}

func (c *CreateContainerRegistryAuthentication) GetGcpGcr() *util.Nullable[RegistryAuthenticationGcpGcr1] {
	if c == nil {
		return nil
	}
	return c.GcpGcr
}

func (c *CreateContainerRegistryAuthentication) SetGcpGcr(gcpGcr util.Nullable[RegistryAuthenticationGcpGcr1]) {
	c.GcpGcr = &gcpGcr
}

func (c *CreateContainerRegistryAuthentication) SetGcpGcrNull() {
	c.GcpGcr = &util.Nullable[RegistryAuthenticationGcpGcr1]{IsNull: true}
}

func (c *CreateContainerRegistryAuthentication) GetAwsEcr() *util.Nullable[RegistryAuthenticationAwsEcr1] {
	if c == nil {
		return nil
	}
	return c.AwsEcr
}

func (c *CreateContainerRegistryAuthentication) SetAwsEcr(awsEcr util.Nullable[RegistryAuthenticationAwsEcr1]) {
	c.AwsEcr = &awsEcr
}

func (c *CreateContainerRegistryAuthentication) SetAwsEcrNull() {
	c.AwsEcr = &util.Nullable[RegistryAuthenticationAwsEcr1]{IsNull: true}
}

func (c *CreateContainerRegistryAuthentication) GetDockerHub() *util.Nullable[RegistryAuthenticationDockerHub1] {
	if c == nil {
		return nil
	}
	return c.DockerHub
}

func (c *CreateContainerRegistryAuthentication) SetDockerHub(dockerHub util.Nullable[RegistryAuthenticationDockerHub1]) {
	c.DockerHub = &dockerHub
}

func (c *CreateContainerRegistryAuthentication) SetDockerHubNull() {
	c.DockerHub = &util.Nullable[RegistryAuthenticationDockerHub1]{IsNull: true}
}

func (c *CreateContainerRegistryAuthentication) GetGcpGar() *util.Nullable[RegistryAuthenticationGcpGar1] {
	if c == nil {
		return nil
	}
	return c.GcpGar
}

func (c *CreateContainerRegistryAuthentication) SetGcpGar(gcpGar util.Nullable[RegistryAuthenticationGcpGar1]) {
	c.GcpGar = &gcpGar
}

func (c *CreateContainerRegistryAuthentication) SetGcpGarNull() {
	c.GcpGar = &util.Nullable[RegistryAuthenticationGcpGar1]{IsNull: true}
}

func (c CreateContainerRegistryAuthentication) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateContainerRegistryAuthentication to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationBasic1 struct {
	Username *string `json:"username,omitempty" required:"true"`
	Password *string `json:"password,omitempty" required:"true"`
}

func (r *RegistryAuthenticationBasic1) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationBasic1) SetUsername(username string) {
	r.Username = &username
}

func (r *RegistryAuthenticationBasic1) GetPassword() *string {
	if r == nil {
		return nil
	}
	return r.Password
}

func (r *RegistryAuthenticationBasic1) SetPassword(password string) {
	r.Password = &password
}

func (r RegistryAuthenticationBasic1) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationBasic1 to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationGcpGcr1 struct {
	ServiceKey *string `json:"service_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationGcpGcr1) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}

func (r *RegistryAuthenticationGcpGcr1) SetServiceKey(serviceKey string) {
	r.ServiceKey = &serviceKey
}

func (r RegistryAuthenticationGcpGcr1) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationGcpGcr1 to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationAwsEcr1 struct {
	AccessKeyId     *string `json:"access_key_id,omitempty" required:"true"`
	SecretAccessKey *string `json:"secret_access_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationAwsEcr1) GetAccessKeyId() *string {
	if r == nil {
		return nil
	}
	return r.AccessKeyId
}

func (r *RegistryAuthenticationAwsEcr1) SetAccessKeyId(accessKeyId string) {
	r.AccessKeyId = &accessKeyId
}

func (r *RegistryAuthenticationAwsEcr1) GetSecretAccessKey() *string {
	if r == nil {
		return nil
	}
	return r.SecretAccessKey
}

func (r *RegistryAuthenticationAwsEcr1) SetSecretAccessKey(secretAccessKey string) {
	r.SecretAccessKey = &secretAccessKey
}

func (r RegistryAuthenticationAwsEcr1) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationAwsEcr1 to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationDockerHub1 struct {
	Username            *string `json:"username,omitempty" required:"true"`
	PersonalAccessToken *string `json:"personal_access_token,omitempty" required:"true"`
}

func (r *RegistryAuthenticationDockerHub1) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationDockerHub1) SetUsername(username string) {
	r.Username = &username
}

func (r *RegistryAuthenticationDockerHub1) GetPersonalAccessToken() *string {
	if r == nil {
		return nil
	}
	return r.PersonalAccessToken
}

func (r *RegistryAuthenticationDockerHub1) SetPersonalAccessToken(personalAccessToken string) {
	r.PersonalAccessToken = &personalAccessToken
}

func (r RegistryAuthenticationDockerHub1) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationDockerHub1 to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationGcpGar1 struct {
	ServiceKey *string `json:"service_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationGcpGar1) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}

func (r *RegistryAuthenticationGcpGar1) SetServiceKey(serviceKey string) {
	r.ServiceKey = &serviceKey
}

func (r RegistryAuthenticationGcpGar1) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationGcpGar1 to string"
	}
	return string(jsonData)
}
