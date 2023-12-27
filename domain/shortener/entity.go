package shortener

type ShortURL struct {
	ID        string `json:"id"`
	Original  string `json:"original"`
	Short     string `json:"short"`
	CreatedAt int64  `json:"created_at"`
}
