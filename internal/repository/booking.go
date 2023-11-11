package repository

import "fww-core/internal/entity"

// FindReminingSeat implements Repository.
func (r *repository) FindReminingSeat(flightID int64) (int, error) {
	query := `SELECT reserved_seat, total_seat FROM flight_reservations WHERE flight_id = $1 AND deleted_at IS NULL`
	var result entity.FlightReservation
	err := r.db.QueryRowx(query, flightID).StructScan(&result)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return result.TotalSeat - result.ReservedSeat, nil
}

// GetBookingByBookingIDCode implements Repository.
func (r *repository) FindBookingByBookingIDCode(bookingIDCode string) (entity.Booking, error) {
	query := `SELECT id, code_booking, booking_date, payment_expired_at, booking_status, case_id, user_id, flight_id FROM bookings WHERE code_booking = $1 AND deleted_at IS NULL`
	var result entity.Booking
	err := r.db.QueryRowx(query, bookingIDCode).StructScan(&result)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return entity.Booking{}, nil
	}
	if err != nil {
		return entity.Booking{}, err
	}
	return result, nil
}

// InsertBooking implements Repository.
func (r *repository) InsertBooking(data *entity.Booking) (int64, error) {
	query := `INSERT INTO bookings (code_booking, booking_date, payment_expired_at, booking_status, case_id, user_id, flight_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var id int64

	// do sqlx transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	// insert booking
	err = tx.QueryRowx(query, data.CodeBooking, data.BookingDate, data.PaymentExpiredAt, data.BookingStatus, data.CaseID, data.UserID, data.FlightID).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return id, nil
}

// InsertBookingDetail implements Repository.
func (r *repository) InsertBookingDetail(data *entity.BookingDetail) (int64, error) {
	query := `INSERT INTO booking_details (passenger_id, seat_number, baggage_capacity, class, booking_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int64

	// do sqlx transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	// insert booking
	err = tx.QueryRowx(query, data.PassengerID, data.SeatNumber, data.BaggageCapacity, data.Class, data.BookingID).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return id, nil
}

// UpdateFlightReservation implements Repository.
func (r *repository) UpdateFlightReservation(data *entity.FlightReservation) (int64, error) {
	// Query Postgres Select for update
	querySelect := `SELECT id, class, reserved_seat, total_seat FROM flight_reservations WHERE flight_id = $1 AND class = $2 AND deleted_at IS NULL FOR UPDATE`

	queryUpdate := `UPDATE flight_reservations SET reserved_seat = $1, updated_at = $2 WHERE flight_id = $3 RETURNING id`

	var id int64
	var result entity.FlightReservation

	// do sqlx transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	// get flight reservation
	err = tx.QueryRowx(querySelect, data.FlightID, data.Class).StructScan(&result)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// update booking reservation
	err = tx.QueryRowx(queryUpdate, data.ReservedSeat, data.UpdatedAt, data.FlightID).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return id, nil
}

// FindBookingDetailByBookingID implements Repository.
func (r *repository) FindBookingDetailByBookingID(bookingID int64) ([]entity.BookingDetail, error) {
	query := `SELECT id, passenger_id, seat_number, baggage_capacity, class, booking_id FROM booking_details WHERE booking_id = $1 AND deleted_at IS NULL`

	var result []entity.BookingDetail
	err := r.db.Select(&result, query, bookingID)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return []entity.BookingDetail{}, nil
	}
	if err != nil {
		return []entity.BookingDetail{}, err
	}

	return result, nil
}
