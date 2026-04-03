package inferenceendpoints

// ListInferenceEndpointsRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListInferenceEndpointsRequestParams struct {
	Page     *int64 `explode:"true" serializationStyle:"form" min:"1" max:"2147483647" queryParam:"page"`
	PageSize *int64 `explode:"true" serializationStyle:"form" min:"1" max:"100" queryParam:"page_size"`
}

// SetPage sets the Page parameter.
func (params *ListInferenceEndpointsRequestParams) SetPage(page int64) {
	params.Page = &page
}

// SetPageSize sets the PageSize parameter.
func (params *ListInferenceEndpointsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}

// ListInferenceEndpointJobsRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListInferenceEndpointJobsRequestParams struct {
	Page     *int64 `explode:"true" serializationStyle:"form" min:"1" max:"2147483647" queryParam:"page"`
	PageSize *int64 `explode:"true" serializationStyle:"form" min:"1" max:"100" queryParam:"page_size"`
}

// SetPage sets the Page parameter.
func (params *ListInferenceEndpointJobsRequestParams) SetPage(page int64) {
	params.Page = &page
}

// SetPageSize sets the PageSize parameter.
func (params *ListInferenceEndpointJobsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
