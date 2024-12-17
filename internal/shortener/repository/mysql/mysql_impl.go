package mysql

import "context"

func (r *repo) Create(ctx context.Context, url, key string) (string, error) {
	if _, err := r.statements.create.ExecContext(ctx, key, url); err != nil {
		return "", err
	}
	return key, nil
}

func (r *repo) Get(ctx context.Context, id string) (string, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	get := tx.StmtContext(ctx, r.statements.get)
	row := get.QueryRowContext(ctx, id)

	var url string
	if err := row.Scan(&url); err != nil {
		return "", err
	}
	if err = tx.Commit(); err != nil {
		return "", err
	}
	return url, nil

}

// BatchCreate urls are map[short][original]
func (r *repo) BatchCreate(ctx context.Context, urls map[string]string) (map[string]string, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// create transaction specific statement
	create := tx.StmtContext(ctx, r.statements.create)

	for key, value := range urls {
		if _, err = create.ExecContext(ctx, key, value); err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}
