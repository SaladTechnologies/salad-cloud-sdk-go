package quotas

// Represents the organization quotas
type Quotas struct {
	ContainerGroupsQuotas *ContainerGroupsQuotas `json:"container_groups_quotas,omitempty" required:"true"`
	// The time the resource was created
	CreateTime    *string        `json:"create_time,omitempty"`
	RecipesQuotas *RecipesQuotas `json:"recipes_quotas,omitempty" required:"true"`
	// The time the resource was last updated
	UpdateTime *string `json:"update_time,omitempty"`
}

func (q *Quotas) SetContainerGroupsQuotas(containerGroupsQuotas ContainerGroupsQuotas) {
	q.ContainerGroupsQuotas = &containerGroupsQuotas
}

func (q *Quotas) GetContainerGroupsQuotas() *ContainerGroupsQuotas {
	if q == nil {
		return nil
	}
	return q.ContainerGroupsQuotas
}

func (q *Quotas) SetCreateTime(createTime string) {
	q.CreateTime = &createTime
}

func (q *Quotas) GetCreateTime() *string {
	if q == nil {
		return nil
	}
	return q.CreateTime
}

func (q *Quotas) SetRecipesQuotas(recipesQuotas RecipesQuotas) {
	q.RecipesQuotas = &recipesQuotas
}

func (q *Quotas) GetRecipesQuotas() *RecipesQuotas {
	if q == nil {
		return nil
	}
	return q.RecipesQuotas
}

func (q *Quotas) SetUpdateTime(updateTime string) {
	q.UpdateTime = &updateTime
}

func (q *Quotas) GetUpdateTime() *string {
	if q == nil {
		return nil
	}
	return q.UpdateTime
}
