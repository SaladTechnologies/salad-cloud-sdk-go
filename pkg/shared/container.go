package shared

import (
	"encoding/json"
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
	touched              map[string]bool
}

func (c *Container) GetImage() *string {
	if c == nil {
		return nil
	}
	return c.Image
}

func (c *Container) SetImage(image string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Image"] = true
	c.Image = &image
}

func (c *Container) SetImageNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Image"] = true
	c.Image = nil
}

func (c *Container) GetResources() *ContainerResourceRequirements {
	if c == nil {
		return nil
	}
	return c.Resources
}

func (c *Container) SetResources(resources ContainerResourceRequirements) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Resources"] = true
	c.Resources = &resources
}

func (c *Container) SetResourcesNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Resources"] = true
	c.Resources = nil
}

func (c *Container) GetCommand() []string {
	if c == nil {
		return nil
	}
	return c.Command
}

func (c *Container) SetCommand(command []string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Command"] = true
	c.Command = command
}

func (c *Container) SetCommandNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Command"] = true
	c.Command = nil
}

func (c *Container) GetPriority() *ContainerGroupPriority {
	if c == nil {
		return nil
	}
	return c.Priority
}

func (c *Container) SetPriority(priority ContainerGroupPriority) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Priority"] = true
	c.Priority = &priority
}

func (c *Container) SetPriorityNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Priority"] = true
	c.Priority = nil
}

func (c *Container) GetSize() *int64 {
	if c == nil {
		return nil
	}
	return c.Size
}

func (c *Container) SetSize(size int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Size"] = true
	c.Size = &size
}

func (c *Container) SetSizeNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Size"] = true
	c.Size = nil
}

func (c *Container) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *Container) SetHash(hash string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Hash"] = true
	c.Hash = &hash
}

func (c *Container) SetHashNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Hash"] = true
	c.Hash = nil
}

func (c *Container) GetEnvironmentVariables() map[string]*string {
	if c == nil {
		return nil
	}
	return c.EnvironmentVariables
}

func (c *Container) SetEnvironmentVariables(environmentVariables map[string]*string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EnvironmentVariables"] = true
	c.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (c *Container) SetEnvironmentVariablesNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EnvironmentVariables"] = true
	c.EnvironmentVariables = nil
}

func (c *Container) GetLogging() *ContainerLogging {
	if c == nil {
		return nil
	}
	return c.Logging
}

func (c *Container) SetLogging(logging ContainerLogging) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Logging"] = true
	c.Logging = &logging
}

func (c *Container) SetLoggingNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Logging"] = true
	c.Logging = nil
}

func (c Container) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Image"] && c.Image == nil {
		data["image"] = nil
	} else if c.Image != nil {
		data["image"] = c.Image
	}

	if c.touched["Resources"] && c.Resources == nil {
		data["resources"] = nil
	} else if c.Resources != nil {
		data["resources"] = c.Resources
	}

	if c.touched["Command"] && c.Command == nil {
		data["command"] = nil
	} else if c.Command != nil {
		data["command"] = c.Command
	}

	if c.touched["Priority"] && c.Priority == nil {
		data["priority"] = nil
	} else if c.Priority != nil {
		data["priority"] = c.Priority
	}

	if c.touched["Size"] && c.Size == nil {
		data["size"] = nil
	} else if c.Size != nil {
		data["size"] = c.Size
	}

	if c.touched["Hash"] && c.Hash == nil {
		data["hash"] = nil
	} else if c.Hash != nil {
		data["hash"] = c.Hash
	}

	if c.touched["EnvironmentVariables"] && c.EnvironmentVariables == nil {
		data["environment_variables"] = nil
	} else if c.EnvironmentVariables != nil {
		data["environment_variables"] = c.EnvironmentVariables
	}

	if c.touched["Logging"] && c.Logging == nil {
		data["logging"] = nil
	} else if c.Logging != nil {
		data["logging"] = c.Logging
	}

	return json.Marshal(data)
}

func (c Container) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: Container to string"
	}
	return string(jsonData)
}

type ContainerLogging struct {
	Axiom    *LoggingAxiom1    `json:"axiom,omitempty"`
	Datadog  *LoggingDatadog1  `json:"datadog,omitempty"`
	NewRelic *LoggingNewRelic1 `json:"new_relic,omitempty"`
	Splunk   *LoggingSplunk1   `json:"splunk,omitempty"`
	Tcp      *LoggingTcp1      `json:"tcp,omitempty"`
	Http     *LoggingHttp1     `json:"http,omitempty"`
	touched  map[string]bool
}

func (c *ContainerLogging) GetAxiom() *LoggingAxiom1 {
	if c == nil {
		return nil
	}
	return c.Axiom
}

func (c *ContainerLogging) SetAxiom(axiom LoggingAxiom1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Axiom"] = true
	c.Axiom = &axiom
}

