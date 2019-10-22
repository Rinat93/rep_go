package data_compress

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"testing"
)

func Test_DataProccessing(t *testing.T) {
	data := []byte("Хахаха232ф")
	buff := new(bytes.Buffer)
	for _, value := range data {
		buff.ReadString(value)
		res := bzip2.NewReader(buff)
		fmt.Printf("%d\n", res)
	}
}
