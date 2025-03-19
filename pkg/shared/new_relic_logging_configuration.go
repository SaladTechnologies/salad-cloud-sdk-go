package shared

import "encoding/json"

// Configuration for sending container logs to New Relic's log management platform.
type NewRelicLoggingConfiguration struct {
	// The New Relic endpoint host for log ingestion (e.g., log-api.newrelic.com).
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// The New Relic license or ingestion key used for authentication and data routing.
	IngestionKey *string `json:"ingestion_key,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
}

func (n *NewRelicLoggingConfiguration) GetHost() *string {
	if n == nil {
		return nil
	}
	return n.Host
}

func (n *NewRelicLoggingConfiguration) SetHost(host string) {
	n.Host = &host
}

func (n *NewRelicLoggingConfiguration) GetIngestionKey() *string {
	if n == nil {
		return nil
	}
	return n.IngestionKey
}

func (n *NewRelicLoggingConfiguration) SetIngestionKey(ingestionKey string) {
	n.IngestionKey = &ingestionKey
}

func (n NewRelicLoggingConfiguration) String() string {
	jsonData, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return "error converting struct: NewRelicLoggingConfiguration to string"
	}
	return string(jsonData)
}
