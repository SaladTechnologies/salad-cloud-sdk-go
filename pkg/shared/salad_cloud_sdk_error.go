package shared

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"net/http"
)

// SaladCloudSdkError wraps API errors with detailed metadata including status code, headers, and raw response.
// It implements the error interface and provides structured access to error information.
type SaladCloudSdkError[T any] struct {
	Err      error
	Data     *T
	Body     []byte
	Raw      *http.Response
	Metadata SaladCloudSdkErrorMetadata
}

// SaladCloudSdkErrorMetadata contains HTTP metadata associated with an error response.
type SaladCloudSdkErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewSaladCloudSdkError creates a new SaladCloudSdkError from an internal transport error.
// It extracts error details, body, status code, and headers into a user-facing error structure.
func NewSaladCloudSdkError[T any](transportError *httptransport.ErrorResponse[T]) *SaladCloudSdkError[T] {
	return &SaladCloudSdkError[T]{
		Err:  transportError.GetError(),
		Data: transportError.Data,
		Body: transportError.GetBody(),
		Raw:  transportError.Raw,
		Metadata: SaladCloudSdkErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

// Error implements the error interface, returning the error message string.
func (e *SaladCloudSdkError[T]) Error() string {
	return e.Err.Error()
}

// GetData returns the deserialized error response data.
// Returns nil if unmarshaling failed or the response body was empty.
func (e *SaladCloudSdkError[T]) GetData() *T {
	return e.Data
}
