package csv

import (
	"coletor-gastos-deputados/data"
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

type Unmarshaler interface {
	Unmarshal(file *os.File) ([]*data.Expense, error)
}

type Manager struct {
}

func NewManager() Manager {
	return Manager{}
}

func (c Manager) Unmarshal(file *os.File) ([]*data.Expense, error) {
	log.Println("starting csv unmarshal")
	var d []*data.Expense
	if err := gocsv.UnmarshalFile(file, &d); err != nil {
		return d, err
	}
	log.Println("finishing csv unmarshal")
	return d, nil
}
