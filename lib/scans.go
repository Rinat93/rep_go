package scans

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ev := input.Text()
		counts[ev]++
		if ev == "exit" {
			break
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
