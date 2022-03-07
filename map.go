package fp

func Map[T any, R any](iteratee func(T, int) R) func([]T) []R {
	return func(xs []T) []R {

		result := make([]R, 0, len(xs))

		for i, x := range xs {
			result = append(result, iteratee(x, i))
		}

		return result
	}
}
