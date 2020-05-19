package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	err := run("./files/*.txt", "files.zip")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(glob, zipName string) error {
	f, err := os.Create(zipName)
	if err != nil {
		return fmt.Errorf("creating zip file: %w", err)
	}
	w := zip.NewWriter(f)
	filenames, err := filepath.Glob(glob)
	if err != nil {
		return fmt.Errorf("globbing files to zip: %w")
	}
	for _, filename := range filenames {
		rawFile, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("opening file to zip: %w", err)
		}
		defer rawFile.Close()
		zipFile, err := w.Create(filename)
		if err != nil {
			return fmt.Errorf("creating file in zip: %w", err)
		}
		_, err = io.Copy(zipFile, rawFile)
		if err != nil {
			return fmt.Errorf("copying file into zip: %w", err)
		}
	}
	err = w.Close()
	if err != nil {
		return fmt.Errorf("closing zip: %w", err)
	}
	return nil
}
