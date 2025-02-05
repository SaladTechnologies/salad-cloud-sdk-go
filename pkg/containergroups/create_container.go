package containergroups

import (
	"encoding/json"
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
	touched                map[string]bool
}

func (c *CreateContainer) GetImage() *string {
	if c == nil {
		return nil
	}
	return c.Image
}

func (c *CreateContainer) SetImage(image string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Image"] = true
	c.Image = &image
}

func (c *CreateContainer) SetImageNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Image"] = true
	c.Image = nil
}

func (c *CreateContainer) GetResources() *shared.ContainerResourceRequirements {
	if c == nil {
		return nil
	}
	return c.Resources
}

func (c *CreateContainer) SetResources(resources shared.ContainerResourceRequirements) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Resources"] = true
	c.Resources = &resources
}

func (c *CreateContainer) SetResourcesNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Resources"] = true
	c.Resources = nil
}

func (c *CreateContainer) GetCommand() []string {
	if c == nil {
		return nil
	}
	return c.Command
}

func (c *CreateContainer) SetCommand(command []string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Command"] = true
	c.Command = command
}

func (c *CreateContainer) SetCommandNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Command"] = true
	c.Command = nil
}

func (c *CreateContainer) GetPriority() *shared.ContainerGroupPriority {
	if c == nil {
		return nil
	}
	return c.Priority
}

func (c *CreateContainer) SetPriority(priority shared.ContainerGroupPriority) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Priority"] = true
	c.Priority = &priority
}

func (c *CreateContainer) SetPriorityNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Priority"] = true
	c.Priority = nil
}

func (c *CreateContainer) GetEnvironmentVariables() map[string]*string {
	if c == nil {
		return nil
	}
	return c.EnvironmentVariables
}

func (c *CreateContainer) SetEnvironmentVariables(environmentVariables map[string]*string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EnvironmentVariables"] = true
	c.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (c *CreateContainer) SetEnvironmentVariablesNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["EnvironmentVariables"] = true
	c.EnvironmentVariables = nil
}

func (c *CreateContainer) GetLogging() *CreateContainerLogging {
	if c == nil {
		return nil
	}
	return c.Logging
}

func (c *CreateContainer) SetLogging(logging CreateContainerLogging) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Logging"] = true
	c.Logging = &logging
}

func (c *CreateContainer) SetLoggingNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Logging"] = true
	c.Logging = nil
}

func (c *CreateContainer) GetRegistryAuthentication() *CreateContainerRegistryAuthentication {
	if c == nil {
		return nil
	}
	return c.RegistryAuthentication
}

func (c *CreateContainer) SetRegistryAuthentication(registryAuthentication CreateContainerRegistryAuthentication) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["RegistryAuthentication"] = true
	c.RegistryAuthentication = &registryAuthentication
}

func (c *CreateContainer) SetRegistryAuthenticationNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["RegistryAuthentication"] = true
	c.RegistryAuthentication = nil
}

func (c CreateContainer) MarshalJSON() ([]byte, error) {
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

	if c.touched["RegistryAuthentication"] && c.RegistryAuthentication == nil {
		data["registry_authentication"] = nil
	} else if c.RegistryAuthentication != nil {
		data["registry_authentication"] = c.RegistryAuthentication
	}

	return json.Marshal(data)
}

func (c CreateContainer) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateContainer to string"
	}
	return string(jsonData)
}

type CreateContainerLogging struct {
	Axiom    *LoggingAxiom2    `json:"axiom,omitempty"`
	Datadog  *LoggingDatadog2  `json:"datadog,omitempty"`
	NewRelic *LoggingNewRelic2 `json:"new_relic,omitempty"`
	Splunk   *LoggingSplunk2   `json:"splunk,omitempty"`
	Tcp      *LoggingTcp2      `json:"tcp,omitempty"`
	Http     *LoggingHttp2     `json:"http,omitempty"`
	touched  map[string]bool
}

