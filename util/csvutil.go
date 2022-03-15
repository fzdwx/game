package util

import (
	"github.com/gocarina/gocsv"
	"os"
)

// CsvParse csv data to out
func CsvParse(filePath string, out any) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, out); err != nil {
		return err
	}

	if _, err := file.Seek(0, 0); err != nil { // Go to the start of the file
		return err
	}

	return nil
}
