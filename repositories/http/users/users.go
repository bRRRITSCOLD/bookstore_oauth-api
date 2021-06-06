package users_http_repository

import (
	http_client "bookstore_oauth-api/clients/http"
	users_domain "bookstore_oauth-api/domains/users"
	"encoding/json"
	"fmt"
	"net/http"

	errors_utils "github.com/bRRRITSCOLD/bookstore_utils-go/errors"
)

const (
	UsersApiBaseUrl                = "http://localhost:3000%s"
	UsersApiUsersLoginPostEndpoint = "/users/login"
)

type UsersHTTPRepository interface {
	LoginUser(users_domain.UserLoginRequest) (*users_domain.User, *errors_utils.APIError)
}

type usersHTTPRepository struct {
}

func NewUsersHTTPRepository() UsersHTTPRepository {
	return &usersHTTPRepository{}
}

func (uhr *usersHTTPRepository) LoginUser(loginRequest users_domain.UserLoginRequest) (*users_domain.User, *errors_utils.APIError) {
	client := http_client.GetHTTPClient()

	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetBody(&loginRequest).
		Post(fmt.Sprintf(UsersApiBaseUrl, UsersApiUsersLoginPostEndpoint))
	if err != nil {
		return nil, &errors_utils.APIError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to login user %s", loginRequest.Email),
		}
	}

	body := resp.Body()

	if resp.StatusCode() > 299 {
		var apiErr errors_utils.APIError
		if body != nil {
			err := json.Unmarshal(body, &apiErr)
			if err != nil {
				return nil, &errors_utils.APIError{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("invalid error response when logging in user %s", loginRequest.Email),
				}
			}
			s := string(body)
			fmt.Println(s) // ABC€
			return nil, &apiErr
		}
	}

	var result users_domain.User
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, &errors_utils.APIError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal user data for user %s", loginRequest.Email),
		}
	}

	return &result, nil
}
