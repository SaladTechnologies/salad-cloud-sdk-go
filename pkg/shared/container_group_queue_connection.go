package shared

import (
	"encoding/json"
)

// Represents container group queue connection
type ContainerGroupQueueConnection struct {
	Path      *string `json:"path,omitempty" required:"true" maxLength:"1024" minLength:"1"`
	Port      *int64  `json:"port,omitempty" required:"true" min:"1" max:"65535"`
	QueueName *string `json:"queue_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	touched   map[string]bool
}

func (c *ContainerGroupQueueConnection) GetPath() *string {
	if c == nil {
		return nil
	}
	return c.Path
}

func (c *ContainerGroupQueueConnection) SetPath(path string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Path"] = true
	c.Path = &path
}

func (c *ContainerGroupQueueConnection) SetPathNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Path"] = true
	c.Path = nil
}

func (c *ContainerGroupQueueConnection) GetPort() *int64 {
	if c == nil {
		return nil
	}
	return c.Port
}

func (c *ContainerGroupQueueConnection) SetPort(port int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = &port
}

func (c *ContainerGroupQueueConnection) SetPortNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Port"] = true
	c.Port = nil
}

func (c *ContainerGroupQueueConnection) GetQueueName() *string {
	if c == nil {
		return nil
	}
	return c.QueueName
}

func (c *ContainerGroupQueueConnection) SetQueueName(queueName string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueueName"] = true
	c.QueueName = &queueName
}

func (c *ContainerGroupQueueConnection) SetQueueNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueueName"] = true
	c.QueueName = nil
}

func (c ContainerGroupQueueConnection) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Path"] && c.Path == nil {
		data["path"] = nil
	} else if c.Path != nil {
		data["path"] = c.Path
	}

	if c.touched["Port"] && c.Port == nil {
		data["port"] = nil
	} else if c.Port != nil {
		data["port"] = c.Port
	}

	if c.touched["QueueName"] && c.QueueName == nil {
		data["queue_name"] = nil
	} else if c.QueueName != nil {
		data["queue_name"] = c.QueueName
	}

	return json.Marshal(data)
}

func (c ContainerGroupQueueConnection) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupQueueConnection to string"
	}
	return string(jsonData)
}
