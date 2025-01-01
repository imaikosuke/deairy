package service

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/imaikosuke/deairy/backend/db_model"
	"github.com/imaikosuke/deairy/backend/graph/model"
	"github.com/imaikosuke/deairy/backend/pkg/ulid"
	"github.com/imaikosuke/deairy/backend/repository"
)

type User struct {
	db       *gorm.DB
	userRepo repository.User
}

func NewUser(db *gorm.DB, userRepo repository.User) User {
	return User{
		db:       db,
		userRepo: userRepo,
	}
}

func (s *User) Get(ctx context.Context, id string) (*db_model.User, error) {
	return s.userRepo.Get(ctx, s.db, id)
}

func (s *User) GetByEmail(ctx context.Context, email string) (*db_model.User, error) {
	return s.userRepo.GetByEmail(ctx, s.db, email)
}

func (s *User) GetMapByIDs(ctx context.Context, ids []string) (map[string]*db_model.User, error) {
	users, err := s.userRepo.GetByIDs(ctx, s.db, ids)
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]*db_model.User)
	for _, user := range users {
		userMap[user.ID] = user
	}
	return userMap, nil
}

func (s *User) List(ctx context.Context) ([]*db_model.User, error) {
	return s.userRepo.List(ctx, s.db)
}

func (s *User) Create(ctx context.Context, input model.CreateUserInput) (*db_model.User, error) {
	user := &db_model.User{
		ID:        ulid.Make(),
		Name:      input.Name,
		Email:     input.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.userRepo.Create(ctx, s.db, user)
}

func (s *User) Update(ctx context.Context, id string, input model.UpdateUserInput) (*db_model.User, error) {
	return s.userRepo.Update(ctx, s.db, id, repository.UpdateUserInput{
		Name:  input.Name,
		Email: input.Email,
	})
}
