package repository

import (
	"fww-core/internal/entity"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db sqlx.DB
}

// FindDetailPassanger implements Repository.
func (r *repository) FindDetailPassanger(id int64) (entity.Passenger, error) {
	result, err := r.db.Queryx("SELECT  FROM passanger WHERE id = ?", id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return entity.Passenger{}, nil
	}

	if err != nil {
		return entity.Passenger{}, err
	}

	// hanldle entity
	var row entity.Passenger
	for result.Next() {
		err := result.StructScan(&row)
		if err != nil {
			return entity.Passenger{}, err
		}
	}

	return row, nil

}

// RegisterPassanger implements Repository.
func (r *repository) RegisterPassanger(data *entity.Passenger) (int64, error) {
	query := `INSERT INTO passengers (full_name, gender, date_of_birth, id_number, id_type, covid_vaccine_status, is_id_verified, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id`

	var id int64
	err := r.db.QueryRowx(query, data.FullName, data.Gender, data.DateOfBirth, data.IDNumber, data.IDType, data.CovidVaccineStatus, data.IsIDVerified, data.CreatedAt, data.UpdatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdatePassanger implements Repository.
func (r *repository) UpdatePassanger(data *entity.Passenger) (int64, error) {
	query := `UPDATE passengers 
SET full_name = $1, gender = $2, date_of_birth = $3, id_number = $4, id_type = $5, covid_vaccine_status = $6, is_id_verified = $7, updated_at = $8
WHERE id = $9
RETURNING id`

	var id int64
	err := r.db.QueryRowx(query, data.FullName, data.Gender, data.DateOfBirth, data.IDNumber, data.IDType, data.CovidVaccineStatus, data.IsIDVerified, data.UpdatedAt, data.ID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

type Repository interface {
	FindDetailPassanger(id int64) (entity.Passenger, error)
	RegisterPassanger(data *entity.Passenger) (int64, error)
	UpdatePassanger(data *entity.Passenger) (int64, error)
}

func New(db sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
