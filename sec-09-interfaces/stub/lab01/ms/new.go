package ms

// NewMemStore allocates storage of size 'cap' bytes and returns a properly constructed object.
func NewMemStore(cap uint) (store *MemStore, err error) {
	// TODO 2 - implement a function that will properly initialize a MemStore and return it

	store = &MemStore{cap: cap, data: make([]byte, cap)}
	if store.data == nil {
		return nil, ErrAlloc
	}

	return
}
