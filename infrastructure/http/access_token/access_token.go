package access_token_http_infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	access_token_service "bookstore_oauth-api/services/access_token"
)

type AccessTokenHandler interface {
	GetByID(c *gin.Context)
}

type accessTokenHandler struct {
	accessTokenService access_token_service.AccessTokenService
}

func NewHandler(accessTokenService access_token_service.AccessTokenService) AccessTokenHandler {
	return &accessTokenHandler{
		accessTokenService: accessTokenService,
	}
}

func (atHandler accessTokenHandler) GetByID(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("accessTokenId"))

	accessToken, err := atHandler.accessTokenService.GetByID(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}
