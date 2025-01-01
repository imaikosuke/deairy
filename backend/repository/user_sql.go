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

type user struct{}

func NewUser() User { return user{} }

func (r user) Get(ctx context.Context, db *gorm.DB, id string) (*db_model.User, error) {
	var user db_model.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		if pkgerrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app.ErrUserNotFound
		}
		return nil, errors.Wrap(err)
	}
	return &user, nil
}

func (r user) GetByEmail(ctx context.Context, db *gorm.DB, email string) (*db_model.User, error) {
	var user db_model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if pkgerrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app.ErrUserNotFound
		}
		return nil, errors.Wrap(err)
	}
	return &user, nil
}

func (r user) GetByIDs(ctx context.Context, db *gorm.DB, ids []string) ([]*db_model.User, error) {
	var users []*db_model.User
	if err := db.Where("id IN (?)", ids).Find(&users).Error; err != nil {
		return nil, errors.Wrap(err)
	}
	return users, nil
}

func (r user) Create(ctx context.Context, db *gorm.DB, user *db_model.User) (*db_model.User, error) {
	// 既存のメールアドレスをチェック
	existing, err := r.GetByEmail(ctx, db, user.Email)
	if err != nil && !pkgerrors.Is(err, app.ErrUserNotFound) {
		return nil, errors.Wrap(err)
	}
	if existing != nil {
		return nil, app.ErrEmailAlreadyExists
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := db.Create(user).Error; err != nil {
		return nil, errors.Wrap(err)
	}
	return user, nil
}

func (r user) List(ctx context.Context, db *gorm.DB) ([]*db_model.User, error) {
	var users []*db_model.User
	if err := db.Find(&users).Error; err != nil {
		return nil, errors.Wrap(err)
	}
	return users, nil
}

func (r user) Update(ctx context.Context, db *gorm.DB, id string, input UpdateUserInput) (*db_model.User, error) {
	user, err := r.Get(ctx, db, id)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	// メールアドレスが変更される場合、重複チェック
	if user.Email != input.Email {
		existing, err := r.GetByEmail(ctx, db, input.Email)
		if err != nil && !pkgerrors.Is(err, app.ErrUserNotFound) {
			return nil, errors.Wrap(err)
		}
		if existing != nil {
			return nil, app.ErrEmailAlreadyExists
		}
	}

	user.Name = input.Name
	user.Email = input.Email
	user.UpdatedAt = time.Now()

	if err := db.Save(user).Error; err != nil {
		return nil, errors.Wrap(err)
	}
	return user, nil
}
