package shared

import "encoding/json"

// Represents a Datadog tag used for container logging metadata.
type DatadogTagForContainerLogging struct {
	// The name of the metadata tag.
	Name *string `json:"name,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// The value of the metadata tag.
	Value *string `json:"value,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
}

func (d *DatadogTagForContainerLogging) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DatadogTagForContainerLogging) SetName(name string) {
	d.Name = &name
}

func (d *DatadogTagForContainerLogging) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

func (d *DatadogTagForContainerLogging) SetValue(value string) {
	d.Value = &value
}

func (d DatadogTagForContainerLogging) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DatadogTagForContainerLogging to string"
	}
	return string(jsonData)
}
