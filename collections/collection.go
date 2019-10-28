package collections

import (
	"io/ioutil"
	"os"
	"reg_go/config"
	"reg_go/exceptions"
)

// инициализация уже существующих коллекции
func initDirectory() ([]Collection, error) {
	data, err := ioutil.ReadDir(config.DIRECTORYCACHE + "/")
	var result []Collection
	if err != nil {
		return []Collection{}, err
	}
	for _, direct := range data {
		if direct.IsDir() {
			collection := Collection{Name: direct.Name()}

			result = append(result, collection)
		}
	}
	return result, nil
}

// New Инициализация коллекции кэша
func (c *Collection) New() *Collection {
	var cacheCol *Collection = new(Collection)
	// c.initDirectory()
	return cacheCol
}

// ReadFile Чтение всех файлов в коллекции и сохранении данных в ОЗУ
func (c *Collection) ReadFile() error {
	files, err := ioutil.ReadDir(config.DIRECTORYCACHE + "/" + c.Name)
	if err != nil {
		return nil
	}
	for _, file := range files {
		if file.IsDir() == false {
			data := &FileData{Name: file.Name()}
			err := data.ReadFile(config.DIRECTORYCACHE + "/" + c.Name)
			if err != nil {
				exceptions.ErrorFy(err)

			}
			c.FileData = append(c.FileData, data)
		}
	}

	return nil
}

// AddAsync Асинхронное добавление
func (c *Collection) AddAsync(name string, dataArr interface{}, res chan<- *Collection) {
	// wg.Add(1)
	c.Add(name, dataArr)
	// wg.Wait()
	res <- c
}

// Add Добавление в коллекцию кэш
func (c *Collection) Add(name string, dataArr interface{}) {
	directoryCollections := config.DIRECTORYCACHE + "/" + c.Name
	// Создаем папку для коллекции кэша если ее нет
	err := os.MkdirAll(directoryCollections, os.ModePerm)
	if os.IsNotExist(err) {
		exceptions.ErrorFy(err)
	}

	file := new(FileData)
	file.Add(name, dataArr)
	/*
		Если файл найден то нужно расширить его структуру а не создавать новую!
	*/
	// for _, k := range c.FileData {
	// 	if k.Name == name {
	// 		k.Data.Data = append(k.Data.Data, file.Data.Data)
	// 	}
	// }
	c.FileData = append(c.FileData, file)

	// defer wg.Done()
}

// Save Сохранение
func (c *Collection) Save() {
	directoryCollections := config.DIRECTORYCACHE + "/" + c.Name
	for _, file := range c.FileData {
		file.Save(directoryCollections, c.FileData)
	}
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
