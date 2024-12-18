package sqlite

import (
	"context"
)

func (r *repo) Create(ctx context.Context, url, key string) (string, error) {
	_, err := r.statements.create.ExecContext(ctx, key, url)
	if err != nil {
		return "", err
	}

	return key, err
}

func (r *repo) Get(ctx context.Context, id string) (string, error) {
	row, err := r.statements.get.QueryContext(ctx)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var res string
	if err := row.Scan(&res); err != nil {
		return "", err
	}

	return res, nil
}

func (r *repo) BatchCreate(ctx context.Context, urls map[string]string) (map[string]string, error) {
	//TODO implement me
	panic("implement me")
}
