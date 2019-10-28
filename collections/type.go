package collections

import (
	"reg_go/lib"
)

// FileData структура для построения кэша данных
type FileData struct {
	Name  string
	Index int64
	Data  *lib.DataProccessing
}

// Transaction Очередь на сохранения в файловую систему
type Transaction struct {
	Name     string
	FileData map[string]*FileData
}

// Collection коллекция кэшей загруженная в память
type Collection struct {
	Name     string
	FileData map[string]*FileData
	pool     map[string][]*Transaction
}
