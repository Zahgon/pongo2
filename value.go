package pongo2

import (
	"reflect"
	"time"
)

type Value struct {
	val  reflect.Value
	safe bool // used to indicate whether a Value needs explicit escaping in the template
}

// AsValue converts any given value to a pongo2.Value
// Usually being used within own functions passed to a template
// through a Context or within filter functions.
//
// Example:
//
//	AsValue("my string")
func AsValue(i any) *Value { _ = "STUB: not implemented"; return nil }

// AsSafeValue works like AsValue, but does not apply the 'escape' filter.
func AsSafeValue(i any) *Value { _ = "STUB: not implemented"; return nil }

func (v *Value) getResolvedValue() reflect.Value {
	_ = "STUB: not implemented"

	// Unwrap pointers and interfaces to get to the underlying value
	return *new(reflect.Value)
}

// IsString checks whether the underlying value is a string
func (v *Value) IsString() bool { _ = "STUB: not implemented"; return false }

// IsBool checks whether the underlying value is a bool
func (v *Value) IsBool() bool { _ = "STUB: not implemented"; return false }

// IsFloat checks whether the underlying value is a float
func (v *Value) IsFloat() bool { _ = "STUB: not implemented"; return false }

// IsInteger checks whether the underlying value is an integer
func (v *Value) IsInteger() bool { _ = "STUB: not implemented"; return false }

// IsNumber checks whether the underlying value is either an integer
// or a float.
func (v *Value) IsNumber() bool { _ = "STUB: not implemented"; return false }

// IsTime checks whether the underlying value is a time.Time.
func (v *Value) IsTime() bool { _ = "STUB: not implemented"; return false }

// IsNil checks whether the underlying value is NIL
func (v *Value) IsNil() bool {
	_ = "STUB: not implemented"
	// fmt.Printf("%+v\n", v.getResolvedValue().Type().String())
	return false
}

// String returns a string for the underlying value. If this value is not
// of type string, pongo2 tries to convert it. Currently the following
// types for underlying values are supported:
//
//  1. string
//  2. int/uint (any size)
//  3. float (any precision)
//  4. bool
//  5. time.Time
//  6. String() will be called on the underlying value if provided
//
// NIL values will lead to an empty string. Unsupported types are leading
// to their respective type name.
func (v *Value) String() string { _ = "STUB: not implemented"; return "" }

// Integer returns the underlying value as an integer (converts the underlying
// value, if necessary). If it's not possible to convert the underlying value,
// it will return 0.
func (v *Value) Integer() int { _ = "STUB: not implemented"; return 0 }

// Try to convert from string to int (base 10)

// Float returns the underlying value as a float (converts the underlying
// value, if necessary). If it's not possible to convert the underlying value,
// it will return 0.0.
func (v *Value) Float() float64 { _ = "STUB: not implemented"; return 0 }

// Try to convert from string to float64 (base 10)

// Bool returns the underlying value as bool. If the value is not bool, false
// will always be returned. If you're looking for true/false-evaluation of the
// underlying value, have a look on the IsTrue()-function.
func (v *Value) Bool() bool { _ = "STUB: not implemented"; return false }

// Time returns the underlying value as time.Time.
// If the underlying value is not a time.Time, it returns the zero value of time.Time.
func (v *Value) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// IsTrue tries to evaluate the underlying value the Pythonic-way:
//
// Returns TRUE in one the following cases:
//
//   - int != 0
//   - uint != 0
//   - float != 0.0
//   - len(array/chan/map/slice/string) > 0
//   - bool == true
//   - underlying value is a struct
//
// Otherwise returns always FALSE.
func (v *Value) IsTrue() bool { _ = "STUB: not implemented"; return false }

// struct instance is always true

// Negate tries to negate the underlying value. It's mainly used for
// the NOT-operator and in conjunction with a call to
// return_value.IsTrue() afterwards.
//
// Example:
//
//	AsValue(1).Negate().IsTrue() == false
func (v *Value) Negate() *Value { _ = "STUB: not implemented"; return nil }

// Len returns the length for an array, chan, map, slice or string.
// Otherwise it will return 0.
func (v *Value) Len() int { _ = "STUB: not implemented"; return 0 }

