package containergroups

import (
	"encoding/json"
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
	touched                map[string]bool
}

func (u *UpdateContainer) GetImage() *string {
	if u == nil {
		return nil
	}
	return u.Image
}

func (u *UpdateContainer) SetImage(image string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Image"] = true
	u.Image = &image
}

func (u *UpdateContainer) SetImageNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Image"] = true
	u.Image = nil
}

func (u *UpdateContainer) GetResources() *Resources {
	if u == nil {
		return nil
	}
	return u.Resources
}

func (u *UpdateContainer) SetResources(resources Resources) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Resources"] = true
	u.Resources = &resources
}

func (u *UpdateContainer) SetResourcesNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Resources"] = true
	u.Resources = nil
}

func (u *UpdateContainer) GetCommand() []string {
	if u == nil {
		return nil
	}
	return u.Command
}

func (u *UpdateContainer) SetCommand(command []string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Command"] = true
	u.Command = command
}

func (u *UpdateContainer) SetCommandNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Command"] = true
	u.Command = nil
}

func (u *UpdateContainer) GetPriority() *shared.ContainerGroupPriority {
	if u == nil {
		return nil
	}
	return u.Priority
}

func (u *UpdateContainer) SetPriority(priority shared.ContainerGroupPriority) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Priority"] = true
	u.Priority = &priority
}

func (u *UpdateContainer) SetPriorityNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Priority"] = true
	u.Priority = nil
}

func (u *UpdateContainer) GetEnvironmentVariables() map[string]*string {
	if u == nil {
		return nil
	}
	return u.EnvironmentVariables
}

func (u *UpdateContainer) SetEnvironmentVariables(environmentVariables map[string]*string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["EnvironmentVariables"] = true
	u.EnvironmentVariables = utils.CloneMap(environmentVariables)
}

func (u *UpdateContainer) SetEnvironmentVariablesNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["EnvironmentVariables"] = true
	u.EnvironmentVariables = nil
}

func (u *UpdateContainer) GetLogging() *UpdateContainerLogging {
	if u == nil {
		return nil
	}
	return u.Logging
}

func (u *UpdateContainer) SetLogging(logging UpdateContainerLogging) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Logging"] = true
	u.Logging = &logging
}

func (u *UpdateContainer) SetLoggingNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Logging"] = true
	u.Logging = nil
}

func (u *UpdateContainer) GetRegistryAuthentication() *UpdateContainerRegistryAuthentication {
	if u == nil {
		return nil
	}
	return u.RegistryAuthentication
}

func (u *UpdateContainer) SetRegistryAuthentication(registryAuthentication UpdateContainerRegistryAuthentication) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["RegistryAuthentication"] = true
	u.RegistryAuthentication = &registryAuthentication
}

func (u *UpdateContainer) SetRegistryAuthenticationNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["RegistryAuthentication"] = true
	u.RegistryAuthentication = nil
}

func (u UpdateContainer) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["Image"] && u.Image == nil {
		data["image"] = nil
	} else if u.Image != nil {
		data["image"] = u.Image
	}

	if u.touched["Resources"] && u.Resources == nil {
		data["resources"] = nil
	} else if u.Resources != nil {
		data["resources"] = u.Resources
	}

	if u.touched["Command"] && u.Command == nil {
		data["command"] = nil
	} else if u.Command != nil {
		data["command"] = u.Command
	}

	if u.touched["Priority"] && u.Priority == nil {
		data["priority"] = nil
	} else if u.Priority != nil {
		data["priority"] = u.Priority
	}

	if u.touched["EnvironmentVariables"] && u.EnvironmentVariables == nil {
		data["environment_variables"] = nil
	} else if u.EnvironmentVariables != nil {
		data["environment_variables"] = u.EnvironmentVariables
	}

	if u.touched["Logging"] && u.Logging == nil {
		data["logging"] = nil
	} else if u.Logging != nil {
		data["logging"] = u.Logging
	}

	if u.touched["RegistryAuthentication"] && u.RegistryAuthentication == nil {
		data["registry_authentication"] = nil
	} else if u.RegistryAuthentication != nil {
		data["registry_authentication"] = u.RegistryAuthentication
	}

	return json.Marshal(data)
}

