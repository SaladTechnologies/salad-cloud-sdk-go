package validation

import (
	"reflect"

	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
)

type validatorFunc = func(fieldType reflect.StructField, fieldValue reflect.Value) error

func ValidateData(data any) error {
	if data == nil {
		return nil
	}

	dataValue := reflect.ValueOf(data)

	// Recursively dereference pointers
	for dataValue.Kind() == reflect.Ptr {
		if dataValue.IsNil() {
			return nil
		}
		dataValue = dataValue.Elem()
	}

	dataType := dataValue.Type()

	// Check if this is a Nullable wrapper
	if isNullableType(dataValue.Type()) {
		if dataValue.FieldByName("IsNull").Bool() {
			return nil
		}
		// Get the actual value from the Nullable wrapper
		dataValue = dataValue.FieldByName("Value")
		dataType = dataValue.Type()
	}

	if utils.GetReflectKind(dataType) == reflect.Struct {
		return validateStruct(dataValue.Interface())
	} else if dataType.Kind() == reflect.Array || dataType.Kind() == reflect.Slice {
		return validateArray(dataValue)
	}

	return nil
}

func validateStruct(data any) error {
	value := reflect.ValueOf(data)

	// Recursively dereference pointers
	for value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
	}

	// Check if this is a Nullable wrapper
	if isNullableType(value.Type()) {
		if value.FieldByName("IsNull").Bool() {
			return nil
		}
		// Get the actual value from the Nullable wrapper
		value = value.FieldByName("Value")
	}

	structValue := utils.GetReflectValue(value)
	for i := 0; i < structValue.NumField(); i++ {
		fieldValue := structValue.Field(i)
		fieldType := structValue.Type().Field(i)

		err := validateField(fieldValue, fieldType)
		if err != nil {
			return err
		}

		// Only check IsNil for types that can be nil
		kind := fieldValue.Kind()
		if (kind == reflect.Ptr || kind == reflect.Interface ||
			kind == reflect.Map || kind == reflect.Slice ||
			kind == reflect.Chan) && fieldValue.IsNil() {
			continue
		}

		kind = utils.GetReflectKind(fieldType.Type)
		if kind == reflect.Struct || kind == reflect.Array || kind == reflect.Slice {
			err := ValidateData(fieldValue.Interface())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func validateArray(value reflect.Value) error {
	// Check if this is a Nullable wrapper
	if isNullableType(value.Type()) {
		if value.FieldByName("IsNull").Bool() {
			return nil
		}
		// Get the actual value from the Nullable wrapper
		value = value.FieldByName("Value")
	}

	arrayValue := utils.GetReflectValue(value)
	for j := 0; j < arrayValue.Len(); j++ {
		err := ValidateData(arrayValue.Index(j).Interface())
		if err != nil {
			return err
		}
	}

	return nil
}

func isNullableType(t reflect.Type) bool {
	// Dereference pointer types until we hit the base type
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Must be a struct
	if t.Kind() != reflect.Struct {
		return false
	}

	// Should have at least the fields "IsNull" (bool) and "Value" (any type)
	isNullField, hasIsNull := t.FieldByName("IsNull")
	_, hasValue := t.FieldByName("Value")

	if !hasIsNull || !hasValue {
		return false
	}

	// Check that IsNull is a bool
	if isNullField.Type.Kind() != reflect.Bool {
		return false
	}

	return true
}

func validateField(fieldValue reflect.Value, fieldType reflect.StructField) error {
	validators := getValidators(fieldType)
	for _, validator := range validators {
		err := validator(fieldType, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func getValidators(fieldType reflect.StructField) []validatorFunc {
	return []validatorFunc{
		validateRequired,
		validatePattern,
		validateMultipleOf,
		validateMin,
		validateMax,
		validateArrayIsUnique,
		validateArrayLength,
	}
}
