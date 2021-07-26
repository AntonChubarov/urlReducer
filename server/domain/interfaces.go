package domain

type LinkStorage interface {
	GetInitialLinkFromStorage(id string) (url string, err error)
	SaveInitialLinkToStorage(url string, id string) error
}

type StringValidator interface {
	Validate(string) bool
}

type Hasher interface {
	Hash(url string, numOfSymbols int) (hash string)
}