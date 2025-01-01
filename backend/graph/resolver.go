package graph

import "github.com/imaikosuke/deairy/backend/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DiarySvc service.Diary
	UserSvc  service.User
}

func NewResolver(
	diaryService service.Diary,
	userService service.User,
) *Resolver {
	return &Resolver{
		DiarySvc: diaryService,
		UserSvc:  userService,
	}
}
