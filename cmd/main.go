package main

import (
	"coletor-gastos-deputados/data"
	"log"
	"os"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dataManager := data.New(homeDir)
	if err := dataManager.Download(); err != nil {
		log.Fatal(err)
	}
	if err := dataManager.Extract(); err != nil {
		log.Fatal(err)
	}

}
