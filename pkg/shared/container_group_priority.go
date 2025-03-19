package shared

// Specifies the priority level for container group execution, which determines resource allocation and scheduling precedence.
type ContainerGroupPriority string

const (
	CONTAINER_GROUP_PRIORITY_HIGH   ContainerGroupPriority = "high"
	CONTAINER_GROUP_PRIORITY_MEDIUM ContainerGroupPriority = "medium"
	CONTAINER_GROUP_PRIORITY_LOW    ContainerGroupPriority = "low"
	CONTAINER_GROUP_PRIORITY_BATCH  ContainerGroupPriority = "batch"
)
