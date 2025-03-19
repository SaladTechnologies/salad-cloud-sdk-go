package shared

import "encoding/json"

// Configuration settings for forwarding container logs to a Splunk instance.
type ContainerLoggingSplunkConfiguration struct {
	// The URL of the Splunk HTTP Event Collector (HEC) endpoint.
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// The authentication token required to send data to the Splunk HEC endpoint.
	Token *string `json:"token,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
}

func (c *ContainerLoggingSplunkConfiguration) GetHost() *string {
	if c == nil {
		return nil
	}
	return c.Host
}

func (c *ContainerLoggingSplunkConfiguration) SetHost(host string) {
	c.Host = &host
}

func (c *ContainerLoggingSplunkConfiguration) GetToken() *string {
	if c == nil {
		return nil
	}
	return c.Token
}

func (c *ContainerLoggingSplunkConfiguration) SetToken(token string) {
	c.Token = &token
}

func (c ContainerLoggingSplunkConfiguration) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerLoggingSplunkConfiguration to string"
	}
	return string(jsonData)
}
