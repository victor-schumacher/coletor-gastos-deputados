package main

import (
	"coletor-gastos-deputados/data"
	"fmt"
	"log"
	"os"
)

func main() {
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// httpClient := &http.Client{
	// 	Timeout: time.Minute * 2,
	// }
	//
	// sm := stream.NewManager()
	// dataManager := data.New(homeDir, httpClient, sm)
	// if err := dataManager.DownloadExtract(data.DatasetDownloadURL); err != nil {
	// 	log.Fatal(err)
	// }

	csv := data.NewCSVManager()
	file, err := os.Open("/home/victor/testdata.csv")
	if err != nil {

	}

	 d, err := csv.UnmarshalCSV(file);
	 if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(d[0])
}
