package unmarshal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func UnmarshalNullable(source []byte, target any) error {
	// Use a temporary map to decode the raw JSON
	var rawMap map[string]json.RawMessage
	if err := json.Unmarshal(source, &rawMap); err != nil {
		return err
	}

	val := reflect.ValueOf(target).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		jsonKey := field.Tag.Get("json")

		split := strings.Split(jsonKey, ",")
		if len(split) > 0 {
			jsonKey = split[0]
		}

		if jsonKey == "" {
			jsonKey = field.Name
		}

		rawVal, ok := rawMap[jsonKey]
		if !ok {
			continue
		}

		fieldVal := val.Field(i)

		// Handle null explicitly
		if string(rawVal) == "null" && fieldVal.Kind() == reflect.Ptr {
			// Create a new instance of the pointer's element (e.g. Nullable[T])
			newNullable := reflect.New(fieldVal.Type().Elem()).Elem()

			// Set the IsNull field to true
			isNullField := newNullable.FieldByName("IsNull")
			if isNullField.IsValid() && isNullField.CanSet() {
				isNullField.SetBool(true)
			}

			// Assign back to the pointer field
			fieldVal.Set(newNullable.Addr())
		} else if fieldVal.CanAddr() {
			if err := json.Unmarshal(rawVal, fieldVal.Addr().Interface()); err != nil {
				return fmt.Errorf("failed to unmarshal field %s: %w", field.Name, err)
			}
		}
	}

	return nil
}

func hasNullableFields(obj any) bool {
	t := reflect.TypeOf(obj)

	// Dereference pointer if needed
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Ensure it's a struct
	if t.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// For pointer to Nullable[T]
		if field.Type.Kind() == reflect.Ptr {
			elem := field.Type.Elem()
			// check if the name starts with "Nullable"
			if strings.HasPrefix(elem.Name(), "Nullable") {
				return true
			}
		}
	}

	return false
}
