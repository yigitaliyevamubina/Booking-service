package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type bookingRepo struct {
	db *sqlx.DB
}

func NewBookingRepo(db *sqlx.DB) *bookingRepo {
	
	return &bookingRepo{db: db}
}
