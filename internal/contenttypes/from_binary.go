package contenttypes

import (
	"fmt"
	"reflect"
)

func FromBinary(data any, target any) error {
	targetValue := reflect.ValueOf(target)

	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}

	targetValue = targetValue.Elem()

	if b, ok := data.([]byte); ok {
		if targetValue.Kind() == reflect.Slice && targetValue.Type().Elem().Kind() == reflect.Uint8 {
			targetValue.Set(reflect.ValueOf(b))
			return nil
		}
	}

	return fmt.Errorf("Failed to convert response body to byte array")
}
