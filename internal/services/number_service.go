package services

import (
	"github.com/jmoiron/sqlx"
)

type TenantNumbersService struct {
	db *sqlx.DB
}

func NewTenantNumbersService(db *sqlx.DB) *TenantNumbersService {
	return &TenantNumbersService{
		db: db,
	}
}
