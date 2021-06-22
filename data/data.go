package data

type Downloader interface {
	DownloadExtract(downloadURL string) error
}
