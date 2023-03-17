package linq

func Select[T any, K any](typeForTest []T, projection func(x T) K) []K {
	result := make([]K, len(typeForTest))
	for i, v := range typeForTest {
		result[i] = projection(v)
	}

	return result
}
