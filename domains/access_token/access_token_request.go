package access_token_domain

import (
	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
)

const (
	GRANT_TYPE_PASSWORD           = "password"
	GRANT_TYPE_CLIENT_CREDENTIALS = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// used for password rant_type
	Username string `json:"username"`
	Password string `json:"password"`

	// used for client_credentials rant_type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (atr *AccessTokenRequest) ValidateAccessTokenRequest() errors_utils.APIError {
	switch atr.GrantType {
	case GRANT_TYPE_CLIENT_CREDENTIALS:
	case GRANT_TYPE_PASSWORD:
		break
	default:
		return errors_utils.NewBadRequestAPIError("invalid grant_type", nil)
	}

	// TODO: validate all parameters

	return nil
}
