package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents an update container object
type UpdateContainer struct {
	Image     *util.Nullable[string]    `json:"image,omitempty" maxLength:"1024" minLength:"1"`
	Resources *util.Nullable[Resources] `json:"resources,omitempty"`
	// Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image.
	Command                *util.Nullable[[]string]                              `json:"command,omitempty" maxItems:"100"`
	Priority               *util.Nullable[shared.ContainerGroupPriority]         `json:"priority,omitempty"`
	EnvironmentVariables   map[string]*string                                    `json:"environment_variables,omitempty"`
	Logging                *util.Nullable[UpdateContainerLogging]                `json:"logging,omitempty"`
	RegistryAuthentication *util.Nullable[UpdateContainerRegistryAuthentication] `json:"registry_authentication,omitempty"`
	ImageCaching           *bool                                                 `json:"image_caching,omitempty"`
}

func (u *UpdateContainer) GetImage() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.Image
}

func (u *UpdateContainer) SetImage(image util.Nullable[string]) {
	u.Image = &image
}

func (u *UpdateContainer) SetImageNull() {
	u.Image = &util.Nullable[string]{IsNull: true}
}

func (u *UpdateContainer) GetResources() *util.Nullable[Resources] {
	if u == nil {
		return nil
	}
	return u.Resources
}

func (u *UpdateContainer) SetResources(resources util.Nullable[Resources]) {
	u.Resources = &resources
}

func (u *UpdateContainer) SetResourcesNull() {
	u.Resources = &util.Nullable[Resources]{IsNull: true}
}

func (u *UpdateContainer) GetCommand() *util.Nullable[[]string] {
	if u == nil {
		return nil
	}
	return u.Command
}

func (u *UpdateContainer) SetCommand(command util.Nullable[[]string]) {
	u.Command = &command
}

func (u *UpdateContainer) SetCommandNull() {
	u.Command = &util.Nullable[[]string]{IsNull: true}
}

func (u *UpdateContainer) GetPriority() *util.Nullable[shared.ContainerGroupPriority] {
	if u == nil {
		return nil
	}
	return u.Priority
}

func (u *UpdateContainer) SetPriority(priority util.Nullable[shared.ContainerGroupPriority]) {
	u.Priority = &priority
}

func (u *UpdateContainer) SetPriorityNull() {
	u.Priority = &util.Nullable[shared.ContainerGroupPriority]{IsNull: true}
}

func (u *UpdateContainer) GetEnvironmentVariables() map[string]*string {
	if u == nil {
		return nil
	}
	return u.EnvironmentVariables
}

func (u *UpdateContainer) SetEnvironmentVariables(environmentVariables map[string]*string) {
	u.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (u *UpdateContainer) GetLogging() *util.Nullable[UpdateContainerLogging] {
	if u == nil {
		return nil
	}
	return u.Logging
}

func (u *UpdateContainer) SetLogging(logging util.Nullable[UpdateContainerLogging]) {
	u.Logging = &logging
}

func (u *UpdateContainer) SetLoggingNull() {
	u.Logging = &util.Nullable[UpdateContainerLogging]{IsNull: true}
}

func (u *UpdateContainer) GetRegistryAuthentication() *util.Nullable[UpdateContainerRegistryAuthentication] {
	if u == nil {
		return nil
	}
	return u.RegistryAuthentication
}

func (u *UpdateContainer) SetRegistryAuthentication(registryAuthentication util.Nullable[UpdateContainerRegistryAuthentication]) {
	u.RegistryAuthentication = &registryAuthentication
}

func (u *UpdateContainer) SetRegistryAuthenticationNull() {
	u.RegistryAuthentication = &util.Nullable[UpdateContainerRegistryAuthentication]{IsNull: true}
}

func (u *UpdateContainer) GetImageCaching() *bool {
	if u == nil {
		return nil
	}
	return u.ImageCaching
}

func (u *UpdateContainer) SetImageCaching(imageCaching bool) {
	u.ImageCaching = &imageCaching
}

func (u UpdateContainer) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateContainer to string"
	}
	return string(jsonData)
}

