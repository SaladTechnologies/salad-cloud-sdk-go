package shared

import (
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
)

type SaladCloudSdkResponse[T any] struct {
	Data     T
	Metadata SaladCloudSdkResponseMetadata
}

type SaladCloudSdkResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewSaladCloudSdkResponse[T any](resp *httptransport.Response[T]) *SaladCloudSdkResponse[T] {
	return &SaladCloudSdkResponse[T]{
		Data: resp.Data,
		Metadata: SaladCloudSdkResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}
