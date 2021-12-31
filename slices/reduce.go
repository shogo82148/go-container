package slices

func Reduce[T any](a []T, f func(T, T) T) T {
	ret := a[0]
	for i := 1; i < len(a); i++ {
		ret = f(ret, a[i])
	}
	return ret
}
