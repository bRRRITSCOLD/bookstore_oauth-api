package access_token_http_infrastructure

import (
	access_token_domain "bookstore_oauth-api/domains/access_token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByID(c *gin.Context)
}

type accessTokenHandler struct {
	accessTokenService access_token_domain.AccessTokenService
}

func NewHandler(accessTokenService access_token_domain.AccessTokenService) AccessTokenHandler {
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
