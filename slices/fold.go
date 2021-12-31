package slices

func Fold[T, U any](a []T, init U, f func(U, T) U) U {
	ret := init
	for _, v := range a {
		ret = f(ret, v)
	}
	return ret
}
