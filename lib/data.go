package lib

import (
	"fmt"
	"hash"
	"hash/crc32"
)

// Data сами данные
type Data struct {
	Text string
	Hash hash.Hash32
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

// Add добавление в структуру данных
func (c *DataProccessing) Add(data string) {
	// Создаем хэш данных
	hashData := crc32.NewIEEE()
	hashData.Write([]byte(data))
	var BlockData *Data = &Data{data, hashData}

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
