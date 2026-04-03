package contenttypes

import (
	"fmt"
	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/unmarshal"
)

// FromJson deserializes JSON data from HTTP response bodies into the target struct.
// Uses the custom unmarshal package to handle complex types and validations.
func FromJson(data []byte, target any) error {
	err := unmarshal.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body into struct: %v", err)
	}
	return nil
}
