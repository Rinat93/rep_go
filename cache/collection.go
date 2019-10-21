package cache

import (
	"os"
	"reg_go/config"
	"reg_go/exceptions"
	// "sync"
)

// var wg sync.WaitGroup

// Collection коллекция кэшей
type Collection struct {
	name  string
	cache []Cache
}

// New Инициализация коллекции кэша
func (c *Collection) New() *Collection {
	var cacheCol *Collection = new(Collection)
	return cacheCol
}

// AddAsync Асинхронное добавление
func (c *Collection) AddAsync(name string, dataArr []string, res chan<- *Collection) {
	// wg.Add(1)
	c.Add(name, dataArr)
	// wg.Wait()
	res <- c
}

// Add Добавление в коллекцию кэш
func (c *Collection) Add(name string, dataArr []string) {
	directoryCollections := config.DIRECTORYCACHE + "/" + c.name
	// Создаем папку для коллекции кэша если ее нет
	err := os.MkdirAll(directoryCollections, os.ModePerm)
	if os.IsNotExist(err) {
		exceptions.ErrorFy(err)
	}

	cache := Cache{}
	cache.add(name, dataArr, directoryCollections)

	c.cache = append(c.cache, cache)
	// defer wg.Done()
}

// Remove Удаление коллекции
func (c *Collection) Remove() error {
	directoryCollections := config.DIRECTORYCACHE + "/" + c.name
	err := os.RemoveAll(directoryCollections)
	if err != nil {
		return err
	}
	return nil
}