func (u UpdateContainer) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateContainer to string"
	}
	return string(jsonData)
}

type Resources struct {
	Cpu           *int64   `json:"cpu,omitempty" min:"1" max:"16"`
	Memory        *int64   `json:"memory,omitempty" min:"1024" max:"61440"`
	GpuClasses    []string `json:"gpu_classes,omitempty"`
	StorageAmount *int64   `json:"storage_amount,omitempty" min:"1073741824" max:"53687091200"`
	touched       map[string]bool
}

func (r *Resources) GetCpu() *int64 {
	if r == nil {
		return nil
	}
	return r.Cpu
}

func (r *Resources) SetCpu(cpu int64) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Cpu"] = true
	r.Cpu = &cpu
}

func (r *Resources) SetCpuNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Cpu"] = true
	r.Cpu = nil
}

func (r *Resources) GetMemory() *int64 {
	if r == nil {
		return nil
	}
	return r.Memory
}

func (r *Resources) SetMemory(memory int64) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Memory"] = true
	r.Memory = &memory
}

func (r *Resources) SetMemoryNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Memory"] = true
	r.Memory = nil
}

func (r *Resources) GetGpuClasses() []string {
	if r == nil {
		return nil
	}
	return r.GpuClasses
}

func (r *Resources) SetGpuClasses(gpuClasses []string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["GpuClasses"] = true
	r.GpuClasses = gpuClasses
}

func (r *Resources) SetGpuClassesNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["GpuClasses"] = true
	r.GpuClasses = nil
}

func (r *Resources) GetStorageAmount() *int64 {
	if r == nil {
		return nil
	}
	return r.StorageAmount
}

func (r *Resources) SetStorageAmount(storageAmount int64) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["StorageAmount"] = true
	r.StorageAmount = &storageAmount
}

func (r *Resources) SetStorageAmountNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["StorageAmount"] = true
	r.StorageAmount = nil
}

func (r Resources) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["Cpu"] && r.Cpu == nil {
		data["cpu"] = nil
	} else if r.Cpu != nil {
		data["cpu"] = r.Cpu
	}

	if r.touched["Memory"] && r.Memory == nil {
		data["memory"] = nil
	} else if r.Memory != nil {
		data["memory"] = r.Memory
	}

	if r.touched["GpuClasses"] && r.GpuClasses == nil {
		data["gpu_classes"] = nil
	} else if r.GpuClasses != nil {
		data["gpu_classes"] = r.GpuClasses
	}

	if r.touched["StorageAmount"] && r.StorageAmount == nil {
		data["storage_amount"] = nil
	} else if r.StorageAmount != nil {
		data["storage_amount"] = r.StorageAmount
	}

	return json.Marshal(data)
}

func (r Resources) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: Resources to string"
	}
	return string(jsonData)
}

type UpdateContainerLogging struct {
	Axiom    *LoggingAxiom3    `json:"axiom,omitempty"`
	Datadog  *LoggingDatadog3  `json:"datadog,omitempty"`
	NewRelic *LoggingNewRelic3 `json:"new_relic,omitempty"`
	Splunk   *LoggingSplunk3   `json:"splunk,omitempty"`
	Tcp      *LoggingTcp3      `json:"tcp,omitempty"`
	Http     *LoggingHttp3     `json:"http,omitempty"`
	touched  map[string]bool
}

func (u *UpdateContainerLogging) GetAxiom() *LoggingAxiom3 {
	if u == nil {
		return nil
	}
	return u.Axiom
}

func (u *UpdateContainerLogging) SetAxiom(axiom LoggingAxiom3) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Axiom"] = true
	u.Axiom = &axiom
}

func (u *UpdateContainerLogging) SetAxiomNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Axiom"] = true
	u.Axiom = nil
}

func (u *UpdateContainerLogging) GetDatadog() *LoggingDatadog3 {
	if u == nil {
		return nil
	}
	return u.Datadog
}

func (u *UpdateContainerLogging) SetDatadog(datadog LoggingDatadog3) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Datadog"] = true
	u.Datadog = &datadog
}

