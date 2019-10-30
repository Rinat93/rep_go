package task

import (
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
		file, err := os.OpenFile(config.DIRECTORYCACHE+"/"+t.Name+".json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
		defer file.Close()
		exceptions.ErrorFy(err)
		jsonData := new(lib.JSONFile)
		jsonSave, err := jsonData.JSONEncode(c.Pool)
		exceptions.ErrorFy(err)
		_, err = file.WriteString(jsonSave + "\n")
		exceptions.ErrorFy(err)
	}
}