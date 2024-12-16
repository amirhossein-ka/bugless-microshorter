package mysql

import "context"

func (r *repo) Create(ctx context.Context, url, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) Get(ctx context.Context, id string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) BatchCreate(ctx context.Context, urls map[string]string) (map[string]string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) BatchGet(ctx context.Context, ids []string) (map[string]string, error) {
	//TODO implement me
	panic("implement me")
}