func (c *ContainerLogging) SetAxiomNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Axiom"] = true
	c.Axiom = nil
}

func (c *ContainerLogging) GetDatadog() *LoggingDatadog1 {
	if c == nil {
		return nil
	}
	return c.Datadog
}

func (c *ContainerLogging) SetDatadog(datadog LoggingDatadog1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Datadog"] = true
	c.Datadog = &datadog
}

func (c *ContainerLogging) SetDatadogNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Datadog"] = true
	c.Datadog = nil
}

func (c *ContainerLogging) GetNewRelic() *LoggingNewRelic1 {
	if c == nil {
		return nil
	}
	return c.NewRelic
}

func (c *ContainerLogging) SetNewRelic(newRelic LoggingNewRelic1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["NewRelic"] = true
	c.NewRelic = &newRelic
}

func (c *ContainerLogging) SetNewRelicNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["NewRelic"] = true
	c.NewRelic = nil
}

func (c *ContainerLogging) GetSplunk() *LoggingSplunk1 {
	if c == nil {
		return nil
	}
	return c.Splunk
}

func (c *ContainerLogging) SetSplunk(splunk LoggingSplunk1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Splunk"] = true
	c.Splunk = &splunk
}

func (c *ContainerLogging) SetSplunkNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Splunk"] = true
	c.Splunk = nil
}

func (c *ContainerLogging) GetTcp() *LoggingTcp1 {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerLogging) SetTcp(tcp LoggingTcp1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Tcp"] = true
	c.Tcp = &tcp
}

func (c *ContainerLogging) SetTcpNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Tcp"] = true
	c.Tcp = nil
}

func (c *ContainerLogging) GetHttp() *LoggingHttp1 {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerLogging) SetHttp(http LoggingHttp1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Http"] = true
	c.Http = &http
}

func (c *ContainerLogging) SetHttpNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Http"] = true
	c.Http = nil
}

func (c ContainerLogging) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Axiom"] && c.Axiom == nil {
		data["axiom"] = nil
	} else if c.Axiom != nil {
		data["axiom"] = c.Axiom
	}

	if c.touched["Datadog"] && c.Datadog == nil {
		data["datadog"] = nil
	} else if c.Datadog != nil {
		data["datadog"] = c.Datadog
	}

	if c.touched["NewRelic"] && c.NewRelic == nil {
		data["new_relic"] = nil
	} else if c.NewRelic != nil {
		data["new_relic"] = c.NewRelic
	}

	if c.touched["Splunk"] && c.Splunk == nil {
		data["splunk"] = nil
	} else if c.Splunk != nil {
		data["splunk"] = c.Splunk
	}

	if c.touched["Tcp"] && c.Tcp == nil {
		data["tcp"] = nil
	} else if c.Tcp != nil {
		data["tcp"] = c.Tcp
	}

	if c.touched["Http"] && c.Http == nil {
		data["http"] = nil
	} else if c.Http != nil {
		data["http"] = c.Http
	}

	return json.Marshal(data)
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
	touched  map[string]bool
}

func (l *LoggingAxiom1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingAxiom1) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingAxiom1) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingAxiom1) GetApiToken() *string {
	if l == nil {
		return nil
	}
	return l.ApiToken
}

func (l *LoggingAxiom1) SetApiToken(apiToken string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiToken"] = true
	l.ApiToken = &apiToken
}

func (l *LoggingAxiom1) SetApiTokenNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiToken"] = true
	l.ApiToken = nil
}

func (l *LoggingAxiom1) GetDataset() *string {
	if l == nil {
		return nil
	}
	return l.Dataset
}

func (l *LoggingAxiom1) SetDataset(dataset string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Dataset"] = true
	l.Dataset = &dataset
}

func (l *LoggingAxiom1) SetDatasetNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Dataset"] = true
	l.Dataset = nil
}

func (l LoggingAxiom1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Host"] && l.Host == nil {
		data["host"] = nil
	} else if l.Host != nil {
		data["host"] = l.Host
	}

	if l.touched["ApiToken"] && l.ApiToken == nil {
		data["api_token"] = nil
	} else if l.ApiToken != nil {
		data["api_token"] = l.ApiToken
	}

	if l.touched["Dataset"] && l.Dataset == nil {
		data["dataset"] = nil
	} else if l.Dataset != nil {
		data["dataset"] = l.Dataset
	}

	return json.Marshal(data)
}

func (l LoggingAxiom1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingAxiom1 to string"
	}
	return string(jsonData)
}

