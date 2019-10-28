package lib

import (
	"compress/bzip2"
	"fmt"
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
	jsonRead, err := new(JSONFile).JSONEncode(data)
	if err != nil {
		return err
	}
	hashData, err := HashUint(jsonRead)
	var BlockData *Data = &Data{data, hashData}

	c.Data = append(c.Data, BlockData)
	c.Hash = hashData
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
