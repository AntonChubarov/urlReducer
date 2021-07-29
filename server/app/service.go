package app

import (
	"server/config"
	"server/domain"
)

type Service struct{
	LinkValidator domain.StringValidator
	IDValidator domain.StringValidator
	Hasher domain.Hasher
	Storage domain.LinkStorage
	Config *config.ServerConfig
}

func NewService(linkValidator domain.StringValidator,
	idValidator domain.StringValidator,
	hasher domain.Hasher,
	storage domain.LinkStorage,
	sConfig *config.ServerConfig,
) *Service {
	return &Service{LinkValidator: linkValidator,
		IDValidator: idValidator,
		Hasher:      hasher,
		Storage:     storage,
		Config:      sConfig,
	}
}

func (s *Service) SaveLink (url string) (string, error) {
	if !s.LinkValidator.Validate(url) {
		return "", domain.ErrorInvalidURL
	}

	id := s.Hasher.Hash(url) // env

	err := s.Storage.SaveInitialLinkToStorage(url, id)
	if err != nil {
		return "", domain.ErrorInternal
	}

	return s.Config.Host.ServerHost + s.Config.Host.ServerStartPort + "/" + id, nil
}

func (s *Service) GetLink (id string) (string, error) {
	if !s.IDValidator.Validate(id) {
		return "", domain.ErrorInvalidShortURL
	}

	initialURL, err := s.Storage.GetInitialLinkFromStorage(id)
	if err != nil {
		return "", domain.ErrorInternal
	}

	return initialURL, nil
}