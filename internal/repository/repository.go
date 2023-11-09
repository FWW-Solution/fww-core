package repository

import (
	"fww-core/internal/data/dto_passanger"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db sqlx.DB
}

// FindDetailPassanger implements Repository.
func (*repository) FindDetailPassanger(id int64) (dto_passanger.ResponseDetail, error) {
	panic("unimplemented")
}

// RegisterPassanger implements Repository.
func (r *repository) RegisterPassanger(data *dto_passanger.RequestRegister) (int64, error) {
	panic("unimplemented")
}

// UpdatePassanger implements Repository.
func (*repository) UpdatePassanger(data *dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error) {
	panic("unimplemented")
}

type Repository interface {
	FindDetailPassanger(id int64) (dto_passanger.ResponseDetail, error)
	RegisterPassanger(data *dto_passanger.RequestRegister) (int64, error)
	UpdatePassanger(data *dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error)
}

func New(db sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
