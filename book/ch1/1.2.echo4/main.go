package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// вывод пример мир
func main() {
	sep := ""

	strSlice := []string{
		"os.Args={",
	}

	for i, arg := range os.Args[1:] {
		strSlice = append(strSlice, sep+strconv.Itoa(i+1)+": "+arg)
		sep = ", "
	}
	strSlice = append(strSlice, "}")
	fmt.Println(strings.Join(strSlice, ""))
}
