package lib

import (
	"fmt"
	"testing"
)

func Test_DataProccessing(t *testing.T) {
	data := new(DataProccessing)
	data.Add("Привет")
	data.Add("Привет")
	data.Add("Привет")
	data.Add("Привет")
	fmt.Println(data)
}
