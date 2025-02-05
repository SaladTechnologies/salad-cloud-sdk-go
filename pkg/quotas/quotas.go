package quotas

import (
	"encoding/json"
)

// Represents the organization quotas
type Quotas struct {
	ContainerGroupsQuotas *ContainerGroupsQuotas `json:"container_groups_quotas,omitempty" required:"true"`
	// The time the resource was created
	CreateTime *string `json:"create_time,omitempty"`
	// The time the resource was last updated
	UpdateTime *string `json:"update_time,omitempty"`
	touched    map[string]bool
}

func (q *Quotas) GetContainerGroupsQuotas() *ContainerGroupsQuotas {
	if q == nil {
		return nil
	}
	return q.ContainerGroupsQuotas
}

func (q *Quotas) SetContainerGroupsQuotas(containerGroupsQuotas ContainerGroupsQuotas) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["ContainerGroupsQuotas"] = true
	q.ContainerGroupsQuotas = &containerGroupsQuotas
}

func (q *Quotas) SetContainerGroupsQuotasNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["ContainerGroupsQuotas"] = true
	q.ContainerGroupsQuotas = nil
}

func (q *Quotas) GetCreateTime() *string {
	if q == nil {
		return nil
	}
	return q.CreateTime
}

func (q *Quotas) SetCreateTime(createTime string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["CreateTime"] = true
	q.CreateTime = &createTime
}

func (q *Quotas) SetCreateTimeNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["CreateTime"] = true
	q.CreateTime = nil
}

func (q *Quotas) GetUpdateTime() *string {
	if q == nil {
		return nil
	}
	return q.UpdateTime
}

func (q *Quotas) SetUpdateTime(updateTime string) {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["UpdateTime"] = true
	q.UpdateTime = &updateTime
}

func (q *Quotas) SetUpdateTimeNil() {
	if q.touched == nil {
		q.touched = map[string]bool{}
	}
	q.touched["UpdateTime"] = true
	q.UpdateTime = nil
}

func (q Quotas) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if q.touched["ContainerGroupsQuotas"] && q.ContainerGroupsQuotas == nil {
		data["container_groups_quotas"] = nil
	} else if q.ContainerGroupsQuotas != nil {
		data["container_groups_quotas"] = q.ContainerGroupsQuotas
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

func (q Quotas) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: Quotas to string"
	}
	return string(jsonData)
}
