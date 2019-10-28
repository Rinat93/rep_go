package collections

import (
	"reg_go/config"
)

/*
	Транзакция сохраняет в файловую систему данные при условии их уникальности по хешу,
	доделать!
*/

// New инициализация
func (c *Transaction) New() *Transaction {
	var trans *Transaction = new(Transaction)
	return trans
}

// Save Сохранить в файловую систему
func (c *Transaction) Save() {
	directoryCollections := config.DIRECTORYCACHE + "/" + c.Name
	for _, file := range c.FileData {
		file.Save(directoryCollections, c.FileData)
	}
}
