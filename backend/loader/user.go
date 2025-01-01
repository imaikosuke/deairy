package loader

import (
	"context"

	app "github.com/imaikosuke/deairy/backend"
	"github.com/imaikosuke/deairy/backend/db_model"
	"github.com/imaikosuke/deairy/backend/pkg/errors"
	"github.com/imaikosuke/deairy/backend/service"
)

func newUserLoader(svc service.User) func(context.Context, []string) ([]*db_model.User, []error) {
	return func(ctx context.Context, userIDs []string) ([]*db_model.User, []error) {
		rowMap, err := svc.GetMapByIDs(ctx, userIDs)
		if err != nil {
			return nil, []error{err}
		}
		users := make([]*db_model.User, len(userIDs))
		errs := make([]error, len(userIDs))
		for i, userID := range userIDs {
			if user, ok := rowMap[userID]; ok {
				users[i] = user
			} else {
				errs[i] = app.ErrUserNotFound
			}
		}
		return users, errs
	}
}

func LoadUsers(ctx context.Context, userIDs []string) ([]*db_model.User, error) {
	rows, err := getLoaders(ctx).UserLoader.LoadAll(ctx, userIDs)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return rows, nil
}
