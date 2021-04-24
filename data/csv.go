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

func NewCSVManager() CSVManager {
	return CSVManager{}
}

func (c CSVManager) UnmarshalCSV(file *os.File) ([]*Data, error) {
	var data []*Data
	if err := gocsv.UnmarshalFile(file, &data); err != nil {
		return data, err
	}
	return data, nil
}
