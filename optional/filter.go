package optional

func Filter[T any](a Optional[T], f func(T) bool) Optional[T] {
	if a.valid && f(a.value) {
		return a
	}
	return Optional[T]{}
}
