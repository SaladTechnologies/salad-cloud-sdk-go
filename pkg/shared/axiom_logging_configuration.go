package shared

import "encoding/json"

// Configuration settings for integrating container logs with the Axiom logging service. When specified, container logs will be forwarded to the Axiom instance defined by these parameters.
type AxiomLoggingConfiguration struct {
	// The Axiom host URL where logs will be sent (e.g. logs.axiom.co)
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// Authentication token for the Axiom API with appropriate write permissions
	ApiToken *string `json:"api_token,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// Name of the Axiom dataset where the container logs will be stored and indexed
	Dataset *string `json:"dataset,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
}

func (a *AxiomLoggingConfiguration) GetHost() *string {
	if a == nil {
		return nil
	}
	return a.Host
}

func (a *AxiomLoggingConfiguration) SetHost(host string) {
	a.Host = &host
}

func (a *AxiomLoggingConfiguration) GetApiToken() *string {
	if a == nil {
		return nil
	}
	return a.ApiToken
}

func (a *AxiomLoggingConfiguration) SetApiToken(apiToken string) {
	a.ApiToken = &apiToken
}

func (a *AxiomLoggingConfiguration) GetDataset() *string {
	if a == nil {
		return nil
	}
	return a.Dataset
}

func (a *AxiomLoggingConfiguration) SetDataset(dataset string) {
	a.Dataset = &dataset
}

func (a AxiomLoggingConfiguration) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AxiomLoggingConfiguration to string"
	}
	return string(jsonData)
}
