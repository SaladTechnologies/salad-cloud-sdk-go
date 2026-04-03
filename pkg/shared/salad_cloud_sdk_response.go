package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"net/http"
)

// SaladCloudSdkResponse is the user-facing wrapper for API responses.
// It contains the deserialized data, raw HTTP response, and metadata like headers and status code.
type SaladCloudSdkResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata SaladCloudSdkResponseMetadata
}

// SaladCloudSdkResponseMetadata contains HTTP metadata from the API response.
// Includes status code and headers for inspection and debugging.
type SaladCloudSdkResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewSaladCloudSdkResponse creates a new response wrapper from an internal transport response.
// Extracts data and metadata into a user-facing structure.
func NewSaladCloudSdkResponse[T any](resp *httptransport.Response[T]) *SaladCloudSdkResponse[T] {
	return &SaladCloudSdkResponse[T]{
		Data: resp.Data,
		Raw:  resp.Raw,
		Metadata: SaladCloudSdkResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

// GetData returns the deserialized response data.
func (r *SaladCloudSdkResponse[T]) GetData() T {
	return r.Data
}

// String returns a JSON representation of the response for debugging.
// Returns an error message if JSON marshaling fails.
func (r SaladCloudSdkResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: SaladCloudSdkResponse to string"
	}
	return string(jsonData)
}
