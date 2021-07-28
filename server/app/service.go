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

func (s *Service) SaveLink (request domain.Request) (domain.Response, error) {
	if !s.LinkValidator.Validate(request.InitialURL) {
		return domain.Response{}, domain.ErrorInvalidURL
	}

	id := s.Hasher.Hash(request.InitialURL) // env

	err := s.Storage.SaveInitialLinkToStorage(request.InitialURL, id)
	if err != nil {
		return domain.Response{}, domain.ErrorInternal
	}

	return domain.Response{
		URL: s.Config.Host.ServerHost + s.Config.Host.ServerStartPort + "/" + id,
	}, nil
}

func (s *Service) GetLink (id string) (domain.Response, error) {
	if !s.IDValidator.Validate(id) {
		return domain.Response{}, domain.ErrorInvalidShortURL
	}

	initialLink, err := s.Storage.GetInitialLinkFromStorage(id)
	if err != nil {
		return domain.Response{}, domain.ErrorInternal
	}

	return domain.Response{URL: initialLink}, nil
}