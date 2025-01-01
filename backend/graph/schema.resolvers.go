package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/imaikosuke/deairy/backend/db_model"
	"github.com/imaikosuke/deairy/backend/graph/model"
	"github.com/imaikosuke/deairy/backend/pkg/slices"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	user, err := r.UserSvc.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return model.FormatUserResponse(user), nil
}

// CreateDiary is the resolver for the createDiary field.
func (r *mutationResolver) CreateDiary(ctx context.Context, input model.CreateDiaryInput) (*model.Diary, error) {
	panic(fmt.Errorf("not implemented: CreateDiary - createDiary"))
}

// UpdateDiary is the resolver for the updateDiary field.
func (r *mutationResolver) UpdateDiary(ctx context.Context, id string, input model.UpdateDiaryInput) (*model.Diary, error) {
	panic(fmt.Errorf("not implemented: UpdateDiary - updateDiary"))
}

// DeleteDiary is the resolver for the deleteDiary field.
func (r *mutationResolver) DeleteDiary(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteDiary - deleteDiary"))
}

// Diary is the resolver for the diary field.
func (r *queryResolver) Diary(ctx context.Context, id string) (*model.Diary, error) {
	panic(fmt.Errorf("not implemented: Diary - diary"))
}

// Diaries is the resolver for the diaries field.
func (r *queryResolver) Diaries(ctx context.Context, userID *string, visibility *model.DiaryVisibility) ([]*model.Diary, error) {
	panic(fmt.Errorf("not implemented: Diaries - diaries"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.UserSvc.List(ctx)
	if err != nil {
		return nil, err
	}
	return slices.Map(users, func(user *db_model.User) *model.User {
		return model.FormatUserResponse(user)
	}), nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }