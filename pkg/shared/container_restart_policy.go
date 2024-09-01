package shared

type ContainerRestartPolicy string

const (
	CONTAINER_RESTART_POLICY_ALWAYS     ContainerRestartPolicy = "always"
	CONTAINER_RESTART_POLICY_ON_FAILURE ContainerRestartPolicy = "on_failure"
	CONTAINER_RESTART_POLICY_NEVER      ContainerRestartPolicy = "never"
)
