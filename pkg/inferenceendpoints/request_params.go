package inferenceendpoints

type ListInferenceEndpointsRequestParams struct {
	Page     *int64 `min:"1" max:"2147483647" queryParam:"page"`
	PageSize *int64 `min:"1" max:"100" queryParam:"page_size"`
}

func (params *ListInferenceEndpointsRequestParams) SetPage(page int64) {
	params.Page = &page
}
func (params *ListInferenceEndpointsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}

type GetInferenceEndpointJobsRequestParams struct {
	Page     *int64 `min:"1" max:"2147483647" queryParam:"page"`
	PageSize *int64 `min:"1" max:"100" queryParam:"page_size"`
}

func (params *GetInferenceEndpointJobsRequestParams) SetPage(page int64) {
	params.Page = &page
}
func (params *GetInferenceEndpointJobsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
