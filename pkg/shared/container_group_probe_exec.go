package shared

type ContainerGroupProbeExec struct {
	Command []string `json:"command,omitempty" required:"true"`
}

func (c *ContainerGroupProbeExec) SetCommand(command []string) {
	c.Command = command
}

func (c *ContainerGroupProbeExec) GetCommand() []string {
	if c == nil {
		return nil
	}
	return c.Command
}
