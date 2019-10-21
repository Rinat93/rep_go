package cache

import (
	"fmt"
	"testing"
)

// Асинхронное создание кэша
func TestCacheSync(t *testing.T) {
	cache := Collection{name: "dirs"}
	var result *Collection

	for i := 0; i < 100; i++ {
		res := make(chan *Collection)
		go cache.AddAsync("test.txt", []string{"Привет", "Lol" + string(i)}, res)
		go cache.AddAsync("test2.txt", []string{"Привет1", "Lol1"}, res)
		go cache.AddAsync("test3.txt", []string{"Привет2", "Lol2"}, res)
		result = <-res
		fmt.Println("--RES--")
		fmt.Printf("%+v\n", result)
		fmt.Println("--END--")
	}

}

// Синхронное создание кэша
func Test_CacheSync(t *testing.T) {
	cache := Collection{name: "dirs2"}
	for i := 0; i < 100; i++ {
		cache.Add("test.txt", []string{"Привет", "Lol" + string(i)})
		cache.Add("test2.txt", []string{"Привет1", "Lol1"})
		cache.Add("test3.txt", []string{"Привет2", "Lol2"})
	}
	err := cache.Remove()
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkCache(t *testing.B) {
	cache := Collection{name: "dirs3"}
	cache.Add("test.txt", []string{"Привет", "Lol"})
	cache.Add("test2.txt", []string{"Привет1", "Lol1"})
	cache.Add("test3.txt", []string{"Привет2", "Lol2"})
	fmt.Printf("%+v\n", cache)
}
