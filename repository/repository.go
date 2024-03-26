package repository

import (
	"RestGolang/models"
	"context"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertGreet(ctx context.Context, greet *models.Greet) error
	GetGreetById(ctx context.Context, id string) (*models.Greet, error)
	UpdateGreet(ctx context.Context, greet *models.Greet) error
	DeleteGreet(ctx context.Context, id string, userId string) error
	ListGreet(ctx context.Context, page uint64) ([]*models.Greet, error)
	Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func Close() error {
	return implementation.Close()
}

func InsertGreet(ctx context.Context, greet *models.Greet) error {
	return implementation.InsertGreet(ctx, greet)
}

func GetGreetById(ctx context.Context, id string) (*models.Greet, error) {
	return implementation.GetGreetById(ctx, id)
}

func UpdateGreet(ctx context.Context, post *models.Greet) error {
	return implementation.UpdateGreet(ctx, post)
}

func DeleteGreet(ctx context.Context, id string, userId string) error {
	return implementation.DeleteGreet(ctx, id, userId)
}

func ListGreet(ctx context.Context, page uint64) ([]*models.Greet, error) {
	return implementation.ListGreet(ctx, page)
}
