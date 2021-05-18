package access_token_domain

import (
	errors_utils "bookstore_oauth-api/utils/errors"
)

type AccessTokenRepository interface {
	GetByID(string) (*AccessToken, *errors_utils.APIError)
}

type AccessTokenService interface {
	GetByID(string) (*AccessToken, *errors_utils.APIError)
}

type accessTokenService struct {
	accessTokenRepository AccessTokenRepository
}

func NewService(atReospitory AccessTokenRepository) AccessTokenService {
	return &accessTokenService{
		accessTokenRepository: atReospitory,
	}
}

func (atService *accessTokenService) GetByID(accessTokenId string) (*AccessToken, *errors_utils.APIError) {
	accessToken, err := atService.accessTokenRepository.GetByID(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
