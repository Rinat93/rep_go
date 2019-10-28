package collections

import (
	"fmt"
	"testing"
)

// Асинхронное создание кэша
func TestFileAsync(t *testing.T) {
	file := Collection{Name: "dirs"}
	file.ReadFile() // Инициализация уже созданных данных
	var result *Collection

	for i := 0; i < 100; i++ {
		res := make(chan *Collection)
		go file.AddAsync("test4.json", []string{"Привет", "Lol" + string(i)}, res)
		go file.AddAsync("test5.json", []string{"Привет1", "Lol1"}, res)
		go file.AddAsync("test6.json", []string{"Привет2", "Lol2"}, res)
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
		file.Add("test.json", []string{"Привет", "Lol" + string(i)})
		file.Add("test2.json", []string{"Привет1", "Lol1"})
		file.Add("test3.json", []string{"Привет2", "Lol2"})
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
