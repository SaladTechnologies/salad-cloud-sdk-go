package shared

import (
	"encoding/json"
)

type ContainerGroupProbeExec struct {
	Command []string `json:"command,omitempty" required:"true"`
	touched map[string]bool
}

func (c *ContainerGroupProbeExec) GetCommand() []string {
	if c == nil {
		return nil
	}
	return c.Command
}

func (c *ContainerGroupProbeExec) SetCommand(command []string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Command"] = true
	c.Command = command
}

func (c *ContainerGroupProbeExec) SetCommandNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Command"] = true
	c.Command = nil
}

func (c ContainerGroupProbeExec) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Command"] && c.Command == nil {
		data["command"] = nil
	} else if c.Command != nil {
		data["command"] = c.Command
	}

	return json.Marshal(data)
}

func (c ContainerGroupProbeExec) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupProbeExec to string"
	}
	return string(jsonData)
}
