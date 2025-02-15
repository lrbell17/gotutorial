package ms

// TODO 3 - implement the method Write(b []byte) (n int, err error)
// TIP: Use copy() built-in function

// Write stores len(b) bytes if there is space in the memory store. Otherwise, it stores
// up to the available space and return how many bytes were stored. If n < len(b), then
// it returns ms.ErrFull

func (store *MemStore) Write(b []byte) (n int, err error) {

	if store == nil {
		return 0, ErrNil
	}

	n = copy(store.data[store.offset:], b) // copy byte to store after offset

	store.offset += n

	if n < len(b) {
		return n, ErrFull
	}

	return
}
