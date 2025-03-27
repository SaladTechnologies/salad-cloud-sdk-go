package shared

// The format in which logs will be delivered
type ContainerLoggingHttpFormat string

const (
	CONTAINER_LOGGING_HTTP_FORMAT_JSON       ContainerLoggingHttpFormat = "json"
	CONTAINER_LOGGING_HTTP_FORMAT_JSON_LINES ContainerLoggingHttpFormat = "json_lines"
)
