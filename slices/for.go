package slices

func For[T any](a []T, f func(int, T) error) error {
	for i, v := range a {
		if err := f(i, v); err != nil {
			return err
		}
	}
	return nil
}
