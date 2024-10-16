// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserInput struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

type DeleteUserResponse struct {
	DeletedUserID string `json:"deletedUserId"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateUserInput struct {
	Name   *string `json:"name,omitempty"`
	Age    *string `json:"age,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Email  *string `json:"email,omitempty"`
	Phone  *string `json:"phone,omitempty"`
}

type User struct {
	ID     string `json:"_id"`
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}
