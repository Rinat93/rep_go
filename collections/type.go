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

// Collection коллекция кэшей
type Collection struct {
	Name     string
	FileData []*FileData
}
