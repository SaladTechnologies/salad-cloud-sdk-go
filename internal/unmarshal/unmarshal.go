package unmarshal

import (
	"fmt"
	"reflect"

	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
)

// Unmarshal deserializes JSON bytes into the target based on its type.
// Handles complex objects (oneOf), regular objects/arrays, and primitive types (string, int, float, bool).
// Target must be a non-nil pointer.
func Unmarshal(source []byte, target any) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}

	if isComplexObject(target) || isObject(target) || isArray(target) {
		return ToObject(source, target)
	} else if isString(targetValue.Elem().Kind()) {
		return ToString(source, targetValue)
	} else if isInteger(targetValue.Elem().Kind()) {
		return ToInt(source, targetValue)
	} else if isFloat(targetValue.Elem().Kind()) {
		return ToFloat(source, targetValue)
	} else if isBool(targetValue.Elem().Kind()) {
		return ToBool(source, targetValue)
	}

	return nil
}

// isArray checks if the target type is an array or slice.
func isArray(target any) bool {
	targetType := reflect.TypeOf(target)
	kind := utils.GetReflectKind(targetType)
	return kind == reflect.Array || kind == reflect.Slice
}

// isObject checks if the target type is a struct.
func isObject(target any) bool {
	targetType := reflect.TypeOf(target)
	return utils.GetReflectKind(targetType) == reflect.Struct
}

// isComplexObject checks if the target is a complex object (oneOf) where all fields have 'oneof' tags.
// Used to identify discriminated union types that require special unmarshaling logic.
func isComplexObject(target any) bool {
	targetType := reflect.TypeOf(target)
	if utils.GetReflectKind(targetType) != reflect.Struct {
		return false
	}

	allFieldsAreOneOf := true

	structValue := utils.GetReflectValue(reflect.ValueOf(target))
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Type().Field(i)
		allFieldsAreOneOf = isOneOfField(field) && allFieldsAreOneOf
	}

	return allFieldsAreOneOf
}

// isOneOfField checks if a struct field has the 'oneof' tag, indicating it's part of a discriminated union.
func isOneOfField(field reflect.StructField) bool {
	_, found := field.Tag.Lookup("oneof")
	return found
}
