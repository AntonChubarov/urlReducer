package domain

type LinkStorage interface {
	GetInitialLinkFromStorage(id string) (url string, err error)
	SaveInitialLinkToStorage(url string, id string) error
}

type Validator interface {
	ValidateInitialURL(url string) bool
	ValidateID(id string) bool
}

type Hasher interface {
	Hash(url string, numOfSymbols int) (hash string)
}
