package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func appendUniq(slice []string, text string) []string {
	for _, arg := range slice {
		if arg == text {
			return slice
		}
	}
	return append(slice, text)
}

func main() {
	counts := make(map[string]int)
	fileLine := make(map[string][]string)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "dup3: %s\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++

			fileLine[line] = appendUniq(fileLine[line], filename)
		}
	}

	for line, n := range counts {
		if n > 1 {
			filenames := fileLine[line]
			fmt.Printf("count=%d\ttext='%s'\nfiles=%s\n\n", n, line, strings.Join(filenames, ", "))
		}
	}
}
