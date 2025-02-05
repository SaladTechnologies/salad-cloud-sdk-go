package quotas

import (
	"encoding/json"
)

type ContainerGroupsQuotas struct {
	MaxCreatedContainerGroups               *int64 `json:"max_created_container_groups,omitempty" required:"true"`
	ContainerInstanceQuota                  *int64 `json:"container_instance_quota,omitempty" required:"true"`
	MaxContainerGroupReallocationsPerMinute *int64 `json:"max_container_group_reallocations_per_minute,omitempty" min:"0"`
	MaxContainerGroupRecreatesPerMinute     *int64 `json:"max_container_group_recreates_per_minute,omitempty" min:"0"`
	MaxContainerGroupRestartsPerMinute      *int64 `json:"max_container_group_restarts_per_minute,omitempty" min:"0"`
	touched                                 map[string]bool
}

func (c *ContainerGroupsQuotas) GetMaxCreatedContainerGroups() *int64 {
	if c == nil {
		return nil
	}
	return c.MaxCreatedContainerGroups
}

func (c *ContainerGroupsQuotas) SetMaxCreatedContainerGroups(maxCreatedContainerGroups int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["MaxCreatedContainerGroups"] = true
	c.MaxCreatedContainerGroups = &maxCreatedContainerGroups
}

func (c *ContainerGroupsQuotas) SetMaxCreatedContainerGroupsNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["MaxCreatedContainerGroups"] = true
	c.MaxCreatedContainerGroups = nil
}

func (c *ContainerGroupsQuotas) GetContainerInstanceQuota() *int64 {
	if c == nil {
		return nil
	}
	return c.ContainerInstanceQuota
}

func (c *ContainerGroupsQuotas) SetContainerInstanceQuota(containerInstanceQuota int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ContainerInstanceQuota"] = true
	c.ContainerInstanceQuota = &containerInstanceQuota
}

func (c *ContainerGroupsQuotas) SetContainerInstanceQuotaNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ContainerInstanceQuota"] = true
	c.ContainerInstanceQuota = nil
}

func (c *ContainerGroupsQuotas) GetMaxContainerGroupReallocationsPerMinute() *int64 {
	if c == nil {
		return nil
	}
	return c.MaxContainerGroupReallocationsPerMinute
}

func (c *ContainerGroupsQuotas) SetMaxContainerGroupReallocationsPerMinute(maxContainerGroupReallocationsPerMinute int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["MaxContainerGroupReallocationsPerMinute"] = true
	c.MaxContainerGroupReallocationsPerMinute = &maxContainerGroupReallocationsPerMinute
}

func (c *ContainerGroupsQuotas) SetMaxContainerGroupReallocationsPerMinuteNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["MaxContainerGroupReallocationsPerMinute"] = true
	c.MaxContainerGroupReallocationsPerMinute = nil
}

func (c *ContainerGroupsQuotas) GetMaxContainerGroupRecreatesPerMinute() *int64 {
	if c == nil {
		return nil
	}
	return c.MaxContainerGroupRecreatesPerMinute
}

func (c *ContainerGroupsQuotas) SetMaxContainerGroupRecreatesPerMinute(maxContainerGroupRecreatesPerMinute int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["MaxContainerGroupRecreatesPerMinute"] = true
	c.MaxContainerGroupRecreatesPerMinute = &maxContainerGroupRecreatesPerMinute
}

func (c *ContainerGroupsQuotas) SetMaxContainerGroupRecreatesPerMinuteNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["MaxContainerGroupRecreatesPerMinute"] = true
	c.MaxContainerGroupRecreatesPerMinute = nil
}

func (c *ContainerGroupsQuotas) GetMaxContainerGroupRestartsPerMinute() *int64 {
	if c == nil {
		return nil
	}
	return c.MaxContainerGroupRestartsPerMinute
}

func (c *ContainerGroupsQuotas) SetMaxContainerGroupRestartsPerMinute(maxContainerGroupRestartsPerMinute int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["MaxContainerGroupRestartsPerMinute"] = true
	c.MaxContainerGroupRestartsPerMinute = &maxContainerGroupRestartsPerMinute
}

func (c *ContainerGroupsQuotas) SetMaxContainerGroupRestartsPerMinuteNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["MaxContainerGroupRestartsPerMinute"] = true
	c.MaxContainerGroupRestartsPerMinute = nil
}

func (c ContainerGroupsQuotas) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["MaxCreatedContainerGroups"] && c.MaxCreatedContainerGroups == nil {
		data["max_created_container_groups"] = nil
	} else if c.MaxCreatedContainerGroups != nil {
		data["max_created_container_groups"] = c.MaxCreatedContainerGroups
	}

	if c.touched["ContainerInstanceQuota"] && c.ContainerInstanceQuota == nil {
		data["container_instance_quota"] = nil
	} else if c.ContainerInstanceQuota != nil {
		data["container_instance_quota"] = c.ContainerInstanceQuota
	}

	if c.touched["MaxContainerGroupReallocationsPerMinute"] && c.MaxContainerGroupReallocationsPerMinute == nil {
		data["max_container_group_reallocations_per_minute"] = nil
	} else if c.MaxContainerGroupReallocationsPerMinute != nil {
		data["max_container_group_reallocations_per_minute"] = c.MaxContainerGroupReallocationsPerMinute
	}

	if c.touched["MaxContainerGroupRecreatesPerMinute"] && c.MaxContainerGroupRecreatesPerMinute == nil {
		data["max_container_group_recreates_per_minute"] = nil
	} else if c.MaxContainerGroupRecreatesPerMinute != nil {
		data["max_container_group_recreates_per_minute"] = c.MaxContainerGroupRecreatesPerMinute
	}

	if c.touched["MaxContainerGroupRestartsPerMinute"] && c.MaxContainerGroupRestartsPerMinute == nil {
		data["max_container_group_restarts_per_minute"] = nil
	} else if c.MaxContainerGroupRestartsPerMinute != nil {
		data["max_container_group_restarts_per_minute"] = c.MaxContainerGroupRestartsPerMinute
	}

	return json.Marshal(data)
}

func (c ContainerGroupsQuotas) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContainerGroupsQuotas to string"
	}
	return string(jsonData)
}
