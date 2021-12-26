package list

func (l *List[T]) Filter(f func(e *Element[T]) bool) *List[T] {
	ret := New[T]()
	for e := l.Front(); e != nil; e = e.Next() {
		if f(e) {
			ret.PushBack(e.Value)
		}
	}
	return ret
}
