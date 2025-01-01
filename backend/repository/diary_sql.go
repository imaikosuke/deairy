package repository

import (
	"context"
	pkgerrors "errors"
	"time"

	"gorm.io/gorm"

	app "github.com/imaikosuke/deairy/backend"
	"github.com/imaikosuke/deairy/backend/db_model"
	"github.com/imaikosuke/deairy/backend/pkg/errors"
)

type diary struct{}

func NewDiary() Diary { return diary{} }

func (r diary) Get(ctx context.Context, db *gorm.DB, id string) (*db_model.Diary, error) {
	var diary db_model.Diary
	if err := db.Where("id = ?", id).First(&diary).Error; err != nil {
		if pkgerrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app.ErrDiaryNotFound
		}
		return nil, errors.Wrap(err)
	}
	return &diary, nil
}

func (r diary) List(ctx context.Context, db *gorm.DB, userID *string, visibility *string) ([]*db_model.Diary, error) {
	query := db.Model(&db_model.Diary{})

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	if visibility != nil {
		query = query.Where("visibility = ?", *visibility)
	}

	var diaries []*db_model.Diary
	if err := query.Find(&diaries).Error; err != nil {
		return nil, errors.Wrap(err)
	}
	return diaries, nil
}

func (r diary) Create(ctx context.Context, db *gorm.DB, diary *db_model.Diary) (*db_model.Diary, error) {
	diary.CreatedAt = time.Now()
	diary.UpdatedAt = time.Now()

	if err := db.Create(diary).Error; err != nil {
		return nil, errors.Wrap(err)
	}
	return diary, nil
}

func (r diary) Update(ctx context.Context, db *gorm.DB, id string, input UpdateDiaryInput) (*db_model.Diary, error) {
	var diary db_model.Diary
	if err := db.Where("id = ?", id).First(&diary).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	diary.Title = input.Title
	diary.Content = input.Content
	diary.Visibility = input.Visibility
	diary.UpdatedAt = time.Now()

	if err := db.Save(&diary).Error; err != nil {
		return nil, errors.Wrap(err)
	}
	return &diary, nil
}

func (r diary) Delete(ctx context.Context, db *gorm.DB, id string) error {
	if err := db.Delete(&db_model.Diary{}, "id = ?", id).Error; err != nil {
		return errors.Wrap(err)
	}
	return nil
}
