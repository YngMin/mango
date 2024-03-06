package sliceutil

func Extract[T, V any](slice []T, extractFunc func(idx int) V) []V {
	values := make([]V, len(slice))
	for idx := range slice {
		values[idx] = extractFunc(idx)
	}
	return values
}

func ExtractIf[T, V any](slice []T, extractIfFunc func(idx int) (extracted V, ok bool)) []V {
	values := make([]V, 0)
	for idx := range slice {
		if v, ok := extractIfFunc(idx); ok {
			values = append(values, v)
		}
	}
	return values
}
