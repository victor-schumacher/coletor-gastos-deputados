package cron

import (
	"coletor-gastos-deputados/data"
	"coletor-gastos-deputados/database/postgres/repository"
	"fmt"
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
		return
	}
	filePath := filepath.Join(dir, data.DataFile)
	fileHandle, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
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

	for r := range c {
		err := repo.Save(r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("salvou de boa")
	}
}
