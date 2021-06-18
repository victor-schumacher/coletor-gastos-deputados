package cron

import (
	"coletor-gastos-deputados/data"
	"coletor-gastos-deputados/data/csv"
	"coletor-gastos-deputados/database/postgres"
	"coletor-gastos-deputados/database/postgres/repository"
	"github.com/go-co-op/gocron"
	"log"
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

	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// httpClient := &http.Client{
	// 	Timeout: time.Minute * 12,
	// 	}
	// sm := stream.NewManager()
	// dataManager := data.New(homeDir, httpClient, sm)
	// here
	// if err := dataManager.DownloadExtract(data.DatasetDownloadURL); err != nil {
	// 	log.Fatal(err)
	// }
	expenseRepo := repository.NewExpense(db)
	if err := csv.Unmarshal(data.DataFile, expenseRepo); err != nil {
		log.Fatal("here" + err.Error())
	}

	//log.Println("starting to save things")
	//for _, datapoint := range d {
	//	if err := er.Save(datapoint); err != nil {
	//		fmt.Println(err)
	//	}
	//}
}
