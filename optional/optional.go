package optional

type Optional[T any] interface {
	Valid() bool
	GetOr(defaults T) T
	GetOrFunc(f func() T) T

	// do not allow the user to implement Optional
	internal()
}

type Some[T any] struct {
	value T
}

func NewSome[T any](v T) Optional[T] {
	return Some[T]{
		value: v,
	}
}

func (v Some[T]) Valid() bool {
	return true
}

func (v Some[T]) Get() T {
	return v.value
}

func (v Some[T]) GetOr(defaults T) T {
	return v.value
}

func (v Some[T]) GetOrFunc(f func() T) T {
	return v.value
}

func (v Some[T]) internal() {
}

type None[T any] struct{}

func NewNone[T any]() Optional[T] {
	return None[T]{}
}

func (v None[T]) Valid() bool {
	return false
}

func (v None[T]) GetOr(defaults T) T {
	return defaults
}

func (v None[T]) GetOrFunc(f func() T) T {
	return f()
}

func (v None[T]) internal() {
}
