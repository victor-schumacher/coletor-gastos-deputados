package data

import (
	"github.com/gocarina/gocsv"
	"os"
)

type CSVUnmarshaler interface {
	UnmarshalCSV(file *os.File) error
}

type CSVManager struct {
}

func NewCSVManager(){
	
}

func (c CSVManager) UnmarshalCSV(file *os.File) error {
	var data []*Data
	if err := gocsv.UnmarshalFile(file, &data); err != nil {
		return err
	}
	return nil
}
