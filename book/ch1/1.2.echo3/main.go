package main

import (
	"fmt"
	"os"
	"strings"
)

func appends(slice []string, list []string) []string {
	for _, args := range list {
		slice = append(slice, args)
	}
	return slice
}

// вывод пример мир
func main() {
	var strSlice []string

	strSlice = appends(strSlice, []string{
		"Args[0]=", os.Args[0],
		"\nArgs=[", strings.Join(os.Args[1:], ", "), "]",
	})

	fmt.Println(strings.Join(strSlice, ""))
}
