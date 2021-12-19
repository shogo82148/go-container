package optional

func Map[T any, U any](v Optional[T], f func(T) U) Optional[U] {
	switch v := v.(type) {
	case Some[T]:
		return Some[U]{f(v.value)}
	case None[T]:
		return None[U]{}
	}
	panic("never reach")
}
