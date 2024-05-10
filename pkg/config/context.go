package config

import (
	"context"
	"errors"
	"github.com/depender/email-sequence-service/constants"
)

func TODOContext() (context.Context, error) {
	ctx := context.TODO()
	ctx, err := SetConfigInContext(ctx)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func SetConfigInContext(ctx context.Context) (context.Context, error) {
	app := getApplication()
	if app == nil {
		return nil, errors.New("config not initialised")
	}

	conf := app.watcher.GetConfig()
	ctx = context.WithValue(ctx, constants.Config, conf)
	return ctx, nil
}
