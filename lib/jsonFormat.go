package lib

import (
	"encoding/json"
)

/*
	ToDo нужно сделать что-бы кодировка в json работа и над синтакс. ошибками JSON, обрабатывала их и приводила
	в необходимый тип
*/

// JSONEncode - Преобразование в json строку
func (c *JSONFile) JSONEncode(data interface{}) (string, error) {
	// b := strings.NewReader(data)
	b, err := json.Marshal(data)
	return string(b), err //bzip2.NewReader(b)
}

// JSONDecode декодирование строки json
func (c *JSONFile) JSONDecode(data string) (interface{}, error) {
	// b := strings.NewReader(data)
	var btJSON []byte = []byte(data)
	var jsonDec interface{}

	err := json.Unmarshal(btJSON, &jsonDec)

	return jsonDec, err //bzip2.NewReader(b)
}

// Декодирование по определенному интерфейсу
func (c *JSONFile) JSONDecodeSerialize(data string, serialize *interface{}) (error) {
	var btJSON []byte = []byte(data)
	err := json.Unmarshal(btJSON, serialize)
	return err //bzip2.NewReader(b)
}