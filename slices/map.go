package slices

// Map converts all items in the slice by using the mapper.
func Map[S ~[]U, T, U any](a []T, mapper func(T) U) S {
	ret := make(S, len(a))
	for i, v := range a {
		ret[i] = mapper(v)
	}
	return ret
}
