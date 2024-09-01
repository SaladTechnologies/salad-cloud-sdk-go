package shared

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
)

type SaladCloudSdkError struct {
	Err      error
	Body     []byte
	Metadata SaladCloudSdkErrorMetadata
}

type SaladCloudSdkErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewSaladCloudSdkError[T any](transportError *httptransport.ErrorResponse[T]) *SaladCloudSdkError {
	return &SaladCloudSdkError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: SaladCloudSdkErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *SaladCloudSdkError) Error() string {
	return e.Err.Error()
}
