package main

import (
	"coletor-gastos-deputados/data/csv"
	"coletor-gastos-deputados/database/postgres"
	"coletor-gastos-deputados/database/postgres/repository"
	"fmt"
	"log"
	"os"
)

func main() {
	db := postgres.NewPgManager()
	db.TestConnection()

	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// httpClient := &http.Client{
	// 	Timeout: time.Minute * 2,
	// }

	// sm := stream.NewManager()
	// dataManager := data.New(homeDir, httpClient, sm)
	// if err := dataManager.DownloadExtract(data.DatasetDownloadURL); err != nil {
	// 	log.Fatal(err)
	// }

	csvManager := csv.NewManager()
	file, err := os.Open("/home/victor/testdata.csv")
	if err != nil {
		log.Fatalln(err)
	}

	d, err := csvManager.Unmarshal(file)
	if err != nil {
		log.Fatalln(err)
	}

	er := repository.NewExpense(db)
	for _, datapoint := range d {
		if err := er.Save(datapoint); err != nil {
			fmt.Println(err)
		}
	}
}
