package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	days := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	var outFile string
	var wg sync.WaitGroup

	if os.Args[1] == "level" {
		outFile = os.Args[2]
		if outFile != "level notice" && outFile != "level error" {
			fmt.Println("[ERROR]: No such command.\n\n[EXAMPLE]:\ngo run main.go level error.log\ngo run main.go level notice.log")
			return
		}
	}

	if os.Args[1] == "date" {
		outFile = os.Args[2]
		for _, v := range days {
			if outFile != v {
				fmt.Println("[ERROR]: no such date.\n\n[ALLOWED]:\nMonday\nTuesday\nWednesday\nThursday\nFriday\nSaturday\nSunday\n\n[EXAMPLE]:\ngo run main.go date Monday")
				return
			}
		}
		wg.Add(1)
		go timeParser(wg, outFile)
	}

	file, err := os.Open("logs.log")
	if err != nil {
		fmt.Println("[ERROR]: cannot open file logs.log")
		return
	}
	defer file.Close()

	result, err := os.OpenFile(outFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("[ERROR]: cannot open file: %s", outFile)
		return
	}
	defer result.Close()

	wg.Wait()

}

// func openFileHelper(input string) {
// 	outFile, err := os.OpenFile(input, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Printf("[ERROR]: cannot open file: %s", input)
// 		return
// 	}
// 	defer outFile.Close()

// }

func errorNoticeParser(file, result *os.File, level string) {

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if level == "error.log" {
			if strings.Contains(line, "[error]") {
				result.WriteString(line + "\n")
			}
		}

		if level == "notice.log" {
			if strings.Contains(line, "notice") {
				result.WriteString(line + "\n")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[ERROR]: failed to read file", err)
	}
}

func timeParser(wg sync.WaitGroup, day string) {
	defer wg.Done()

}