type Resources struct {
	Cpu           *util.Nullable[int64]    `json:"cpu,omitempty" min:"1" max:"16"`
	Memory        *util.Nullable[int64]    `json:"memory,omitempty" min:"1024" max:"61440"`
	GpuClasses    *util.Nullable[[]string] `json:"gpu_classes,omitempty"`
	StorageAmount *util.Nullable[int64]    `json:"storage_amount,omitempty" min:"1073741824" max:"53687091200"`
}

func (r *Resources) GetCpu() *util.Nullable[int64] {
	if r == nil {
		return nil
	}
	return r.Cpu
}

func (r *Resources) SetCpu(cpu util.Nullable[int64]) {
	r.Cpu = &cpu
}

func (r *Resources) SetCpuNull() {
	r.Cpu = &util.Nullable[int64]{IsNull: true}
}

func (r *Resources) GetMemory() *util.Nullable[int64] {
	if r == nil {
		return nil
	}
	return r.Memory
}

func (r *Resources) SetMemory(memory util.Nullable[int64]) {
	r.Memory = &memory
}

func (r *Resources) SetMemoryNull() {
	r.Memory = &util.Nullable[int64]{IsNull: true}
}

func (r *Resources) GetGpuClasses() *util.Nullable[[]string] {
	if r == nil {
		return nil
	}
	return r.GpuClasses
}

func (r *Resources) SetGpuClasses(gpuClasses util.Nullable[[]string]) {
	r.GpuClasses = &gpuClasses
}

func (r *Resources) SetGpuClassesNull() {
	r.GpuClasses = &util.Nullable[[]string]{IsNull: true}
}

func (r *Resources) GetStorageAmount() *util.Nullable[int64] {
	if r == nil {
		return nil
	}
	return r.StorageAmount
}

func (r *Resources) SetStorageAmount(storageAmount util.Nullable[int64]) {
	r.StorageAmount = &storageAmount
}

func (r *Resources) SetStorageAmountNull() {
	r.StorageAmount = &util.Nullable[int64]{IsNull: true}
}

func (r Resources) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: Resources to string"
	}
	return string(jsonData)
}

type UpdateContainerLogging struct {
	Axiom    *util.Nullable[LoggingAxiom3]    `json:"axiom,omitempty"`
	Datadog  *util.Nullable[LoggingDatadog3]  `json:"datadog,omitempty"`
	NewRelic *util.Nullable[LoggingNewRelic3] `json:"new_relic,omitempty"`
	Splunk   *util.Nullable[LoggingSplunk3]   `json:"splunk,omitempty"`
	Tcp      *util.Nullable[LoggingTcp3]      `json:"tcp,omitempty"`
	Http     *util.Nullable[LoggingHttp3]     `json:"http,omitempty"`
}

func (u *UpdateContainerLogging) GetAxiom() *util.Nullable[LoggingAxiom3] {
	if u == nil {
		return nil
	}
	return u.Axiom
}

func (u *UpdateContainerLogging) SetAxiom(axiom util.Nullable[LoggingAxiom3]) {
	u.Axiom = &axiom
}

func (u *UpdateContainerLogging) SetAxiomNull() {
	u.Axiom = &util.Nullable[LoggingAxiom3]{IsNull: true}
}

func (u *UpdateContainerLogging) GetDatadog() *util.Nullable[LoggingDatadog3] {
	if u == nil {
		return nil
	}
	return u.Datadog
}

func (u *UpdateContainerLogging) SetDatadog(datadog util.Nullable[LoggingDatadog3]) {
	u.Datadog = &datadog
}

func (u *UpdateContainerLogging) SetDatadogNull() {
	u.Datadog = &util.Nullable[LoggingDatadog3]{IsNull: true}
}

