package cron

import (
	"coletor-gastos-deputados/data"
	"coletor-gastos-deputados/data/csv"
	"coletor-gastos-deputados/database/postgres"
	"coletor-gastos-deputados/database/postgres/repository"
	"coletor-gastos-deputados/stream"
	"github.com/go-co-op/gocron"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	s := gocron.NewScheduler(time.UTC)
	if _, err := s.Every(7).Days().Do(sync); err != nil {
		log.Println("cannot start cron")
	}
	s.StartAsync()
}

func sync() {
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
	expenseRepo := repository.NewExpense(db)
	if err := csv.Unmarshal(data.DataFile, expenseRepo); err != nil {
		log.Fatal("here" + err.Error())
	}
}
