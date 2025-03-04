package shared

import "encoding/json"

type ContainerGroupProbeExec struct {
	Command []string `json:"command,omitempty" required:"true"`
}

func (c *ContainerGroupProbeExec) GetCommand() []string {
	if c == nil {
		return nil
	}
	return c.Command
}

func (c *ContainerGroupProbeExec) SetCommand(command []string) {
	c.Command = command
}

func (c ContainerGroupProbeExec) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupProbeExec to string"
	}
	return string(jsonData)
}
