package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"fww-core/internal/container/infrastructure/redis"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/entity"
	"time"
)

// RequestBooking implements UseCase.
func (u *useCase) RequestBooking(data *dto_booking.Request, bookingIDCode string) error {
	ctx := context.Background()
	// Check Booking ID Code
	resultBooking, err := u.repository.FindBookingByBookingIDCode(bookingIDCode)
	if err != nil {
		return err
	}

	if resultBooking.ID != 0 {
		return errors.New("booking id code already exist")
	}

	// Check Remining Sea
	flightIDReminingSeat := fmt.Sprintf("flight-%d-seat", data.FlightID)
	result := u.redis.Get(ctx, flightIDReminingSeat)
	if result.Err() != nil {
		resultSeat, err := u.repository.FindReminingSeat(data.FlightID)
		if err != nil {
			return err
		}
		if resultSeat <= 0 {
			return errors.New("no remaning seat")
		}
		u.redis.Set(ctx, flightIDReminingSeat, resultSeat, 0)
	}

	resultSeat, err := result.Int()
	if err != nil {
		return err
	}

	if resultSeat <= 0 {
		return errors.New("no remaning seat")
	}

	flightIDKey := fmt.Sprintf("flight-%d", data.FlightID)

	// Lock Transaction Redis
	rc := redis.InitMutex(flightIDKey)
	redis.LockMutex(rc)
	defer redis.UnlockMutex(rc)
	// Check Remining Seat

	bookingDate := time.Now().Round(time.Minute)
	paymentExpiredAt := time.Now().Add(time.Hour * 24).Round(time.Minute)

	// Insert Booking
	bookingEntity := &entity.Booking{
		CodeBooking:      bookingIDCode,
		BookingDate:      bookingDate,
		PaymentExpiredAt: paymentExpiredAt,
		BookingStatus:    "pending",
		CaseID:           0,
		UserID:           data.UserID,
		FlightID:         data.FlightID,
	}

	bookingID, err := u.repository.InsertBooking(bookingEntity)
	if err != nil {
		return err
	}

	// Update Flight Reservation
	entityReservation := &entity.FlightReservation{
		Class: data.BookDetails[0].Class,
		// ReservedSeat: (172 - resultSeat) + len(data.BookDetails),
		ReservedSeat: (172 - 0) + len(data.BookDetails),
		TotalSeat:    172,
		UpdatedAt: sql.NullTime{
			Time:  time.Now().Round(time.Minute),
			Valid: true,
		},
		FlightID: data.FlightID,
	}
	_, err = u.repository.UpdateFlightReservation(entityReservation)
	if err != nil {
		return err
	}

	// Insert Booking Detail
	for _, v := range data.BookDetails {

		entityBookingDetail := entity.BookingDetail{
			BookingID:      bookingID,
			PassengerID:    v.PassangerID,
			SeatNumber:     v.SeatNumber,
			BagageCapacity: v.Baggage,
			Class:          v.Class,
		}

		_, err := u.repository.InsertBookingDetail(&entityBookingDetail)
		if err != nil {
			return err
		}
	}

	return err

}
