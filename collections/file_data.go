package collections

import (
	"io/ioutil"
	"os"
	"reg_go/exceptions"
	"reg_go/lib"
)

// New Инициализация кэша
func (c *FileData) New() *FileData {
	var file *FileData = new(FileData)
	file.Index = 0
	return file
}

// AddAsync - асинхронное добавление
func (c *FileData) AddAsync(name string, dataArr []string, res chan<- *FileData) {
	c.Add(name, dataArr)
	res <- c
}

/*
Add - Добавить на сохранение кэша
direcory string - Директория где будет находится кэш
name string - имя файла/кэша
dataArr []string - Массив строк, данные.
*/
func (c *FileData) Add(name string, dataArr interface{}) {
	c.Name = name

	// Сохраняем данные для сохранения в память
	data := new(lib.DataProccessing)

	data.Add(dataArr)

	c.Data = data
	c.Index++
	// c.Save(directory)
}

// ReadFile Чтение файлов
func (c *FileData) ReadFile(directory string) error {
	fileOpen, err := os.Open(directory + "/" + c.Name)
	if err != nil {
		exceptions.ErrorFy(err)
	}
	defer fileOpen.Close()

	text, err := ioutil.ReadAll(fileOpen)

	if err != nil {
		exceptions.ErrorFy(err)
	}

	jsonData := new(lib.JSONFile)
	jsondec, err := jsonData.JSONDecode(string(text))
	if err != nil {
		return err
	}

	result := new(lib.DataProccessing)
	result.Add(jsondec)

	c.Data = result

	return nil
}

// Save Сохранить в файл
func (c *FileData) Save(directory string, dataArr interface{}) {

	// Создаем файл
	file, err := os.OpenFile(directory+"/"+c.Name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	exceptions.ErrorFy(err)
	defer file.Close()

	// Кодируем в json формат
	jsonData := new(lib.JSONFile)
	jsonSave, err := jsonData.JSONEncode(dataArr)
	exceptions.ErrorFy(err)

	// Записывает в файл
	_, err = file.WriteString(jsonSave + "\n")
	exceptions.ErrorFy(err)
	// Записываем данные
	// for _, value := range c.Data.Data {
	// 	_, err2 := file.WriteString(value.Text + "\n")
	// 	exceptions.ErrorFy(err2)
	// }
}
