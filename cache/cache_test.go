package cache

import (
	"fmt"
	"testing"
)

func Test_Cache(t *testing.T) {
	cache := Collection{name: "dirs"}
	for i := 0; i < 100; i++ {
		res := make(chan *Collection, 1)
		go cache.AddAsync("test.txt", []string{"Привет", "Lol" + string(i)}, res)
		go cache.AddAsync("test2.txt", []string{"Привет1", "Lol1"}, res)
		go cache.AddAsync("test3.txt", []string{"Привет2", "Lol2"}, res)
		result := <-res
		fmt.Scanln("%+v\n", result)
	}
}

func BenchmarkCache(t *testing.B) {
	cache := Collection{name: "dirs"}
	cache.Add("test.txt", []string{"Привет", "Lol"})
	cache.Add("test2.txt", []string{"Привет1", "Lol1"})
	cache.Add("test3.txt", []string{"Привет2", "Lol2"})
	fmt.Printf("%+v\n", cache)
}
