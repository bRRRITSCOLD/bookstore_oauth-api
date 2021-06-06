package app

import (
	cassandra_client "bookstore_oauth-api/clients/cassandra"
	access_token_http_infrastructure "bookstore_oauth-api/infrastructure/http/access_token"
	access_token_database_repository "bookstore_oauth-api/repositories/database/access_token"
	users_http_repository "bookstore_oauth-api/repositories/http/users"
	access_token_service "bookstore_oauth-api/services/access_token"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	cassandraClientSession, cassandraClientErr := cassandra_client.GetSession()
	if cassandraClientErr != nil {
		panic(cassandraClientErr)
	}
	cassandraClientSession.Close()

	// repos
	atDbRepository := access_token_database_repository.NewAccessTokenRepository()
	uHttpRepository := users_http_repository.NewUsersHTTPRepository()

	// services
	atService := access_token_service.NewService(atDbRepository, uHttpRepository)

	// handlers
	atHttpInfrastructureHandler := access_token_http_infrastructure.NewHandler(atService)

	// routes
	router.POST("/oauth/access_token", atHttpInfrastructureHandler.CreateAccessToken)
	router.GET("/oauth/access_token/:accessTokenId", atHttpInfrastructureHandler.GetAccessTokenByID)
	// router.PUT("/oauth/access-token/:accessTokenId", atHttpInfrastructureHandler.GetAccessTokenByID)

	if err := router.Run(":3000"); err != nil {
		panic(err)
	}
}
