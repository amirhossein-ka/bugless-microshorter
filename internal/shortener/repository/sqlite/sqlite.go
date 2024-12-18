package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"ush/internal/pkg/config"
	"ush/internal/shortener/repository"

	_ "github.com/mattn/go-sqlite3"
)

type repo struct {
	db         *sql.DB
	statements *statements
}

// you may ask why I did this. here is the answer:
// I was reading about prepared statements, so I thought: How can I use them ?,
// so I added this struct to initialize all queries and store them somewhere that I can use easily
type statements struct {
	// get (short)
	get *sql.Stmt
	// create (short, original)
	create *sql.Stmt
}

func (s *statements) initialize(db *sql.DB) error {
	var err error
	s.get, err = db.Prepare(`SELECT original_url FROM urls WHERE short_url = ?`)
	if err != nil {
		return err
	}

	s.create, err = db.Prepare(`INSERT INTO urls (short_url, original_url) VALUES (?, ?)`)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) createTables(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
			CREATE TABLE IF NOT EXISTS urls (
				id INT AUTO_INCREMENT PRIMARY KEY,
				short_url varchar(10) UNIQUE NOT NULL,
				original_url TEXT NOT NULL
			)
`)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Stop(ctx context.Context) error {
	var errs []error
	if err := r.statements.get.Close(); err != nil {
		errs = append(errs, err)
	}

	if err := r.statements.create.Close(); err != nil {
		errs = append(errs, err)
	}

	if err := r.db.Close(); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors during cleanup: %v", errs)
	}

	return nil
}

func New(cfg *config.ShortenerConfig) (repository.Repository, error) {
	db, err := sql.Open("sqlite", cfg.DSN)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	r := &repo{
		db:         db,
		statements: &statements{},
	}

	// prepare sql queries
	if err = r.statements.initialize(db); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if err := r.createTables(ctx); err != nil {
		return nil, err
	}

	return r, nil
}
