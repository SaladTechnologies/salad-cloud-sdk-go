package quotas

type RecipesQuotas struct {
	MaxCreatedRecipeDeployments *int64 `json:"max_created_recipe_deployments,omitempty" required:"true"`
	RecipeInstanceQuota         *int64 `json:"recipe_instance_quota,omitempty" required:"true"`
}

func (r *RecipesQuotas) SetMaxCreatedRecipeDeployments(maxCreatedRecipeDeployments int64) {
	r.MaxCreatedRecipeDeployments = &maxCreatedRecipeDeployments
}

func (r *RecipesQuotas) GetMaxCreatedRecipeDeployments() *int64 {
	if r == nil {
		return nil
	}
	return r.MaxCreatedRecipeDeployments
}

func (r *RecipesQuotas) SetRecipeInstanceQuota(recipeInstanceQuota int64) {
	r.RecipeInstanceQuota = &recipeInstanceQuota
}

func (r *RecipesQuotas) GetRecipeInstanceQuota() *int64 {
	if r == nil {
		return nil
	}
	return r.RecipeInstanceQuota
}
