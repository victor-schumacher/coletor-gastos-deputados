package data

import (
	"bytes"
	"context"
	"github.com/codeclysm/extract"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

const (
	writePermission = 0700
	downloadURL     = "https://data.brasil.io/dataset/gastos-deputados/cota_parlamentar.csv.gz"
	fileToExtract   = "data.csv.gz"
	dataFile        = "data.csv"
)

type Manager struct {
	homePath string
}

func New(homePath string) Manager {
	return Manager{homePath: homePath}
}

// Download downloads a .gz with the data csv file inside,
// and write it to home path passed on constructor
func (m Manager) Download() error {
	resp, err := http.Get(downloadURL)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(m.homePath, body, writePermission); err != nil {
		return err
	}
	return nil
}

// Extract extracts the file downloaded by Download and write it as a .csv
func (m Manager) Extract() error {
	fileToExtractPath := filepath.Join(m.homePath, fileToExtract)
	data, err := ioutil.ReadFile(fileToExtractPath)
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(data)
	dataFilePath := filepath.Join(m.homePath, dataFile)
	if err = extract.Gz(context.TODO(), buffer, dataFilePath, nil); err != nil {
		return err
	}

	return nil
}