func (u *UpdateContainerLogging) SetDatadogNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Datadog"] = true
	u.Datadog = nil
}

func (u *UpdateContainerLogging) GetNewRelic() *LoggingNewRelic3 {
	if u == nil {
		return nil
	}
	return u.NewRelic
}

func (u *UpdateContainerLogging) SetNewRelic(newRelic LoggingNewRelic3) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["NewRelic"] = true
	u.NewRelic = &newRelic
}

func (u *UpdateContainerLogging) SetNewRelicNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["NewRelic"] = true
	u.NewRelic = nil
}

func (u *UpdateContainerLogging) GetSplunk() *LoggingSplunk3 {
	if u == nil {
		return nil
	}
	return u.Splunk
}

func (u *UpdateContainerLogging) SetSplunk(splunk LoggingSplunk3) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Splunk"] = true
	u.Splunk = &splunk
}

func (u *UpdateContainerLogging) SetSplunkNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Splunk"] = true
	u.Splunk = nil
}

func (u *UpdateContainerLogging) GetTcp() *LoggingTcp3 {
	if u == nil {
		return nil
	}
	return u.Tcp
}

func (u *UpdateContainerLogging) SetTcp(tcp LoggingTcp3) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Tcp"] = true
	u.Tcp = &tcp
}

func (u *UpdateContainerLogging) SetTcpNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Tcp"] = true
	u.Tcp = nil
}

func (u *UpdateContainerLogging) GetHttp() *LoggingHttp3 {
	if u == nil {
		return nil
	}
	return u.Http
}

func (u *UpdateContainerLogging) SetHttp(http LoggingHttp3) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Http"] = true
	u.Http = &http
}

func (u *UpdateContainerLogging) SetHttpNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Http"] = true
	u.Http = nil
}

func (u UpdateContainerLogging) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["Axiom"] && u.Axiom == nil {
		data["axiom"] = nil
	} else if u.Axiom != nil {
		data["axiom"] = u.Axiom
	}

	if u.touched["Datadog"] && u.Datadog == nil {
		data["datadog"] = nil
	} else if u.Datadog != nil {
		data["datadog"] = u.Datadog
	}

	if u.touched["NewRelic"] && u.NewRelic == nil {
		data["new_relic"] = nil
	} else if u.NewRelic != nil {
		data["new_relic"] = u.NewRelic
	}

	if u.touched["Splunk"] && u.Splunk == nil {
		data["splunk"] = nil
	} else if u.Splunk != nil {
		data["splunk"] = u.Splunk
	}

	if u.touched["Tcp"] && u.Tcp == nil {
		data["tcp"] = nil
	} else if u.Tcp != nil {
		data["tcp"] = u.Tcp
	}

	if u.touched["Http"] && u.Http == nil {
		data["http"] = nil
	} else if u.Http != nil {
		data["http"] = u.Http
	}

	return json.Marshal(data)
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
	touched  map[string]bool
}

func (l *LoggingAxiom3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingAxiom3) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingAxiom3) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingAxiom3) GetApiToken() *string {
	if l == nil {
		return nil
	}
	return l.ApiToken
}

func (l *LoggingAxiom3) SetApiToken(apiToken string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiToken"] = true
	l.ApiToken = &apiToken
}

func (l *LoggingAxiom3) SetApiTokenNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiToken"] = true
	l.ApiToken = nil
}

func (l *LoggingAxiom3) GetDataset() *string {
	if l == nil {
		return nil
	}
	return l.Dataset
}

func (l *LoggingAxiom3) SetDataset(dataset string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Dataset"] = true
	l.Dataset = &dataset
}

func (l *LoggingAxiom3) SetDatasetNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Dataset"] = true
	l.Dataset = nil
}

func (l LoggingAxiom3) MarshalJSON() ([]byte, error) {
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

func (l LoggingAxiom3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingAxiom3 to string"
	}
	return string(jsonData)
}

type LoggingDatadog3 struct {
	Host    *string        `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	ApiKey  *string        `json:"api_key,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Tags    []DatadogTags3 `json:"tags,omitempty" maxItems:"1000"`
	touched map[string]bool
}

