package logs

import "encoding/json"

// Represents a query for logs
type LogEntryQuery struct {
	// The end time of the time range
	EndTime *string `json:"end_time,omitempty" required:"true"`
	// The maximum number of items per page.
	PageSize *int64 `json:"page_size,omitempty" min:"1" max:"100"`
	// The query string for filtering logs
	Query *string `json:"query,omitempty" required:"true" maxLength:"20000"`
	// The sort order of the log entries. `asc` will sort the log entries in chronological order. `desc` will sort the log entries in reverse chronological order.
	SortOrder *LogEntryQuerySortOrder `json:"sort_order,omitempty"`
	// The start time of the time range
	StartTime *string `json:"start_time,omitempty" required:"true"`
}

func (l *LogEntryQuery) GetEndTime() *string {
	if l == nil {
		return nil
	}
	return l.EndTime
}

func (l *LogEntryQuery) SetEndTime(endTime string) {
	l.EndTime = &endTime
}

func (l *LogEntryQuery) GetPageSize() *int64 {
	if l == nil {
		return nil
	}
	return l.PageSize
}

func (l *LogEntryQuery) SetPageSize(pageSize int64) {
	l.PageSize = &pageSize
}

func (l *LogEntryQuery) GetQuery() *string {
	if l == nil {
		return nil
	}
	return l.Query
}

func (l *LogEntryQuery) SetQuery(query string) {
	l.Query = &query
}

func (l *LogEntryQuery) GetSortOrder() *LogEntryQuerySortOrder {
	if l == nil {
		return nil
	}
	return l.SortOrder
}

func (l *LogEntryQuery) SetSortOrder(sortOrder LogEntryQuerySortOrder) {
	l.SortOrder = &sortOrder
}

func (l *LogEntryQuery) GetStartTime() *string {
	if l == nil {
		return nil
	}
	return l.StartTime
}

func (l *LogEntryQuery) SetStartTime(startTime string) {
	l.StartTime = &startTime
}

func (l LogEntryQuery) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LogEntryQuery to string"
	}
	return string(jsonData)
}
