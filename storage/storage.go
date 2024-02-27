package storage

import (
	"Booking-service/storage/postgres"
	r "Booking-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Booking() r.BookingStorageI
}

type storagePg struct {
	db          *sqlx.DB
	bookingRepo r.BookingStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		bookingRepo: postgres.NewBookingRepo(db),
	}
}

func (s storagePg) Booking() r.BookingStorageI {
	return s.bookingRepo
}
