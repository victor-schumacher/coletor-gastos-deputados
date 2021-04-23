package data

import (
	"coletor-gastos-deputados/internal/mock"
	"coletor-gastos-deputados/stream"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	table := []struct {
		name        string
		client      *http.Client
		stream      stream.ReadExtractWriter
		URL         string
		wantedError string
	}{
		{
			name:   "full success download",
			client: mock.Server.Client(),
			stream: mock.StreamMock{},
			URL: DatasetDownloadURL,
		},
		{
			name:   "bad request from server",
			client: mock.BadRequestServer.Client(),
			URL:    mock.BadRequestServer.URL,
			stream: mock.StreamMock{},
			wantedError: "download URL is not responding with status 200",
		},
		{
			name:   "bad url",
			client: mock.BadRequestServer.Client(),
			URL:    "",
			stream: mock.StreamMock{},
			wantedError: `Get "": unsupported protocol scheme ""`,
		},
	}
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			data := New("", tt.client, tt.stream)
			err := data.DownloadExtract(tt.URL)

			if err != nil {
				assert.EqualError(t, err, tt.wantedError)
				return
			}
			assert.Equal(t, nil, err)
		})
	}
}
