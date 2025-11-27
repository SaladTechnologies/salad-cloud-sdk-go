package shared

import (
	"encoding/json"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
	"net/http"
)

type SaladCloudSdkResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata SaladCloudSdkResponseMetadata
}

type SaladCloudSdkResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

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
