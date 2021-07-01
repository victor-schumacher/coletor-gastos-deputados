package brasil_io

import (
	"coletor-gastos-deputados/stream"
	"errors"
	"github.com/rs/zerolog/log"

	"net/http"
	"path/filepath"
)

const (
	DatasetDownloadURL = "https://brasil_io.brasil.io/dataset/gastos-deputados/cota_parlamentar.csv.gz"
	datasetFile        = "brasil_io.csv.gz"
	DataFile           = "brasil_io.csv"
)

type Downloader interface {
	DownloadExtract(downloadURL string) error
}

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
	log.Info().Msg("starting download and extract")
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

	fileToExtractPath := filepath.Join(m.homePath, datasetFile)
	if err := m.io.FileWrite(fileToExtractPath, body); err != nil {
		return err
	}

	pathToExtract := filepath.Join(m.homePath, DataFile)
	if err := m.io.FileExtract(fileToExtractPath, pathToExtract); err != nil {
		return err
	}
	log.Info().Msg("finishing download and extract")

	return nil
}
