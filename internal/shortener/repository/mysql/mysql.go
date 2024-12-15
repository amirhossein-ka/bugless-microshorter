package mysql

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
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

	return &repo{
		db: db,
	}, nil
}
