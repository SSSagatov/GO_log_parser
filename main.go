package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("logs.log")
	if err != nil {
		fmt.Println("[ERROR]: cannot open file logs.log")
		return
	}
	defer file.Close()

	outFile := os.Args[1]

	result, err := os.OpenFile(outFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("[ERROR]: cannot open file: %s", outFile)
		return
	}
	defer result.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if outFile == "error.log" {
			if strings.Contains(line, "[error]") {
				result.WriteString(line + "\n")
			}
		}

		if outFile == "notice.log" {
			if strings.Contains(line, "[notice]") {
				result.WriteString(line + "\n")
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[ERROR]: reading file", err)
	}
}
