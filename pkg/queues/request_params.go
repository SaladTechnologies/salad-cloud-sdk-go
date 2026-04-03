package queues

// ListQueueJobsRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListQueueJobsRequestParams struct {
	Page     *int64 `explode:"true" serializationStyle:"form" min:"1" max:"2147483647" queryParam:"page"`
	PageSize *int64 `explode:"true" serializationStyle:"form" min:"1" max:"100" queryParam:"page_size"`
}

// SetPage sets the Page parameter.
func (params *ListQueueJobsRequestParams) SetPage(page int64) {
	params.Page = &page
}

// SetPageSize sets the PageSize parameter.
func (params *ListQueueJobsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
