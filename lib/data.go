package lib

import (
	"crypto/sha1"
	"fmt"
	"hash"
)

// Data
type Data struct {
	text string
	hash hash.Hash
}

// DataProccessing ... Данные находятся тут
type DataProccessing struct {
	data  []Data
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
	fmt.Println("Save data")
	// Создаем хэш данных
	hashData := sha1.New()
	hashData.Write([]byte(data))
	var BlockData Data = Data{data, hashData}

	c.data = append(c.data, BlockData)
	c.depth++
}

// Search data
func (c *DataProccessing) Search(data string) {
	fmt.Println("Search ...")

}
