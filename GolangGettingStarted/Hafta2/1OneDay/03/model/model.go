package model

type Ata struct {
	Isim    string `json:"isim"`
	Message string `json:"message"`
	Sehir   string `json:"sehir"`
	Sayi    int    `json:"sayi"`
}

type Nine struct {
	Gorev string `json:"operation"`
}
