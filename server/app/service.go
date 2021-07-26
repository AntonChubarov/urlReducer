package app

import (
	"server/domain"
)

type Service struct{
	LinkValidator domain.StringValidator
	IDValidator domain.StringValidator
	Hasher domain.Hasher
	Storage domain.LinkStorage
}

func NewService(linkValidator domain.StringValidator,
	idValidator domain.StringValidator,
	hasher domain.Hasher,
	storage domain.LinkStorage,
) *Service {
	return &Service{LinkValidator: linkValidator,
		IDValidator: idValidator,
		Hasher: hasher,
		Storage: storage,
	}
}

func (s *Service) SaveLink (request domain.Request) (domain.Response, error) {
	if !s.LinkValidator.Validate(request.InitialURL) {
		return domain.Response{}, domain.ErrorInvalidURL
	}

	id := s.Hasher.Hash(request.InitialURL, 7) // env

	err := s.Storage.SaveInitialLinkToStorage(request.InitialURL, id)
	if err != nil {
		return domain.Response{}, domain.ErrorInternal
	}

	return domain.Response{
		URL: domain.WebHost + "/" + id,
	}, nil
}

func (s *Service) GetLink (id string) (string, error) {
	if !s.IDValidator.Validate(id) {
		return "", domain.ErrorInvalidShortURL
	}

	initialLink, err := s.Storage.GetInitialLinkFromStorage(id)
	if err != nil {
		return "", domain.ErrorInternal
	}

	return initialLink, nil
}