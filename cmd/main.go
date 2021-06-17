package main

import (
	"coletor-gastos-deputados/data"
	"coletor-gastos-deputados/data/csv"
	"coletor-gastos-deputados/database/postgres"
	"coletor-gastos-deputados/database/postgres/repository"
	"coletor-gastos-deputados/stream"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	r := mux.NewRouter()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	go log.Fatal(http.ListenAndServe(port, r))

	db := postgres.NewPgManager()
	db.TestConnection()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	httpClient := &http.Client{
		Timeout: time.Minute * 12,
	}
	sm := stream.NewManager()
	dataManager := data.New(homeDir, httpClient, sm)
	if err := dataManager.DownloadExtract(data.DatasetDownloadURL); err != nil {
		log.Fatal(err)
	}

	csvManager := csv.NewManager()
	path := filepath.Join(homeDir, data.DataFile)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	d, err := csvManager.Unmarshal(file)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("starting to save things")
	er := repository.NewExpense(db)
	for _, datapoint := range d {
		if err := er.Save(datapoint); err != nil {
			fmt.Println(err)
		}
	}
	log.Println("done")

}
