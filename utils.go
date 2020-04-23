package main

import (
	"bufio"
	"io"
	"os"
)

type fileReader struct{}

func (r *fileReader) ReadLinesFromFile(fileName string) []string {
	lines := []string{}

	if f, err := os.Open(fileName); err == nil {
		defer f.Close()
		reader := bufio.NewReader(f)

		var line string = ""

		for {
			if bytes, isPrefix, err := reader.ReadLine(); err == nil {
				line += string(bytes)

				if !isPrefix {
					lines = append(lines, line)
					line = ""
				}
			} else {
				if err == io.EOF {
					break
				} else {
					panic(err)
				}
			}
		}

	} else {
		panic(err)
	}

	return lines
}
