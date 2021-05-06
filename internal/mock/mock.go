package mock

import (
	"io"
	"net/http"
	"net/http/httptest"
)

var (
	BadRequestServer = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusBadRequest)
	}))
	Server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))
)

type StreamMock struct {
}

func (StreamMock) FileExtract(fileToExtractPath, pathToExtract string) error {
	return nil
}
func (StreamMock) FileWrite(path string, content []byte) error {
	return nil
}
func (StreamMock) Read(r io.Reader) ([]byte, error) {
	return []byte("something"), nil
}