func (u *UpdateContainerLogging) GetNewRelic() *util.Nullable[LoggingNewRelic3] {
	if u == nil {
		return nil
	}
	return u.NewRelic
}

func (u *UpdateContainerLogging) SetNewRelic(newRelic util.Nullable[LoggingNewRelic3]) {
	u.NewRelic = &newRelic
}

func (u *UpdateContainerLogging) SetNewRelicNull() {
	u.NewRelic = &util.Nullable[LoggingNewRelic3]{IsNull: true}
}

func (u *UpdateContainerLogging) GetSplunk() *util.Nullable[LoggingSplunk3] {
	if u == nil {
		return nil
	}
	return u.Splunk
}

func (u *UpdateContainerLogging) SetSplunk(splunk util.Nullable[LoggingSplunk3]) {
	u.Splunk = &splunk
}

func (u *UpdateContainerLogging) SetSplunkNull() {
	u.Splunk = &util.Nullable[LoggingSplunk3]{IsNull: true}
}

func (u *UpdateContainerLogging) GetTcp() *util.Nullable[LoggingTcp3] {
	if u == nil {
		return nil
	}
	return u.Tcp
}

func (u *UpdateContainerLogging) SetTcp(tcp util.Nullable[LoggingTcp3]) {
	u.Tcp = &tcp
}

func (u *UpdateContainerLogging) SetTcpNull() {
	u.Tcp = &util.Nullable[LoggingTcp3]{IsNull: true}
}

func (u *UpdateContainerLogging) GetHttp() *util.Nullable[LoggingHttp3] {
	if u == nil {
		return nil
	}
	return u.Http
}

func (u *UpdateContainerLogging) SetHttp(http util.Nullable[LoggingHttp3]) {
	u.Http = &http
}

func (u *UpdateContainerLogging) SetHttpNull() {
	u.Http = &util.Nullable[LoggingHttp3]{IsNull: true}
}

func (u UpdateContainerLogging) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateContainerLogging to string"
	}
	return string(jsonData)
}

type LoggingAxiom3 struct {
	Host     *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiToken *string `json:"api_token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Dataset  *string `json:"dataset,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingAxiom3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingAxiom3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingAxiom3) GetApiToken() *string {
	if l == nil {
		return nil
	}
	return l.ApiToken
}

func (l *LoggingAxiom3) SetApiToken(apiToken string) {
	l.ApiToken = &apiToken
}

func (l *LoggingAxiom3) GetDataset() *string {
	if l == nil {
		return nil
	}
	return l.Dataset
}

func (l *LoggingAxiom3) SetDataset(dataset string) {
	l.Dataset = &dataset
}

func (l LoggingAxiom3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingAxiom3 to string"
	}
	return string(jsonData)
}

type LoggingDatadog3 struct {
	Host   *string                        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiKey *string                        `json:"api_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Tags   *util.Nullable[[]DatadogTags3] `json:"tags,omitempty" maxItems:"1000"`
}

func (l *LoggingDatadog3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingDatadog3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingDatadog3) GetApiKey() *string {
	if l == nil {
		return nil
	}
	return l.ApiKey
}

func (l *LoggingDatadog3) SetApiKey(apiKey string) {
	l.ApiKey = &apiKey
}

func (l *LoggingDatadog3) GetTags() *util.Nullable[[]DatadogTags3] {
	if l == nil {
		return nil
	}
	return l.Tags
}

func (l *LoggingDatadog3) SetTags(tags util.Nullable[[]DatadogTags3]) {
	l.Tags = &tags
}

func (l *LoggingDatadog3) SetTagsNull() {
	l.Tags = &util.Nullable[[]DatadogTags3]{IsNull: true}
}

func (l LoggingDatadog3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingDatadog3 to string"
	}
	return string(jsonData)
}

type DatadogTags3 struct {
	Name  *string `json:"name,omitempty" required:"true"`
	Value *string `json:"value,omitempty" required:"true"`
}

func (d *DatadogTags3) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DatadogTags3) SetName(name string) {
	d.Name = &name
}

func (d *DatadogTags3) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

func (d *DatadogTags3) SetValue(value string) {
	d.Value = &value
}

func (d DatadogTags3) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DatadogTags3 to string"
	}
	return string(jsonData)
}

type LoggingNewRelic3 struct {
	Host         *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	IngestionKey *string `json:"ingestion_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingNewRelic3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingNewRelic3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingNewRelic3) GetIngestionKey() *string {
	if l == nil {
		return nil
	}
	return l.IngestionKey
}

