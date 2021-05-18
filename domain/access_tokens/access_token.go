package action_tokens_domain

import "time"

const (
	EXPIRATION_TIME = 24
)

type AccessToken struct {
	AccessToken string `json:"accessToken"`
	UserID      int64  `json:"userId"`
	ClientID    int64  `json:"clientId"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(EXPIRATION_TIME * time.Hour).Unix(),
	}
}

func (at AccessToken) IsAccessTokenExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)

	return expirationTime.Before(now)
}
