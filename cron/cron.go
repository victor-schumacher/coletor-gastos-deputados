package cron

import (
	"coletor-gastos-deputados/brasil_io"
	"coletor-gastos-deputados/database/postgres"
	"coletor-gastos-deputados/database/postgres/repository"
	"github.com/go-co-op/gocron"
	"github.com/gocarina/gocsv"
	"github.com/rs/zerolog/log"

	"os"
	"path/filepath"
	"time"
)

type Cron struct {
	r repository.Manager
	d brasil_io.Downloader
}

func New(r repository.Manager, d brasil_io.Downloader) Cron {
	return Cron{
		r: r,
		d: d,
	}
}

func (c Cron) Start() {
	s := gocron.NewScheduler(time.UTC)
	if _, err := s.Every(7).Days().Do(c.sync); err != nil {
		log.Err(err).Msg("cannot start cron")
	}
	s.StartAsync()
}

func (c Cron) sync() {
	if err := c.d.DownloadExtract(brasil_io.DatasetDownloadURL); err != nil {
		log.Fatal().Err(err)
	}

	log.Info().Msg("starting to read file and csv unmarshal")
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal().Err(err)
	}
	filePath := filepath.Join(dir, brasil_io.DataFile)
	fileHandle, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal().Err(err)
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
			log.Fatal().Err(err)
		}
	}()
	var expenses []repository.Expense
	for expense := range c {
		expenses = append(expenses, expense)
		if len(expenses)%postgres.BatchSize == 0 {
			repo.Save(expenses)
		}
		expenses = nil
	}
}
