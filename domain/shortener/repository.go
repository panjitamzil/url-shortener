package shortener

type ShortURLRepository interface {
	Find(id string) (*ShortURL, error)
	FindByShort(short string) (*ShortURL, error)
	Save(url *ShortURL) error
}
