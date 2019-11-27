package books

import (
	"fmt"
	"testing"
)

func BenchmarkArrays(b *testing.B) {
	var x [150000]int

	// x := Contains("Привет", "Ну и Привет ыфвыф афыфыавыфа")
	fmt.Println(x)
}
