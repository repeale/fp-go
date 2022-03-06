package fp

func Map[T any, R any](iteratee func(T, int) R) func([]T) []R {
	return func(collection []T) []R {

		result := make([]R, 0, len(collection))

		for i, item := range collection {
			result = append(result, iteratee(item, i))
		}

		return result
	}
}
