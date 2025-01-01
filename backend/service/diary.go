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

type Diary struct {
	db        *gorm.DB
	diaryRepo repository.Diary
}

func NewDiary(db *gorm.DB, diaryRepo repository.Diary) Diary {
	return Diary{
		db:        db,
		diaryRepo: diaryRepo,
	}
}

func (s *Diary) Get(ctx context.Context, id string) (*db_model.Diary, error) {
	return s.diaryRepo.Get(ctx, s.db, id)
}

func (s *Diary) List(ctx context.Context, userID *string, visibility *string) ([]*db_model.Diary, error) {
	return s.diaryRepo.List(ctx, s.db, userID, visibility)
}

func (s *Diary) Create(ctx context.Context, input model.CreateDiaryInput, userID string) (*db_model.Diary, error) {
	diary := &db_model.Diary{
		ID:         ulid.Make(),
		UserID:     userID,
		Title:      input.Title,
		Content:    input.Content,
		Visibility: string(input.Visibility),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return s.diaryRepo.Create(ctx, s.db, diary)
}

func (s *Diary) Update(ctx context.Context, id string, input model.UpdateDiaryInput) (*db_model.Diary, error) {
	return s.diaryRepo.Update(ctx, s.db, id, repository.UpdateDiaryInput{
		Title:      input.Title,
		Content:    input.Content,
		Visibility: string(input.Visibility),
		Tags:       input.Tags,
	})
}

func (s *Diary) Delete(ctx context.Context, id string) error {
	return s.diaryRepo.Delete(ctx, s.db, id)
}