func (c *CreateContainerLogging) GetAxiom() *LoggingAxiom2 {
	if c == nil {
		return nil
	}
	return c.Axiom
}

func (c *CreateContainerLogging) SetAxiom(axiom LoggingAxiom2) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Axiom"] = true
	c.Axiom = &axiom
}

func (c *CreateContainerLogging) SetAxiomNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Axiom"] = true
	c.Axiom = nil
}

func (c *CreateContainerLogging) GetDatadog() *LoggingDatadog2 {
	if c == nil {
		return nil
	}
	return c.Datadog
}

func (c *CreateContainerLogging) SetDatadog(datadog LoggingDatadog2) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Datadog"] = true
	c.Datadog = &datadog
}

func (c *CreateContainerLogging) SetDatadogNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Datadog"] = true
	c.Datadog = nil
}

func (c *CreateContainerLogging) GetNewRelic() *LoggingNewRelic2 {
	if c == nil {
		return nil
	}
	return c.NewRelic
}

func (c *CreateContainerLogging) SetNewRelic(newRelic LoggingNewRelic2) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["NewRelic"] = true
	c.NewRelic = &newRelic
}

func (c *CreateContainerLogging) SetNewRelicNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["NewRelic"] = true
	c.NewRelic = nil
}

func (c *CreateContainerLogging) GetSplunk() *LoggingSplunk2 {
	if c == nil {
		return nil
	}
	return c.Splunk
}

func (c *CreateContainerLogging) SetSplunk(splunk LoggingSplunk2) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Splunk"] = true
	c.Splunk = &splunk
}

func (c *CreateContainerLogging) SetSplunkNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Splunk"] = true
	c.Splunk = nil
}

func (c *CreateContainerLogging) GetTcp() *LoggingTcp2 {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *CreateContainerLogging) SetTcp(tcp LoggingTcp2) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Tcp"] = true
	c.Tcp = &tcp
}

func (c *CreateContainerLogging) SetTcpNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Tcp"] = true
	c.Tcp = nil
}

func (c *CreateContainerLogging) GetHttp() *LoggingHttp2 {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *CreateContainerLogging) SetHttp(http LoggingHttp2) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Http"] = true
	c.Http = &http
}

func (c *CreateContainerLogging) SetHttpNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Http"] = true
	c.Http = nil
}

func (c CreateContainerLogging) MarshalJSON() ([]byte, error) {
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
	touched  map[string]bool
}

func (l *LoggingAxiom2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingAxiom2) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingAxiom2) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingAxiom2) GetApiToken() *string {
	if l == nil {
		return nil
	}
	return l.ApiToken
}

func (l *LoggingAxiom2) SetApiToken(apiToken string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiToken"] = true
	l.ApiToken = &apiToken
}

func (l *LoggingAxiom2) SetApiTokenNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiToken"] = true
	l.ApiToken = nil
}

func (l *LoggingAxiom2) GetDataset() *string {
	if l == nil {
		return nil
	}
	return l.Dataset
}

func (l *LoggingAxiom2) SetDataset(dataset string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Dataset"] = true
	l.Dataset = &dataset
}

func (l *LoggingAxiom2) SetDatasetNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Dataset"] = true
	l.Dataset = nil
}

func (l LoggingAxiom2) MarshalJSON() ([]byte, error) {
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

func (l LoggingAxiom2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingAxiom2 to string"
	}
	return string(jsonData)
}

