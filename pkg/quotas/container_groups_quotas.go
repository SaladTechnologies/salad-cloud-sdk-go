package quotas

import "encoding/json"

// Represents the organization quotas for container groups
type ContainerGroupsQuotas struct {
	// The maximum number of replicas that can be created for a container group
	ContainerReplicasQuota *int64 `json:"container_replicas_quota,omitempty" required:"true" min:"0" max:"500"`
	// The number of replicas that are currently in use
	ContainerReplicasUsed *int64 `json:"container_replicas_used,omitempty" required:"true" min:"0" max:"500"`
	// The maximum number of container group reallocations per minute
	MaxContainerGroupReallocationsPerMinute *int64 `json:"max_container_group_reallocations_per_minute,omitempty" min:"0" max:"100"`
	// The maximum number of container group recreates per minute
	MaxContainerGroupRecreatesPerMinute *int64 `json:"max_container_group_recreates_per_minute,omitempty" min:"0" max:"100"`
	// The maximum number of container group restarts per minute
	MaxContainerGroupRestartsPerMinute *int64 `json:"max_container_group_restarts_per_minute,omitempty" min:"0" max:"100"`
}

func (c *ContainerGroupsQuotas) GetContainerReplicasQuota() *int64 {
	if c == nil {
		return nil
	}
	return c.ContainerReplicasQuota
}

func (c *ContainerGroupsQuotas) SetContainerReplicasQuota(containerReplicasQuota int64) {
	c.ContainerReplicasQuota = &containerReplicasQuota
}

func (c *ContainerGroupsQuotas) GetContainerReplicasUsed() *int64 {
	if c == nil {
		return nil
	}
	return c.ContainerReplicasUsed
}

func (c *ContainerGroupsQuotas) SetContainerReplicasUsed(containerReplicasUsed int64) {
	c.ContainerReplicasUsed = &containerReplicasUsed
}

func (c *ContainerGroupsQuotas) GetMaxContainerGroupReallocationsPerMinute() *int64 {
	if c == nil {
		return nil
	}
	return c.MaxContainerGroupReallocationsPerMinute
}

func (c *ContainerGroupsQuotas) SetMaxContainerGroupReallocationsPerMinute(maxContainerGroupReallocationsPerMinute int64) {
	c.MaxContainerGroupReallocationsPerMinute = &maxContainerGroupReallocationsPerMinute
}

func (c *ContainerGroupsQuotas) GetMaxContainerGroupRecreatesPerMinute() *int64 {
	if c == nil {
		return nil
	}
	return c.MaxContainerGroupRecreatesPerMinute
}

func (c *ContainerGroupsQuotas) SetMaxContainerGroupRecreatesPerMinute(maxContainerGroupRecreatesPerMinute int64) {
	c.MaxContainerGroupRecreatesPerMinute = &maxContainerGroupRecreatesPerMinute
}

func (c *ContainerGroupsQuotas) GetMaxContainerGroupRestartsPerMinute() *int64 {
	if c == nil {
		return nil
	}
	return c.MaxContainerGroupRestartsPerMinute
}

func (c *ContainerGroupsQuotas) SetMaxContainerGroupRestartsPerMinute(maxContainerGroupRestartsPerMinute int64) {
	c.MaxContainerGroupRestartsPerMinute = &maxContainerGroupRestartsPerMinute
}

func (c ContainerGroupsQuotas) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupsQuotas to string"
	}
	return string(jsonData)
}
