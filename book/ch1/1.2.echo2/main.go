package main

import (
	"fmt"
	"os"
)

// вывод пример мир
func main() {
	fmt.Println("app run: " + os.Args[0])

	// для строк это аналогично:
	// var s, sep string
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ", "
	}

	fmt.Println("os.Args=[" + s + "]")
}
