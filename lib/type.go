package lib

import (
	"hash"
)

//JSONFile json файл
type JSONFile struct{}

// Data сами данные
type Data struct {
	Text string
	Hash hash.Hash32
}

// DataProccessing ... Массив Данных находятся тут
type DataProccessing struct {
	Data  []*Data
	depth uint64
}
