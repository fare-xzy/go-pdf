package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse_pdf(path string) {
	var lines []string
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
