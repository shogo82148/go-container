package list

func Map[T, U any](l *List[T], f func(v T) U) *List[U] {
	ret := New[U]()
	for e := l.Front(); e != nil; e = e.Next() {
		ret.PushBack(f(e.Value))
	}
	return ret
}
