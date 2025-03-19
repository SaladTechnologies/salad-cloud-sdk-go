package quotas

import "encoding/json"

// Represents the organization quotas
type Quotas struct {
	// Represents the organization quotas for container groups
	ContainerGroupsQuotas *ContainerGroupsQuotas `json:"container_groups_quotas,omitempty" required:"true"`
	// The time the resource was created
	CreateTime *string `json:"create_time,omitempty"`
	// The time the resource was last updated
	UpdateTime *string `json:"update_time,omitempty"`
}

func (q *Quotas) GetContainerGroupsQuotas() *ContainerGroupsQuotas {
	if q == nil {
		return nil
	}
	return q.ContainerGroupsQuotas
}

func (q *Quotas) SetContainerGroupsQuotas(containerGroupsQuotas ContainerGroupsQuotas) {
	q.ContainerGroupsQuotas = &containerGroupsQuotas
}

func (q *Quotas) GetCreateTime() *string {
	if q == nil {
		return nil
	}
	return q.CreateTime
}

func (q *Quotas) SetCreateTime(createTime string) {
	q.CreateTime = &createTime
}

func (q *Quotas) GetUpdateTime() *string {
	if q == nil {
		return nil
	}
	return q.UpdateTime
}

func (q *Quotas) SetUpdateTime(updateTime string) {
	q.UpdateTime = &updateTime
}

func (q Quotas) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: Quotas to string"
	}
	return string(jsonData)
}