type LoggingDatadog1 struct {
	Host    *string        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiKey  *string        `json:"api_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Tags    []DatadogTags1 `json:"tags,omitempty" maxItems:"1000"`
	touched map[string]bool
}

func (l *LoggingDatadog1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingDatadog1) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingDatadog1) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingDatadog1) GetApiKey() *string {
	if l == nil {
		return nil
	}
	return l.ApiKey
}

func (l *LoggingDatadog1) SetApiKey(apiKey string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiKey"] = true
	l.ApiKey = &apiKey
}

func (l *LoggingDatadog1) SetApiKeyNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiKey"] = true
	l.ApiKey = nil
}

func (l *LoggingDatadog1) GetTags() []DatadogTags1 {
	if l == nil {
		return nil
	}
	return l.Tags
}

func (l *LoggingDatadog1) SetTags(tags []DatadogTags1) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Tags"] = true
	l.Tags = tags
}

func (l *LoggingDatadog1) SetTagsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Tags"] = true
	l.Tags = nil
}

func (l LoggingDatadog1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Host"] && l.Host == nil {
		data["host"] = nil
	} else if l.Host != nil {
		data["host"] = l.Host
	}

	if l.touched["ApiKey"] && l.ApiKey == nil {
		data["api_key"] = nil
	} else if l.ApiKey != nil {
		data["api_key"] = l.ApiKey
	}

	if l.touched["Tags"] && l.Tags == nil {
		data["tags"] = nil
	} else if l.Tags != nil {
		data["tags"] = l.Tags
	}

	return json.Marshal(data)
}

func (l LoggingDatadog1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingDatadog1 to string"
	}
	return string(jsonData)
}

type DatadogTags1 struct {
	Name    *string `json:"name,omitempty" required:"true"`
	Value   *string `json:"value,omitempty" required:"true"`
	touched map[string]bool
}

func (d *DatadogTags1) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DatadogTags1) SetName(name string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = &name
}

func (d *DatadogTags1) SetNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = nil
}

func (d *DatadogTags1) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

func (d *DatadogTags1) SetValue(value string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Value"] = true
	d.Value = &value
}

func (d *DatadogTags1) SetValueNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Value"] = true
	d.Value = nil
}

func (d DatadogTags1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["Name"] && d.Name == nil {
		data["name"] = nil
	} else if d.Name != nil {
		data["name"] = d.Name
	}

	if d.touched["Value"] && d.Value == nil {
		data["value"] = nil
	} else if d.Value != nil {
		data["value"] = d.Value
	}

	return json.Marshal(data)
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
	touched      map[string]bool
}

func (l *LoggingNewRelic1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingNewRelic1) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingNewRelic1) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingNewRelic1) GetIngestionKey() *string {
	if l == nil {
		return nil
	}
	return l.IngestionKey
}

func (l *LoggingNewRelic1) SetIngestionKey(ingestionKey string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["IngestionKey"] = true
	l.IngestionKey = &ingestionKey
}

func (l *LoggingNewRelic1) SetIngestionKeyNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["IngestionKey"] = true
	l.IngestionKey = nil
}

func (l LoggingNewRelic1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Host"] && l.Host == nil {
		data["host"] = nil
	} else if l.Host != nil {
		data["host"] = l.Host
	}

	if l.touched["IngestionKey"] && l.IngestionKey == nil {
		data["ingestion_key"] = nil
	} else if l.IngestionKey != nil {
		data["ingestion_key"] = l.IngestionKey
	}

	return json.Marshal(data)
}

func (l LoggingNewRelic1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingNewRelic1 to string"
	}
	return string(jsonData)
}

type LoggingSplunk1 struct {
	Host    *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Token   *string `json:"token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	touched map[string]bool
}

func (l *LoggingSplunk1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingSplunk1) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingSplunk1) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingSplunk1) GetToken() *string {
	if l == nil {
		return nil
	}
	return l.Token
}

func (l *LoggingSplunk1) SetToken(token string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Token"] = true
	l.Token = &token
}

func (l *LoggingSplunk1) SetTokenNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Token"] = true
	l.Token = nil
}

func (l LoggingSplunk1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Host"] && l.Host == nil {
		data["host"] = nil
	} else if l.Host != nil {
		data["host"] = l.Host
	}

	if l.touched["Token"] && l.Token == nil {
		data["token"] = nil
	} else if l.Token != nil {
		data["token"] = l.Token
	}

	return json.Marshal(data)
}

func (l LoggingSplunk1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingSplunk1 to string"
	}
	return string(jsonData)
}

type LoggingTcp1 struct {
	Host    *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port    *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	touched map[string]bool
}

func (l *LoggingTcp1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingTcp1) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingTcp1) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingTcp1) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingTcp1) SetPort(port int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = &port
}

func (l *LoggingTcp1) SetPortNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = nil
}

func (l LoggingTcp1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Host"] && l.Host == nil {
		data["host"] = nil
	} else if l.Host != nil {
		data["host"] = l.Host
	}

	if l.touched["Port"] && l.Port == nil {
		data["port"] = nil
	} else if l.Port != nil {
		data["port"] = l.Port
	}

	return json.Marshal(data)
}

