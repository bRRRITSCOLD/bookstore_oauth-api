package access_token_service

import (
	access_token_domain "bookstore_oauth-api/domains/access_token"
	users_domain "bookstore_oauth-api/domains/users"
	access_token_database_repository "bookstore_oauth-api/repositories/database/access_token"
	users_http_repository "bookstore_oauth-api/repositories/http/users"
	"errors"
	"strings"

	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
)

type AccessTokenService interface {
	GetAccessTokenByID(string) (*access_token_domain.AccessToken, errors_utils.APIError)
	CreateAccessToken(access_token_domain.AccessTokenRequest) (*access_token_domain.AccessToken, errors_utils.APIError)
	UpdateAccessTokenExpiresByID(access_token_domain.AccessToken) errors_utils.APIError
}

type accessTokenService struct {
	accessTokenDatabaseRepository access_token_database_repository.AccessTokenDatabaseRepository
	usersHttpRepository           users_http_repository.UsersHTTPRepository
}

func NewService(
	atDbReospitory access_token_database_repository.AccessTokenDatabaseRepository,
	uHttpReospitory users_http_repository.UsersHTTPRepository,
) AccessTokenService {
	return &accessTokenService{
		accessTokenDatabaseRepository: atDbReospitory,
		usersHttpRepository:           uHttpReospitory,
	}
}

func (atService *accessTokenService) GetAccessTokenByID(accessTokenId string) (*access_token_domain.AccessToken, errors_utils.APIError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors_utils.NewBadRequestAPIError("invalid access token id", errors.New("validation error"))
	}

	accessToken, err := atService.accessTokenDatabaseRepository.GetAccessTokenByID(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (atService *accessTokenService) CreateAccessToken(accessTokenRequest access_token_domain.AccessTokenRequest) (*access_token_domain.AccessToken, errors_utils.APIError) {
	validateAccessTokenRequestErr := accessTokenRequest.ValidateAccessTokenRequest()
	if validateAccessTokenRequestErr != nil {
		return nil, validateAccessTokenRequestErr
	}

	// TODO: support noth client_credentials and password

	user, loginUserErr := atService.usersHttpRepository.LoginUser(users_domain.UserLoginRequest{
		Email:    accessTokenRequest.Username,
		Password: accessTokenRequest.Password,
	})
	if loginUserErr != nil {
		return nil, loginUserErr
	}

	// Generate a new access token:
	accessToken := access_token_domain.GetNewAccessToken(user.UserID)
	accessToken.GenerateAccessToken()

	createAccessTokenErr := atService.accessTokenDatabaseRepository.CreateAccessToken(accessToken)
	if createAccessTokenErr != nil {
		return nil, createAccessTokenErr
	}

	return &accessToken, nil
}

func (atService *accessTokenService) UpdateAccessTokenExpiresByID(accessToken access_token_domain.AccessToken) errors_utils.APIError {
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
