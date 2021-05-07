package csv

import (
	"coletor-gastos-deputados/data"
	"github.com/gocarina/gocsv"
	"os"
)

type Unmarshaler interface {
	Unmarshal(file *os.File) ([]*data.Data, error)
}

type Manager struct {
}

func NewManager() Manager {
	return Manager{}
}

func (c Manager) Unmarshal(file *os.File) ([]*data.Data, error) {
	var d []*data.Data
	if err := gocsv.UnmarshalFile(file, &d); err != nil {
		return d, err
	}
	return d, nil
}
