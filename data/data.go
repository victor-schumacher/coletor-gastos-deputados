package data

import "time"

type Data struct {
	Date            *time.Time
	Legislatura     string
	Partido         string
	NomeParlamentar string
	CPFCNPJ         string
	Description     string
	Provider        string
	Value           float32
}

type Downloader interface {
	DownloadExtract(downloadURL string) error