func (l *LoggingDatadog3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingDatadog3) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingDatadog3) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingDatadog3) GetApiKey() *string {
	if l == nil {
		return nil
	}
	return l.ApiKey
}

func (l *LoggingDatadog3) SetApiKey(apiKey string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiKey"] = true
	l.ApiKey = &apiKey
}

func (l *LoggingDatadog3) SetApiKeyNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["ApiKey"] = true
	l.ApiKey = nil
}

func (l *LoggingDatadog3) GetTags() []DatadogTags3 {
	if l == nil {
		return nil
	}
	return l.Tags
}

func (l *LoggingDatadog3) SetTags(tags []DatadogTags3) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Tags"] = true
	l.Tags = tags
}

func (l *LoggingDatadog3) SetTagsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Tags"] = true
	l.Tags = nil
}

func (l LoggingDatadog3) MarshalJSON() ([]byte, error) {
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

func (l LoggingDatadog3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingDatadog3 to string"
	}
	return string(jsonData)
}

type DatadogTags3 struct {
	Name    *string `json:"name,omitempty" required:"true"`
	Value   *string `json:"value,omitempty" required:"true"`
	touched map[string]bool
}

func (d *DatadogTags3) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DatadogTags3) SetName(name string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = &name
}

func (d *DatadogTags3) SetNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = nil
}

func (d *DatadogTags3) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

func (d *DatadogTags3) SetValue(value string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Value"] = true
	d.Value = &value
}

func (d *DatadogTags3) SetValueNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Value"] = true
	d.Value = nil
}

func (d DatadogTags3) MarshalJSON() ([]byte, error) {
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
	touched      map[string]bool
}

func (l *LoggingNewRelic3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingNewRelic3) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingNewRelic3) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingNewRelic3) GetIngestionKey() *string {
	if l == nil {
		return nil
	}
	return l.IngestionKey
}

func (l *LoggingNewRelic3) SetIngestionKey(ingestionKey string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["IngestionKey"] = true
	l.IngestionKey = &ingestionKey
}

func (l *LoggingNewRelic3) SetIngestionKeyNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["IngestionKey"] = true
	l.IngestionKey = nil
}

func (l LoggingNewRelic3) MarshalJSON() ([]byte, error) {
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

func (l LoggingNewRelic3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingNewRelic3 to string"
	}
	return string(jsonData)
}

type LoggingSplunk3 struct {
	Host    *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Token   *string `json:"token,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	touched map[string]bool
}

func (l *LoggingSplunk3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingSplunk3) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingSplunk3) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingSplunk3) GetToken() *string {
	if l == nil {
		return nil
	}
	return l.Token
}

func (l *LoggingSplunk3) SetToken(token string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Token"] = true
	l.Token = &token
}

func (l *LoggingSplunk3) SetTokenNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Token"] = true
	l.Token = nil
}

func (l LoggingSplunk3) MarshalJSON() ([]byte, error) {
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

func (l LoggingSplunk3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingSplunk3 to string"
	}
	return string(jsonData)
}

type LoggingTcp3 struct {
	Host    *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1"`
	Port    *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	touched map[string]bool
}

func (l *LoggingTcp3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingTcp3) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingTcp3) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingTcp3) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingTcp3) SetPort(port int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = &port
}

func (l *LoggingTcp3) SetPortNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = nil
}

func (l LoggingTcp3) MarshalJSON() ([]byte, error) {
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

func (l LoggingTcp3) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoggingTcp3 to string"
	}
	return string(jsonData)
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
	touched     map[string]bool
}

func (l *LoggingHttp3) GetHost() *string {
	if l == nil {
		return nil
	}
	return l.Host
}

func (l *LoggingHttp3) SetHost(host string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = &host
}

func (l *LoggingHttp3) SetHostNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Host"] = true
	l.Host = nil
}

func (l *LoggingHttp3) GetPort() *int64 {
	if l == nil {
		return nil
	}
	return l.Port
}

func (l *LoggingHttp3) SetPort(port int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = &port
}

