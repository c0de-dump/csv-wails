package pkg

import (
	"encoding/csv"
	"fmt"
	"os"
)

type fileInfo struct {
	ext     string
	path    string
	content [][]string // TODO: stream content
}

func loadFile(path string, fileType string) error {
	switch fileType {
	case "csv":
		{
			file := fileInfo{"csv", path, make([][]string, 0)}
			err := loadCSV(file)
			if err != nil {
				fmt.Printf("LoadFile Error: %s", err.Error())
				return err
			}
		}
	}

	panic("No valid fileType found")
}

func loadCSV(c fileInfo) error {
	if c.ext != "csv" {
		panic("Expecting csv file type only")
	}

	file, err := os.Open(c.path)
	if err != nil {
		panic("os.Open() failed with error:" + err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	c.content, err = reader.ReadAll()

	return err
}

func exportCSV(c fileInfo) error {
	if c.ext != "csv" {
		panic("Expecting csv file type only")
	}

	file, err := os.OpenFile(c.path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic("os.Open() failed with error:" + err.Error())
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(c.content)

	return err
}
