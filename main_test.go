package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	var z string = "as"
	for _, i := range z {
		fmt.Println(i)
	}
}
