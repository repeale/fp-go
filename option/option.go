package opt

import (
	"github.com/repeale/fp-go"
)

// Base option struct
type Option[T any] struct {
	value    T
	hasValue bool
}

// Constructor for Option with a value
func Some[T any](value T) Option[T] {
	return Option[T]{value, true}
}

// Constructor for Option without a value
func None[T any]() Option[T] {
	return Option[T]{}
}

// Constructor for Option from a pointer. nil pointer == None, otherwise Some
func FromPtr[T any](ptr *T) Option[T] {
	if ptr == nil {
		return None[T]()
	}

	return Some(*ptr)
}

// Helper to check if the Option has a value
func IsSome[T any](option Option[T]) bool {
	return option.hasValue
}

// Helper to check if the Option has a value and if that value satisfies a predicate
func IsSomeAnd[T any](pred fp.Pred[T]) func(Option[T]) bool {
	return func(option Option[T]) bool {
		return option.hasValue && pred(option.value)
	}
}

// Helper to check if the Option is missing the value
func IsNone[T any](option Option[T]) bool {
	return !option.hasValue
}

// Extracts the value out of the Option, if it exists. Otherwise returns the function with a default value
func GetOrElse[T any](onNone fp.Lazy[T]) func(Option[T]) T {
	return func(option Option[T]) T {

		if IsNone(option) {
			return onNone()
		}

		return option.value
	}
}

// Extracts the value out of the Option, if it exists. Otherwise panics
func Get[T any](option Option[T]) T {
	if IsNone(option) {
		panic("Can't extract a value out of None")
	}

	return option.value

}

// Extracts the value out of the Option as a pointer, if it exists. Otherwise returns a nil pointer
func ToPtr[T any](option Option[T]) *T {
	if IsNone(option) {
		return nil
	}

	return &option.value
}

// Extracts the value out of the Option, if it exists, with a function. Otherwise returns the function with a default value
func Match[T, R any](onNone func() R, onSome func(T) R) func(Option[T]) R {
	return func(option Option[T]) R {

		if IsNone(option) {
			return onNone()
		}

		return onSome(option.value)
	}
}

// Execute the function on the Option value if it exists. Otherwise return the empty Option itself
func Map[T, R any](fn func(T) R) func(o Option[T]) Option[R] {
	return func(option Option[T]) Option[R] {

		if IsNone(option) {
			return None[R]()
		}

		return Some(fn(option.value))
	}
}

// Execute a function that returns an Option on the Option value if it exists. Otherwise return the empty Option itself
func Chain[A any, B any](fn func(a A) Option[B]) func(Option[A]) Option[B] {
	return func(a Option[A]) Option[B] {

		if IsNone(a) {
			return None[B]()
		}

		return fn(a.value)
	}
}

// Execute a predicate on the Option value if it exists.
// If the result is false or the Option is empty, return the empty Option.
// Otherwise, return the option itself
func Filter[T any](fn fp.Pred[T]) func(o Option[T]) Option[T] {
	return func(option Option[T]) Option[T] {
		if IsNone(option) {
			return None[T]()
		}

		if !fn(option.value) {
			return None[T]()
		}

		return option
	}
}

// Removes one level of nesting at a time. Option[Option[T]] -> Option[T]
func Flat[T any](option Option[Option[T]]) Option[T] {
	if IsNone(option) {
		return None[T]()
	}

	return option.value
}