func (l *LoggingHttp3) SetPortNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Port"] = true
	l.Port = nil
}

func (l *LoggingHttp3) GetUser() *string {
	if l == nil {
		return nil
	}
	return l.User
}

func (l *LoggingHttp3) SetUser(user string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["User"] = true
	l.User = &user
}

func (l *LoggingHttp3) SetUserNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["User"] = true
	l.User = nil
}

func (l *LoggingHttp3) GetPassword() *string {
	if l == nil {
		return nil
	}
	return l.Password
}

func (l *LoggingHttp3) SetPassword(password string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Password"] = true
	l.Password = &password
}

func (l *LoggingHttp3) SetPasswordNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Password"] = true
	l.Password = nil
}

func (l *LoggingHttp3) GetPath() *string {
	if l == nil {
		return nil
	}
	return l.Path
}

func (l *LoggingHttp3) SetPath(path string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Path"] = true
	l.Path = &path
}

func (l *LoggingHttp3) SetPathNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Path"] = true
	l.Path = nil
}

func (l *LoggingHttp3) GetFormat() *HttpFormat3 {
	if l == nil {
		return nil
	}
	return l.Format
}

func (l *LoggingHttp3) SetFormat(format HttpFormat3) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Format"] = true
	l.Format = &format
}

func (l *LoggingHttp3) SetFormatNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Format"] = true
	l.Format = nil
}

func (l *LoggingHttp3) GetHeaders() []HttpHeaders4 {
	if l == nil {
		return nil
	}
	return l.Headers
}

func (l *LoggingHttp3) SetHeaders(headers []HttpHeaders4) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Headers"] = true
	l.Headers = headers
}

func (l *LoggingHttp3) SetHeadersNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Headers"] = true
	l.Headers = nil
}

func (l *LoggingHttp3) GetCompression() *HttpCompression3 {
	if l == nil {
		return nil
	}
	return l.Compression
}

func (l *LoggingHttp3) SetCompression(compression HttpCompression3) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Compression"] = true
	l.Compression = &compression
}

func (l *LoggingHttp3) SetCompressionNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Compression"] = true
	l.Compression = nil
}

func (l LoggingHttp3) MarshalJSON() ([]byte, error) {
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
	Name    *string `json:"name,omitempty" required:"true"`
	Value   *string `json:"value,omitempty" required:"true"`
	touched map[string]bool
}

func (h *HttpHeaders4) GetName() *string {
	if h == nil {
		return nil
	}
	return h.Name
}

func (h *HttpHeaders4) SetName(name string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Name"] = true
	h.Name = &name
}

func (h *HttpHeaders4) SetNameNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Name"] = true
	h.Name = nil
}

func (h *HttpHeaders4) GetValue() *string {
	if h == nil {
		return nil
	}
	return h.Value
}

func (h *HttpHeaders4) SetValue(value string) {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Value"] = true
	h.Value = &value
}

func (h *HttpHeaders4) SetValueNil() {
	if h.touched == nil {
		h.touched = map[string]bool{}
	}
	h.touched["Value"] = true
	h.Value = nil
}

func (h HttpHeaders4) MarshalJSON() ([]byte, error) {
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
	Basic     *RegistryAuthenticationBasic2     `json:"basic,omitempty"`
	GcpGcr    *RegistryAuthenticationGcpGcr2    `json:"gcp_gcr,omitempty"`
	AwsEcr    *RegistryAuthenticationAwsEcr2    `json:"aws_ecr,omitempty"`
	DockerHub *RegistryAuthenticationDockerHub2 `json:"docker_hub,omitempty"`
	GcpGar    *RegistryAuthenticationGcpGar2    `json:"gcp_gar,omitempty"`
	touched   map[string]bool
}

func (u *UpdateContainerRegistryAuthentication) GetBasic() *RegistryAuthenticationBasic2 {
	if u == nil {
		return nil
	}
	return u.Basic
}

func (u *UpdateContainerRegistryAuthentication) SetBasic(basic RegistryAuthenticationBasic2) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Basic"] = true
	u.Basic = &basic
}

func (u *UpdateContainerRegistryAuthentication) SetBasicNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Basic"] = true
	u.Basic = nil
}

