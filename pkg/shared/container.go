package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Represents a container
type Container struct {
	Image *string `json:"image,omitempty" required:"true" maxLength:"1024" minLength:"1"`
	// Represents a container resource requirements
	Resources            *ContainerResourceRequirements         `json:"resources,omitempty" required:"true"`
	Command              []string                               `json:"command,omitempty" required:"true" maxItems:"100"`
	Priority             *util.Nullable[ContainerGroupPriority] `json:"priority,omitempty"`
	Size                 *int64                                 `json:"size,omitempty"`
	Hash                 *string                                `json:"hash,omitempty"`
	EnvironmentVariables map[string]*string                     `json:"environment_variables,omitempty"`
	Logging              *util.Nullable[ContainerLogging]       `json:"logging,omitempty"`
	ImageCaching         *bool                                  `json:"image_caching,omitempty"`
}

func (c *Container) GetImage() *string {
	if c == nil {
		return nil
	}
	return c.Image
}

func (c *Container) SetImage(image string) {
	c.Image = &image
}

func (c *Container) GetResources() *ContainerResourceRequirements {
	if c == nil {
		return nil
	}
	return c.Resources
}

func (c *Container) SetResources(resources ContainerResourceRequirements) {
	c.Resources = &resources
}

func (c *Container) GetCommand() []string {
	if c == nil {
		return nil
	}
	return c.Command
}

func (c *Container) SetCommand(command []string) {
	c.Command = command
}

func (c *Container) GetPriority() *util.Nullable[ContainerGroupPriority] {
	if c == nil {
		return nil
	}
	return c.Priority
}

func (c *Container) SetPriority(priority util.Nullable[ContainerGroupPriority]) {
	c.Priority = &priority
}

func (c *Container) SetPriorityNull() {
	c.Priority = &util.Nullable[ContainerGroupPriority]{IsNull: true}
}

func (c *Container) GetSize() *int64 {
	if c == nil {
		return nil
	}
	return c.Size
}

func (c *Container) SetSize(size int64) {
	c.Size = &size
}

func (c *Container) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *Container) SetHash(hash string) {
	c.Hash = &hash
}

func (c *Container) GetEnvironmentVariables() map[string]*string {
	if c == nil {
		return nil
	}
	return c.EnvironmentVariables
}

func (c *Container) SetEnvironmentVariables(environmentVariables map[string]*string) {
	c.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (c *Container) GetLogging() *util.Nullable[ContainerLogging] {
	if c == nil {
		return nil
	}
	return c.Logging
}

func (c *Container) SetLogging(logging util.Nullable[ContainerLogging]) {
	c.Logging = &logging
}

func (c *Container) SetLoggingNull() {
	c.Logging = &util.Nullable[ContainerLogging]{IsNull: true}
}

func (c *Container) GetImageCaching() *bool {
	if c == nil {
		return nil
	}
	return c.ImageCaching
}

func (c *Container) SetImageCaching(imageCaching bool) {
	c.ImageCaching = &imageCaching
}

func (c Container) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: Container to string"
	}
	return string(jsonData)
}

type ContainerLogging struct {
	Axiom    *util.Nullable[LoggingAxiom1]    `json:"axiom,omitempty"`
	Datadog  *util.Nullable[LoggingDatadog1]  `json:"datadog,omitempty"`
	NewRelic *util.Nullable[LoggingNewRelic1] `json:"new_relic,omitempty"`
	Splunk   *util.Nullable[LoggingSplunk1]   `json:"splunk,omitempty"`
	Tcp      *util.Nullable[LoggingTcp1]      `json:"tcp,omitempty"`
	Http     *util.Nullable[LoggingHttp1]     `json:"http,omitempty"`
}

func (c *ContainerLogging) GetAxiom() *util.Nullable[LoggingAxiom1] {
	if c == nil {
		return nil
	}
	return c.Axiom
}

func (c *ContainerLogging) SetAxiom(axiom util.Nullable[LoggingAxiom1]) {
	c.Axiom = &axiom
}

func (c *ContainerLogging) SetAxiomNull() {
	c.Axiom = &util.Nullable[LoggingAxiom1]{IsNull: true}
}

func (c *ContainerLogging) GetDatadog() *util.Nullable[LoggingDatadog1] {
	if c == nil {
		return nil
	}
	return c.Datadog
}

func (c *ContainerLogging) SetDatadog(datadog util.Nullable[LoggingDatadog1]) {
	c.Datadog = &datadog
}

