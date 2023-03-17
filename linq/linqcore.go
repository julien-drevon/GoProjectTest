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

func Select[T any, K any](typeForTest []T, projection func(x T) K) []K {
	result := make([]K, len(typeForTest))
	for i, v := range typeForTest {
		result[i] = projection(v)
	}

	return result
}
func Any[T any](typeForTest []T, any func(T) bool) bool {

	for _, v := range typeForTest {
		if any(v) {
			return true
		}
	}
	return false
}
