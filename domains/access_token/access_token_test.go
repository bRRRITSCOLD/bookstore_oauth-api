package access_token_domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	// Init

	// Execution

	// Validation
	assert.EqualValues(t, EXPIRATION_TIME, 24)

	// Teardown

}

func TestGetNewAccessToken(t *testing.T) {
	// Init

	// Execution
	at := GetNewAccessToken(0)

	// Validation
	assert.NotNil(t, at)
	assert.NotNil(t, at.Expires)
	assert.EqualValues(t, at.IsAccessTokenExpired(), false)
	assert.Empty(t, at.AccessToken)
	assert.EqualValues(t, at.UserID, 0)
	assert.EqualValues(t, at.ClientID, 0)

	// Teardown
}

func TestIsAccessTokenExpired(t *testing.T) {
	// Init

	// Execution
	at := AccessToken{}

	// Validation
	assert.NotNil(t, at)
	assert.NotNil(t, at.Expires)
	assert.EqualValues(t, at.IsAccessTokenExpired(), true)

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	assert.EqualValues(t, at.IsAccessTokenExpired(), false)

	// Teardown
}

func TestGenerateAccessToken(t *testing.T) {
	// Init

	// Execution
	at := AccessToken{
		UserID: 0,
	}
	at.GenerateAccessToken()

	// Validation
	assert.NotNil(t, at)
	assert.NotNil(t, at.Expires)
	assert.NotEmpty(t, at.AccessToken)
	assert.EqualValues(t, at.UserID, 0)
	assert.EqualValues(t, at.ClientID, 0)

	// Teardown
}
