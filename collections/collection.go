package collections

import (
	"os"
	"reg_go/config"
	"reg_go/exceptions"
	// "sync"
)

// var wg sync.WaitGroup

// Collection коллекция кэшей
type Collection struct {
	Name     string
	FileData []*FileData
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
	directoryCollections := config.DIRECTORYCACHE + "/" + c.Name
	// Создаем папку для коллекции кэша если ее нет
	err := os.MkdirAll(directoryCollections, os.ModePerm)
	if os.IsNotExist(err) {
		exceptions.ErrorFy(err)
	}

	file := new(FileData)
	file.Add(name, dataArr, directoryCollections)

	c.FileData = append(c.FileData, file)
	// defer wg.Done()
}

// Remove Удаление коллекции и всех вложенных файлов/папок
func (c *Collection) Remove() error {
	directoryCollections := config.DIRECTORYCACHE + "/" + c.Name
	err := os.RemoveAll(directoryCollections)
	if err != nil {
		return err
	}
	return nil
}
