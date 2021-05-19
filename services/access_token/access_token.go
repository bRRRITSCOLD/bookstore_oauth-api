package access_token_domain

import (
	access_token_domain "bookstore_oauth-api/domains/access_token"
	errors_utils "bookstore_oauth-api/utils/errors"
)

type AccessTokenDatabaseRepository interface {
	GetByID(string) (*access_token_domain.AccessToken, *errors_utils.APIError)
}

type AccessTokenService interface {
	GetByID(string) (*access_token_domain.AccessToken, *errors_utils.APIError)
}

type accessTokenService struct {
	accessTokenDatabaseRepository AccessTokenDatabaseRepository
}

func NewService(atDbReospitory AccessTokenDatabaseRepository) AccessTokenService {
	return &accessTokenService{
		accessTokenDatabaseRepository: atDbReospitory,
	}
}

func (atService *accessTokenService) GetByID(accessTokenId string) (*access_token_domain.AccessToken, *errors_utils.APIError) {
	accessToken, err := atService.accessTokenDatabaseRepository.GetByID(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