func (l LoggingTcp1) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingTcp1 to string"
	}
	return string(jsonData)
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
	touched     map[string]bool
}

func (l *LoggingHttp1) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingHttp1) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingHttp1) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingHttp1) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingHttp1) SetPort(port int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = &port
}

func (l *LoggingHttp1) SetPortNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = nil
}

func (l *LoggingHttp1) GetUser() *string {
	if l == nil {
		return nil
	}
	return l.User
}

func (l *LoggingHttp1) SetUser(user string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["User"] = true
	l.User = &user
}

func (l *LoggingHttp1) SetUserNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["User"] = true
	l.User = nil
}

func (l *LoggingHttp1) GetPassword() *string {
	if l == nil {
		return nil
	}
	return l.Password
}

func (l *LoggingHttp1) SetPassword(password string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Password"] = true
	l.Password = &password
}

func (l *LoggingHttp1) SetPasswordNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Password"] = true
	l.Password = nil
}

func (l *LoggingHttp1) GetPath() *string {
	if l == nil {
		return nil
	}
	return l.Path
}

func (l *LoggingHttp1) SetPath(path string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Path"] = true
	l.Path = &path
}

func (l *LoggingHttp1) SetPathNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Path"] = true
	l.Path = nil
}

func (l *LoggingHttp1) GetFormat() *HttpFormat1 {
	if l == nil {
		return nil
	}
	return l.Format
}

func (l *LoggingHttp1) SetFormat(format HttpFormat1) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Format"] = true
	l.Format = &format
}

func (l *LoggingHttp1) SetFormatNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Format"] = true
	l.Format = nil
}

func (l *LoggingHttp1) GetHeaders() []HttpHeaders1 {
	if l == nil {
		return nil
	}
	return l.Headers
}

func (l *LoggingHttp1) SetHeaders(headers []HttpHeaders1) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Headers"] = true
	l.Headers = headers
}

func (l *LoggingHttp1) SetHeadersNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Headers"] = true
	l.Headers = nil
}

func (l *LoggingHttp1) GetCompression() *HttpCompression1 {
	if l == nil {
		return nil
	}
	return l.Compression
}

func (l *LoggingHttp1) SetCompression(compression HttpCompression1) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Compression"] = true
	l.Compression = &compression
}

func (l *LoggingHttp1) SetCompressionNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Compression"] = true
	l.Compression = nil
}

func (l LoggingHttp1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Host"] && l.Host == nil {
		data["host"] = nil
	} else if l.Host != nil {
		data["host"] = l.Host
	}

	if l.touched["Port"] && l.Port == nil {
		data["port"] = nil
	} else if l.Port != nil {
		data["port"] = l.Port
	}

	if l.touched["User"] && l.User == nil {
		data["user"] = nil
	} else if l.User != nil {
		data["user"] = l.User
	}

	if l.touched["Password"] && l.Password == nil {
		data["password"] = nil
	} else if l.Password != nil {
		data["password"] = l.Password
	}

	if l.touched["Path"] && l.Path == nil {
		data["path"] = nil
	} else if l.Path != nil {
		data["path"] = l.Path
	}

	if l.touched["Format"] && l.Format == nil {
		data["format"] = nil
	} else if l.Format != nil {
		data["format"] = l.Format
	}

	if l.touched["Headers"] && l.Headers == nil {
		data["headers"] = nil
	} else if l.Headers != nil {
		data["headers"] = l.Headers
	}

	if l.touched["Compression"] && l.Compression == nil {
		data["compression"] = nil
	} else if l.Compression != nil {
		data["compression"] = l.Compression
	}

	return json.Marshal(data)
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
	Name    *string `json:"name,omitempty" required:"true"`
	Value   *string `json:"value,omitempty" required:"true"`
	touched map[string]bool
}

func (h *HttpHeaders1) GetName() *string {
	if h == nil {
		return nil
	}
	return h.Name
}

func (h *HttpHeaders1) SetName(name string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Name"] = true
	h.Name = &name
}

func (h *HttpHeaders1) SetNameNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Name"] = true
	h.Name = nil
}

func (h *HttpHeaders1) GetValue() *string {
	if h == nil {
		return nil
	}
	return h.Value
}

func (h *HttpHeaders1) SetValue(value string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Value"] = true
	h.Value = &value
}

func (h *HttpHeaders1) SetValueNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Value"] = true
	h.Value = nil
}

func (h HttpHeaders1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if h.touched["Name"] && h.Name == nil {
		data["name"] = nil
	} else if h.Name != nil {
		data["name"] = h.Name
	}

	if h.touched["Value"] && h.Value == nil {
		data["value"] = nil
	} else if h.Value != nil {
		data["value"] = h.Value
	}

	return json.Marshal(data)
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