func (l *LoggingNewRelic3) SetIngestionKey(ingestionKey string) {
	l.IngestionKey = &ingestionKey
}

func (l LoggingNewRelic3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingNewRelic3 to string"
	}
	return string(jsonData)
}

type LoggingSplunk3 struct {
	Host  *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Token *string `json:"token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingSplunk3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingSplunk3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingSplunk3) GetToken() *string {
	if l == nil {
		return nil
	}
	return l.Token
}

func (l *LoggingSplunk3) SetToken(token string) {
	l.Token = &token
}

func (l LoggingSplunk3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingSplunk3 to string"
	}
	return string(jsonData)
}

type LoggingTcp3 struct {
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
}

func (l *LoggingTcp3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingTcp3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingTcp3) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingTcp3) SetPort(port int64) {
	l.Port = &port
}

func (l LoggingTcp3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingTcp3 to string"
	}
	return string(jsonData)
}

type LoggingHttp3 struct {
	Host        *string                        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port        *int64                         `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	User        *util.Nullable[string]         `json:"user,omitempty"`
	Password    *util.Nullable[string]         `json:"password,omitempty"`
	Path        *util.Nullable[string]         `json:"path,omitempty"`
	Format      *HttpFormat3                   `json:"format,omitempty" required:"true"`
	Headers     *util.Nullable[[]HttpHeaders4] `json:"headers,omitempty" maxItems:"1000"`
	Compression *HttpCompression3              `json:"compression,omitempty" required:"true"`
}

func (l *LoggingHttp3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingHttp3) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingHttp3) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingHttp3) SetPort(port int64) {
	l.Port = &port
}

func (l *LoggingHttp3) GetUser() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.User
}

func (l *LoggingHttp3) SetUser(user util.Nullable[string]) {
	l.User = &user
}

func (l *LoggingHttp3) SetUserNull() {
	l.User = &util.Nullable[string]{IsNull: true}
}

func (l *LoggingHttp3) GetPassword() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.Password
}

func (l *LoggingHttp3) SetPassword(password util.Nullable[string]) {
	l.Password = &password
}

func (l *LoggingHttp3) SetPasswordNull() {
	l.Password = &util.Nullable[string]{IsNull: true}
}

func (l *LoggingHttp3) GetPath() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.Path
}

func (l *LoggingHttp3) SetPath(path util.Nullable[string]) {
	l.Path = &path
}

func (l *LoggingHttp3) SetPathNull() {
	l.Path = &util.Nullable[string]{IsNull: true}
}

func (l *LoggingHttp3) GetFormat() *HttpFormat3 {
	if l == nil {
		return nil
	}
	return l.Format
}

func (l *LoggingHttp3) SetFormat(format HttpFormat3) {
	l.Format = &format
}

func (l *LoggingHttp3) GetHeaders() *util.Nullable[[]HttpHeaders4] {
	if l == nil {
		return nil
	}
	return l.Headers
}

func (l *LoggingHttp3) SetHeaders(headers util.Nullable[[]HttpHeaders4]) {
	l.Headers = &headers
}

func (l *LoggingHttp3) SetHeadersNull() {
	l.Headers = &util.Nullable[[]HttpHeaders4]{IsNull: true}
}

