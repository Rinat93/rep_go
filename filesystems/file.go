package filesystems

import (
	"compress/bzip2"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reg_go/config"
	"reg_go/exceptions"
	"reg_go/lib"
	"strings"
)

func (c *Files) compressingFile(data string) io.Reader {
	b := strings.NewReader(data)
	return bzip2.NewReader(b)
}

// Сохранить в файл
func (c *Files) SaveFile() {
	err := os.MkdirAll(config.DIRECTORYCACHE, os.ModePerm)
	if os.IsNotExist(err) {
		exceptions.ErrorFy(err)
	}
	file, err := os.OpenFile(config.DIRECTORYCACHE+"/"+c.Name+".json", os.O_CREATE|os.O_RDWR, 0755)
	old,err := ioutil.ReadAll(file)
	jsonData := new(lib.JSONFile)

	if string(old) != "" && err == nil {
		fileR := []*DataBody{}
		err := json.Unmarshal(old, &fileR)
		exceptions.ErrorFy(err)
		c.Data = append(c.Data,fileR...)
	}

	jsonSave, err := jsonData.JSONEncode(c.Data)
	exceptions.ErrorFy(err)
	file.Close()

	_,err = file.WriteAt([]byte(jsonSave),0)
	exceptions.ErrorFy(err)
}

// Выгрузка данных из файла
func ReadFile(directory string) ([]*Files,error){
	var poll []*Files = []*Files{}
	files, err := ioutil.ReadDir(config.DIRECTORYCACHE)
	if err != nil {
		return nil,err
	}
	for _, file := range files {
		if file.IsDir() == false {
			file, err := os.OpenFile(config.DIRECTORYCACHE+"/"+directory+"/"+file.Name(), os.O_RDWR, 0755)
			if err != nil {
				return  nil,err
			}
			data,err := ioutil.ReadAll(file)
			exceptions.ErrorFy(err)
			err = json.Unmarshal(data, &poll)
			exceptions.ErrorFy(err)
		}
	}

	return poll,nil
}
