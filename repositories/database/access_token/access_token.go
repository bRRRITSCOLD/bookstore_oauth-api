package access_token_database_repository

import (
	access_token_domain "bookstore_oauth-api/domains/access_token"
	errors_utils "bookstore_oauth-api/utils/errors"
)

type AccessTokenDatabaseRepository interface {
	GetByID(string) (*access_token_domain.AccessToken, *errors_utils.APIError)
}

type accessTokenDatabaseRepository struct {
}

func NewAccessTokenRepository() AccessTokenDatabaseRepository {
	return &accessTokenDatabaseRepository{}
}

func (atDbRepo *accessTokenDatabaseRepository) GetByID(string) (*access_token_domain.AccessToken, *errors_utils.APIError) {
	return nil, errors_utils.NewInternalServerAPIError("database connection not implemented")
}
