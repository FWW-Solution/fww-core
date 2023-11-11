package repository

import "fww-core/internal/entity"

// FindReminingSeat implements Repository.
func (*repository) FindReminingSeat(flightID int64) (int, error) {
	panic("unimplemented")
}

// GetBookingByBookingIDCode implements Repository.
func (*repository) FindBookingByBookingIDCode(bookingIDCode string) (entity.Booking, error) {
	panic("unimplemented")
}

// InsertBooking implements Repository.
func (*repository) InsertBooking(data *entity.Booking) (int64, error) {
	panic("unimplemented")
}

// InsertBookingDetail implements Repository.
func (*repository) InsertBookingDetail(data *entity.BookingDetail) (int64, error) {
	panic("unimplemented")
}

// UpdateFlightReservation implements Repository.
func (*repository) UpdateFlightReservation(data *entity.FlightReservation) (int64, error) {
	panic("unimplemented")
}
