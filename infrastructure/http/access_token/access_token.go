package access_token_http_infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	access_token_domain "bookstore_oauth-api/domains/access_token"
	access_token_service "bookstore_oauth-api/services/access_token"
	errors_utils "bookstore_oauth-api/utils/errors"
)

type AccessTokenHandler interface {
	GetAccessTokenByID(c *gin.Context)
	CreateAccessToken(c *gin.Context)
	// UpdateAccessTokenExpiresByID(c *gin.Context)
}

type accessTokenHandler struct {
	accessTokenService access_token_service.AccessTokenService
}

func NewHandler(accessTokenService access_token_service.AccessTokenService) AccessTokenHandler {
	return &accessTokenHandler{
		accessTokenService: accessTokenService,
	}
}

func (atHandler accessTokenHandler) GetAccessTokenByID(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("accessTokenId"))

	accessToken, err := atHandler.accessTokenService.GetAccessTokenByID(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (atHandler accessTokenHandler) CreateAccessToken(c *gin.Context) {
	var accessToken access_token_domain.AccessToken
	if shouldBindJSONErr := c.ShouldBindJSON(&accessToken); shouldBindJSONErr != nil {
		apiError := errors_utils.NewBadRequestAPIError("invalid json body")
		c.JSON(apiError.Status, apiError)
		return
	}

	if createAccessTokenErr := atHandler.accessTokenService.CreateAccessToken(accessToken); createAccessTokenErr != nil {
		c.JSON(createAccessTokenErr.Status, createAccessTokenErr)
		return
	}

	c.JSON(http.StatusCreated, accessToken)
}

// func (atHandler accessTokenHandler) UpdateAccessTokenExpiresByID(c *gin.Context) {
// 	accessTokenId := strings.TrimSpace(c.Param("accessTokenId"))

// 	accessToken, err := atHandler.accessTokenService.GetAccessTokenByID(accessTokenId)
// 	if err != nil {
// 		c.JSON(err.Status, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, accessToken)
// }