func (u *UpdateContainerRegistryAuthentication) GetGcpGcr() *RegistryAuthenticationGcpGcr2 {
	if u == nil {
		return nil
	}
	return u.GcpGcr
}

func (u *UpdateContainerRegistryAuthentication) SetGcpGcr(gcpGcr RegistryAuthenticationGcpGcr2) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["GcpGcr"] = true
	u.GcpGcr = &gcpGcr
}

func (u *UpdateContainerRegistryAuthentication) SetGcpGcrNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["GcpGcr"] = true
	u.GcpGcr = nil
}

func (u *UpdateContainerRegistryAuthentication) GetAwsEcr() *RegistryAuthenticationAwsEcr2 {
	if u == nil {
		return nil
	}
	return u.AwsEcr
}

func (u *UpdateContainerRegistryAuthentication) SetAwsEcr(awsEcr RegistryAuthenticationAwsEcr2) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["AwsEcr"] = true
	u.AwsEcr = &awsEcr
}

func (u *UpdateContainerRegistryAuthentication) SetAwsEcrNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["AwsEcr"] = true
	u.AwsEcr = nil
}

func (u *UpdateContainerRegistryAuthentication) GetDockerHub() *RegistryAuthenticationDockerHub2 {
	if u == nil {
		return nil
	}
	return u.DockerHub
}

func (u *UpdateContainerRegistryAuthentication) SetDockerHub(dockerHub RegistryAuthenticationDockerHub2) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["DockerHub"] = true
	u.DockerHub = &dockerHub
}

func (u *UpdateContainerRegistryAuthentication) SetDockerHubNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["DockerHub"] = true
	u.DockerHub = nil
}

func (u *UpdateContainerRegistryAuthentication) GetGcpGar() *RegistryAuthenticationGcpGar2 {
	if u == nil {
		return nil
	}
	return u.GcpGar
}

func (u *UpdateContainerRegistryAuthentication) SetGcpGar(gcpGar RegistryAuthenticationGcpGar2) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["GcpGar"] = true
	u.GcpGar = &gcpGar
}

func (u *UpdateContainerRegistryAuthentication) SetGcpGarNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["GcpGar"] = true
	u.GcpGar = nil
}

func (u UpdateContainerRegistryAuthentication) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["Basic"] && u.Basic == nil {
		data["basic"] = nil
	} else if u.Basic != nil {
		data["basic"] = u.Basic
	}

	if u.touched["GcpGcr"] && u.GcpGcr == nil {
		data["gcp_gcr"] = nil
	} else if u.GcpGcr != nil {
		data["gcp_gcr"] = u.GcpGcr
	}

	if u.touched["AwsEcr"] && u.AwsEcr == nil {
		data["aws_ecr"] = nil
	} else if u.AwsEcr != nil {
		data["aws_ecr"] = u.AwsEcr
	}

	if u.touched["DockerHub"] && u.DockerHub == nil {
		data["docker_hub"] = nil
	} else if u.DockerHub != nil {
		data["docker_hub"] = u.DockerHub
	}

	if u.touched["GcpGar"] && u.GcpGar == nil {
		data["gcp_gar"] = nil
	} else if u.GcpGar != nil {
		data["gcp_gar"] = u.GcpGar
	}

	return json.Marshal(data)
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
	touched  map[string]bool
}

func (r *RegistryAuthenticationBasic2) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationBasic2) SetUsername(username string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Username"] = true
	r.Username = &username
}

func (r *RegistryAuthenticationBasic2) SetUsernameNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Username"] = true
	r.Username = nil
}

func (r *RegistryAuthenticationBasic2) GetPassword() *string {
	if r == nil {
		return nil
	}
	return r.Password
}

func (r *RegistryAuthenticationBasic2) SetPassword(password string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Password"] = true
	r.Password = &password
}

func (r *RegistryAuthenticationBasic2) SetPasswordNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Password"] = true
	r.Password = nil
}

