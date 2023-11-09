package repository

import "github.com/jmoiron/sqlx"

type repository struct {
	db sqlx.DB
}

// FindDetailPassanger implements Repository.
func (*repository) FindDetailPassanger(id int64) (interface{}, error) {
	panic("unimplemented")
}

// RegisterPassanger implements Repository.
func (*repository) RegisterPassanger(data interface{}) (interface{}, error) {
	panic("unimplemented")
}

// UpdatePassanger implements Repository.
func (*repository) UpdatePassanger(data interface{}) (interface{}, error) {
	panic("unimplemented")
}

type Repository interface {
	FindDetailPassanger(id int64) (interface{}, error)
	RegisterPassanger(data interface{}) (interface{}, error)
	UpdatePassanger(data interface{}) (interface{}, error)
}

func New(db sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
