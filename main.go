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

	outFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[ERROR]: cannot open file errors.log")
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "[error]") {
			outFile.WriteString(line + "\n")
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[ERROR]: reading file", err)
	}
}