type LoggingDatadog2 struct {
	Host    *string        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiKey  *string        `json:"api_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Tags    []DatadogTags2 `json:"tags,omitempty" maxItems:"1000"`
	touched map[string]bool
}

func (l *LoggingDatadog2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingDatadog2) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingDatadog2) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingDatadog2) GetApiKey() *string {
	if l == nil {
		return nil
	}
	return l.ApiKey
}

func (l *LoggingDatadog2) SetApiKey(apiKey string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiKey"] = true
	l.ApiKey = &apiKey
}

func (l *LoggingDatadog2) SetApiKeyNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiKey"] = true
	l.ApiKey = nil
}

func (l *LoggingDatadog2) GetTags() []DatadogTags2 {
	if l == nil {
		return nil
	}
	return l.Tags
}

func (l *LoggingDatadog2) SetTags(tags []DatadogTags2) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Tags"] = true
	l.Tags = tags
}

func (l *LoggingDatadog2) SetTagsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Tags"] = true
	l.Tags = nil
}

func (l LoggingDatadog2) MarshalJSON() ([]byte, error) {
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

func (l LoggingDatadog2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingDatadog2 to string"
	}
	return string(jsonData)
}

type DatadogTags2 struct {
	Name    *string `json:"name,omitempty" required:"true"`
	Value   *string `json:"value,omitempty" required:"true"`
	touched map[string]bool
}

func (d *DatadogTags2) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DatadogTags2) SetName(name string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = &name
}

func (d *DatadogTags2) SetNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = nil
}

func (d *DatadogTags2) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

func (d *DatadogTags2) SetValue(value string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Value"] = true
	d.Value = &value
}

func (d *DatadogTags2) SetValueNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Value"] = true
	d.Value = nil
}

func (d DatadogTags2) MarshalJSON() ([]byte, error) {
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
	touched      map[string]bool
}

func (l *LoggingNewRelic2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingNewRelic2) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingNewRelic2) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingNewRelic2) GetIngestionKey() *string {
	if l == nil {
		return nil
	}
	return l.IngestionKey
}

func (l *LoggingNewRelic2) SetIngestionKey(ingestionKey string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["IngestionKey"] = true
	l.IngestionKey = &ingestionKey
}

func (l *LoggingNewRelic2) SetIngestionKeyNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["IngestionKey"] = true
	l.IngestionKey = nil
}

func (l LoggingNewRelic2) MarshalJSON() ([]byte, error) {
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

func (l LoggingNewRelic2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingNewRelic2 to string"
	}
	return string(jsonData)
}

type LoggingSplunk2 struct {
	Host    *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Token   *string `json:"token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	touched map[string]bool
}

func (l *LoggingSplunk2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingSplunk2) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingSplunk2) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingSplunk2) GetToken() *string {
	if l == nil {
		return nil
	}
	return l.Token
}

func (l *LoggingSplunk2) SetToken(token string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Token"] = true
	l.Token = &token
}

func (l *LoggingSplunk2) SetTokenNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Token"] = true
	l.Token = nil
}

func (l LoggingSplunk2) MarshalJSON() ([]byte, error) {
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

func (l LoggingSplunk2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingSplunk2 to string"
	}
	return string(jsonData)
}

type LoggingTcp2 struct {
	Host    *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port    *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	touched map[string]bool
}

func (l *LoggingTcp2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingTcp2) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingTcp2) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingTcp2) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingTcp2) SetPort(port int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = &port
}

func (l *LoggingTcp2) SetPortNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = nil
}

func (l LoggingTcp2) MarshalJSON() ([]byte, error) {
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

func (l LoggingTcp2) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingTcp2 to string"
	}
	return string(jsonData)
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
	touched     map[string]bool
}

func (l *LoggingHttp2) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingHttp2) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingHttp2) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingHttp2) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingHttp2) SetPort(port int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = &port
}

func (l *LoggingHttp2) SetPortNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = nil
}

func (l *LoggingHttp2) GetUser() *string {
	if l == nil {
		return nil
	}
	return l.User
}

func (l *LoggingHttp2) SetUser(user string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["User"] = true
	l.User = &user
}

func (l *LoggingHttp2) SetUserNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["User"] = true
	l.User = nil
}

func (l *LoggingHttp2) GetPassword() *string {
	if l == nil {
		return nil
	}
	return l.Password
}

func (l *LoggingHttp2) SetPassword(password string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Password"] = true
	l.Password = &password
}

func (l *LoggingHttp2) SetPasswordNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Password"] = true
	l.Password = nil
}

