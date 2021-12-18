package slice

func Map[T any, U any](a []T, f func(T) U) []U {
	ret := make([]U, len(a))
	for i, v := range a {
		ret[i] = f(v)
	}
	return ret
}
