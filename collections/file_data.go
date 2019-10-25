package collections

import (
	"os"
	"reg_go/exceptions"
	"reg_go/lib"
)

// FileData структура для построения кэша данных
type FileData struct {
	Name  string
	Index int64
	Data  *lib.DataProccessing
}

// New Инициализация кэша
func (c *FileData) New() *FileData {
	var file *FileData = new(FileData)
	file.Index = 0
	return file
}

// AddAsync - асинхронное добавление
func (c *FileData) AddAsync(name string, dataArr []string, directory string, res chan<- *FileData) {
	c.Add(name, dataArr, directory)
	res <- c
}

/*
	Add - Добавить на сохранение кэша
	direcory string - Директория где будет находится кэш
	name string - имя файла/кэша
	dataArr []string - Массив строк, данные.
*/
func (c *FileData) Add(name string, dataArr []string, directory string) {
	c.Name = name

	// Сохраняем данные для сохранения в память
	data := new(lib.DataProccessing)
	for _, value := range dataArr {
		data.Add(value)
	}

	c.Data = data
	c.Index++
	c.Save(directory)
}

// Save Сохранить в файл
func (c *FileData) Save(directory string) {

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
