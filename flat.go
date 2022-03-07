package fp

func Flat[T any](xs [][]T) []T {
	result := []T{}

	for _, item := range xs {
		result = append(result, item...)
	}

	return result
}
