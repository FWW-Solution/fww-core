package repository

import (
	"fww-core/internal/entity"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

type Repository interface {
	FindDetailPassanger(id int64) (entity.Passenger, error)
	RegisterPassanger(data *entity.Passenger) (int64, error)
	UpdatePassanger(data *entity.Passenger) (int64, error)

	// Airport
	FindAirport(city string, province string, iata string) ([]entity.Airport, error)
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
