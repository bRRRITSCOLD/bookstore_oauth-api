package users_domain

import "time"

type User struct {
	UserID      int64     `json:"userId"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"dateCreated"`
	Status      string    `json:"status"`
}
