package fp

func Some[T any](predicate func(T) bool) func([]T) bool {
	return func(xs []T) bool {

		for _, x := range xs {
			if predicate(x) {
				return true
			}
		}

		return false
	}
}

func SomeWithIndex[T any](predicate func(T, int) bool) func([]T) bool {
	return func(xs []T) bool {

		for i, x := range xs {
			if predicate(x, i) {
				return true
			}
		}

		return false
	}
}

func SomeWithSlice[T any](predicate func(T, int, []T) bool) func([]T) bool {
	return func(xs []T) bool {

		for i, x := range xs {
			if predicate(x, i, xs) {
				return true
			}
		}

		return false
	}
}
