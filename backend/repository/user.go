package repository

import (
	"context"

	"github.com/imaikosuke/deairy/backend/db_model"
	"gorm.io/gorm"
)

type User interface {
	Get(ctx context.Context, db *gorm.DB, id string) (*db_model.User, error)
	GetByEmail(ctx context.Context, db *gorm.DB, email string) (*db_model.User, error)
	GetByIDs(ctx context.Context, db *gorm.DB, ids []string) ([]*db_model.User, error)
	Create(ctx context.Context, db *gorm.DB, user *db_model.User) (*db_model.User, error)
	List(ctx context.Context, db *gorm.DB) ([]*db_model.User, error)
	Update(ctx context.Context, db *gorm.DB, id string, input UpdateUserInput) (*db_model.User, error)
}

type UpdateUserInput struct {
	Name  string
	Email string
}
