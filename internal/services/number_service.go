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

type EstateNumbersService struct {
	db *sqlx.DB
}

func NewEstateNumbersService(db *sqlx.DB) *EstateNumbersService {
	return &EstateNumbersService{
		db: db,
	}
}
