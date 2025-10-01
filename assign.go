package utils

// Assign assigns a value to a pointer and returns the assigned value.
func Assign[T any](dst *T, src T) T {
	if dst != nil {
		*dst = src
	}
	return *dst
}
