package fp

func Map[T any, R any](callback func(T) R) func([]T) []R {
	return func(xs []T) []R {

		result := make([]R, 0, len(xs))

		for _, x := range xs {
			result = append(result, callback(x))
		}

		return result
	}
}

func MapWithIndex[T any, R any](callback func(T, int) R) func([]T) []R {
	return func(xs []T) []R {

		result := make([]R, 0, len(xs))

		for i, x := range xs {
			result = append(result, callback(x, i))
		}

		return result
	}
}
