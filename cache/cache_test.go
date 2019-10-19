package cache

import (
	"fmt"
	"testing"
)

func Test_Cache(t *testing.T) {
	cache := Collection{name: "dirs"}
	cache.Add("test.txt", []string{"Привет", "Lol"})
	cache.Add("test2.txt", []string{"Привет1", "Lol1"})
	cache.Add("test3.txt", []string{"Привет2", "Lol2"})
	fmt.Printf("%+v\n", cache)
}
