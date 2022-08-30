package either

import opt "github.com/repeale/fp-go/option"

// Base struct
type Either[L, R any] struct {
	isLeft bool
	left   L
	right  R
}

func Left[L, R any](value L) Either[L, R] {
	return Either[L, R]{
		isLeft: true,
		left:   value,
	}
}

func Right[L, R any](value R) Either[L, R] {
	return Either[L, R]{
		isLeft: false,
		right:  value,
	}
}

func IsLeft[L, R any](e Either[L, R]) bool {
	return e.isLeft
}

func IsRight[L, R any](e Either[L, R]) bool {
	return !e.isLeft
}

func Exists[L, R any](predicate func(right R) bool) func(Either[L, R]) bool {
	return func(e Either[L, R]) bool {

		if IsLeft(e) {
			return false
		}

		return predicate(e.right)
	}
}

func Flatten[L, R any](e Either[L, Either[L, R]]) Either[L, R] {

	if IsLeft(e) {
		return Left[L, R](e.left)
	}

	return e.right
}

func FromError[R any](value R, e error) Either[error, R] {

	if e != nil {
		return Left[error, R](e)
	}

	return Right[error](value)
}

func FromErrorFn[R any](fn func() (value R, e error)) Either[error, R] {

	val, err := fn()

	if err != nil {
		return Left[error, R](err)
	}

	return Right[error](val)
}

func FromOption[L, R any](onNone func() L) func(o opt.Option[R]) Either[L, R] {
	return func(o opt.Option[R]) Either[L, R] {

		if opt.IsNone(o) {
			return Left[L, R](onNone())
		}

		return Right[L](o.Value)

	}
}

func FromPredicate[L, R any](predicate func(value R) bool, onLeft func() L) func(R) Either[L, R] {
	return func(value R) Either[L, R] {

		if predicate(value) {
			return Right[L](value)
		}

		return Left[L, R](onLeft())
	}
}

func GetOrElse[L, R any](onLeft func(left L) R) func(Either[L, R]) R {
	return func(e Either[L, R]) R {

		if IsLeft(e) {
			return onLeft(e.left)
		}

		return e.right
	}
}

func Map[L, R, T any](onRight func(right R) T) func(Either[L, R]) Either[L, T] {
	return func(e Either[L, R]) Either[L, T] {

		if IsLeft(e) {
			return Left[L, T](e.left)
		}

		return Right[L](onRight(e.right))
	}
}

func MapLeft[L, R, T any](fn func(left L) T) func(Either[L, R]) Either[T, R] {
	return func(e Either[L, R]) Either[T, R] {

		if IsLeft(e) {
			return Left[T, R](fn(e.left))
		}

		return Right[T](e.right)
	}
}

func Match[L, R, T any](onLeft func(left L) T, onRight func(right R) T) func(Either[L, R]) T {
	return func(e Either[L, R]) T {

		if IsLeft(e) {
			return onLeft(e.left)
		}

		return onRight(e.right)
	}
}