func (c *ContainerLogging) SetDatadogNull() {
	c.Datadog = &util.Nullable[LoggingDatadog1]{IsNull: true}
}

func (c *ContainerLogging) GetNewRelic() *util.Nullable[LoggingNewRelic1] {
	if c == nil {
		return nil
	}
	return c.NewRelic
}

func (c *ContainerLogging) SetNewRelic(newRelic util.Nullable[LoggingNewRelic1]) {
	c.NewRelic = &newRelic
}

func (c *ContainerLogging) SetNewRelicNull() {
	c.NewRelic = &util.Nullable[LoggingNewRelic1]{IsNull: true}
}

func (c *ContainerLogging) GetSplunk() *util.Nullable[LoggingSplunk1] {
	if c == nil {
		return nil
	}
	return c.Splunk
}

func (c *ContainerLogging) SetSplunk(splunk util.Nullable[LoggingSplunk1]) {
	c.Splunk = &splunk
}

func (c *ContainerLogging) SetSplunkNull() {
	c.Splunk = &util.Nullable[LoggingSplunk1]{IsNull: true}
}

func (c *ContainerLogging) GetTcp() *util.Nullable[LoggingTcp1] {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerLogging) SetTcp(tcp util.Nullable[LoggingTcp1]) {
	c.Tcp = &tcp
}

func (c *ContainerLogging) SetTcpNull() {
	c.Tcp = &util.Nullable[LoggingTcp1]{IsNull: true}
}

func (c *ContainerLogging) GetHttp() *util.Nullable[LoggingHttp1] {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerLogging) SetHttp(http util.Nullable[LoggingHttp1]) {
	c.Http = &http
}

func (c *ContainerLogging) SetHttpNull() {
	c.Http = &util.Nullable[LoggingHttp1]{IsNull: true}
}

func (c ContainerLogging) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerLogging to string"
	}
	return string(jsonData)
}

type LoggingAxiom1 struct {
	Host     *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiToken *string `json:"api_token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Dataset  *string `json:"dataset,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingAxiom1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingAxiom1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingAxiom1) GetApiToken() *string {
	if l == nil {
		return nil
	}
	return l.ApiToken
}

func (l *LoggingAxiom1) SetApiToken(apiToken string) {
	l.ApiToken = &apiToken
}

func (l *LoggingAxiom1) GetDataset() *string {
	if l == nil {
		return nil
	}
	return l.Dataset
}

func (l *LoggingAxiom1) SetDataset(dataset string) {
	l.Dataset = &dataset
}

func (l LoggingAxiom1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingAxiom1 to string"
	}
	return string(jsonData)
}

type LoggingDatadog1 struct {
	Host   *string                        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiKey *string                        `json:"api_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Tags   *util.Nullable[[]DatadogTags1] `json:"tags,omitempty" maxItems:"1000"`
}

func (l *LoggingDatadog1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingDatadog1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingDatadog1) GetApiKey() *string {
	if l == nil {
		return nil
	}
	return l.ApiKey
}

func (l *LoggingDatadog1) SetApiKey(apiKey string) {
	l.ApiKey = &apiKey
}

func (l *LoggingDatadog1) GetTags() *util.Nullable[[]DatadogTags1] {
	if l == nil {
		return nil
	}
	return l.Tags
}

func (l *LoggingDatadog1) SetTags(tags util.Nullable[[]DatadogTags1]) {
	l.Tags = &tags
}

func (l *LoggingDatadog1) SetTagsNull() {
	l.Tags = &util.Nullable[[]DatadogTags1]{IsNull: true}
}

func (l LoggingDatadog1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingDatadog1 to string"
	}
	return string(jsonData)
}

type DatadogTags1 struct {
	Name  *string `json:"name,omitempty" required:"true"`
	Value *string `json:"value,omitempty" required:"true"`
}

func (d *DatadogTags1) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DatadogTags1) SetName(name string) {
	d.Name = &name
}

func (d *DatadogTags1) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

func (d *DatadogTags1) SetValue(value string) {
	d.Value = &value
}

func (d DatadogTags1) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DatadogTags1 to string"
	}
	return string(jsonData)
}

type LoggingNewRelic1 struct {
	Host         *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	IngestionKey *string `json:"ingestion_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingNewRelic1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingNewRelic1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingNewRelic1) GetIngestionKey() *string {
	if l == nil {
		return nil
	}
	return l.IngestionKey
}

