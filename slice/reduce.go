package slice

func Reduce[T any](a []T, f func(T, T) T) T {
	if len(a) == 0 {
		var zero T
		return zero
	}
	ret := a[0]
	for _, v := range a[1:] {
		ret = f(ret, v)
	}
	return ret
}

func ReduceInit[T any](a []T, init T, f func(T, T) T) T {
	ret := init
	for _, v := range a {
		ret = f(ret, v)
	}
	return ret
}
