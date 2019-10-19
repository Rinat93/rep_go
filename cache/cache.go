package cache

import (
	"os"
	"reg_go/exceptions"
	"reg_go/lib"
)

// Cache структура для построения кэша данных
type Cache struct {
	Name  string
	Index int64
	Data  lib.DataProccessing
}

// New Инициализация кэша
func (c *Cache) New() *Cache {
	var cache *Cache = new(Cache)
	cache.Index = 0
	return cache
}

// Add Добавить на сохранение кэша
func (c *Cache) add(name string, dataArr []string, directory string) {
	c.Name = name

	// Сохраняем данные для сохранения в память
	data := lib.DataProccessing{}
	for _, value := range dataArr {
		data.Add(value)
	}

	c.Data = data
	c.Index++
	c.Save(directory)
}

// Save Сохранить в файл
func (c *Cache) Save(directory string) {

	// Создаем файл
	file, err := os.OpenFile(directory+"/"+c.Name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	exceptions.ErrorFy(err)
	defer file.Close()

	// Записываем данные
	for _, value := range c.Data.Data {
		_, err2 := file.WriteString(value.Text + "\n")
		exceptions.ErrorFy(err2)
	}
}