func (r RegistryAuthenticationBasic2) MarshalJSON() ([]byte, error) {
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

func (r RegistryAuthenticationBasic2) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationBasic2 to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationGcpGcr2 struct {
	ServiceKey *string `json:"service_key,omitempty" required:"true"`
	touched    map[string]bool
}

func (r *RegistryAuthenticationGcpGcr2) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}

func (r *RegistryAuthenticationGcpGcr2) SetServiceKey(serviceKey string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["ServiceKey"] = true
	r.ServiceKey = &serviceKey
}

func (r *RegistryAuthenticationGcpGcr2) SetServiceKeyNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["ServiceKey"] = true
	r.ServiceKey = nil
}

func (r RegistryAuthenticationGcpGcr2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["ServiceKey"] && r.ServiceKey == nil {
		data["service_key"] = nil
	} else if r.ServiceKey != nil {
		data["service_key"] = r.ServiceKey
	}

	return json.Marshal(data)
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
	touched         map[string]bool
}

func (r *RegistryAuthenticationAwsEcr2) GetAccessKeyId() *string {
	if r == nil {
		return nil
	}
	return r.AccessKeyId
}

func (r *RegistryAuthenticationAwsEcr2) SetAccessKeyId(accessKeyId string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["AccessKeyId"] = true
	r.AccessKeyId = &accessKeyId
}

func (r *RegistryAuthenticationAwsEcr2) SetAccessKeyIdNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["AccessKeyId"] = true
	r.AccessKeyId = nil
}

func (r *RegistryAuthenticationAwsEcr2) GetSecretAccessKey() *string {
	if r == nil {
		return nil
	}
	return r.SecretAccessKey
}

func (r *RegistryAuthenticationAwsEcr2) SetSecretAccessKey(secretAccessKey string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["SecretAccessKey"] = true
	r.SecretAccessKey = &secretAccessKey
}

func (r *RegistryAuthenticationAwsEcr2) SetSecretAccessKeyNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["SecretAccessKey"] = true
	r.SecretAccessKey = nil
}

func (r RegistryAuthenticationAwsEcr2) MarshalJSON() ([]byte, error) {
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
	touched             map[string]bool
}

func (r *RegistryAuthenticationDockerHub2) GetUsername() *string {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RegistryAuthenticationDockerHub2) SetUsername(username string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Username"] = true
	r.Username = &username
}

func (r *RegistryAuthenticationDockerHub2) SetUsernameNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Username"] = true
	r.Username = nil
}

func (r *RegistryAuthenticationDockerHub2) GetPersonalAccessToken() *string {
	if r == nil {
		return nil
	}
	return r.PersonalAccessToken
}

func (r *RegistryAuthenticationDockerHub2) SetPersonalAccessToken(personalAccessToken string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["PersonalAccessToken"] = true
	r.PersonalAccessToken = &personalAccessToken
}

func (r *RegistryAuthenticationDockerHub2) SetPersonalAccessTokenNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["PersonalAccessToken"] = true
	r.PersonalAccessToken = nil
}

func (r RegistryAuthenticationDockerHub2) MarshalJSON() ([]byte, error) {
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

func (r RegistryAuthenticationDockerHub2) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationDockerHub2 to string"
	}
	return string(jsonData)
}

type RegistryAuthenticationGcpGar2 struct {
	ServiceKey *string `json:"service_key,omitempty" required:"true"`
	touched    map[string]bool
}

func (r *RegistryAuthenticationGcpGar2) GetServiceKey() *string {
	if r == nil {
		return nil
	}
	return r.ServiceKey
}

func (r *RegistryAuthenticationGcpGar2) SetServiceKey(serviceKey string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["ServiceKey"] = true
	r.ServiceKey = &serviceKey
}

func (r *RegistryAuthenticationGcpGar2) SetServiceKeyNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["ServiceKey"] = true
	r.ServiceKey = nil
}

func (r RegistryAuthenticationGcpGar2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["ServiceKey"] && r.ServiceKey == nil {
		data["service_key"] = nil
	} else if r.ServiceKey != nil {
		data["service_key"] = r.ServiceKey
	}

	return json.Marshal(data)
}

func (r RegistryAuthenticationGcpGar2) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegistryAuthenticationGcpGar2 to string"
	}
	return string(jsonData)
}
