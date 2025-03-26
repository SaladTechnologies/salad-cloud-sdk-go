package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Configuration options for directing container logs to a logging provider. This schema enables you to specify a single logging destination for container output, supporting monitoring, debugging, and analytics use cases. Each provider has its own configuration parameters defined in the referenced schemas. Only one logging provider can be selected at a time.
type UpdateContainerLogging struct {
	// Configuration settings for integrating container logs with the Axiom logging service. When specified, container logs will be forwarded to the Axiom instance defined by these parameters.
	Axiom *shared.AxiomLoggingConfiguration `json:"axiom,omitempty"`
	// Configuration for forwarding container logs to Datadog monitoring service.
	Datadog *shared.DatadogLoggingConfiguration `json:"datadog,omitempty"`
	// Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.
	Http *shared.ContainerLoggingConfigurationHttp1 `json:"http,omitempty"`
	// Configuration for sending container logs to New Relic's log management platform.
	NewRelic *shared.NewRelicLoggingConfiguration `json:"new_relic,omitempty"`
	// Configuration settings for forwarding container logs to a Splunk instance.
	Splunk *shared.ContainerLoggingSplunkConfiguration `json:"splunk,omitempty"`
	// Configuration for forwarding container logs to a remote TCP endpoint
	Tcp *shared.TcpLoggingConfiguration `json:"tcp,omitempty"`
}

func (u *UpdateContainerLogging) GetAxiom() *shared.AxiomLoggingConfiguration {
	if u == nil {
		return nil
	}
	return u.Axiom
}

func (u *UpdateContainerLogging) SetAxiom(axiom shared.AxiomLoggingConfiguration) {
	u.Axiom = &axiom
}

func (u *UpdateContainerLogging) GetDatadog() *shared.DatadogLoggingConfiguration {
	if u == nil {
		return nil
	}
	return u.Datadog
}

func (u *UpdateContainerLogging) SetDatadog(datadog shared.DatadogLoggingConfiguration) {
	u.Datadog = &datadog
}

func (u *UpdateContainerLogging) GetHttp() *shared.ContainerLoggingConfigurationHttp1 {
	if u == nil {
		return nil
	}
	return u.Http
}

func (u *UpdateContainerLogging) SetHttp(http shared.ContainerLoggingConfigurationHttp1) {
	u.Http = &http
}

func (u *UpdateContainerLogging) GetNewRelic() *shared.NewRelicLoggingConfiguration {
	if u == nil {
		return nil
	}
	return u.NewRelic
}

func (u *UpdateContainerLogging) SetNewRelic(newRelic shared.NewRelicLoggingConfiguration) {
	u.NewRelic = &newRelic
}

func (u *UpdateContainerLogging) GetSplunk() *shared.ContainerLoggingSplunkConfiguration {
	if u == nil {
		return nil
	}
	return u.Splunk
}

func (u *UpdateContainerLogging) SetSplunk(splunk shared.ContainerLoggingSplunkConfiguration) {
	u.Splunk = &splunk
}

func (u *UpdateContainerLogging) GetTcp() *shared.TcpLoggingConfiguration {
	if u == nil {
		return nil
	}
	return u.Tcp
}

func (u *UpdateContainerLogging) SetTcp(tcp shared.TcpLoggingConfiguration) {
	u.Tcp = &tcp
}

func (u UpdateContainerLogging) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateContainerLogging to string"
	}
	return string(jsonData)
}
