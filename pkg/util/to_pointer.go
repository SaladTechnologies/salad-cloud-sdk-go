package util

// ToPointer converts a value to a pointer to that value.
// Useful for creating pointers to literals or converting values to optional fields.
func ToPointer[T any](v T) *T {
	return &v
}
