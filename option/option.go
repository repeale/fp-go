package option

import (
	"github.com/repeale/fp-go"
)

type Option[T any] struct {
	value    T
	hasValue bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{value, true}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func IsSome[T any](option Option[T]) bool {
	return option.hasValue
}

func IsNone[T any](option Option[T]) bool {
	return !option.hasValue
}

func GetOrElse[T any](onNone fp.Lazy[T]) func(Option[T]) T {
	return func(option Option[T]) T {

		if IsSome(option) {
			return option.value
		}

		return onNone()
	}
}

func Match[T any](onNone fp.Lazy[T], onSome fp.LazyVal[T]) func(Option[T]) T {
	return func(option Option[T]) T {

		if IsSome(option) {
			return onSome(option.value)
		}

		return onNone()
	}
}

func Map[T any](fn fp.LazyVal[T]) func(o Option[T]) Option[T] {
	return func(option Option[T]) Option[T] {

		if IsSome(option) {
			return Some(fn(option.value))
		}

		return None[T]()
	}
}

func Chain[A any, B any](fn func(a A) Option[B]) func(Option[A]) Option[B] {
	return func(a Option[A]) Option[B] {
		if IsNone(a) {
			return None[B]()
		}

		return fn(a.value)
	}
}
