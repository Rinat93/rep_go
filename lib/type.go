package lib

//JSONFile json файл
type JSONFile struct{}

// Data сами данные
type Data struct {
	Text string
	Hash uint32
}

// DataProccessing ... Массив Данных находятся тут
type DataProccessing struct {
	Data  []*Data
	depth uint64
	Hash  uint32
}
