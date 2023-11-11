package usecase_test

import (
	"database/sql"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRequestBooking(t *testing.T) {
	setup()
	t.Run("success", func(t *testing.T) {
		ID := int64(1)
		bookingIDCode := "123qwe"
		flightID := int64(1)
		BookDetails := []dto_booking.BookDetail{
			{
				Baggage:     20,
				SeatNumber:  "A1",
				PassangerID: 1,
				Class:       "Economy",
			},
		}

		req := &dto_booking.Request{
			FlightID:    flightID,
			BookDetails: BookDetails,
			UserID:      1,
		}

		entityBookingNull := entity.Booking{}

		bookingDate := time.Now().Round(time.Minute)
		paymentExpiredAt := time.Now().Add(time.Hour * 24).Round(time.Minute)

		entityBooking := &entity.Booking{
			FlightID:         flightID,
			CodeBooking:      bookingIDCode,
			BookingDate:      bookingDate,
			PaymentExpiredAt: paymentExpiredAt,
			BookingStatus:    "pending",
			CaseID:           0,
			UserID:           1,
		}

		entityBookingDetail := &entity.BookingDetail{
			PassengerID:    1,
			SeatNumber:     "A1",
			BagageCapacity: 20,
			Class:          "Economy",
			BookingID:      ID,
		}

		entityReservation := &entity.FlightReservation{
			FlightID:     flightID,
			Class:        "Economy",
			ReservedSeat: 172 - 0 + len(BookDetails),
			TotalSeat:    172,
			UpdatedAt: sql.NullTime{
				Time:  time.Now().Round(time.Minute),
				Valid: true,
			},
		}

		repositoryMock.On("FindBookingByBookingIDCode", mock.Anything).Return(entityBookingNull, nil)
		redisMock.ExpectGet("flight-1-seat").SetVal("172")
		repositoryMock.On("FindReminingSeat", req.FlightID).Return(172, nil)
		repositoryMock.On("UpdateFlightReservation", entityReservation).Return(ID, nil)
		repositoryMock.On("InsertBooking", entityBooking).Return(ID, nil)
		repositoryMock.On("InsertBookingDetail", entityBookingDetail).Return(ID, nil)

		err := uc.RequestBooking(req, bookingIDCode)

		assert.NoError(t, err)

	})
}
