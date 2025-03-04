package shared

import "encoding/json"

// Represents container group queue connection
type ContainerGroupQueueConnection struct {
	Path      *string `json:"path,omitempty" required:"true" maxLength:"1024" minLength:"1"`
	Port      *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
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
