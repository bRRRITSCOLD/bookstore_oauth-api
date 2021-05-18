package app

import (
	access_token_http_controllers "bookstore_oauth-api/controllers/http/access_token"
	access_token_domain "bookstore_oauth-api/domains/access_token"
	access_token_database_repository "bookstore_oauth-api/repositories/database/access_token"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	atRepository := access_token_database_repository.NewAccessTokenRepository()
	atService := access_token_domain.NewService(atRepository)
	atHandler := access_token_http_controllers.NewHandler(atService)

	// ping
	router.GET("/oauth/access-token/:accessTokenId", atHandler.GetByID)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
