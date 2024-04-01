package db

import (
	"database/sql"
	"fmt"
	"surl-server/internal/config"

	_ "github.com/lib/pq"
)

func Connect(c *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DbHost,
		c.DbPort,
		c.DbUser,
		c.DbPass,
		c.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
