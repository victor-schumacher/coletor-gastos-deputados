package main

import (
	"coletor-gastos-deputados/data"
	"coletor-gastos-deputados/data/csv"
	"coletor-gastos-deputados/database/postgres"
	"coletor-gastos-deputados/stream"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	db := postgres.NewPgManager()
	db.TestConnection()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{
		Timeout: time.Minute * 2,
	}

	sm := stream.NewManager()
	dataManager := data.New(homeDir, httpClient, sm)
	if err := dataManager.DownloadExtract(data.DatasetDownloadURL); err != nil {
		log.Fatal(err)
	}

	csvManager := csv.NewManager()
	file, err := os.Open("/home/victor/testdata.csv")
	if err != nil {

	}

	d, err := csvManager.Unmarshal(file)
	if err != nil {
		log.Fatalln(err)
	}
	for _, datapoint := range d {
		fmt.Println(datapoint)
	}
}
