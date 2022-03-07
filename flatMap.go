package fp

func FlatMap[T any, R any](callback func(T) []R) func([]T) []R {
	return func(xs []T) []R {

		result := []R{}

		for _, x := range xs {
			result = append(result, callback(x)...)
		}

		return result
	}
}

func FlatMapWithIndex[T any, R any](callback func(T, int) []R) func([]T) []R {
	return func(xs []T) []R {

		result := []R{}

		for i, x := range xs {
			result = append(result, callback(x, i)...)
		}

		return result
	}
}

func FlatMapWithSlice[T any, R any](callback func(T, int, []T) []R) func([]T) []R {
	return func(xs []T) []R {

		result := []R{}

		for i, x := range xs {
			result = append(result, callback(x, i, xs)...)
		}

		return result
	}
}
