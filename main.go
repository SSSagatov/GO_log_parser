package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	days := map[int]string{
		1: "Mon",
		2: "Tue",
		3: "Wed",
		4: "Thu",
		5: "Fri",
		6: "Sat",
		7: "Sun",
	}

	var wg sync.WaitGroup

	modifier := os.Args[1]
	outFile := os.Args[2]

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

	if modifier == "level" {
		if outFile != "notice.log" && outFile != "error.log" {
			fmt.Println("[ERROR]: No such command.\n\n[EXAMPLE]:\ngo run main.go level error.log\ngo run main.go level notice.log")
			return
		}
		wg.Add(1)
		go errorNoticeParser(&wg, file, result, outFile)
	}

	if modifier == "date" {
		for _, v := range days {
			if outFile != v {
				fmt.Println("[ERROR]: no such date.\n\n[ALLOWED]:\nMonday\nTuesday\nWednesday\nThursday\nFriday\nSaturday\nSunday\n\n[EXAMPLE]:\ngo run main.go date Monday")
				return
			}
		}
		wg.Add(1)
		go dayParser(&wg, file, result, outFile)
	}

	wg.Wait()

	fmt.Println("[SUCCESS]: file successfuly parsed!")
}

func errorNoticeParser(wg *sync.WaitGroup, file, result *os.File, level string) {
	defer wg.Done()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if level == "error.log" {
			if strings.Contains(line, "[error]") {
				result.WriteString(line + "\n")
			}
		}

		if level == "notice.log" {
			if strings.Contains(line, "[notice]") {
				result.WriteString(line + "\n")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[ERROR]: failed to read file", err)
	}
}

func dayParser(wg *sync.WaitGroup, file, result *os.File, day string) {
	defer wg.Done()

	days := map[int]string{
		1: "Mon",
		2: "Tue",
		3: "Wed",
		4: "Thu",
		5: "Fri",
		6: "Sat",
		7: "Sun",
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for _, v := range days {
			if v == day {
				if strings.Contains(line, day) {
					result.WriteString(line + "\n")
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[ERROR]: failed to read file", err)
	}
}

func yearParser(wg *sync.WaitGroup, file, result *os.File, year int) {
	strYear := strconv.Itoa(year)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, strYear) {
			result.WriteString(line + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[ERROR]: failed to read file", err)
	}
}
