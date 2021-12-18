package slice

func Filter[T any](a []T, f func(T) bool) []T {
	ret := make([]T, 0, cap(a))
	for _, v := range a {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
