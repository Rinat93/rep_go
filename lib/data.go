package lib

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"strings"
)

// Data сами данные
type Data struct {
	Text    string
	Hash    hash.Hash32
	Readers io.Reader
}

// DataProccessing ... Массив Данных находятся тут
type DataProccessing struct {
	Data  []*Data
	depth uint64
}

// New Data
func (c *DataProccessing) New() *DataProccessing {
	var data *DataProccessing = new(DataProccessing)
	data.depth = 0
	return data
}

// CompressData - Сжатие данных
func (c *DataProccessing) CompressData(data string) io.Reader {

	b := strings.NewReader(data)
	return bzip2.NewReader(b)

}

// Add добавление в структуру данных
func (c *DataProccessing) Add(data string) {
	// Создаем хэш данных
	hashData := crc32.NewIEEE()
	hashData.Write([]byte(data))
	readerIo := c.CompressData(data)
	var BlockData *Data = &Data{data, hashData, readerIo}

	c.Data = append(c.Data, BlockData)
	c.depth++
}

// Search data
func (c *DataProccessing) Search(data string) {
	fmt.Println("Search ...")
}

// func (c *DataProccessing) Views() []byte {
// 	return []byte(c.text)
// }
