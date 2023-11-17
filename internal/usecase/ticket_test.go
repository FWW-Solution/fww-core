package usecase_test

import (
	"database/sql"
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

		// entityTicket := entity.Ticket{
		// 	ID:                 1,
		// 	CodeTicket:         "",
		// 	IsBoardingPass:     false,
		// 	IsEligibleToFlight: false,
		// 	BookingID:          1,
		// }

		// entityTicketUpdate := entity.Ticket{
		// 	ID:                 1,
		// 	CodeTicket:         "",
		// 	IsBoardingPass:     false,
		// 	IsEligibleToFlight: false,
		// 	BookingID:          1,
		// }

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

		result, err := uc.RedeemTicket(codeBookingUUID)

		assert.NoError(t, err)
		assert.Equal(t, expectBordingTicket, result.BordingTime)

	})
}
