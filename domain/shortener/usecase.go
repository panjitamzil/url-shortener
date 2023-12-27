package shortener

type ShortURLUseCase interface {
	ShortenURL(original string) (string, error)
	ResolveURL(short string) (string, error)
}
