package main

import (
	"fmt"
	"os"
)

// вывод пример мир
func main() {
	fmt.Println("app run: " + os.Args[0])

	// для строк это аналогично:
	// s := ""
	// sep := ""
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ", "
	}
	fmt.Println("os.Args=[" + s + "]")
}
