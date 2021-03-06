package stream

import (
	"bytes"
	"context"
	"github.com/codeclysm/extract"
	"io"
	"io/ioutil"
	"os"
)

type FileExtractor interface {
	FileExtract(fileToExtractPath, pathToExtract string) error
}

type FileWriter interface {
	FileWrite(path string, content []byte) error
}

type FileReader interface {
	FileRead(path string) ([]byte, error)
}

type ReadExtractWriter interface {
	FileExtract(fileToExtractPath, pathToExtract string) error
	FileWrite(path string, content []byte) error
	Read(r io.Reader) ([]byte, error)
}

type Reader interface {
	Read(r io.Reader) ([]byte, error)
}

type Manager struct {
}

func NewManager() Manager {
	return Manager{}
}

func (m Manager) FileExtract(fileToExtractPath, pathToExtract string) error {
	data, err := ioutil.ReadFile(fileToExtractPath)
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(data)
	if err = extract.Gz(context.TODO(), buffer, pathToExtract, nil); err != nil {
		return err
	}
	return nil
}

func (m Manager) FileWrite(path string, content []byte) error {
	return ioutil.WriteFile(path, content, os.ModePerm)
}

func (m Manager) FileRead(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, err
}

func (m Manager) Read(r io.Reader) ([]byte, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return b, err
}
