package containergroups

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Configuration options for directing container logs to a logging provider. This schema enables you to specify a single logging destination for container output, supporting monitoring, debugging, and analytics use cases. Each provider has its own configuration parameters defined in the referenced schemas. Only one logging provider can be selected at a time.
type ContainerConfigurationLogging struct {
	// Configuration settings for integrating container logs with the Axiom logging service. When specified, container logs will be forwarded to the Axiom instance defined by these parameters.
	Axiom *shared.AxiomLoggingConfiguration `json:"axiom,omitempty"`
	// Configuration for forwarding container logs to Datadog monitoring service.
	Datadog *shared.DatadogLoggingConfiguration `json:"datadog,omitempty"`
	// Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.
	Http *ContainerLoggingConfigurationHttp2 `json:"http,omitempty"`
	// Configuration for sending container logs to New Relic's log management platform.
	NewRelic *shared.NewRelicLoggingConfiguration `json:"new_relic,omitempty"`
	// Configuration settings for forwarding container logs to a Splunk instance.
	Splunk *shared.ContainerLoggingSplunkConfiguration `json:"splunk,omitempty"`
	// Configuration for forwarding container logs to a remote TCP endpoint
	Tcp *shared.TcpLoggingConfiguration `json:"tcp,omitempty"`
}

func (c *ContainerConfigurationLogging) GetAxiom() *shared.AxiomLoggingConfiguration {
	if c == nil {
		return nil
	}
	return c.Axiom
}

func (c *ContainerConfigurationLogging) SetAxiom(axiom shared.AxiomLoggingConfiguration) {
	c.Axiom = &axiom
}

func (c *ContainerConfigurationLogging) GetDatadog() *shared.DatadogLoggingConfiguration {
	if c == nil {
		return nil
	}
	return c.Datadog
}

func (c *ContainerConfigurationLogging) SetDatadog(datadog shared.DatadogLoggingConfiguration) {
	c.Datadog = &datadog
}

func (c *ContainerConfigurationLogging) GetHttp() *ContainerLoggingConfigurationHttp2 {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerConfigurationLogging) SetHttp(http ContainerLoggingConfigurationHttp2) {
	c.Http = &http
}

func (c *ContainerConfigurationLogging) GetNewRelic() *shared.NewRelicLoggingConfiguration {
	if c == nil {
		return nil
	}
	return c.NewRelic
}

func (c *ContainerConfigurationLogging) SetNewRelic(newRelic shared.NewRelicLoggingConfiguration) {
	c.NewRelic = &newRelic
}

func (c *ContainerConfigurationLogging) GetSplunk() *shared.ContainerLoggingSplunkConfiguration {
	if c == nil {
		return nil
	}
	return c.Splunk
}

func (c *ContainerConfigurationLogging) SetSplunk(splunk shared.ContainerLoggingSplunkConfiguration) {
	c.Splunk = &splunk
}

func (c *ContainerConfigurationLogging) GetTcp() *shared.TcpLoggingConfiguration {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerConfigurationLogging) SetTcp(tcp shared.TcpLoggingConfiguration) {
	c.Tcp = &tcp
}

func (c ContainerConfigurationLogging) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerConfigurationLogging to string"
	}
	return string(jsonData)
}