// Slice slices an array, slice or string. Otherwise it will
// return nil.
func (v *Value) Slice(i, j int) *Value { _ = "STUB: not implemented"; return nil }

// Index gets the i-th item of an array, slice or string. Otherwise
// it will return NIL.
func (v *Value) Index(i int) *Value { _ = "STUB: not implemented"; return nil }

// Contains checks whether the underlying value (which must be of type struct, map,
// string, array or slice) contains of another Value (e. g. used to check
// whether a struct contains of a specific field or a map contains a specific key).
//
// Example:
//
//	AsValue("Hello, World!").Contains(AsValue("World")) == true
func (v *Value) Contains(other *Value) bool { _ = "STUB: not implemented"; return false }

// We can't check against invalid types

// Ensure that map key type is equal to the resolved other type.

// CanSlice checks whether the underlying value is of type array, slice or string.
// You normally would use CanSlice() before using the Slice() operation.
func (v *Value) CanSlice() bool { _ = "STUB: not implemented"; return false }

// IsSliceOrArray returns true if the value is a slice or array (not a string)
func (v *Value) IsSliceOrArray() bool { _ = "STUB: not implemented"; return false }

// IsMap checks whether the underlying value is a map
func (v *Value) IsMap() bool { _ = "STUB: not implemented"; return false }

// IsStruct checks whether the underlying value is a struct
func (v *Value) IsStruct() bool { _ = "STUB: not implemented"; return false }

// GetItem retrieves a value from a map by key or a field from a struct by name.
// For maps, it attempts to convert the key to the map's key type.
// For structs, it uses the key's string representation as the field name.
// Returns nil Value if the key/field doesn't exist or the type doesn't support item access.
func (v *Value) GetItem(key *Value) *Value { _ = "STUB: not implemented"; return nil }

// Try to get the map value using appropriate key type

// Try direct conversion if the key type matches

// Iterate iterates over a map, array, slice or a string. It calls the
// function's first argument for every value with the following arguments:
//
//	idx      current 0-index
//	count    total amount of items
//	key      *Value for the key or item
//	value    *Value (only for maps, the respective value for a specific key)
//
// If the underlying value has no items or is not one of the types above,
// the empty function (function's second argument) will be called.
func (v *Value) Iterate(fn func(idx, count int, key, value *Value) bool, empty func()) {
	_ = "STUB: not implemented"
	return
}

// IterateOrder behaves like Value.Iterate, but can iterate through an array/slice/string in reverse. Does
// not affect the iteration through a map because maps don't have any particular order.
// However, you can force an order using the `sorted` keyword (and even use `reversed sorted`).
func (v *Value) IterateOrder(fn func(idx, count int, key, value *Value) bool, empty func(), reverse bool, sorted bool) {
	_ = "STUB: not implemented"
	return
}

// done

// done

// done

// Interface gives you access to the underlying value.
func (v *Value) Interface() any { _ = "STUB: not implemented"; return *new(any) }

// EqualValueTo checks whether two values are containing the same value or object (if comparable).
func (v *Value) EqualValueTo(other *Value) bool {
	_ = "STUB: not implemented"
	// Handle numeric comparison: float vs int should compare by value (e.g., 8.0 == 8)
	// Also handles uint vs int comparison (see issue #64)
	return false
}

// If either is a float, compare as floats

// Both are integers (includes uint vs int)

// Handle nil/undefined values (see issue #341)
// Two nil values are considered equal

// One nil and one non-nil are not equal

// Note: reflect.Value.Equal() and Value.Comparable() (Go 1.20+) were considered
// but benchmarking showed they are slower. Type().Comparable() and
// Interface() == Interface() is faster due to Go's interface comparison optimization.

type sortedKeys []reflect.Value

func (sk sortedKeys) Len() int { _ = "STUB: not implemented"; return 0 }

func (sk sortedKeys) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (sk sortedKeys) Swap(i, j int) { _ = "STUB: not implemented"; return }

type valuesList []*Value

func (vl valuesList) Len() int { _ = "STUB: not implemented"; return 0 }

func (vl valuesList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (vl valuesList) Swap(i, j int) { _ = "STUB: not implemented"; return }
