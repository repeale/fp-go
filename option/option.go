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

// Helper to check if the Option has a value
func IsSome[T any](option Option[T]) bool {
	return option.hasValue
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

// Extracts the value out of the Option, if it exists, with a function. Otherwise returns the function with a default value
func Match[T any](onNone fp.Lazy[T], onSome fp.LazyVal[T]) func(Option[T]) T {
	return func(option Option[T]) T {

		if IsNone(option) {
			return onNone()
		}

		return onSome(option.value)
	}
}

// Execute the function on the Option value if it exists. Otherwise return the empty Option itself
func Map[T any](fn fp.LazyVal[T]) func(o Option[T]) Option[T] {
	return func(option Option[T]) Option[T] {

		if IsNone(option) {
			return None[T]()
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
