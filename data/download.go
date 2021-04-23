package data

import (
	"coletor-gastos-deputados/stream"
	"errors"
	"net/http"
	"path/filepath"
)

const (
	DatasetDownloadURL = "https://data.brasil.io/dataset/gastos-deputados/cota_parlamentar.csv.gz"
	DatasetFile        = "data.csv.gz"
	dataFile           = "data.csv"
)

type Manager struct {
	homePath string
	client   *http.Client
	io       stream.ReadExtractWriter
}

func New(
	homePath string,
	client *http.Client,
	io stream.ReadExtractWriter,
) Manager {
	return Manager{
		homePath: homePath,
		client:   client,
		io:       io,
	}
}

// DownloadExtract downloads from the URL passed on parameter
//and extracts it to the constructor path
func (m Manager) DownloadExtract(downloadURL string) error {
	resp, err := m.client.Get(downloadURL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("download URL is not responding with status 200")
	}

	body, err := m.io.Read(resp.Body)
	if err != nil {
		return err
	}

	fileToExtractPath := filepath.Join(m.homePath, DatasetFile)
	if err := m.io.FileWrite(fileToExtractPath, body); err != nil {
		return err
	}

	pathToExtract := filepath.Join(m.homePath, dataFile)
	if err := m.io.FileExtract(fileToExtractPath, pathToExtract); err != nil {
		return err
	}

	return nil
}