func (l *LoggingNewRelic1) SetIngestionKey(ingestionKey string) {
	l.IngestionKey = &ingestionKey
}

func (l LoggingNewRelic1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingNewRelic1 to string"
	}
	return string(jsonData)
}

type LoggingSplunk1 struct {
	Host  *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Token *string `json:"token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
}

func (l *LoggingSplunk1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingSplunk1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingSplunk1) GetToken() *string {
	if l == nil {
		return nil
	}
	return l.Token
}

func (l *LoggingSplunk1) SetToken(token string) {
	l.Token = &token
}

func (l LoggingSplunk1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingSplunk1 to string"
	}
	return string(jsonData)
}

type LoggingTcp1 struct {
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
}

func (l *LoggingTcp1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingTcp1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingTcp1) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingTcp1) SetPort(port int64) {
	l.Port = &port
}

func (l LoggingTcp1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingTcp1 to string"
	}
	return string(jsonData)
}

type LoggingHttp1 struct {
	Host        *string                        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port        *int64                         `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	User        *util.Nullable[string]         `json:"user,omitempty"`
	Password    *util.Nullable[string]         `json:"password,omitempty"`
	Path        *util.Nullable[string]         `json:"path,omitempty"`
	Format      *HttpFormat1                   `json:"format,omitempty" required:"true"`
	Headers     *util.Nullable[[]HttpHeaders1] `json:"headers,omitempty" maxItems:"1000"`
	Compression *HttpCompression1              `json:"compression,omitempty" required:"true"`
}

func (l *LoggingHttp1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingHttp1) SetHost(host string) {
	l.Host = &host
}

func (l *LoggingHttp1) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingHttp1) SetPort(port int64) {
	l.Port = &port
}

func (l *LoggingHttp1) GetUser() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.User
}

func (l *LoggingHttp1) SetUser(user util.Nullable[string]) {
	l.User = &user
}

func (l *LoggingHttp1) SetUserNull() {
	l.User = &util.Nullable[string]{IsNull: true}
}

func (l *LoggingHttp1) GetPassword() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.Password
}

func (l *LoggingHttp1) SetPassword(password util.Nullable[string]) {
	l.Password = &password
}

func (l *LoggingHttp1) SetPasswordNull() {
	l.Password = &util.Nullable[string]{IsNull: true}
}

func (l *LoggingHttp1) GetPath() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.Path
}

func (l *LoggingHttp1) SetPath(path util.Nullable[string]) {
	l.Path = &path
}

func (l *LoggingHttp1) SetPathNull() {
	l.Path = &util.Nullable[string]{IsNull: true}
}

func (l *LoggingHttp1) GetFormat() *HttpFormat1 {
	if l == nil {
		return nil
	}
	return l.Format
}

func (l *LoggingHttp1) SetFormat(format HttpFormat1) {
	l.Format = &format
}

func (l *LoggingHttp1) GetHeaders() *util.Nullable[[]HttpHeaders1] {
	if l == nil {
		return nil
	}
	return l.Headers
}

func (l *LoggingHttp1) SetHeaders(headers util.Nullable[[]HttpHeaders1]) {
	l.Headers = &headers
}

func (l *LoggingHttp1) SetHeadersNull() {
	l.Headers = &util.Nullable[[]HttpHeaders1]{IsNull: true}
}

func (l *LoggingHttp1) GetCompression() *HttpCompression1 {
	if l == nil {
		return nil
	}
	return l.Compression
}

func (l *LoggingHttp1) SetCompression(compression HttpCompression1) {
	l.Compression = &compression
}

func (l LoggingHttp1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingHttp1 to string"
	}
	return string(jsonData)
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

func (h *HttpHeaders1) GetName() *string {
	if h == nil {
		return nil
	}
	return h.Name
}

func (h *HttpHeaders1) SetName(name string) {
	h.Name = &name
}

func (h *HttpHeaders1) GetValue() *string {
	if h == nil {
		return nil
	}
	return h.Value
}

func (h *HttpHeaders1) SetValue(value string) {
	h.Value = &value
}

func (h HttpHeaders1) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: HttpHeaders1 to string"
	}
	return string(jsonData)
}

type HttpCompression1 string

const (
	HTTP_COMPRESSION1_NONE HttpCompression1 = "none"
	HTTP_COMPRESSION1_GZIP HttpCompression1 = "gzip"
)
