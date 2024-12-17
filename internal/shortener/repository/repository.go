package repository

import "context"

type (
	Repository interface {
		Create(ctx context.Context, url, key string) (string, error)
		Get(ctx context.Context, id string) (string, error)
		BatchCreate(ctx context.Context, urls map[string]string) (map[string]string, error)

		Stop(ctx context.Context) error
	}
)
