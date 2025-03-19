package shared

import "encoding/json"

// Configuration for forwarding container logs to a remote TCP endpoint
type TcpLoggingConfiguration struct {
	// The hostname or IP address of the remote TCP logging endpoint
	Host *string `json:"host,omitempty" required:"true" maxLength:"1000" minLength:"1" pattern:"^.*$"`
	// The port number on which the TCP logging endpoint is listening
	Port *int64 `json:"port,omitempty" required:"true" min:"1" max:"65535"`
}

func (t *TcpLoggingConfiguration) GetHost() *string {
	if t == nil {
		return nil
	}
	return t.Host
}

func (t *TcpLoggingConfiguration) SetHost(host string) {
	t.Host = &host
}

func (t *TcpLoggingConfiguration) GetPort() *int64 {
	if t == nil {
		return nil
	}
	return t.Port
}

func (t *TcpLoggingConfiguration) SetPort(port int64) {
	t.Port = &port
}

func (t TcpLoggingConfiguration) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TcpLoggingConfiguration to string"
	}
	return string(jsonData)
}
