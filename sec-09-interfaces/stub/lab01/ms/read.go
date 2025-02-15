package ms

// TODO 4 - implement the method Read(b []byte) (n int, err error)
// TIP: Use copy() built-in function

// Read stores at most len(b) bytes in b from the memory store. Otherwise, it stores
// up to the available bytes and return how many bytes were stored. If n < len(b), then
// it returns ms.ErrEmpty

func (store *MemStore) Read(b []byte) (n int, err error) {

	if store == nil {
		return 0, ErrNil
	}

	n = copy(b, store.data[:store.offset])                // copy bytes to b
	shift := copy(store.data, store.data[n:store.offset]) // shift data in mem store to beginning

	store.offset = shift

	if n < len(b) {
		return n, ErrEmpty
	}

	return
}
