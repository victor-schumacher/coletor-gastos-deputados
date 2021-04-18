package main

import (
	"coletor-gastos-deputados/data"
	 "coletor-gastos-deputados/stream"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{
		Timeout: time.Minute * 2,
	}

	sm := stream.NewManager()
	dataManager := data.New(homeDir, httpClient, sm)
	if err := dataManager.DownloadFile(data.DatasetDownloadURL); err != nil {
		log.Fatal(err)
	}

}