func (l *LoggingHttp2) GetPath() *string {
	if l == nil {
		return nil
	}
	return l.Path
}

func (l *LoggingHttp2) SetPath(path string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Path"] = true
	l.Path = &path
}

func (l *LoggingHttp2) SetPathNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Path"] = true
	l.Path = nil
}

func (l *LoggingHttp2) GetFormat() *HttpFormat2 {
	if l == nil {
		return nil
	}
	return l.Format
}

func (l *LoggingHttp2) SetFormat(format HttpFormat2) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Format"] = true
	l.Format = &format
}

func (l *LoggingHttp2) SetFormatNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Format"] = true
	l.Format = nil
}

func (l *LoggingHttp2) GetHeaders() []HttpHeaders3 {
	if l == nil {
		return nil
	}
	return l.Headers
}

func (l *LoggingHttp2) SetHeaders(headers []HttpHeaders3) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Headers"] = true
	l.Headers = headers
}

func (l *LoggingHttp2) SetHeadersNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Headers"] = true
	l.Headers = nil
}

func (l *LoggingHttp2) GetCompression() *HttpCompression2 {
	if l == nil {
		return nil
	}
	return l.Compression
}

func (l *LoggingHttp2) SetCompression(compression HttpCompression2) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Compression"] = true
	l.Compression = &compression
}

func (l *LoggingHttp2) SetCompressionNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Compression"] = true
	l.Compression = nil
}

func (l LoggingHttp2) MarshalJSON() ([]byte, error) {
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
	Name    *string `json:"name,omitempty" required:"true"`
	Value   *string `json:"value,omitempty" required:"true"`
	touched map[string]bool
}

func (h *HttpHeaders3) GetName() *string {
	if h == nil {
		return nil
	}
	return h.Name
}

func (h *HttpHeaders3) SetName(name string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Name"] = true
	h.Name = &name
}

func (h *HttpHeaders3) SetNameNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Name"] = true
	h.Name = nil
}

func (h *HttpHeaders3) GetValue() *string {
	if h == nil {
		return nil
	}
	return h.Value
}

func (h *HttpHeaders3) SetValue(value string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Value"] = true
	h.Value = &value
}

func (h *HttpHeaders3) SetValueNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Value"] = true
	h.Value = nil
}

func (h HttpHeaders3) MarshalJSON() ([]byte, error) {
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
	Basic     *RegistryAuthenticationBasic1     `json:"basic,omitempty"`
	GcpGcr    *RegistryAuthenticationGcpGcr1    `json:"gcp_gcr,omitempty"`
	AwsEcr    *RegistryAuthenticationAwsEcr1    `json:"aws_ecr,omitempty"`
	DockerHub *RegistryAuthenticationDockerHub1 `json:"docker_hub,omitempty"`
	GcpGar    *RegistryAuthenticationGcpGar1    `json:"gcp_gar,omitempty"`
	touched   map[string]bool
}

func (c *CreateContainerRegistryAuthentication) GetBasic() *RegistryAuthenticationBasic1 {
	if c == nil {
		return nil
	}
	return c.Basic
}

func (c *CreateContainerRegistryAuthentication) SetBasic(basic RegistryAuthenticationBasic1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Basic"] = true
	c.Basic = &basic
}

func (c *CreateContainerRegistryAuthentication) SetBasicNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Basic"] = true
	c.Basic = nil
}

func (c *CreateContainerRegistryAuthentication) GetGcpGcr() *RegistryAuthenticationGcpGcr1 {
	if c == nil {
		return nil
	}
	return c.GcpGcr
}

func (c *CreateContainerRegistryAuthentication) SetGcpGcr(gcpGcr RegistryAuthenticationGcpGcr1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["GcpGcr"] = true
	c.GcpGcr = &gcpGcr
}

func (c *CreateContainerRegistryAuthentication) SetGcpGcrNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["GcpGcr"] = true
	c.GcpGcr = nil
}

func (c *CreateContainerRegistryAuthentication) GetAwsEcr() *RegistryAuthenticationAwsEcr1 {
	if c == nil {
		return nil
	}
	return c.AwsEcr
}

