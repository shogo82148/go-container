package sets

type Set[T comparable] map[T]struct{}

// New returns a new set that contains items.
func New[T comparable](items ...T) Set[T] {
	set := make(Set[T], len(items))
	for _, v := range items {
		set[v] = struct{}{}
	}
	return set
}

// Add adds items into the set.
func (set Set[T]) Add(items ...T) Set[T] {
	for _, v := range items {
		set[v] = struct{}{}
	}
	return set
}

// Remove removes items from the set.
func (set Set[T]) Remove(items ...T) Set[T] {
	for _, v := range items {
		delete(set, v)
	}
	return set
}

// Contains reports whether the set contains the item.
func (set Set[T]) Contains(item T) bool {
	_, ok := set[item]
	return ok
}
