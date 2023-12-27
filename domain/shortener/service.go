package shortener

type ShortURLService interface {
	Shorten(original string) (string, error)
	Resolve(short string) (string, error)
}