func (l *LoggingHttp3) GetCompression() *HttpCompression3 {
	if l == nil {
		return nil
	}
	return l.Compression
}

func (l *LoggingHttp3) SetCompression(compression HttpCompression3) {
	l.Compression = &compression
}

func (l LoggingHttp3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingHttp3 to string"
	}
	return string(jsonData)
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

func (h *HttpHeaders4) GetName() *string {
	if h == nil {
		return nil
	}
	return h.Name
}

func (h *HttpHeaders4) SetName(name string) {
	h.Name = &name
}

func (h *HttpHeaders4) GetValue() *string {
	if h == nil {
		return nil
	}
	return h.Value
}

func (h *HttpHeaders4) SetValue(value string) {
	h.Value = &value
}

func (h HttpHeaders4) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: HttpHeaders4 to string"
	}
	return string(jsonData)
}

type HttpCompression3 string

const (
	HTTP_COMPRESSION3_NONE HttpCompression3 = "none"
	HTTP_COMPRESSION3_GZIP HttpCompression3 = "gzip"
)

type UpdateContainerRegistryAuthentication struct {
	Basic     *util.Nullable[RegistryAuthenticationBasic2]     `json:"basic,omitempty"`
	GcpGcr    *util.Nullable[RegistryAuthenticationGcpGcr2]    `json:"gcp_gcr,omitempty"`
	AwsEcr    *util.Nullable[RegistryAuthenticationAwsEcr2]    `json:"aws_ecr,omitempty"`
	DockerHub *util.Nullable[RegistryAuthenticationDockerHub2] `json:"docker_hub,omitempty"`
	GcpGar    *util.Nullable[RegistryAuthenticationGcpGar2]    `json:"gcp_gar,omitempty"`
}

func (u *UpdateContainerRegistryAuthentication) GetBasic() *util.Nullable[RegistryAuthenticationBasic2] {
	if u == nil {
		return nil
	}
	return u.Basic
}

func (u *UpdateContainerRegistryAuthentication) SetBasic(basic util.Nullable[RegistryAuthenticationBasic2]) {
	u.Basic = &basic
}

func (u *UpdateContainerRegistryAuthentication) SetBasicNull() {
	u.Basic = &util.Nullable[RegistryAuthenticationBasic2]{IsNull: true}
}

func (u *UpdateContainerRegistryAuthentication) GetGcpGcr() *util.Nullable[RegistryAuthenticationGcpGcr2] {
	if u == nil {
		return nil
	}
	return u.GcpGcr
}

func (u *UpdateContainerRegistryAuthentication) SetGcpGcr(gcpGcr util.Nullable[RegistryAuthenticationGcpGcr2]) {
	u.GcpGcr = &gcpGcr
}

func (u *UpdateContainerRegistryAuthentication) SetGcpGcrNull() {
	u.GcpGcr = &util.Nullable[RegistryAuthenticationGcpGcr2]{IsNull: true}
}

func (u *UpdateContainerRegistryAuthentication) GetAwsEcr() *util.Nullable[RegistryAuthenticationAwsEcr2] {
	if u == nil {
		return nil
	}
	return u.AwsEcr
}

func (u *UpdateContainerRegistryAuthentication) SetAwsEcr(awsEcr util.Nullable[RegistryAuthenticationAwsEcr2]) {
	u.AwsEcr = &awsEcr
}

func (u *UpdateContainerRegistryAuthentication) SetAwsEcrNull() {
	u.AwsEcr = &util.Nullable[RegistryAuthenticationAwsEcr2]{IsNull: true}
}

func (u *UpdateContainerRegistryAuthentication) GetDockerHub() *util.Nullable[RegistryAuthenticationDockerHub2] {
	if u == nil {
		return nil
	}
	return u.DockerHub
}

func (u *UpdateContainerRegistryAuthentication) SetDockerHub(dockerHub util.Nullable[RegistryAuthenticationDockerHub2]) {
	u.DockerHub = &dockerHub
}

