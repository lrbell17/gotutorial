package rw

import (
	"errors"
	"fmt"
)

func (rec *RecordWriter) Write(b []byte) (n int, err error) {
	if rec == nil {
		return n, errors.New("Invalid parameter, nil receiver")
	}

	// TODO 3 - complete implementation of this function
	header := []byte(fmt.Sprintf(headerFmt, rec.count))
	footer := []byte(fmt.Sprintf(footerFmt, rec.count))

	b = append(header, b...)
	b = append(b, footer...)

	rec.count++
	return rec.writer.Write(b)
}
