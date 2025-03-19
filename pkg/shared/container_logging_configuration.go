package shared

import "encoding/json"

// Configuration options for directing container logs to a logging provider. This schema enables you to specify a single logging destination for container output, supporting monitoring, debugging, and analytics use cases. Each provider has its own configuration parameters defined in the referenced schemas. Only one logging provider can be selected at a time.
type ContainerLoggingConfiguration struct {
	// Configuration settings for integrating container logs with the Axiom logging service. When specified, container logs will be forwarded to the Axiom instance defined by these parameters.
	Axiom *AxiomLoggingConfiguration `json:"axiom,omitempty"`
	// Configuration for forwarding container logs to Datadog monitoring service.
	Datadog *DatadogLoggingConfiguration `json:"datadog,omitempty"`
	// Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.
	Http *ContainerHttpLoggingConfiguration `json:"http,omitempty"`
	// Configuration for sending container logs to New Relic's log management platform.
	NewRelic *NewRelicLoggingConfiguration `json:"new_relic,omitempty"`
	// Configuration settings for forwarding container logs to a Splunk instance.
	Splunk *ContainerLoggingSplunkConfiguration `json:"splunk,omitempty"`
	// Configuration for forwarding container logs to a remote TCP endpoint
	Tcp *TcpLoggingConfiguration `json:"tcp,omitempty"`
}

func (c *ContainerLoggingConfiguration) GetAxiom() *AxiomLoggingConfiguration {
	if c == nil {
		return nil
	}
	return c.Axiom
}

func (c *ContainerLoggingConfiguration) SetAxiom(axiom AxiomLoggingConfiguration) {
	c.Axiom = &axiom
}

func (c *ContainerLoggingConfiguration) GetDatadog() *DatadogLoggingConfiguration {
	if c == nil {
		return nil
	}
	return c.Datadog
}

func (c *ContainerLoggingConfiguration) SetDatadog(datadog DatadogLoggingConfiguration) {
	c.Datadog = &datadog
}

func (c *ContainerLoggingConfiguration) GetHttp() *ContainerHttpLoggingConfiguration {
	if c == nil {
		return nil
	}
	return c.Http
}

func (c *ContainerLoggingConfiguration) SetHttp(http ContainerHttpLoggingConfiguration) {
	c.Http = &http
}

func (c *ContainerLoggingConfiguration) GetNewRelic() *NewRelicLoggingConfiguration {
	if c == nil {
		return nil
	}
	return c.NewRelic
}

func (c *ContainerLoggingConfiguration) SetNewRelic(newRelic NewRelicLoggingConfiguration) {
	c.NewRelic = &newRelic
}

func (c *ContainerLoggingConfiguration) GetSplunk() *ContainerLoggingSplunkConfiguration {
	if c == nil {
		return nil
	}
	return c.Splunk
}

func (c *ContainerLoggingConfiguration) SetSplunk(splunk ContainerLoggingSplunkConfiguration) {
	c.Splunk = &splunk
}

func (c *ContainerLoggingConfiguration) GetTcp() *TcpLoggingConfiguration {
	if c == nil {
		return nil
	}
	return c.Tcp
}

func (c *ContainerLoggingConfiguration) SetTcp(tcp TcpLoggingConfiguration) {
	c.Tcp = &tcp
}

func (c ContainerLoggingConfiguration) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerLoggingConfiguration to string"
	}
	return string(jsonData)
}
