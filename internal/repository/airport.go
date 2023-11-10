package repository

import (
	"fww-core/internal/entity"
)

// FindAirport implements Repository.
func (r *repository) FindAirport(city string, province string, iata string) ([]entity.Airport, error) {
	var result []entity.Airport

	query := `SELECT id, city, province, name, iata, icao, created_at, updated_at FROM airports WHERE deleted_at IS NULL`

	if city != "" {
		query += " AND city LIKE '" + city + "'"
	}

	if province != "" {
		query += " AND province LIKE '" + province + "'"
	}

	if iata != "" {
		query += " AND iata LIKE '" + iata + "'"
	}

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row entity.Airport
		err := rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	return result, nil
}
