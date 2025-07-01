package logs

// The sort order of the log entries. `asc` will sort the log entries in chronological order. `desc` will sort the log entries in reverse chronological order.
type LogEntryQuerySortOrder string

const (
	LOG_ENTRY_QUERY_SORT_ORDER_DESC LogEntryQuerySortOrder = "desc"
	LOG_ENTRY_QUERY_SORT_ORDER_ASC  LogEntryQuerySortOrder = "asc"
)
