package slices

func Map[S ~[]U, T, U any](a []T, f func(T) U) S {
	ret := make(S, len(a))
	for i, v := range a {
		ret[i] = f(v)
	}
	return ret
}
