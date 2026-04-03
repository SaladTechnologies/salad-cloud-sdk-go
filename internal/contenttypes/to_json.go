package contenttypes

import (
	"bytes"
	"encoding/json"
)

// ToJson serializes data to JSON format for HTTP request bodies.
// Returns a bytes.Reader containing the JSON-encoded data, or nil if data is nil.
func ToJson(data any) (*bytes.Reader, error) {
	if data == nil {
		return nil, nil
	}

	marshalledBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(marshalledBody), nil
}
