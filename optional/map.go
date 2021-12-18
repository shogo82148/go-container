package optional

func Map[T any, U any](v Optional[T], f func(T) U) Optional[U] {
	if v.valid {
		return Optional[U]{
			value: f(v.value),
			valid: true,
		}
	} else {
		return Optional[U]{}
	}
}
