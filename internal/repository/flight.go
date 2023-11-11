package repository

import (
	"fmt"
	"fww-core/internal/entity"
)

// FindFlightByID implements Repository.
func (r *repository) FindFlightByID(id int64) (entity.Flight, error) {
	query := `SELECT id, code_flight, departure_time, arrival_time, arrival_airport_name, departure_airport_name, departure_airport_id, arrival_airport_id, created_at, updated_at FROM flights WHERE id = $1 AND deleted_at IS NULL`
	var result entity.Flight
	err := r.db.QueryRowx(query, id).StructScan(&result)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return entity.Flight{}, nil
	}
	if err != nil {
		return entity.Flight{}, err
	}
	return result, nil
}

// FindFlightPriceByID implements Repository.
func (r *repository) FindFlightPriceByID(id int64) (entity.FlightPrice, error) {
	query := `SELECT id, flight_id, price, created_at, updated_at FROM flight_prices WHERE flight_id = $1 AND deleted_at IS NULL`
	var result entity.FlightPrice
	err := r.db.QueryRowx(query, id).StructScan(&result)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return entity.FlightPrice{}, nil
	}
	if err != nil {
		return entity.FlightPrice{}, err
	}
	return result, nil
}

// FindFlightReservationByID implements Repository.
func (r *repository) FindFlightReservationByID(flightID int64) (entity.FlightReservation, error) {
	query := `SELECT id, class, reserved_seat, total_seat, flight_id, created_at, updated_at FROM flight_reservations WHERE flight_id = $1 AND deleted_at IS NULL`
	var result entity.FlightReservation
	err := r.db.QueryRowx(query, flightID).StructScan(&result)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return entity.FlightReservation{}, nil
	}
	if err != nil {
		return entity.FlightReservation{}, err
	}
	return result, nil
}

// FindFlights implements Repository.
func (r *repository) FindFlights(departureTime string, arrivalTime string, limit int, offset int) ([]entity.Flight, error) {
	var result []entity.Flight

	query := `SELECT id, code_flight, departure_time, arrival_time, arrival_airport_name, departure_airport_name, departure_airport_id, arrival_airport_id, created_at, updated_at FROM flights WHERE deleted_at IS NULL`

	if departureTime != "" {
		query += fmt.Sprintf(" AND DATE(departure_time) = '%s'", departureTime)
	}

	if arrivalTime != "" {
		query += fmt.Sprintf(" AND DATE(arrival_time) = '%s'", arrivalTime)
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row entity.Flight
		err := rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	return result, nil
}
