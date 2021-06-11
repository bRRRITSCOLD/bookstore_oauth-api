package users_http_repository

import (
	http_client "bookstore_oauth-api/clients/http"
	users_domain "bookstore_oauth-api/domains/users"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	httpClient := http_client.GetHTTPClient()
	httpmock.ActivateNonDefault(httpClient.GetClient())

	os.Exit(m.Run())
}

func TestLoginUser(t *testing.T) {
	// init
	httpmock.Reset()
	responder := httpmock.NewStringResponder(200, `{"userId":14,"firstName":"Eric2","lastName":"Cartman2","email":"test2@email.com","dateCreated":"2021-05-16T18:01:04Z","status":"active"}`)
	httpmock.RegisterResponder("POST", fmt.Sprintf(UsersApiBaseUrl, UsersApiUsersLoginPostEndpoint), responder)

	uR := usersHTTPRepository{}

	// test
	user, err := uR.LoginUser(users_domain.UserLoginRequest{
		Email:    "test@test.com",
		Password: "password1",
	})

	// validation
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 14, user.UserID)
	assert.EqualValues(t, "Eric2", user.FirstName)
	assert.EqualValues(t, "Cartman2", user.LastName)
	assert.EqualValues(t, "test2@email.com", user.Email)
	assert.EqualValues(t, "active", user.Status)
}

func TestLoginHttpClientError(t *testing.T) {
	// init
	httpmock.Reset()
	responder := httpmock.NewErrorResponder(errors.New("Client Error"))
	httpmock.RegisterResponder("POST", fmt.Sprintf(UsersApiBaseUrl, UsersApiUsersLoginPostEndpoint), responder)

	uR := usersHTTPRepository{}

	// test
	user, err := uR.LoginUser(users_domain.UserLoginRequest{
		Email:    "test@test.com",
		Password: "password1",
	})

	// validation
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "Unable to login user test@test.com", err.Message())
}

func TestLoginInvalidError(t *testing.T) {
	// init
	httpmock.Reset()
	responder, _ := httpmock.NewJsonResponder(http.StatusGatewayTimeout, 1)
	httpmock.RegisterResponder("POST", fmt.Sprintf(UsersApiBaseUrl, UsersApiUsersLoginPostEndpoint), responder)

	uR := usersHTTPRepository{}

	// test
	user, err := uR.LoginUser(users_domain.UserLoginRequest{
		Email:    "test@test.com",
		Password: "password1",
	})

	// validation
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid error response when logging in user test@test.com", err.Message())
}

func TestLoginUserNotFound(t *testing.T) {
	// init
	httpmock.Reset()
	responder := httpmock.NewStringResponder(404, `{"message": "mysql: no rows found", "error": "not_found", "status": 404, "cause": []}`)
	httpmock.RegisterResponder("POST", fmt.Sprintf(UsersApiBaseUrl, UsersApiUsersLoginPostEndpoint), responder)

	uR := usersHTTPRepository{}

	// test
	user, err := uR.LoginUser(users_domain.UserLoginRequest{
		Email:    "test@test.com",
		Password: "password1",
	})

	// validation
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "mysql: no rows found", err.Message())
}

func TestLoginUserInvalidResponse(t *testing.T) {
	// init
	httpmock.Reset()
	responder := httpmock.NewStringResponder(200, `{"userId": "1" }`)
	// responder := httpmock.NewStringResponder(http.StatusOK, )
	httpmock.RegisterResponder("POST", fmt.Sprintf(UsersApiBaseUrl, UsersApiUsersLoginPostEndpoint), responder)

	uR := usersHTTPRepository{}

	// test
	user, err := uR.LoginUser(users_domain.UserLoginRequest{
		Email:    "test@test.com",
		Password: "password1",
	})

	// validation
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "error when trying to unmarshal user data for user test@test.com", err.Message())
}
