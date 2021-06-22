package main

import (
	"coletor-gastos-deputados/cron"
	"coletor-gastos-deputados/data"
	"coletor-gastos-deputados/database/postgres"
	"coletor-gastos-deputados/database/postgres/repository"
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
		Timeout: time.Minute * 12,
	}
	sm := stream.NewManager()
	dataManager := data.New(homeDir, httpClient, sm)
	expenseRepo := repository.NewExpense(db)
	cron.New(expenseRepo, dataManager).Start()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(port, nil))
}
