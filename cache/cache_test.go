package cache

import (
	"testing"
)

func Test_Cachr(t *testing.T) {
	cache := new(Cache)
	cache.Add("test.txt", []string{"Привет", "Lol"})
	cache.Save()
}
