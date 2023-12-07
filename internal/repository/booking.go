package repository

import (
	"fmt"
	"fww-core/internal/entity"
	"log"
)

// FindReminingSeat implements Repository.
func (r *repository) FindReminingSeat(flightID int64) (int, error) {
	// query := `SELECT remining_seat, total_seat FROM flight_reservations WHERE flight_id = $1 AND deleted_at IS NULL`
	query := fmt.Sprintf(`SELECT remining_seat, total_seat FROM flight_reservations WHERE flight_id = %d AND deleted_at IS NULL`, flightID)
	fmt.Println(query)
	var result entity.FlightReservation
	err := r.db.QueryRowx(query).StructScan(&result)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return result.ReminingSeat, nil
}

// GetBookingByBookingIDCode implements Repository.
func (r *repository) FindBookingByBookingIDCode(bookingIDCode string) (entity.Booking, error) {
	query := `SELECT id, code_booking, booking_date, payment_expired_at, booking_expired_at, booking_status, case_id, user_id, flight_id FROM bookings WHERE code_booking = $1 AND deleted_at IS NULL`
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
	query := `INSERT INTO bookings (code_booking, booking_date, payment_expired_at, booking_expired_at, booking_status, case_id, user_id, flight_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var id int64

	// do sqlx transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	// insert booking
	err = tx.QueryRowx(query, data.CodeBooking, data.BookingDate, data.PaymentExpiredAt, data.BookingExpiredAt, data.BookingStatus, data.CaseID, data.UserID, data.FlightID).Scan(&id)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
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
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

// UpdateFlightReservation implements Repository.
func (r *repository) UpdateFlightReservation(data *entity.FlightReservation) (int64, error) {
	// // Query Postgres Select for update
	// querySelect := fmt.Sprintf(`SELECT id, class, remining_seat, total_seat FROM flight_reservations WHERE flight_id = %d AND class = %s AND deleted_at IS NULL FOR UPDATE`, data.FlightID, data.Class)

	queryUpdate := fmt.Sprintf(`UPDATE flight_reservations SET remining_seat = %d, updated_at = NOW() WHERE flight_id = %d RETURNING id`, data.ReminingSeat, data.FlightID)

	var id int64
	// var result entity.FlightReservation

	// do sqlx transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	// // get flight reservation
	// err = tx.QueryRowx(querySelect).StructScan(&result)
	// if err != nil {
	// 	err = tx.Rollback()
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// 	return 0, err
	// }

	// update booking reservation
	err = tx.QueryRowx(queryUpdate).Scan(&id)
	if err != nil {
		log.Println(err)
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		log.Print(err)
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
	}

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

// FindBookingByID implements Repository.
func (r *repository) FindBookingByID(id int64) (entity.Booking, error) {
	query := `SELECT id, code_booking, booking_date, payment_expired_at, booking_status, case_id, user_id, flight_id FROM bookings WHERE id = $1 AND deleted_at IS NULL`
	var result entity.Booking
	err := r.db.QueryRowx(query, id).StructScan(&result)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return entity.Booking{}, nil
	}
	if err != nil {
		return entity.Booking{}, err
	}

	return result, nil
}

// UpdateBooking implements Repository.
func (r *repository) UpdateBooking(data *entity.Booking) (int64, error) {
	query := `UPDATE bookings SET case_id = $1, booking_status = $2, updated_at = NOW() WHERE id = $3`

	// do sqlx transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	// update booking
	_, err = tx.Exec(query, data.CaseID, data.BookingStatus, data.ID)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
	}

	return data.ID, nil
}

// FindBookingDetailByID implements Repository.
func (r *repository) FindBookingDetailByID(id int64) (entity.BookingDetail, error) {
	query := `SELECT id, passenger_id, seat_number, baggage_capacity, class, booking_id FROM booking_details WHERE id = $1 AND deleted_at IS NULL`

	var result entity.BookingDetail
	err := r.db.QueryRowx(query, id).StructScan(&result)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return entity.BookingDetail{}, nil
	}
	if err != nil {
		return entity.BookingDetail{}, err
	}

	return result, nil
}

// UpdateBookingDetail implements Repository.
func (r *repository) UpdateBookingDetail(data *entity.BookingDetail) (int64, error) {
	query := `UPDATE booking_details SET seat_number = $1, baggage_capacity = $2, class = $3, is_eligible_to_flight = $4, updated_at = NOW() WHERE id = $5 RETURNING id`

	// do sqlx transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	// update booking
	var id int64
	err = tx.QueryRowx(query, data.SeatNumber, data.BaggageCapacity, data.Class, data.IsEligibleToFlight, data.ID).Scan(&id)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return id, nil
}
