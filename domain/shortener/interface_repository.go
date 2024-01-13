package shortener

type RepositoryInterface interface {
	SetKey(key, value string) error
	Exist(key string) (int64, error)
	GetKey(key string) (string, error)
}
