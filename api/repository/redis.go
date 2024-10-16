package repository

import (
	"context"
	"encoding/json"
	"gqlgen-todos/graph/model"
	"gqlgen-todos/infrastructure"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

// AnswerRepository database structure
type UserRepository struct {
	rdb infrastructure.RedisClient
}

// NewAnswerRepository creates a new Answer repository
func NewUserRepository(rdb infrastructure.RedisClient) UserRepository {
	return UserRepository{
		rdb: rdb,
	}
}

// Create Answer
func (r UserRepository) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	id := uuid.New().String()
	user := &model.User{
		ID:     id,
		Name:   input.Name,
		Email:  input.Email,
		Age:    input.Age,
		Phone:  input.Phone,
		Gender: input.Gender,
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = r.rdb.RDB.Set(ctx, "user:"+id, userJSON, 0).Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r UserRepository) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*model.User, error) {
	userJSON, err := r.rdb.RDB.Get(ctx, "user:"+id).Bytes()
	if err != nil {
		return nil, err
	}

	var user model.User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Email != nil {
		user.Email = *input.Email
	}

	updatedUserJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = r.rdb.RDB.Set(ctx, "user:"+id, updatedUserJSON, 0).Err()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r UserRepository) DeleteUser(ctx context.Context, id string) (*model.DeleteUserResponse, error) {
	result, err := r.rdb.RDB.Del(ctx, "user:"+id).Result()
	if err != nil {
		return nil, err
	}

	if result == 0 {
		return &model.DeleteUserResponse{DeletedUserID: id}, nil
	}

	return &model.DeleteUserResponse{DeletedUserID: id}, nil
}

// Users is the resolver for the users field.
func (r UserRepository) Users(ctx context.Context) ([]*model.User, error) {
	keys, err := r.rdb.RDB.Keys(ctx, "user:*").Result()
	if err != nil {
		return nil, err
	}

	var users []*model.User

	for _, key := range keys {
		userJSON, err := r.rdb.RDB.Get(ctx, key).Bytes()
		if err != nil {
			return nil, err
		}

		var user model.User
		err = json.Unmarshal(userJSON, &user)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// User is the resolver for the user field.
func (r UserRepository) User(ctx context.Context, id string) (*model.User, error) {
	userJSON, err := r.rdb.RDB.Get(ctx, "user:"+id).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // User not found
		}
		return nil, err
	}

	var user model.User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