func (c *CreateContainerRegistryAuthentication) SetAwsEcr(awsEcr RegistryAuthenticationAwsEcr1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AwsEcr"] = true
	c.AwsEcr = &awsEcr
}

func (c *CreateContainerRegistryAuthentication) SetAwsEcrNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AwsEcr"] = true
	c.AwsEcr = nil
}

func (c *CreateContainerRegistryAuthentication) GetDockerHub() *RegistryAuthenticationDockerHub1 {
	if c == nil {
		return nil
	}
	return c.DockerHub
}

func (c *CreateContainerRegistryAuthentication) SetDockerHub(dockerHub RegistryAuthenticationDockerHub1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["DockerHub"] = true
	c.DockerHub = &dockerHub
}

func (c *CreateContainerRegistryAuthentication) SetDockerHubNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["DockerHub"] = true
	c.DockerHub = nil
}

func (c *CreateContainerRegistryAuthentication) GetGcpGar() *RegistryAuthenticationGcpGar1 {
	if c == nil {
		return nil
	}
	return c.GcpGar
}

func (c *CreateContainerRegistryAuthentication) SetGcpGar(gcpGar RegistryAuthenticationGcpGar1) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["GcpGar"] = true
	c.GcpGar = &gcpGar
}

func (c *CreateContainerRegistryAuthentication) SetGcpGarNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["GcpGar"] = true
	c.GcpGar = nil
}

func (c CreateContainerRegistryAuthentication) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Basic"] && c.Basic == nil {
		data["basic"] = nil
	} else if c.Basic != nil {
		data["basic"] = c.Basic
	}

	if c.touched["GcpGcr"] && c.GcpGcr == nil {
		data["gcp_gcr"] = nil
	} else if c.GcpGcr != nil {
		data["gcp_gcr"] = c.GcpGcr
	}

	if c.touched["AwsEcr"] && c.AwsEcr == nil {
		data["aws_ecr"] = nil
	} else if c.AwsEcr != nil {
		data["aws_ecr"] = c.AwsEcr
	}

	if c.touched["DockerHub"] && c.DockerHub == nil {
		data["docker_hub"] = nil
	} else if c.DockerHub != nil {
		data["docker_hub"] = c.DockerHub
	}

	if c.touched["GcpGar"] && c.GcpGar == nil {
		data["gcp_gar"] = nil
	} else if c.GcpGar != nil {
		data["gcp_gar"] = c.GcpGar
	}

	return json.Marshal(data)
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
	touched  map[string]bool
}

func (r *RegistryAuthenticationBasic1) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationBasic1) SetUsername(username string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Username"] = true
	r.Username = &username
}

func (r *RegistryAuthenticationBasic1) SetUsernameNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Username"] = true
	r.Username = nil
}

func (r *RegistryAuthenticationBasic1) GetPassword() *string {
	if r == nil {
		return nil
	}
	return r.Password
}

func (r *RegistryAuthenticationBasic1) SetPassword(password string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Password"] = true
	r.Password = &password
}

func (r *RegistryAuthenticationBasic1) SetPasswordNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Password"] = true
	r.Password = nil
}

func (r RegistryAuthenticationBasic1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["Username"] && r.Username == nil {
		data["username"] = nil
	} else if r.Username != nil {
		data["username"] = r.Username
	}

	if r.touched["Password"] && r.Password == nil {
		data["password"] = nil
	} else if r.Password != nil {
		data["password"] = r.Password
	}

	return json.Marshal(data)
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
	touched    map[string]bool
}

func (r *RegistryAuthenticationGcpGcr1) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}

func (r *RegistryAuthenticationGcpGcr1) SetServiceKey(serviceKey string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["ServiceKey"] = true
	r.ServiceKey = &serviceKey
}

func (r *RegistryAuthenticationGcpGcr1) SetServiceKeyNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["ServiceKey"] = true
	r.ServiceKey = nil
}

