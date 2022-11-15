package storage

import (
	"database/sql"
)

type SitemapRepository struct {
	db *sql.DB
}

func NewSitemapRepository(db *sql.DB) *SitemapRepository {
	return &SitemapRepository{db: db}
}
