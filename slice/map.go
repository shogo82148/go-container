package slice

func Map[T any, U any](f func(T) U, a []T) []U {
	ret := make([]U, len(a))
	for i, v := range a {
		ret[i] = f(v)
	}
	return ret
}
