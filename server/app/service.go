package app

import (
	"server/domain"
	"server/infrastructure/dal"
)

func SaveLink (request domain.Request) (domain.Response, error) {
	linkValidator := NewLinkValidator()

	if !linkValidator.ValidateInitialURL(request.InitialURL) {
		return domain.Response{}, domain.ErrorInvalidURL
	}

	linkHasher := NewLinkHasher()

	id := linkHasher.Hash(request.InitialURL, 7)

	if !linkValidator.ValidateID(id) {
		return domain.Response{}, domain.ErrorInternal
	}

	storage := dal.NewDatabaseConnector()

	err := storage.SaveInitialLinkToStorage(request.InitialURL, id)
	if err != nil {
		return domain.Response{}, domain.ErrorInternal
	}

	return domain.Response{
		domain.WebHost + "/" + id,
	}, nil
}

func GetLink (id string) (domain.Response, error) {
	linkValidator := NewLinkValidator()

	if !linkValidator.ValidateID(id) {
		return domain.Response{}, domain.ErrorInvalidShortURL
	}

	storage := dal.NewDatabaseConnector()

	initialLink, err := storage.GetInitialLinkFromStorage(id)
	if err != nil {
		return domain.Response{}, domain.ErrorInternal
	}

	if !linkValidator.ValidateInitialURL(initialLink) {
		return domain.Response{}, domain.ErrorInternal
	}

	return domain.Response{
		initialLink,
	}, nil
}