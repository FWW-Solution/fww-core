package usecase_test

import (
	"database/sql"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/data/dto_ticket"
	"fww-core/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRedeemTicket(t *testing.T) {
	setup()
	t.Run("success", func(t *testing.T) {
		codeBookingUUID := "123e4567-e89b-12d3-a456-426614174000"
		// codeTicketUUID := "123e4567-e89b-12d3-a456-2341235r31324"

		ticketID64 := int64(1)

		timePayment := time.Now().Add(time.Hour * 6)

		entityBooking := entity.Booking{
			ID:               1,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now(),
			PaymentExpiredAt: timePayment,
			BookingExpiredAt: timePayment.Add(time.Hour * 6),
			BookingStatus:    "redeemed",
			CaseID:           0,
			CreatedAt:        time.Now(),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		entityBookingDetails := []entity.BookingDetail{
			{
				ID:          1,
				BookingID:   1,
				PassengerID: 1,
				SeatNumber:  "1A",
				Class:       "economy",
			},
		}

		entityPassenger := entity.Passenger{
			ID:       1,
			IDNumber: "1234567890",
			IDType:   "KTP",
			FullName: "John Doe",
		}

		boardingTime := entityBooking.BookingExpiredAt.Add((24 * time.Hour) - (time.Minute * 30))

		expect := dto_ticket.Response{
			BordingTime: boardingTime.Format("2006-01-02 15:04:05"),
		}

		expectBordingTicket := expect.BordingTime

		repositoryMock.On("FindBookingByBookingIDCode", codeBookingUUID).Return(entityBooking, nil)
		// Validate Booking date
		repositoryMock.On("UpdateBooking", &entityBooking).Return(entityBooking.ID, nil)
		// Check Peduli Lindungi
		repositoryMock.On("UpsertTicket", mock.Anything).Return(ticketID64, nil)

		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetails, nil)

		repositoryMock.On("FindDetailPassanger", entityBookingDetails[0].ID).Return(entityPassenger, nil).Once()

		adapterMock.On("RedeemTicket", mock.Anything).Return(nil)

		result, err := uc.RedeemTicket(codeBookingUUID)

		assert.NoError(t, err)
		assert.Equal(t, expectBordingTicket, result.BordingTime)

	})
}

func TestUpdateTicket(t *testing.T) {
	setup()
	t.Run("success", func(t *testing.T) {
		// codeBookingUUID := "123e4567-e89b-12d3-a456-426614174000"
		codeTicketUUID := "123e4567-e89b-12d3-a456-2341235r31324"
		req := dto_ticket.RequestUpdateTicket{
			CodeTicket:         codeTicketUUID,
			BookingDetailID:    1,
			IsEligibleToFlight: false,
		}

		entityTicket := entity.Ticket{
			ID:                 1,
			CodeTicket:         codeTicketUUID,
			IsBoardingPass:     false,
			IsEligibleToFlight: false,
			BookingID:          1,
		}

		specUpdate := dto_booking.BookDetailRequest{
			BookingDetailID:    req.BookingDetailID,
			IsEligibleToFlight: req.IsEligibleToFlight,
		}

		entityBookingDetail := entity.BookingDetail{
			ID:                 1,
			PassengerID:        1,
			SeatNumber:         "A1",
			BaggageCapacity:    20,
			Class:              "economy",
			IsEligibleToFlight: false,
			CreatedAt:          time.Now().Round(time.Minute),
			BookingID:          1,
		}

		entityRequest := &entity.BookingDetail{
			ID:                 specUpdate.BookingDetailID,
			PassengerID:        1,
			SeatNumber:         "A1",
			BaggageCapacity:    20,
			Class:              "economy",
			IsEligibleToFlight: specUpdate.IsEligibleToFlight,
			CreatedAt:          time.Now().Round(time.Minute),
			BookingID:          1,
		}

		repositoryMock.On("FindTicketByCodeTicket", codeTicketUUID).Return(entityTicket, nil)
		repositoryMock.On("FindBookingDetailByID", req.BookingDetailID).Return(entityBookingDetail, nil)
		repositoryMock.On("UpdateBookingDetail", entityRequest).Return(int64(1), nil)

		err := uc.UpdateTicket(&req)
		assert.NoError(t, err)

	})
}
