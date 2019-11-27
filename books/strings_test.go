package books

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	x := Contains("Ну и Привет ыфвыф афыфыавыфа", "Привет")
	// x := Contains("Привет", "Ну и Привет ыфвыф афыфыавыфа")
	fmt.Println(x)
}

func BenchmarkContains(b *testing.B) {
	x := Contains("Ну и Привет ыфвыф афыфыавыфа", "Привет")
	// x := Contains("Привет", "Ну и Привет ыфвыф афыфыавыфа")
	fmt.Println(x)
}
