package collections

import (
	"fmt"
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
			if c.FileData == nil {
				c.FileData = map[string]*FileData{}
				c.FileData[file.Name()] = data
			} else if c.FileData[file.Name()] == nil {
				c.FileData[file.Name()] = data
			} else {
				c.FileData[file.Name()].Data.Data = append(c.FileData[file.Name()].Data.Data, data.Data.Data...)
			}
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

	if c.FileData == nil {
		// Инициализируем
		c.FileData = map[string]*FileData{}
		c.FileData[name] = file
	} else if c.FileData[name] == nil {
		// Если нет ключа то создаем и присваиваем
		c.FileData[name] = file
	} else {
		// Если этот файл уже есть то расшираем его но перед добавлением проверяем уникальность по хэшу
		if c.FileData[name].Data.Hash != c.FileData[name].Data.Hash {
			c.FileData[name].Data.Data = append(c.FileData[name].Data.Data, file.Data.Data...)
		}
	}
}

// Save Сохранение
func (c *Collection) Save() {
	directoryCollections := config.DIRECTORYCACHE + "/" + c.Name
	for _, file := range c.FileData {
		fmt.Println(file)
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
