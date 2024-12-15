package repository

import "context"

type (
	Repository interface {
		Create(ctx context.Context, url string) (string, error)
		Get(ctx context.Context, id string) (string, error)
		BatchCreate(ctx context.Context, urls []string) (map[string]string, error)
		BatchGet(ctx context.Context, ids []string) (map[string]string, error)
	}
)
