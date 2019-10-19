package cache

import (
	"os"
	"reg_go/exceptions"
)

// Collection коллекция кэшей
type Collection struct {
	name  string
	cache []Cache
}

// New Инициализация коллекции кэша
func (c *Collection) New() *Collection {
	var cacheCol *Collection = new(Collection)
	cacheCol.name = "data"
	return cacheCol
}

// Add Добавление в коллекцию кэш
func (c *Collection) Add(name string, dataArr []string) {

	// Создаем папку для коллекции кэша если ее нет
	err := os.MkdirAll(c.name, os.ModePerm)
	if os.IsNotExist(err) {
		exceptions.ErrorFy(err)
	}

	cache := Cache{}
	cache.add(name, dataArr, c.name)

	c.cache = append(c.cache, cache)
}
