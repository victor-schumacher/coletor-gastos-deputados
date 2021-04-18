package data

type Downloader interface {
	DownloadFile(downloadURL string) error
}
