package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/util"
)

// Configuration for forwarding container logs to Datadog monitoring service.
type DatadogLoggingConfiguration struct {
	// The Datadog intake server host URL where logs will be sent.
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// The Datadog API key used for authentication when sending logs.
	ApiKey *string `json:"api_key,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// Optional metadata tags to attach to logs for filtering and categorization in Datadog.
	Tags *util.Nullable[[]DatadogTagForContainerLogging] `json:"tags,omitempty" required:"true" maxItems:"1000"`
}

func (d *DatadogLoggingConfiguration) GetHost() *string {
	if d == nil {
		return nil
	}
	return d.Host
}

func (d *DatadogLoggingConfiguration) SetHost(host string) {
	d.Host = &host
}

func (d *DatadogLoggingConfiguration) GetApiKey() *string {
	if d == nil {
		return nil
	}
	return d.ApiKey
}

func (d *DatadogLoggingConfiguration) SetApiKey(apiKey string) {
	d.ApiKey = &apiKey
}

func (d *DatadogLoggingConfiguration) GetTags() *util.Nullable[[]DatadogTagForContainerLogging] {
	if d == nil {
		return nil
	}
	return d.Tags
}

func (d *DatadogLoggingConfiguration) SetTags(tags util.Nullable[[]DatadogTagForContainerLogging]) {
	d.Tags = &tags
}

func (d *DatadogLoggingConfiguration) SetTagsNull() {
	d.Tags = &util.Nullable[[]DatadogTagForContainerLogging]{IsNull: true}
}

func (d DatadogLoggingConfiguration) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DatadogLoggingConfiguration to string"
	}
	return string(jsonData)
}
