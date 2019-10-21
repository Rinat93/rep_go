package collections

import (
	"fmt"
	"testing"
)

// Асинхронное создание кэша
func TestFileAsync(t *testing.T) {
	file := Collection{Name: "dirs"}
	var result *Collection

	for i := 0; i < 100; i++ {
		res := make(chan *Collection)
		go file.AddAsync("test.txt", []string{"Привет", "Lol" + string(i)}, res)
		go file.AddAsync("test2.txt", []string{"Привет1", "Lol1"}, res)
		go file.AddAsync("test3.txt", []string{"Привет2", "Lol2"}, res)
		result = <-res
		fmt.Println("--RES--")
		fmt.Printf("%+v\n", result)
		fmt.Println("--END--")
	}

}

// Синхронное создание кэша
func Test_FileSync(t *testing.T) {
	file := Collection{Name: "dirs2"}
	for i := 0; i < 100; i++ {
		file.Add("test.txt", []string{"Привет", "Lol" + string(i)})
		file.Add("test2.txt", []string{"Привет1", "Lol1"})
		file.Add("test3.txt", []string{"Привет2", "Lol2"})
	}
	err := file.Remove()
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkFile(t *testing.B) {
	file := Collection{Name: "dirs3"}
	file.Add("test.txt", []string{"Привет", "Lol"})
	file.Add("test2.txt", []string{"Привет1", "Lol1"})
	file.Add("test3.txt", []string{"Привет2", "Lol2"})
	fmt.Printf("%+v\n", file)
}
