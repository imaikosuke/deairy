package model

import (
	"time"

	"github.com/imaikosuke/deairy/backend/db_model"
)

func FormatUserResponse(row *db_model.User) *User {
	return &User{
		ID:    row.ID,
		Name:  row.Name,
		Email: row.Email,
	}
}

func FormatDiaryResponse(row *db_model.Diary) *Diary {
	if row == nil {
		return nil
	}
	return &Diary{
		ID:         row.ID,
		Title:      row.Title,
		Content:    row.Content,
		Visibility: DiaryVisibility(row.Visibility),
		CreatedAt:  row.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  row.UpdatedAt.Format(time.RFC3339),
	}
}
