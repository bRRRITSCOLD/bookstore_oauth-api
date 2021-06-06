package access_token_database_repository

import (
	cassandra_client "bookstore_oauth-api/clients/cassandra"
	access_token_domain "bookstore_oauth-api/domains/access_token"

	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
)

const (
	QUERY_GET_ACCESS_TOKEN_BY_ID            = "SELECT * FROM access_tokens WHERE access_token=:id;"
	QUERY_INSERT_ACCESS_TOKEN               = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (:access_token, :user_id, :client_id, :expires);"
	QUERY_UPDATE_ACCESS_TOKEN_EXPIRES_BY_ID = "UPDATE access_tokes SET expires=:expires WHERE access_token=:id;"
)

type AccessTokenDatabaseRepository interface {
	GetAccessTokenByID(string) (*access_token_domain.AccessToken, *errors_utils.APIError)
	CreateAccessToken(access_token_domain.AccessToken) *errors_utils.APIError
	UpdateAccessTokenExpiresByID(access_token_domain.AccessToken) *errors_utils.APIError
}

type accessTokenDatabaseRepository struct {
}

func NewAccessTokenRepository() AccessTokenDatabaseRepository {
	return &accessTokenDatabaseRepository{}
}

func (atDbRepo *accessTokenDatabaseRepository) GetAccessTokenByID(id string) (*access_token_domain.AccessToken, *errors_utils.APIError) {
	cassandraClientSession, err := cassandra_client.GetSession()
	if err != nil {
		return nil, errors_utils.NewInternalServerAPIError(err.Error())
	}
	defer cassandraClientSession.Close()

	var accessToken access_token_domain.AccessToken
	q := cassandraClientSession.Query(QUERY_GET_ACCESS_TOKEN_BY_ID, []string{"id"}).Bind(id)
	if q.Iter().NumRows() == 0 {
		return nil, errors_utils.NewNotFoundAPIError("access token not found")
	}
	q.Iter().StructScan(&accessToken)
	return &accessToken, nil
}

func (atDbRepo *accessTokenDatabaseRepository) CreateAccessToken(at access_token_domain.AccessToken) *errors_utils.APIError {
	cassandraClientSession, err := cassandra_client.GetSession()
	if err != nil {
		return errors_utils.NewInternalServerAPIError(err.Error())
	}
	defer cassandraClientSession.Close()

	q := cassandraClientSession.Query(QUERY_INSERT_ACCESS_TOKEN, []string{"access_token", "user_id", "client_id", "expires"}).BindStruct(&at)
	if execReleaseErr := q.ExecRelease(); execReleaseErr != nil {
		return errors_utils.NewInternalServerAPIError(execReleaseErr.Error())
	}

	return nil
}

func (atDbRepo *accessTokenDatabaseRepository) UpdateAccessTokenExpiresByID(at access_token_domain.AccessToken) *errors_utils.APIError {
	cassandraClientSession, err := cassandra_client.GetSession()
	if err != nil {
		return errors_utils.NewInternalServerAPIError(err.Error())
	}
	defer cassandraClientSession.Close()

	q := cassandraClientSession.Query(QUERY_UPDATE_ACCESS_TOKEN_EXPIRES_BY_ID, []string{"expires", "accessToken"}).BindStruct(&at.AccessToken)
	if execReleaseErr := q.ExecRelease(); execReleaseErr != nil {
		return errors_utils.NewInternalServerAPIError(execReleaseErr.Error())
	}

	return nil
}
