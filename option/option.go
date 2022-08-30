package opt

import (
	"github.com/repeale/fp-go"
)

// Base struct
type Option[T any] struct {
	Value    T
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
func IsSome[T any](o Option[T]) bool {
	return o.hasValue
}

// Helper to check if the Option is missing the value
func IsNone[T any](o Option[T]) bool {
	return !o.hasValue
}

// Execute a function that returns an Option on the Option value if it exists. Otherwise return the empty Option itself
func Chain[A, B any](fn func(a A) Option[B]) func(Option[A]) Option[B] {
	return func(a Option[A]) Option[B] {

		if IsNone(a) {
			return None[B]()
		}

		return fn(a.Value)
	}
}

// Extracts the value out of the Option, if it exists. Otherwise returns the function with a default value
func GetOrElse[T any](onNone fp.Lazy[T]) func(Option[T]) T {
	return func(o Option[T]) T {

		if IsNone(o) {
			return onNone()
		}

		return o.Value
	}
}

// Execute the function on the Option value if it exists. Otherwise return the empty Option itself
func Map[T, R any](fn func(value T) R) func(o Option[T]) Option[R] {
	return func(o Option[T]) Option[R] {

		if IsNone(o) {
			return None[R]()
		}

		return Some(fn(o.Value))
	}
}

// Extracts the value out of the Option, if it exists, with a function. Otherwise returns the function with a default value
func Match[T, R any](onNone fp.Lazy[R], onSome func(value T) R) func(Option[T]) R {
	return func(o Option[T]) R {

		if IsNone(o) {
			return onNone()
		}

		return onSome(o.Value)
	}
}
