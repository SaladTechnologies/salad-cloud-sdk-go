package validation

import (
	"fmt"
	"reflect"
)

// validateRequired checks if a required field has a non-nil value.
// Returns an error if the field has a 'required:"true"' tag and is nil.
func validateRequired(fieldType reflect.StructField, fieldValue reflect.Value) error {
	if IsRequiredField(fieldType) && fieldValue.IsNil() {
		return fmt.Errorf("field %s is required", fieldType.Name)
	}

	return nil
}

// IsRequiredField checks if a struct field has the 'required:"true"' tag.
func IsRequiredField(fieldType reflect.StructField) bool {
	required, found := fieldType.Tag.Lookup("required")
	return found && required == "true"
}

// IsOptionalField checks if a struct field is optional (not required or no tag).
func IsOptionalField(fieldType reflect.StructField) bool {
	required, found := fieldType.Tag.Lookup("required")
	return !found || required == "" || required == "false"
}
