package int64file

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"os"
)

type Writer struct {
	fd  *os.File
	gz  *gzip.Writer
	buf *bufio.Writer
}

func (w *Writer) WriteInt64(i int64) (err error) {
	b := make([]byte, 10)
	l := binary.PutVarint(b, i)
	_, err = w.buf.Write(b[0:l])
	return
}

func (w *Writer) Close() (err error) {
	var bufErr, gzErr, fdErr error

	bufErr = w.buf.Flush()
	if w.gz != nil {
		gzErr = w.gz.Close()
	}
	fdErr = w.fd.Close()

	switch {
	case bufErr != nil:
		err = bufErr
	case gzErr != nil:
		err = gzErr
	case fdErr != nil:
		err = fdErr
	}
	return
}

func NewWriter(fd *os.File, gz bool) (w *Writer, err error) {
	w = new(Writer)
	w.fd = fd
	if gz {
		gzWriter := gzip.NewWriter(fd)
		w.gz = gzWriter
		w.buf = bufio.NewWriter(gzWriter)
	} else {
		w.buf = bufio.NewWriter(fd)
	}
	return
}
