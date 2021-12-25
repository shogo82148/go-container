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

// Equal reports the other has same items.
func (set Set[T]) Equal(other Set[T]) bool {
	if len(set) != len(other) {
		return false
	}

	for v := range set {
		if _, ok := other[v]; !ok {
			return false
		}
	}
	return true
}

// Map converts all items in the set by using the mapper.
func Map[T, U comparable](set Set[T], mapper func(v T) U) Set[U] {
	ret := make(Set[U], len(set))
	for v := range set {
		ret[mapper(v)] = struct{}{}
	}
	return ret
}
