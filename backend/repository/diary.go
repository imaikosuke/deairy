package repository

import (
	"context"

	"github.com/imaikosuke/deairy/backend/db_model"
	"gorm.io/gorm"
)

type Diary interface {
	Get(ctx context.Context, db *gorm.DB, id string) (*db_model.Diary, error)
	List(ctx context.Context, db *gorm.DB, userID *string, visibility *string) ([]*db_model.Diary, error)
	Create(ctx context.Context, db *gorm.DB, diary *db_model.Diary) (*db_model.Diary, error)
	Update(ctx context.Context, db *gorm.DB, id string, input UpdateDiaryInput) (*db_model.Diary, error)
	Delete(ctx context.Context, db *gorm.DB, id string) error
}

type UpdateDiaryInput struct {
	Title      string
	Content    string
	Visibility string
	Tags       []string
}
