package shared

// The compression algorithm to apply to logs before transmission
type ContainerLoggingHttpCompression string

const (
	CONTAINER_LOGGING_HTTP_COMPRESSION_NONE ContainerLoggingHttpCompression = "none"
	CONTAINER_LOGGING_HTTP_COMPRESSION_GZIP ContainerLoggingHttpCompression = "gzip"
)
