package fp

func Reduce[T any, R any](callback func(R, T) R, initial R) func([]T) R {
	return func(xs []T) R {

		for _, x := range xs {
			initial = callback(initial, x)
		}

		return initial
	}
}

func ReduceWithIndex[T any, R any](callback func(R, T, int) R, initial R) func([]T) R {
	return func(xs []T) R {

		for i, x := range xs {
			initial = callback(initial, x, i)
		}

		return initial
	}
}
