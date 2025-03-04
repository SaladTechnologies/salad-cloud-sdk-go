package quotas

import "encoding/json"

type ContainerGroupsQuotas struct {
	MaxCreatedContainerGroups               *int64 `json:"max_created_container_groups,omitempty" required:"true"`
	ContainerInstanceQuota                  *int64 `json:"container_instance_quota,omitempty" required:"true"`
	MaxContainerGroupReallocationsPerMinute *int64 `json:"max_container_group_reallocations_per_minute,omitempty" min:"0"`
	MaxContainerGroupRecreatesPerMinute     *int64 `json:"max_container_group_recreates_per_minute,omitempty" min:"0"`
	MaxContainerGroupRestartsPerMinute      *int64 `json:"max_container_group_restarts_per_minute,omitempty" min:"0"`
}

func (c *ContainerGroupsQuotas) GetMaxCreatedContainerGroups() *int64 {
	if c == nil {
		return nil
	}
	return c.MaxCreatedContainerGroups
}

func (c *ContainerGroupsQuotas) SetMaxCreatedContainerGroups(maxCreatedContainerGroups int64) {
	c.MaxCreatedContainerGroups = &maxCreatedContainerGroups
}

func (c *ContainerGroupsQuotas) GetContainerInstanceQuota() *int64 {
	if c == nil {
		return nil
	}
	return c.ContainerInstanceQuota
}

func (c *ContainerGroupsQuotas) SetContainerInstanceQuota(containerInstanceQuota int64) {
	c.ContainerInstanceQuota = &containerInstanceQuota
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
