package loader

import (
	"context"

	"github.com/vikstrous/dataloadgen"

	"github.com/imaikosuke/deairy/backend/db_model"
	"github.com/imaikosuke/deairy/backend/service"
)

type Loaders struct {
	UserLoader *dataloadgen.Loader[string, *db_model.User]
}

func New(userSvc service.User, diarySvc service.Diary) *Loaders {
	return &Loaders{
		UserLoader: dataloadgen.NewLoader(newUserLoader(userSvc)),
	}
}

type loadersKey struct{}

func (l *Loaders) Attach(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, loadersKey{}, l)
	return ctx
}

func getLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey{}).(*Loaders)
}
