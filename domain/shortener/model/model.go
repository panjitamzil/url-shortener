package model

type Payload struct {
	URL string `json:"url"`
}

type ShortURL struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}
