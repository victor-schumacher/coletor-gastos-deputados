package csv

import (
	"coletor-gastos-deputados/database"
	"coletor-gastos-deputados/database/postgres/repository"
	"fmt"
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"path/filepath"
)



func Unmarshal(fileName string, repo repository.ExpenseRepo) error {
	log.Println("starting to read file and csv unmarshal")
	dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := filepath.Join(dir, fileName)
	fileHandle, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	c := make(chan database.Expense)

	go func() {
		err = gocsv.UnmarshalToChan(fileHandle, c)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for r := range c {
		err = repo.Save(r); if err != nil {
			fmt.Println(err)
		}
		fmt.Println("salvou de boa")
	}

	return nil
}
