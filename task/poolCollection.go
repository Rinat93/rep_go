package task

import (
	"encoding/json"

	"io/ioutil"
	"os"
	"reg_go/config"
	"reg_go/exceptions"
	"reg_go/lib"
)

func (c *PoolCollection) New() *PoolCollection {
	var poolCollection *PoolCollection = new(PoolCollection)
	return poolCollection
}

/*
Add добавление
namePool имя пула который следует добавить в коллекцию
nameTaskGroup - Имя группы тасков
task - Массив задач который должен будет выполнить данная коллекция пулов
*/
func (c *PoolCollection) Add(namePool string, nameTaskGroup string, task []*Task) {
	pool := new(Pool).New()
	pool.Name = namePool
	for _,t := range task {
		err := pool.Add(nameTaskGroup, t)
		exceptions.ErrorFy(err)
	}
	c.Pool = append(c.Pool, pool)
}

// Сохранить результат в файл
func (c *PoolCollection) SaveFile() {
	err := os.MkdirAll(config.DIRECTORYCACHE, os.ModePerm)
	if os.IsNotExist(err) {
		exceptions.ErrorFy(err)
	}
	for _,t := range c.Pool {
		// Все старые и нынешниу пулы будут находится тут
		var poll []*Pool = []*Pool{}
		poll = append(poll,t)

		//  Открываем файл данно пула
		file, err := os.OpenFile(config.DIRECTORYCACHE+"/"+t.Name+".json", os.O_CREATE|os.O_RDWR, 0755)
		exceptions.ErrorFy(err)

		// Считываем сохраненые пулы с файла
		old,err := ioutil.ReadAll(file)

		exceptions.ErrorFy(err)

		jsonData := new(lib.JSONFile)
		if string(old) != "" && err == nil{
			pol2 := []*Pool{}
			err := json.Unmarshal(old, &pol2)
			exceptions.ErrorFy(err)
			poll = append(poll,pol2...)
		}

		jsonSave, err := jsonData.JSONEncode(poll)
		exceptions.ErrorFy(err)
		file.WriteAt([]byte(jsonSave),0)
		//_, err = io.WriteString(file, jsonSave )
		exceptions.ErrorFy(err)

		file.Close()
	}
}