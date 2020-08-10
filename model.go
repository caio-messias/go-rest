package main

// User is a domain model with username, email and wether it is enabled or not
type User struct {
	UserName  string
	Email     string
	IsEnabled bool
}

// UserPayload is the payload received when creating a new user
type UserPayload struct {
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	IsEnabled bool   `json:"isEnabled"`
}

// CreateUserFromPayload maps a UserPayload from a request to a User domain model
func CreateUserFromPayload(userPayload UserPayload) *User {
	return &User{userPayload.UserName, userPayload.Email, userPayload.IsEnabled}
}
