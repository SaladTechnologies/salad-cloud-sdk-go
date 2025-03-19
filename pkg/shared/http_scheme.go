package shared

// The protocol scheme used for HTTP probe requests in container health checks.
type HttpScheme string

const (
	HTTP_SCHEME_HTTP  HttpScheme = "http"
	HTTP_SCHEME_HTTPS HttpScheme = "https"
)
