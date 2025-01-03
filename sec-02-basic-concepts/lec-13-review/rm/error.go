package rm

import "fmt"

type error string

const NoErr = error("")

func (e error) String() string {
	return fmt.Sprintf("MyErr: %v", string(e))
}

func New(s string) error {
	return error(s)
}
