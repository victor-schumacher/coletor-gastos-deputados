package main

import (
	"coletor-gastos-deputados/brasil_io"
	"coletor-gastos-deputados/cron"
	"coletor-gastos-deputados/database/postgres"
	"coletor-gastos-deputados/database/postgres/repository"
	"coletor-gastos-deputados/stream"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

func main() {
	db := postgres.NewPgManager()
	db.TestConnection()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal().Err(err)
	}
	httpClient := &http.Client{
		Timeout: time.Minute * 12,
	}
	sm := stream.NewManager()
	dataManager := brasil_io.New(homeDir, httpClient, sm)
	expenseRepo := repository.NewExpense(db)
	cron.New(expenseRepo, dataManager).Start()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal().Err(err)
	}

}
