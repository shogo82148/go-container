package list

// For calls the f for each items in the list.
// If the f returns an error, For stops the iteration and return the error.
func (l *List[T]) For(f func(e *Element[T]) error) error {
	for e := l.Front(); e != nil; e = e.Next() {
		if err := f(e); err != nil {
			return err
		}
	}
	return nil
}
