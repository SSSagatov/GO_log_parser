package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	Openfile("logs.log")
}

func Openfile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		e := fmt.Sprintf("[ERROR]: cannot found file %s", file)
		return errors.New(e)
	}

}
