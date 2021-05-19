package app

import (
	cassandra_client "bookstore_oauth-api/clients/cassandra"
	access_token_http_infrastructure "bookstore_oauth-api/infrastructure/http/access_token"
	access_token_database_repository "bookstore_oauth-api/repositories/database/access_token"
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

	atRepository := access_token_database_repository.NewAccessTokenRepository()
	atService := access_token_service.NewService(atRepository)
	atHttpInfrastructureHandler := access_token_http_infrastructure.NewHandler(atService)

	// ping
	// ping
	router.POST("/oauth/access-token", atHttpInfrastructureHandler.CreateAccessToken)
	router.GET("/oauth/access-token/:accessTokenId", atHttpInfrastructureHandler.GetAccessTokenByID)
	// router.PUT("/oauth/access-token/:accessTokenId", atHttpInfrastructureHandler.GetAccessTokenByID)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
