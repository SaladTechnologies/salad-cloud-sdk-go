package queues

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/shared"
)

// Represents a queue.
type Queue struct {
	// The queue identifier. This is automatically generated and assigned when the queue is created.
	Id *string `json:"id,omitempty" required:"true"`
	// The queue name. This must be unique within the project.
	Name *string `json:"name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The display name. This may be used as a more human-readable name.
	DisplayName *string `json:"display_name,omitempty" required:"true" maxLength:"63" minLength:"2" pattern:"^[ ,-.0-9A-Za-z]+$"`
	// The description. This may be used as a space for notes or other information about the queue.
	Description     *string                 `json:"description,omitempty" maxLength:"500"`
	ContainerGroups []shared.ContainerGroup `json:"container_groups,omitempty" required:"true" maxItems:"100"`
	// The date and time the queue was created.
	CreateTime *string `json:"create_time,omitempty" required:"true"`
	// The date and time the queue was last updated.
	UpdateTime *string `json:"update_time,omitempty" required:"true"`
}

func (q *Queue) SetId(id string) {
	q.Id = &id
}

func (q *Queue) GetId() *string {
	if q == nil {
		return nil
	}
	return q.Id
}

func (q *Queue) SetName(name string) {
	q.Name = &name
}

func (q *Queue) GetName() *string {
	if q == nil {
		return nil
	}
	return q.Name
}

func (q *Queue) SetDisplayName(displayName string) {
	q.DisplayName = &displayName
}

func (q *Queue) GetDisplayName() *string {
	if q == nil {
		return nil
	}
	return q.DisplayName
}

func (q *Queue) SetDescription(description string) {
	q.Description = &description
}

func (q *Queue) GetDescription() *string {
	if q == nil {
		return nil
	}
	return q.Description
}

func (q *Queue) SetContainerGroups(containerGroups []shared.ContainerGroup) {
	q.ContainerGroups = containerGroups
}

func (q *Queue) GetContainerGroups() []shared.ContainerGroup {
	if q == nil {
		return nil
	}
	return q.ContainerGroups
}

func (q *Queue) SetCreateTime(createTime string) {
	q.CreateTime = &createTime
}

func (q *Queue) GetCreateTime() *string {
	if q == nil {
		return nil
	}
	return q.CreateTime
}

func (q *Queue) SetUpdateTime(updateTime string) {
	q.UpdateTime = &updateTime
}

func (q *Queue) GetUpdateTime() *string {
	if q == nil {
		return nil
	}
	return q.UpdateTime
}
