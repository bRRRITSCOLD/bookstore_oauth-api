package access_token_service

import (
	access_token_domain "bookstore_oauth-api/domains/access_token"
	errors_utils "bookstore_oauth-api/utils/errors"
	"strings"
)

type AccessTokenDatabaseRepository interface {
	GetAccessTokenByID(string) (*access_token_domain.AccessToken, *errors_utils.APIError)
	CreateAccessToken(access_token_domain.AccessToken) *errors_utils.APIError
	UpdateAccessTokenExpiresByID(access_token_domain.AccessToken) *errors_utils.APIError
}

type AccessTokenService interface {
	GetAccessTokenByID(string) (*access_token_domain.AccessToken, *errors_utils.APIError)
	CreateAccessToken(access_token_domain.AccessToken) *errors_utils.APIError
	UpdateAccessTokenExpiresByID(access_token_domain.AccessToken) *errors_utils.APIError
}

type accessTokenService struct {
	accessTokenDatabaseRepository AccessTokenDatabaseRepository
}

func NewService(atDbReospitory AccessTokenDatabaseRepository) AccessTokenService {
	return &accessTokenService{
		accessTokenDatabaseRepository: atDbReospitory,
	}
}

func (atService *accessTokenService) GetAccessTokenByID(accessTokenId string) (*access_token_domain.AccessToken, *errors_utils.APIError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors_utils.NewBadRequestAPIError("invalid access token id")
	}

	accessToken, err := atService.accessTokenDatabaseRepository.GetAccessTokenByID(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (atService *accessTokenService) CreateAccessToken(accessToken access_token_domain.AccessToken) *errors_utils.APIError {
	validateAccessTokenErr := accessToken.ValidateAccessToken()
	if validateAccessTokenErr != nil {
		return validateAccessTokenErr
	}

	err := atService.accessTokenDatabaseRepository.CreateAccessToken(accessToken)
	if err != nil {
		return err
	}

	return nil
}

func (atService *accessTokenService) UpdateAccessTokenExpiresByID(accessToken access_token_domain.AccessToken) *errors_utils.APIError {
	validateAccessTokenErr := accessToken.ValidateAccessToken()
	if validateAccessTokenErr != nil {
		return validateAccessTokenErr
	}

	err := atService.accessTokenDatabaseRepository.UpdateAccessTokenExpiresByID(accessToken)
	if err != nil {
		return err
	}

	return nil
}
