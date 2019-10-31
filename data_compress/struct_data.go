package data_compress

import (
	"compress/bzip2"
	"io"
	"strings"
)

// Сжатие данных
func CompressString(data string) io.Reader {
	b := strings.NewReader(data)
	return bzip2.NewReader(b)
}

