package shared

// Represents the current operational state of a container group within the Salad platform.
type ContainerGroupStatus string

const (
	CONTAINER_GROUP_STATUS_PENDING   ContainerGroupStatus = "pending"
	CONTAINER_GROUP_STATUS_RUNNING   ContainerGroupStatus = "running"
	CONTAINER_GROUP_STATUS_STOPPED   ContainerGroupStatus = "stopped"
	CONTAINER_GROUP_STATUS_SUCCEEDED ContainerGroupStatus = "succeeded"
	CONTAINER_GROUP_STATUS_FAILED    ContainerGroupStatus = "failed"
	CONTAINER_GROUP_STATUS_DEPLOYING ContainerGroupStatus = "deploying"
)
