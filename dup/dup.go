package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filesContaining := make(map[string][]string)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			filesContaining[line] = append(filesContaining[line], filename)
		}
	}

	for line, files := range filesContaining {
		if len(files) > 1 {
			fmt.Printf("%d\t%s:\n", len(files), line)
			for _, file := range files {
				fmt.Printf("%s ", file);
			}
			fmt.Println()
		}
	}
}
