package repository

import "fww-core/internal/entity"

// FindAirport implements Repository.
func (r *repository) FindAirport(city string, province string, iata string) ([]entity.Airport, error) {
	panic("unimplemented")
}
