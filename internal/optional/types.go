package optional

import "encoding/json"

// Optional is a type that represents an optional value of any type.
//
// It is used to represent a value that may or may not be set.
//
// If IsSet is false, the Value is not set.
// If IsSet is true, the Value is set.
type Optional[T any] struct {
	IsSet bool
	Value T
}

func (o Optional[T]) IsZero() bool {
	return !o.IsSet
}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.IsSet {
		return []byte("null"), nil
	}
	return json.Marshal(o.Value)
}

func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	var value T

	if len(b) == 0 || string(b) == "null" {
		o.IsSet = false
		o.Value = value
		return nil
	}

	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	o.IsSet = true
	o.Value = value
	return nil
}

// String is an optional string.
//
// It is used to represent a string that may or may not be set.
//
// If IsSet is false, the Value is not set.
// If IsSet is true, the Value is set.
type String = Optional[string]

// Int is an optional int.
//
// It is used to represent an int that may or may not be set.
//
// If IsSet is false, the Value is not set.
// If IsSet is true, the Value is set.
type Int = Optional[int]

// Float64 is an optional float64.
//
// It is used to represent a float64 that may or may not be set.
//
// If IsSet is false, the Value is not set.
// If IsSet is true, the Value is set.
type Float64 = Optional[float64]

// Bool is an optional bool.
//
// It is used to represent a bool that may or may not be set.
//
// If IsSet is false, the Value is not set.
// If IsSet is true, the Value is set.
type Bool = Optional[bool]

// Any is an optional any type.
//
// It is used to represent a value of any type that may or may not be set.
//
// If IsSet is false, the Value is not set.
// If IsSet is true, the Value is set.
type Any = Optional[any]

// MapStringAny is an optional map of string to any type.
//
// It is used to represent a map of string to any type that may or may not be set.
//
// If IsSet is false, the Value is not set.
// If IsSet is true, the Value is set.
type MapStringAny = Optional[map[string]any]

// MapStringString is an optional map of string to string.
//
// It is used to represent a map of string to string that may or may not be set.
//
// If IsSet is false, the Value is not set.
// If IsSet is true, the Value is set.
type MapStringString = Optional[map[string]string]

// MapIntInt is an optional map of int to int.
//
// It is used to represent a map of int to int that may or may not be set.
//
// If IsSet is false, the Value is not set.
// If IsSet is true, the Value is set.
type MapIntInt = Optional[map[int]int]
