package shared

import "encoding/json"

// Configuration for connecting a container group to a message queue system, enabling asynchronous communication between services.
type ContainerGroupQueueConnection struct {
	// The endpoint path for accessing the queue service, relative to the base URL of the queue server.
	Path *string `json:"path,omitempty" required:"true" maxLength:"1024" minLength:"1" pattern:"^.*$"`
	// The network port number used to connect to the queue service. Must be a valid TCP/IP port between 1 and 65535.
	Port *int64 `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	// Unique identifier for the queue. Must start with a lowercase letter, can contain lowercase letters, numbers, and hyphens, and must end with a letter or number.
	QueueName *string `json:"queue_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
}

func (c *ContainerGroupQueueConnection) GetPath() *string {
	if c == nil {
		return nil
	}
	return c.Path
}

func (c *ContainerGroupQueueConnection) SetPath(path string) {
	c.Path = &path
}

func (c *ContainerGroupQueueConnection) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupQueueConnection) SetPort(port int64) {
	c.Port = &port
}

func (c *ContainerGroupQueueConnection) GetQueueName() *string {
	if c == nil {
		return nil
	}
	return c.QueueName
}

func (c *ContainerGroupQueueConnection) SetQueueName(queueName string) {
	c.QueueName = &queueName
}

func (c ContainerGroupQueueConnection) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupQueueConnection to string"
	}
	return string(jsonData)
}
