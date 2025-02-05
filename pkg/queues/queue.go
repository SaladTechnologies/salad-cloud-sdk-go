package queues

import (
	"encoding/json"
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
	touched    map[string]bool
}

func (q *Queue) GetId() *string {
	if q == nil {
		return nil
	}
	return q.Id
}

func (q *Queue) SetId(id string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Id"] = true
	q.Id = &id
}

func (q *Queue) SetIdNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Id"] = true
	q.Id = nil
}

func (q *Queue) GetName() *string {
	if q == nil {
		return nil
	}
	return q.Name
}

func (q *Queue) SetName(name string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Name"] = true
	q.Name = &name
}

func (q *Queue) SetNameNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Name"] = true
	q.Name = nil
}

func (q *Queue) GetDisplayName() *string {
	if q == nil {
		return nil
	}
	return q.DisplayName
}

func (q *Queue) SetDisplayName(displayName string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["DisplayName"] = true
	q.DisplayName = &displayName
}

func (q *Queue) SetDisplayNameNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["DisplayName"] = true
	q.DisplayName = nil
}

func (q *Queue) GetDescription() *string {
	if q == nil {
		return nil
	}
	return q.Description
}

func (q *Queue) SetDescription(description string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Description"] = true
	q.Description = &description
}

func (q *Queue) SetDescriptionNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["Description"] = true
	q.Description = nil
}

func (q *Queue) GetContainerGroups() []shared.ContainerGroup {
	if q == nil {
		return nil
	}
	return q.ContainerGroups
}

func (q *Queue) SetContainerGroups(containerGroups []shared.ContainerGroup) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["ContainerGroups"] = true
	q.ContainerGroups = containerGroups
}

func (q *Queue) SetContainerGroupsNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["ContainerGroups"] = true
	q.ContainerGroups = nil
}

func (q *Queue) GetCreateTime() *string {
	if q == nil {
		return nil
	}
	return q.CreateTime
}

func (q *Queue) SetCreateTime(createTime string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["CreateTime"] = true
	q.CreateTime = &createTime
}

func (q *Queue) SetCreateTimeNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["CreateTime"] = true
	q.CreateTime = nil
}

func (q *Queue) GetUpdateTime() *string {
	if q == nil {
		return nil
	}
	return q.UpdateTime
}

func (q *Queue) SetUpdateTime(updateTime string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["UpdateTime"] = true
	q.UpdateTime = &updateTime
}

func (q *Queue) SetUpdateTimeNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["UpdateTime"] = true
	q.UpdateTime = nil
}

func (q Queue) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if q.touched["Id"] && q.Id == nil {
		data["id"] = nil
	} else if q.Id != nil {
		data["id"] = q.Id
	}

	if q.touched["Name"] && q.Name == nil {
		data["name"] = nil
	} else if q.Name != nil {
		data["name"] = q.Name
	}

	if q.touched["DisplayName"] && q.DisplayName == nil {
		data["display_name"] = nil
	} else if q.DisplayName != nil {
		data["display_name"] = q.DisplayName
	}

	if q.touched["Description"] && q.Description == nil {
		data["description"] = nil
	} else if q.Description != nil {
		data["description"] = q.Description
	}

	if q.touched["ContainerGroups"] && q.ContainerGroups == nil {
		data["container_groups"] = nil
	} else if q.ContainerGroups != nil {
		data["container_groups"] = q.ContainerGroups
	}

	if q.touched["CreateTime"] && q.CreateTime == nil {
		data["create_time"] = nil
	} else if q.CreateTime != nil {
		data["create_time"] = q.CreateTime
	}

	if q.touched["UpdateTime"] && q.UpdateTime == nil {
		data["update_time"] = nil
	} else if q.UpdateTime != nil {
		data["update_time"] = q.UpdateTime
	}

	return json.Marshal(data)
}

func (q Queue) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: Queue to string"
	}
	return string(jsonData)
}
