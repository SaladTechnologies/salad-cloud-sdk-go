package queues

type ListQueueJobsRequestParams struct {
	Page     *int64 `min:"1" max:"2147483647" queryParam:"page"`
	PageSize *int64 `min:"1" max:"100" queryParam:"page_size"`
}

func (params *ListQueueJobsRequestParams) SetPage(page int64) {
	params.Page = &page
}
func (params *ListQueueJobsRequestParams) SetPageSize(pageSize int64) {
	params.PageSize = &pageSize
}
