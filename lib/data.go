package lib

import (
	"compress/bzip2"
	"fmt"
	"hash/crc32"
	"io"
	"strings"
)

// New Data
func (c *DataProccessing) New() *DataProccessing {
	var data *DataProccessing = new(DataProccessing)
	data.depth = 0
	return data
}

// Сжатие данных
func (c *DataProccessing) compressingFile(data string) io.Reader {
	b := strings.NewReader(data)

	return bzip2.NewReader(b)
}

// Add добавление в структуру данныхa ...interface{}
func (c *DataProccessing) Add(data interface{}) error {
	// Создаем хэш данных
	hashData := crc32.NewIEEE()
	jsonRead, err := new(JSONFile).JSONEncode(data)
	if err != nil {
		return err
	}
	hashData.Write([]byte(jsonRead))
	var BlockData *Data = &Data{jsonRead, hashData.Sum32()}

	c.Data = append(c.Data, BlockData)
	c.Hash = hashData.Sum32()
	c.depth++
	return nil
}

// Search data
func (c *DataProccessing) Search(data string) {
	fmt.Println("Search ...")
}

// func (c *DataProccessing) Views() []byte {
// 	return []byte(c.text)
// }
