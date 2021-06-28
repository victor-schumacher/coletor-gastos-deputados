package cron

import (
	"coletor-gastos-deputados/data"
	"coletor-gastos-deputados/database/postgres/repository"
	"github.com/go-co-op/gocron"
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Cron struct {
	r repository.Manager
	d data.Downloader
}

func New(r repository.Manager, d data.Downloader) Cron {
	return Cron{
		r: r,
		d: d,
	}
}

func (c Cron) Start() {
	s := gocron.NewScheduler(time.UTC)
	if _, err := s.Every(7).Days().Do(c.sync); err != nil {
		log.Println("cannot start cron")
	}
	s.StartAsync()
}

func (c Cron) sync() {
	if err := c.d.DownloadExtract(data.DatasetDownloadURL); err != nil {
		log.Fatal(err)
	}

	log.Println("starting to read file and csv unmarshal")
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filePath := filepath.Join(dir, data.DataFile)
	fileHandle, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fileHandle.Close()
	readAndSave(c.r, fileHandle)
}

func readAndSave(repo repository.Manager, file *os.File) {
	c := make(chan repository.Expense)
	go func() {
		err := gocsv.UnmarshalToChan(file, c)
		if err != nil {
			log.Fatal(err)
		}
	}()
	var expenses []repository.Expense
	for expense := range c {
		expenses = append(expenses, expense)
		m := len(expenses) % 500
		if m == 0 {
			if err := repo.Save(expenses); err != nil {
				log.Println(err)
				return
			}
			expenses = nil
			log.Println("successfully saved")

		}
	}
}
