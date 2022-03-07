package fp

func Filter[T any](predicate func(T) bool) func([]T) []T {
	return func(xs []T) []T {

		result := []T{}

		for _, x := range xs {
			if predicate(x) {
				result = append(result, x)
			}
		}

		return result
	}
}

func FilterWithIndex[T any](predicate func(T, int) bool) func([]T) []T {
	return func(xs []T) []T {

		result := []T{}

		for i, x := range xs {
			if predicate(x, i) {
				result = append(result, x)
			}
		}

		return result
	}
}
