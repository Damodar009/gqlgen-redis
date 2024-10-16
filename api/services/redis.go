package services

import (
	"context"
	"gqlgen-todos/api/repository"
	"gqlgen-todos/graph/model"
)

// ArticleService -> struct
type UserService struct {
	repository repository.UserRepository
}

// NewArticleService  -> creates a new Articleservice
func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

// CreateArticle -> call to create the Article
func (c UserService) CreateUser(ctx context.Context, user model.CreateUserInput) (*model.User, error) {
	return c.repository.Create(ctx, user)
}

// UpdateUser is the resolver for the updateUser field.
func (r UserService) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*model.User, error) {
	return r.repository.UpdateUser(ctx, id, input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r UserService) DeleteUser(ctx context.Context, id string) (*model.DeleteUserResponse, error) {
	return r.repository.DeleteUser(ctx, id)
}

// Users is the resolver for the users field.
func (r UserService) Users(ctx context.Context) ([]*model.User, error) {
	return r.repository.Users(ctx)
}

// User is the resolver for the user field.
func (r UserService) User(ctx context.Context, id string) (*model.User, error) {
	return r.repository.User(ctx, id)
}
