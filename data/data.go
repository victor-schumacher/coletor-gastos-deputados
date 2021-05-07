package data

type Datas []Data

type Data struct {
	Date            string  `csv:"datemissao"`
	Legislatura     string  `csv:"nulegislatura"`
	Partido         string  `csv:"sgpartido"`
	NomeParlamentar string  `csv:"txnomeparlamentar"`
	CPFCNPJ         string  `csv:"txtcnpjcpf"`
	Description     string  `csv:"txtdescricao"`
	Provider        string  `csv:"txtfornecedor"`
	Value           float32 `csv:"vlrliquido"`
}

type Downloader interface {
	DownloadExtract(downloadURL string) error
}
