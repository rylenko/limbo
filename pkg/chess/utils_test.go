package chess

// Ptr returns pointer to passed value.
//
// Useful to create a pointer to constant value.
func Ptr[T any](value T) *T {
	return &value
}
