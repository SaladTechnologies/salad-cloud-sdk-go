package containergroups

// The state of the container group instance
type TheContainerGroupInstanceState string

const (
	THE_CONTAINER_GROUP_INSTANCE_STATE_ALLOCATING  TheContainerGroupInstanceState = "allocating"
	THE_CONTAINER_GROUP_INSTANCE_STATE_DOWNLOADING TheContainerGroupInstanceState = "downloading"
	THE_CONTAINER_GROUP_INSTANCE_STATE_CREATING    TheContainerGroupInstanceState = "creating"
	THE_CONTAINER_GROUP_INSTANCE_STATE_RUNNING     TheContainerGroupInstanceState = "running"
	THE_CONTAINER_GROUP_INSTANCE_STATE_STOPPING    TheContainerGroupInstanceState = "stopping"
)
