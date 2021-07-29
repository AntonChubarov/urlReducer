package domain

//go:generate mockgen -source=interfaces.go -destination=mocks.go -package=domain

type LinkStorage interface {
	GetInitialLinkFromStorage(id string) (url string, err error)
	SaveInitialLinkToStorage(url string, id string) error
}

type StringValidator interface {
	Validate(string) bool
}

type Hasher interface {
	Hash(url string) (hash string)
}

type IService interface {
	SaveLink (string) (string, error)
	GetLink (string) (string, error)
}