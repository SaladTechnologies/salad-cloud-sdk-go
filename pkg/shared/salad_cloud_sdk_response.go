package shared

import (
	"encoding/json"

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

func (r *SaladCloudSdkResponse[T]) GetData() T {
	return r.Data
}

func (r SaladCloudSdkResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: SaladCloudSdkResponse to string"
	}
	return string(jsonData)
}
