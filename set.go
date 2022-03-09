package set

import "encoding/json"

// Set is a data structure that contains a set of comparable elements. It is
// implemented using an internal map.
type Set[T comparable] struct {
	// m is the internal map. Its keys are the elements of the set and the
	// values are booleans that indicates whether the key is present or not.
	m map[T]bool
}

var (
	_ json.Marshaler   = (*Set[int])(nil)
	_ json.Unmarshaler = (*Set[int])(nil)
)

// Empty initializes a new Set without any elements inside it.
func Empty[T comparable]() Set[T] {
	return Set[T]{m: make(map[T]bool)}
}

// Of initializes a new Set and appends the given values to it.
func Of[T comparable](values ...T) Set[T] {
	s := Empty[T]()
	s.Append(values...)
	return s
}

// Len returns the number of elements that s contains.
func (s Set[T]) Len() int {
	return len(s.m)
}

// Contains reports whether s contains val.
func (s Set[T]) Contains(val T) bool {
	return s.m[val]
}

// Append adds the values to s. If any value is already present, the value does
// not impact the set.
func (s Set[T]) Append(values ...T) {
	for _, v := range values {
		s.m[v] = true
	}
}

// Delete removes the elements of values from s.
func (s Set[T]) Delete(values ...T) {
	for _, v := range values {
		delete(s.m, v)
	}
}

// Clear removes all elements from s.
func (s Set[T]) Clear() {
	s.m = make(map[T]bool)
}

// Slice converts s into a slice.
func (s Set[T]) Slice() []T {
	values := make([]T, len(s.m))
	for t := range s.m {
		values = append(values, t)
	}
	return values
}

// MarshalJSON marshals s into a JSON array.
func (s Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Slice())
}

// UnmarshalJSON unmarshals a JSON array into s. The previous content of s is
// cleared.
func (s Set[T]) UnmarshalJSON(data []byte) error {
	s.Clear()
	var values []T
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	s.Append(values...)
	return nil
}