func (r RegistryAuthenticationGcpGcr1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["ServiceKey"] && r.ServiceKey == nil {
		data["service_key"] = nil
	} else if r.ServiceKey != nil {
		data["service_key"] = r.ServiceKey
	}

	return json.Marshal(data)
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
	touched         map[string]bool
}

func (r *RegistryAuthenticationAwsEcr1) GetAccessKeyId() *string {
	if r == nil {
		return nil
	}
	return r.AccessKeyId
}

func (r *RegistryAuthenticationAwsEcr1) SetAccessKeyId(accessKeyId string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["AccessKeyId"] = true
	r.AccessKeyId = &accessKeyId
}

func (r *RegistryAuthenticationAwsEcr1) SetAccessKeyIdNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["AccessKeyId"] = true
	r.AccessKeyId = nil
}

func (r *RegistryAuthenticationAwsEcr1) GetSecretAccessKey() *string {
	if r == nil {
		return nil
	}
	return r.SecretAccessKey
}

func (r *RegistryAuthenticationAwsEcr1) SetSecretAccessKey(secretAccessKey string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["SecretAccessKey"] = true
	r.SecretAccessKey = &secretAccessKey
}

func (r *RegistryAuthenticationAwsEcr1) SetSecretAccessKeyNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["SecretAccessKey"] = true
	r.SecretAccessKey = nil
}

func (r RegistryAuthenticationAwsEcr1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["AccessKeyId"] && r.AccessKeyId == nil {
		data["access_key_id"] = nil
	} else if r.AccessKeyId != nil {
		data["access_key_id"] = r.AccessKeyId
	}

	if r.touched["SecretAccessKey"] && r.SecretAccessKey == nil {
		data["secret_access_key"] = nil
	} else if r.SecretAccessKey != nil {
		data["secret_access_key"] = r.SecretAccessKey
	}

	return json.Marshal(data)
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
	touched             map[string]bool
}

func (r *RegistryAuthenticationDockerHub1) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationDockerHub1) SetUsername(username string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Username"] = true
	r.Username = &username
}

func (r *RegistryAuthenticationDockerHub1) SetUsernameNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Username"] = true
	r.Username = nil
}

func (r *RegistryAuthenticationDockerHub1) GetPersonalAccessToken() *string {
	if r == nil {
		return nil
	}
	return r.PersonalAccessToken
}

func (r *RegistryAuthenticationDockerHub1) SetPersonalAccessToken(personalAccessToken string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["PersonalAccessToken"] = true
	r.PersonalAccessToken = &personalAccessToken
}

func (r *RegistryAuthenticationDockerHub1) SetPersonalAccessTokenNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["PersonalAccessToken"] = true
	r.PersonalAccessToken = nil
}

func (r RegistryAuthenticationDockerHub1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["Username"] && r.Username == nil {
		data["username"] = nil
	} else if r.Username != nil {
		data["username"] = r.Username
	}

	if r.touched["PersonalAccessToken"] && r.PersonalAccessToken == nil {
		data["personal_access_token"] = nil
	} else if r.PersonalAccessToken != nil {
		data["personal_access_token"] = r.PersonalAccessToken
	}

	return json.Marshal(data)
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
	touched    map[string]bool
}

func (r *RegistryAuthenticationGcpGar1) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}

func (r *RegistryAuthenticationGcpGar1) SetServiceKey(serviceKey string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["ServiceKey"] = true
	r.ServiceKey = &serviceKey
}

func (r *RegistryAuthenticationGcpGar1) SetServiceKeyNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["ServiceKey"] = true
	r.ServiceKey = nil
}

func (r RegistryAuthenticationGcpGar1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["ServiceKey"] && r.ServiceKey == nil {
		data["service_key"] = nil
	} else if r.ServiceKey != nil {
		data["service_key"] = r.ServiceKey
	}

	return json.Marshal(data)
}

func (r RegistryAuthenticationGcpGar1) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationGcpGar1 to string"
	}
	return string(jsonData)
}
