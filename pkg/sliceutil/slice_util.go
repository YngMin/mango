package sliceutil

func Map[T, V any](slice []T, extractFunc func(idx int) V) []V {
	values := make([]V, len(slice))
	for idx := range slice {
		values[idx] = extractFunc(idx)
	}
	return values
}
