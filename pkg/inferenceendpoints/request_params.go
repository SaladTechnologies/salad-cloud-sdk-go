package inferenceendpoints

type ListInferenceEndpointsRequestParams struct {
	Page     *int64 `explode:"true" serializationStyle:"form" min:"1" max:"2147483647" queryParam:"page"`
	PageSize *int64 `explode:"true" serializationStyle:"form" min:"1" max:"100" queryParam:"page_size"`
}

func (params *ListInferenceEndpointsRequestParams) SetPage(page int64) {
	params.Page = &page
}
func (params *ListInferenceEndpointsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}

type ListInferenceEndpointJobsRequestParams struct {
	Page     *int64 `explode:"true" serializationStyle:"form" min:"1" max:"2147483647" queryParam:"page"`
	PageSize *int64 `explode:"true" serializationStyle:"form" min:"1" max:"100" queryParam:"page_size"`
}

func (params *ListInferenceEndpointJobsRequestParams) SetPage(page int64) {
	params.Page = &page
}
func (params *ListInferenceEndpointJobsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
