package graph

import (
	"gqlgen-todos/api/services"
)

type Resolver struct {
	UserService services.UserService
}
