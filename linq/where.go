package linq

func Where[T any](typeForTest []T, where func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range typeForTest {
		if where(v) {
			result = append(result, v)
		}
	}
	return result
}
