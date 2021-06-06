package access_token_domain

import (
	"fmt"
	"strings"
	"time"

	crypto_utils "github.com/bRRRITSCOLD/bookstore_utils-go/crypto"
	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
)

const (
	EXPIRATION_TIME = 24
)

type AccessToken struct {
	AccessToken string `json:"accessToken" cassandra:"access_token"`
	UserID      int64  `json:"userId" cassandra:"user_id"`
	ClientID    int64  `json:"clientId" cassandra:"client_id"`
	Expires     int64  `json:"expires" cassandra:"expires"`
}

func (at *AccessToken) ValidateAccessToken() *errors_utils.APIError {
	accessTokenId := strings.TrimSpace(at.AccessToken)
	if len(accessTokenId) == 0 {
		return errors_utils.NewBadRequestAPIError("invalid access token id")
	}

	if at.UserID <= 0 {
		return errors_utils.NewBadRequestAPIError("invalid user id")
	}

	if at.ClientID <= 0 {
		return errors_utils.NewBadRequestAPIError("invalid client id")
	}

	if at.Expires <= 0 {
		return errors_utils.NewBadRequestAPIError("invalid expiration time")
	}

	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserID:  userId,
		Expires: time.Now().UTC().Add(EXPIRATION_TIME * time.Hour).Unix(),
	}
}

func (at AccessToken) IsAccessTokenExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)

	return expirationTime.Before(now)
}

func (at *AccessToken) GenerateAccessToken() {
	at.AccessToken = crypto_utils.MD5Hash(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
