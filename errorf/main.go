package main

import (
	"errors"
	"fmt"
	"os"
)

var osCreate = os.Create

func init() {
	// Replacing vars for demo purposes
	osCreate = func(filename string) (*os.File, error) {
		if filename == "file1.txt" {
			return &os.File{}, nil
		}
		return nil, errors.New("hard drive is full")
	}
}

func main() {
	err := Zip()
	fmt.Println(err)

	err = CreateDemo()
	fmt.Println(err)
}

func Zip() error {
	_, err := osCreate("files.zip")
	if err != nil {
		return fmt.Errorf("creating zip file: %w", err)
	}
	return nil
}

func CreateDemo() error {
	errorf := func(err error) error { return fmt.Errorf("creating demo: %w", err) }

	_, err := osCreate("file1.txt")
	if err != nil {
		return errorf(err)
	}

	// ...

	err = Zip()
	if err != nil {
		return errorf(err)
	}
	return nil
}
