package int64file

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"os"
)

type Reader struct {
	fd  *os.File
	gz  *gzip.Reader
	buf *bufio.Reader
}

func (r *Reader) ReadInt64() (i int64, err error) {
	i, err = binary.ReadVarint(r.buf)
	return
}

func (r *Reader) Close() (err error) {
	if r.gz != nil {
		err = r.gz.Close()
		if err != nil {
			r.fd.Close()
			return
		}
	}
	err = r.fd.Close()
	return
}

func NewReader(fd *os.File, gz bool) (r *Reader, err error) {
	r = new(Reader)
	r.fd = fd
	if gz {
		var gzReader *gzip.Reader
		gzReader, err = gzip.NewReader(fd)
		r.gz = gzReader
		r.buf = bufio.NewReader(gzReader)
	} else {
		r.buf = bufio.NewReader(fd)
	}
	return
}