func (u *UpdateContainerRegistryAuthentication) SetDockerHubNull() {
	u.DockerHub = &util.Nullable[RegistryAuthenticationDockerHub2]{IsNull: true}
}

func (u *UpdateContainerRegistryAuthentication) GetGcpGar() *util.Nullable[RegistryAuthenticationGcpGar2] {
	if u == nil {
		return nil
	}
	return u.GcpGar
}

func (u *UpdateContainerRegistryAuthentication) SetGcpGar(gcpGar util.Nullable[RegistryAuthenticationGcpGar2]) {
	u.GcpGar = &gcpGar
}

func (u *UpdateContainerRegistryAuthentication) SetGcpGarNull() {
	u.GcpGar = &util.Nullable[RegistryAuthenticationGcpGar2]{IsNull: true}
}

func (u UpdateContainerRegistryAuthentication) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateContainerRegistryAuthentication to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationBasic2 struct {
	Username *string `json:"username,omitempty" required:"true"`
	Password *string `json:"password,omitempty" required:"true"`
}

func (r *RegistryAuthenticationBasic2) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationBasic2) SetUsername(username string) {
	r.Username = &username
}

func (r *RegistryAuthenticationBasic2) GetPassword() *string {
	if r == nil {
		return nil
	}
	return r.Password
}

func (r *RegistryAuthenticationBasic2) SetPassword(password string) {
	r.Password = &password
}

func (r RegistryAuthenticationBasic2) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationBasic2 to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationGcpGcr2 struct {
	ServiceKey *string `json:"service_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationGcpGcr2) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}

func (r *RegistryAuthenticationGcpGcr2) SetServiceKey(serviceKey string) {
	r.ServiceKey = &serviceKey
}

func (r RegistryAuthenticationGcpGcr2) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationGcpGcr2 to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationAwsEcr2 struct {
	AccessKeyId     *string `json:"access_key_id,omitempty" required:"true"`
	SecretAccessKey *string `json:"secret_access_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationAwsEcr2) GetAccessKeyId() *string {
	if r == nil {
		return nil
	}
	return r.AccessKeyId
}

func (r *RegistryAuthenticationAwsEcr2) SetAccessKeyId(accessKeyId string) {
	r.AccessKeyId = &accessKeyId
}

func (r *RegistryAuthenticationAwsEcr2) GetSecretAccessKey() *string {
	if r == nil {
		return nil
	}
	return r.SecretAccessKey
}

func (r *RegistryAuthenticationAwsEcr2) SetSecretAccessKey(secretAccessKey string) {
	r.SecretAccessKey = &secretAccessKey
}

func (r RegistryAuthenticationAwsEcr2) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationAwsEcr2 to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationDockerHub2 struct {
	Username            *string `json:"username,omitempty" required:"true"`
	PersonalAccessToken *string `json:"personal_access_token,omitempty" required:"true"`
}

func (r *RegistryAuthenticationDockerHub2) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationDockerHub2) SetUsername(username string) {
	r.Username = &username
}

func (r *RegistryAuthenticationDockerHub2) GetPersonalAccessToken() *string {
	if r == nil {
		return nil
	}
	return r.PersonalAccessToken
}

func (r *RegistryAuthenticationDockerHub2) SetPersonalAccessToken(personalAccessToken string) {
	r.PersonalAccessToken = &personalAccessToken
}

func (r RegistryAuthenticationDockerHub2) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationDockerHub2 to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationGcpGar2 struct {
	ServiceKey *string `json:"service_key,omitempty" required:"true"`
}

func (r *RegistryAuthenticationGcpGar2) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}

func (r *RegistryAuthenticationGcpGar2) SetServiceKey(serviceKey string) {
	r.ServiceKey = &serviceKey
}

func (r RegistryAuthenticationGcpGar2) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationGcpGar2 to string"
	}
	return string(jsonData)
}
