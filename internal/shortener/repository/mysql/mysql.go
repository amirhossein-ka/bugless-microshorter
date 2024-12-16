package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
	"ush/internal/pkg/config"
	"ush/internal/shortener/repository"
)

type repo struct {
	db *sql.DB
}

func New(cfg *config.ShortenerConfig) (repository.Repository, error) {
	m := mysql.Config{
		User:   cfg.DBUser,
		Passwd: cfg.DBPassword,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort),
		DBName: cfg.DBName,
	}
	dsn := m.FormatDSN()

	db, err := sql.Open(dsn, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	r := &repo{
		db: db,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if err := r.createTables(ctx); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *repo) createTables(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
CREATE TABLE IF NOT EXISTS urls (
    id INT AUTO_INCREMENT PRIMARY KEY,
    short_url varchar(10) UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    access_count INT DEFAULT 0,
    INDEX (created_at)
)
`)
	if err != nil {
		return err
	}

	return nil
}
