package shared

import "encoding/json"

// Defines the exec action for a probe in a container group. This is used to execute a command inside a container for health checks.
type ContainerGroupProbeExec struct {
	// The command to execute inside the container. Exit status of 0 is considered successful, any other exit status is considered failure.
	Command []string `json:"command,omitempty" required:"true" minItems:"1" maxItems:"100"`
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
