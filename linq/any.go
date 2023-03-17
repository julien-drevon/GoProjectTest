package linq

func Any[T any](typeForTest []T, any func(T) bool) bool {

	for _, v := range typeForTest {
		if any(v) {
			return true
		}
	}
	return false
}
