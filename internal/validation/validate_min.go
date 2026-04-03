package validation

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/utils"
)

// validateMin validates that a numeric field value is greater than or equal to the 'min' tag value.
// Supports int and float types. Nil values are skipped.
func validateMin(field reflect.StructField, value reflect.Value) error {
	minValue, found := field.Tag.Lookup("min")
	if !found || minValue == "" || value.IsNil() {
		return nil
	}

	min, err := strconv.Atoi(minValue)
	if err != nil {
		return err
	}

	val := utils.GetReflectValue(value)

	if val.CanInt() {
		if val.Int() < int64(min) {
			return fmt.Errorf("field %s is less than min value", field.Name)
		}
	} else if val.CanFloat() {
		if val.Float() < float64(min) {
			return fmt.Errorf("field %s is less than min value", field.Name)
		}
	} else {
		return fmt.Errorf("field %s is not a number", field.Name)
	}

	return nil
}
