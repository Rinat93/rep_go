package cache

import (
	"fmt"
	"os"
	"reg_go/lib"
)

// Cache
type Cache struct {
	name  string
	index int64
	Data  lib.DataProccessing
}

// New Инициализация
func (c *Cache) New() *Cache {
	var cache *Cache = new(Cache)
	return cache
}

// Add Добавить на сохранение
func (c *Cache) Add(name string, dataArr []string) {
	c.name = name
	data := lib.DataProccessing{}
	for _, value := range dataArr {
		data.Add(value)
	}
	c.Data = data
	c.index++
}

// Save Сохранить в файл
func (c *Cache) Save() {
	file, err := os.OpenFile("./data/"+c.name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		os.Create("./data/" + c.name)
		fmt.Println("CREATE")
	}
	defer file.Close()
	for _, value := range c.Data.Data {
		_, err := file.WriteString(value.Text + "\n")
		if err != nil {
			panic(fmt.Errorf(err.Error()))
		}
	}
}
